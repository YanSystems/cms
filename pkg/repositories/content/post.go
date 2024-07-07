package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/YanSystems/cms/pkg/models"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *ContentRepository) CreateContent(coll string, content *models.Content) (string, error) {
	if err := validateContent(content); err != nil {
		return "", err
	}

	var existingContent models.Content
	err := r.DB.Collection(coll).FindOne(context.TODO(), bson.M{"id": content.Id}).Decode(&existingContent)
	if err == nil {
		return "", errors.New("content with this ID already exists")
	}

	_, err = r.DB.Collection(coll).InsertOne(
		context.TODO(),
		content,
	)
	if err != nil {
		return "", err
	}

	return content.Id, nil
}

func validateUUID(fl validator.FieldLevel) bool {
	_, err := uuid.Parse(fl.Field().String())
	return err == nil
}

func validateContent(content *models.Content) error {
	validate := validator.New()
	validate.RegisterValidation("uuid", validateUUID)
	err := validate.Struct(content)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return fmt.Errorf("field %s is not valid: %v", err.StructField(), err)
		}
	}
	return nil
}
