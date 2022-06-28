package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"entproject/database"
	"entproject/ent"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ErrorData          = "database error"
	ErrorInternal      = "internal server error"
	ErrorNotFound      = "resource not found"
	ErrorBadIdentifier = "invalid service identifier"
	ErrorInvalidPage   = "invalid page"
)

type Server struct {
	db  database.Database
	log *log.Logger
}

type GetServiceReply struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Versions    []int  `json:"versions"`
}

type ListServicesReply struct {
	Services []GetServiceReply `json:"services"`
	Count    int               `json:"count"`
	Prev     string            `json:"prev"`
	Next     string            `json:"next"`
}

type AddServiceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type AddServiceReply struct {
	ID string `json:"id"`
}

// NewServer creates a new server using the given database implementation.
func NewServer(db database.Database, log *log.Logger) *Server {
	return &Server{db: db, log: log}
}

// writeJSONError returns a JSON error on the ResponseWriter
func (s *Server) writeJSONError(w http.ResponseWriter, status int, err string) {
	type JSONError struct {
		Status int    `json:"status"`
		Error  string `json:"error"`
	}
	s.log.Printf("writing error: %v", err)
	s.writeJSON(w, status, JSONError{Status: status, Error: err})
}

// writeJSONError returns a JSON message on the ResponseWriter
func (s *Server) writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		s.log.Printf("error marshaling JSON: %v", err)
		http.Error(w, ErrorInternal, http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	w.Write(b)
}

func (s *Server) ListServicesHandler(w http.ResponseWriter, r *http.Request) {
	filter := r.URL.Query().Get("filter")
	sortBy := r.URL.Query().Get("sort_by")
	orderBy := r.URL.Query().Get("order_by")

	// Offset pagination with sane limits and defaults
	var pageNo, pageSize int
	if r.URL.Query().Get("page") != "" {
		pageNo, _ = strconv.Atoi(r.URL.Query().Get("page"))
	}
	if r.URL.Query().Get("count") != "" {
		pageSize, _ = strconv.Atoi(r.URL.Query().Get("count"))
	}
	if pageSize < 0 || pageSize > 50 {
		pageSize = 10
	}
	totalCount, err := s.db.GetServicesCount(r.Context(), filter)
	if err != nil {
		s.log.Printf("Error getting services count: %v", err)
		s.writeJSONError(w, http.StatusInternalServerError, ErrorInternal)
		return
	}
	// Out of bounds
	if pageSize*pageNo >= totalCount {
		s.writeJSONError(w, http.StatusNotFound, ErrorInvalidPage)
		return
	}

	prev := ""
	if pageNo > 0 {
		prev = fmt.Sprintf("%d", pageNo-1)
	}

	next := ""
	if pageSize*pageNo <= totalCount {
		next = fmt.Sprintf("%d", pageNo+1)
	}

	services, err := s.db.GetServices(r.Context(), pageSize*pageNo, pageSize, sortBy, orderBy, filter)
	if err != nil {
		s.log.Printf("Error getting services: %v", err)
		s.writeJSONError(w, http.StatusInternalServerError, ErrorInternal)
		return
	}

	res := ListServicesReply{
		Services: []GetServiceReply{},
		Count:    totalCount,
		Prev:     prev,
		Next:     next,
	}

	for _, s := range services {
		holder := GetServiceReply{
			ID:          s.ID.String(),
			Name:        s.Name,
			Description: s.Description,
			Versions:    []int{},
		}
		for _, version := range s.Edges.ServiceVersions {
			holder.Versions = append(holder.Versions, version.Version)
		}

		res.Services = append(res.Services, holder)
	}

	s.writeJSON(w, http.StatusOK, res)
}

func (s *Server) AddServiceHandler(w http.ResponseWriter, r *http.Request) {
	var data AddServiceRequest

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		s.writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Perform some basic validation
	if data.Name == "" {
		s.writeJSONError(w, http.StatusBadRequest, "Service name is required")
		return
	}
	if len(data.Name) > 25 {
		s.writeJSONError(w, http.StatusBadRequest, "Service name limited to 25 chars")
		return
	}

	id, err := s.db.AddService(r.Context(), ent.Service{
		Name:        data.Name,
		Description: data.Description,
	})
	if err != nil {
		s.writeJSONError(w, http.StatusInternalServerError, ErrorInternal)
		return
	}

	s.writeJSON(w, http.StatusCreated, AddServiceReply{
		ID: id.String(),
	})
}

func (s *Server) GetServiceHandler(w http.ResponseWriter, r *http.Request) {
	// Request must contain an ID in the path, "/services/<id>".
	path := strings.Trim(r.URL.Path, "/")
	pathParts := strings.Split(path, "/")
	if len(pathParts) < 2 {
		s.writeJSONError(w, http.StatusBadRequest, ErrorBadIdentifier)
		return
	}

	id, err := uuid.Parse(pathParts[1])
	if err != nil {
		s.writeJSONError(w, http.StatusBadRequest, ErrorBadIdentifier)
		return
	}

	service, err := s.db.GetService(r.Context(), id)
	if err != nil {
		if ent.IsNotFound(err) {
			s.writeJSONError(w, http.StatusNotFound, ErrorNotFound)
			return
		}
		s.writeJSONError(w, http.StatusInternalServerError, ErrorInternal)
		return
	}

	if service == nil {
		s.writeJSONError(w, http.StatusNotFound, ErrorNotFound)
		return
	}

	res := GetServiceReply{
		ID:          service.ID.String(),
		Name:        service.Name,
		Description: service.Description,
		Versions:    []int{},
	}

	for _, version := range service.Edges.ServiceVersions {
		res.Versions = append(res.Versions, version.Version)
	}

	s.writeJSON(w, http.StatusOK, res)
}

func main() {
	db := database.InitDb()
	s := NewServer(db, log.Default())

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)

	r.Route("/services", func(r chi.Router) {
		r.Get("/", s.ListServicesHandler)
		r.Post("/", s.AddServiceHandler)

		r.Route("/{serviceID}", func(r chi.Router) {
			r.Get("/", s.GetServiceHandler)
		})
	})

	http.ListenAndServe(":8080", r)
}
