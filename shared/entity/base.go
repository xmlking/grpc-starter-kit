package entity

import (
	"time"

	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
)

// Base contains common columns for all tables.
// CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public; ?
type Base struct {
	ID        *resource.Identifier `orm:"type:uuid;primary_key;"`
	CreatedAt *time.Time           `json:"created_at"`
	UpdatedAt *time.Time           `json:"update_at"`
	DeletedAt *time.Time           `sql:"index" json:"deleted_at"`
}
