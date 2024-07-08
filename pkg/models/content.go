package models

import (
	"time"
)

type ReadContent struct {
	Id          string    `bson:"id" json:"id" validate:"required,uuid"`
	Class       string    `bson:"class" json:"class" validate:"required"`
	Title       string    `bson:"title" json:"title"`
	Description string    `bson:"description" json:"description"`
	Body        string    `bson:"body" json:"body"`
	IsPublic    bool      `bson:"is_public" json:"is_public"`
	Views       int       `bson:"views" json:"views" validate:"gte=0"`
	CreatorId   string    `bson:"creator_id" json:"creator_id" validate:"required,uuid"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at" validate:"required"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at" validate:"required"`
}

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

type UpdateContent struct {
	Class       *string    `bson:"class" json:"class,omitempty"`
	Title       *string    `bson:"title" json:"title,omitempty"`
	Description *string    `bson:"description" json:"description,omitempty"`
	Body        *string    `bson:"body" json:"body,omitempty"`
	IsPublic    *bool      `bson:"is_public" json:"is_public,omitempty"`
	Views       *int       `bson:"views" json:"views,omitempty" validate:"gte=0"`
	CreatorId   *string    `bson:"creator_id" json:"creator_id,omitempty" validate:"uuid"`
	UpdatedAt   *time.Time `bson:"updated_at" json:"updated_at,omitempty"`
	CreatedAt   *time.Time `bson:"created_at" json:"created_at,omitempty"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
