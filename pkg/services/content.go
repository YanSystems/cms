package services

import (
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/YanSystems/cms/pkg/models"
	repositories "github.com/YanSystems/cms/pkg/repositories/content"
	"github.com/YanSystems/cms/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type ContentService struct {
	DB *mongo.Database
}

func (s *ContentService) HandleCreateContent(w http.ResponseWriter, r *http.Request) {
	coll := chi.URLParam(r, "collection")

	var c models.Content
	err := utils.ReadJSON(w, r, &c)
	if err != nil {
		slog.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	c.Id = uuid.New().String()
	c.UpdatedAt = time.Now().UTC()
	c.CreatedAt = time.Now().UTC()

	if c.Class == "" || c.Title == "" || c.Description == "" || c.Body == "" || c.Views < 0 || c.CreatorId == "" {
		err := errors.New("missing fields in request payload")
		slog.Error("Validation error", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	repo := repositories.ContentRepository{DB: s.DB}
	id, err := repo.CreateContent(coll, &c)
	if err != nil {
		slog.Error("Failed to create content", "error", err)
		utils.ErrorJSON(w, err)
		return
	}

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully created content",
		Data:    id,
	}

	utils.WriteJSON(w, http.StatusCreated, responsePayload)
}

func (s *ContentService) HandleGetContent(w http.ResponseWriter, r *http.Request)       {}
func (s *ContentService) HandleGetCollection(w http.ResponseWriter, r *http.Request)    {}
func (s *ContentService) HandleGetClass(w http.ResponseWriter, r *http.Request)         {}
func (s *ContentService) HandleUpdateContent(w http.ResponseWriter, r *http.Request)    {}
func (s *ContentService) HandleDeleteContent(w http.ResponseWriter, r *http.Request)    {}
func (s *ContentService) HandleDeleteClass(w http.ResponseWriter, r *http.Request)      {}
func (s *ContentService) HandleDeleteCollection(w http.ResponseWriter, r *http.Request) {}
