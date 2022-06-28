package database

import (
	"context"
	"entproject/ent"
	"entproject/ent/service"
	"entproject/ent/serviceversion"
	"log"

	"github.com/google/uuid"
)

type Database interface {
	AddService(ctx context.Context, service ent.Service) (uuid.UUID, error)
	GetService(ctx context.Context, id uuid.UUID) (*ent.Service, error)
	GetServices(ctx context.Context, offset int, count int, sortBy string, orderBy string, filter string) ([]ent.Service, error)
	GetServicesCount(ctx context.Context, filter string) (int, error)
}

type EntDatabase struct {
	Client *ent.Client
}

func (d *EntDatabase) AddService(ctx context.Context, service ent.Service) (uuid.UUID, error) {
	new, err := d.Client.Service.Create().SetName(service.Name).SetDescription(service.Description).Save(ctx)
	if err != nil {
		return uuid.UUID{}, err
	}
	return new.ID, nil
}

func (d *EntDatabase) GetService(ctx context.Context, id uuid.UUID) (*ent.Service, error) {
	service, err := d.Client.Service.Query().
		Where(service.ID(id)).WithServiceVersions(func(query *ent.ServiceVersionQuery) {
		query.Select(serviceversion.FieldVersion)
	}).Only(ctx)
	if err != nil {
		return nil, err
	}
	return service, nil
}

// GetServices returns a paginated list of services from the database. Supports optional sorting and filter by valid field names
func (d *EntDatabase) GetServices(ctx context.Context, offset int, count int, sortBy string, orderBy string, filter string) ([]ent.Service, error) {
	// Query builder
	query := d.Client.Service.Query().
		WithServiceVersions(func(query *ent.ServiceVersionQuery) {
			query.Select(serviceversion.FieldVersion)
		})

	// Paginate
	query.Limit(count).Offset(offset)

	// Filter
	if filter != "" {
		query = query.Where(service.Or(
			service.DescriptionContains(filter),
			service.NameContains(filter),
		))
	}
	// Sort - Default by field name. Probably should be date created
	var sortKey string
	switch sortBy {
	case "id":
		sortKey = service.FieldID
	case "desc":
		sortKey = service.FieldDescription
	default:
		sortKey = service.FieldName
	}
	// Order
	switch orderBy {
	case "asc":
		query = query.Order(ent.Asc(sortKey))
	default:
		query = query.Order(ent.Desc(sortKey))
	}

	services, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	var resp []ent.Service
	for _, service := range services {
		resp = append(resp, *service)
	}
	return resp, nil
}

// GetServicesCount returns a total count for the list of services from the database. Supports optional filter
func (d *EntDatabase) GetServicesCount(ctx context.Context, filter string) (int, error) {
	// Query builder
	query := d.Client.Service.Query()

	// Filter
	if filter != "" {
		query = query.Where(service.Or(
			service.DescriptionContains(filter),
			service.NameContains(filter),
		))
	}

	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// InitDb creates a new ent database connection
func InitDb() *EntDatabase {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening db connection: %v", err)
	}

	log.Println("Connected to database")

	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Seed the database for testing
	ctx := context.Background()
	names := []string{"Grafana", "Prometheus", "Lightyear"}
	bulk := make([]*ent.ServiceCreate, len(names))
	for i, name := range names {
		v1 := client.ServiceVersion.Create().SetVersion(1).SetConfig("placeholder").SaveX(ctx)
		v2 := client.ServiceVersion.Create().SetVersion(2).SetConfig("placeholder").SaveX(ctx)
		bulk[i] = client.Service.Create().SetName(name).AddServiceVersions(v1, v2)
	}

	if _, err := client.Service.CreateBulk(bulk...).Save(ctx); err != nil {
		log.Fatalf("Seeding database failed")
	}

	return &EntDatabase{Client: client}
}
