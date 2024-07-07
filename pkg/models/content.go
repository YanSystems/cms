package models

import (
	"time"
)

type Content struct {
	Id          string    `bson:"id" json:"id,omitempty"`
	Class       string    `bson:"class" json:"class,omitempty"`
	Title       string    `bson:"title" json:"title,omitempty"`
	Description string    `bson:"description" json:"description,omitempty"`
	Body        string    `bson:"body" json:"body,omitempty"`
	IsPublic    bool      `bson:"is_public" json:"is_public,omitempty"`
	Views       int       `bson:"views" json:"views,omitempty"`
	CreatorId   string    `bson:"creator_id" json:"creator_id,omitempty"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at,omitempty"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at,omitempty"`
}
