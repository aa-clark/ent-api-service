package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ServiceVersion holds the schema definition for the ServiceVersion entity.
type ServiceVersion struct {
	ent.Schema
}

// Fields of the ServiceVersion.
func (ServiceVersion) Fields() []ent.Field {
	return []ent.Field{
		field.Int("version"),
		field.String("config"),
	}
}

// Edges of the ServiceVersion.
func (ServiceVersion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("service", Service.Type).
			Ref("serviceVersions").
			Unique(),
	}
}
