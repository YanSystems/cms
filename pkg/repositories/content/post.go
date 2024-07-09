package repositories

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/YanSystems/cms/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ContentRepository) CreateContent(coll string, content *models.Content) (string, error) {
	slog.Debug("CreateContent called", "collection", coll, "contentID", content.Id)

	if err := validateContent(content); err != nil {
		slog.Error("Content validation failed", "error", err)
		return "", err
	}
	slog.Info("Content validation passed", "contentID", content.Id)

	var existingContent models.Content
	err := r.DB.Collection(coll).FindOne(context.TODO(), bson.M{"id": content.Id}).Decode(&existingContent)
	if err == nil {
		err := errors.New("content with this ID already exists")
		slog.Error("Content with this ID already exists", "contentID", content.Id, "error", err)
		return "", err
	}
	slog.Debug("No existing content with this ID found", "contentID", content.Id)

	_, err = r.DB.Collection(coll).InsertOne(
		context.TODO(),
		content,
	)
	if err != nil {
		slog.Error("Failed to insert content", "collection", coll, "contentID", content.Id, "error", err)
		return "", err
	}
	slog.Info("Content inserted successfully", "collection", coll, "contentID", content.Id)

	return content.Id, nil
}

func validateUUID(fl validator.FieldLevel) bool {
	_, err := uuid.Parse(fl.Field().String())
	return err == nil
}

func validateContent(content *models.Content) error {
	slog.Debug("Validating content", "contentID", content.Id)
	validate := validator.New()
	validate.RegisterValidation("uuid", validateUUID)
	err := validate.Struct(content)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validationError := fmt.Errorf("field %s is not valid: %v", err.StructField(), err)
			slog.Error("Validation error", "field", err.StructField(), "error", validationError)
			return validationError
		}
	}
	slog.Debug("Content validation passed", "contentID", content.Id)
	return nil
}
