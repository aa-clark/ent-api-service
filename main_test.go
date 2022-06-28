package main

import (
	"bytes"
	"encoding/json"
	"entproject/database"
	"entproject/ent/enttest"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func newMockServer(t *testing.T) *Server {
	// Using a persistent database here is bad practice as it will grow when new services are added
	// TODO: Add proper golden seed data
	// Ent appears to have solutions for auto migrations that could be useful https://entgo.io/blog/2022/03/14/announcing-versioned-migrations/
	client := enttest.Open(t, "sqlite3", "file:test/ent_test.db?_fk=1")
	db := database.EntDatabase{Client: client}
	s := NewServer(&db, log.Default())
	return s
}

func TestAddServiceHandler(t *testing.T) {
	s := newMockServer(t)

	req, err := http.NewRequest("POST", "/service", bytes.NewBuffer([]byte(`{
		"name": "Robots second Service",
		"description": "Very apt description"
	}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(s.AddServiceHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetServiceHandler(t *testing.T) {
	s := newMockServer(t)

	t.Run("Invalid service id", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services/123", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.GetServiceHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("Service not found", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services/780a41d7-01f2-4cdb-abe8-c8d4b1aaa57f", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.GetServiceHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})

	t.Run("Success", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services/0777f355-d215-4ff5-99f7-5f089d7730b2", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.GetServiceHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Assert that response matches service from seed data
	})
}

func TestListServicesHandler(t *testing.T) {
	s := newMockServer(t)

	t.Run("Success with no options", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.ListServicesHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		var data ListServicesReply
		err = json.NewDecoder(rr.Body).Decode(&data)
		if err != nil {
			t.Error("failed to marshall list services response")
		}
		if len(data.Services) == 0 {
			t.Error("No services returned")
		}
	})

	t.Run("Sort", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services?sort=name", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.ListServicesHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		var data ListServicesReply
		err = json.NewDecoder(rr.Body).Decode(&data)
		if err != nil {
			t.Error("failed to marshall list services response")
		}
		if len(data.Services) == 0 {
			t.Error("No services returned")
		}
	})

	t.Run("Filter with results", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services?filter=Grafana", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.ListServicesHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		var data ListServicesReply
		err = json.NewDecoder(rr.Body).Decode(&data)
		if err != nil {
			t.Error("failed to marshall list services response")
		}
		if len(data.Services) == 0 {
			t.Errorf("Should have results, got %v", data.Services)
		}
	})

	t.Run("Filter with no results", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services?filter=DEADBEEF", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.ListServicesHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		var data ListServicesReply
		err = json.NewDecoder(rr.Body).Decode(&data)
		if err != nil {
			t.Error("failed to marshall list services response")
		}
		if len(data.Services) != 0 {
			t.Errorf("Should have no results, got %v", data.Services)
		}
	})

	t.Run("Pagination", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services?count=1&page=1", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.ListServicesHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		var data ListServicesReply
		err = json.NewDecoder(rr.Body).Decode(&data)
		if err != nil {
			t.Error("failed to marshall list services response")
		}
		if len(data.Services) != 1 {
			t.Errorf("Should have page of 1 results, got %v", data.Services)
		}
	})

	t.Run("Page out of bounds defaults to 0", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/services?count=1&page=9999", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(s.ListServicesHandler)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		var data ListServicesReply
		err = json.NewDecoder(rr.Body).Decode(&data)
		if err != nil {
			t.Error("failed to marshall list services response")
		}
		if len(data.Services) != 1 {
			t.Errorf("Should have page of 1 results, got %v", data.Services)
		}
	})
}
