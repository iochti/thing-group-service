package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// ThingGroup represents a group of thing
type ThingGroup struct {
	ID          bson.ObjectId
	AccountID   string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
