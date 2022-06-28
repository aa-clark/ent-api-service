package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("name").MaxLen(25).
			NotEmpty(),
		field.String("description").Optional(),
	}
}

func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("serviceVersions", ServiceVersion.Type),
	}
}
