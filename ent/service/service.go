// Code generated by entc, DO NOT EDIT.

package service

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the service type in the database.
	Label = "service"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// EdgeServiceVersions holds the string denoting the serviceversions edge name in mutations.
	EdgeServiceVersions = "serviceVersions"
	// ServiceVersionFieldID holds the string denoting the ID field of the ServiceVersion.
	ServiceVersionFieldID = "id"
	// Table holds the table name of the service in the database.
	Table = "services"
	// ServiceVersionsTable is the table that holds the serviceVersions relation/edge.
	ServiceVersionsTable = "service_versions"
	// ServiceVersionsInverseTable is the table name for the ServiceVersion entity.
	// It exists in this package in order to avoid circular dependency with the "serviceversion" package.
	ServiceVersionsInverseTable = "service_versions"
	// ServiceVersionsColumn is the table column denoting the serviceVersions relation/edge.
	ServiceVersionsColumn = "service_service_versions"
)

// Columns holds all SQL columns for service fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)