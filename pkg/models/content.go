package models

import (
	"time"
)

type Content struct {
	Id          string    `bson:"id" json:"id,omitempty" validate:"required,uuid"`
	Class       string    `bson:"class" json:"class,omitempty" validate:"required"`
	Title       string    `bson:"title" json:"title,omitempty"`
	Description string    `bson:"description" json:"description,omitempty"`
	Body        string    `bson:"body" json:"body,omitempty"`
	IsPublic    bool      `bson:"is_public" json:"is_public,omitempty"`
	Views       int       `bson:"views" json:"views,omitempty" validate:"gte=0"`
	CreatorId   string    `bson:"creator_id" json:"creator_id,omitempty" validate:"required,uuid"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at,omitempty" validate:"required"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at,omitempty" validate:"required"`
}
