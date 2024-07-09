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
	slog.Debug("HandleCreateContent called")
	coll := chi.URLParam(r, "collection")
	slog.Debug("Collection parameter extracted", "collection", coll)

	var c models.Content
	err := utils.ReadJSON(w, r, &c)
	if err != nil {
		slog.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Debug("JSON request body read successfully", "content", c)

	c.Id = uuid.New().String()
	c.UpdatedAt = time.Now().UTC()
	c.CreatedAt = time.Now().UTC()
	slog.Debug("Generated new UUID and timestamps", "id", c.Id, "created_at", c.CreatedAt, "updated_at", c.UpdatedAt)

	if c.Class == "" || c.Title == "" || c.Description == "" || c.Body == "" || c.Views < 0 || c.CreatorId == "" {
		err := errors.New("missing fields in request payload")
		slog.Error("Validation error", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Request payload validated successfully")

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Creating content...")
	id, err := repo.CreateContent(coll, &c)
	if err != nil {
		slog.Error("Failed to create content", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Content created successfully", "id", id)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully created content",
		Data:    id,
	}

	utils.WriteJSON(w, http.StatusCreated, responsePayload)
	slog.Info("Response sent for HandleCreateContent", "status", http.StatusCreated)
}

func (s *ContentService) HandleGetContent(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleGetContent called")
	coll := chi.URLParam(r, "collection")
	id := chi.URLParam(r, "id")
	slog.Debug("Collection and ID parameters extracted", "collection", coll, "id", id)

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Getting content of id " + id)
	content, err := repo.GetContent(coll, id)
	if err != nil {
		slog.Error("Failed to get content", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Content retrieved successfully", "content", content)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully retrieved content",
		Data:    content,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent for HandleGetContent", "status", http.StatusOK)
}

func (s *ContentService) HandleGetCollection(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleGetCollection called")
	coll := chi.URLParam(r, "collection")
	slog.Debug("Collection parameter extracted", "collection", coll)

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Getting collection...")
	collection, err := repo.GetCollection(coll)
	if err != nil {
		slog.Error("Failed to get collection", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Collection retrieved successfully", "collection", collection)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully retrieved collection",
		Data:    collection,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent for HandleGetCollection", "status", http.StatusOK)
}

func (s *ContentService) HandleGetClass(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleGetClass called")
	coll := chi.URLParam(r, "collection")
	class := chi.URLParam(r, "class")
	slog.Debug("Collection and Class parameters extracted", "collection", coll, "class", class)

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Getting class...")
	contents, err := repo.GetClass(coll, class)
	if err != nil {
		slog.Error("Failed to get class", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Class contents retrieved successfully", "contents", contents)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully retrieved class",
		Data:    contents,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent for HandleGetClass", "status", http.StatusOK)
}

func (s *ContentService) HandleUpdateContent(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleUpdateContent called")
	var c models.UpdateContent
	err := utils.ReadJSON(w, r, &c)
	if err != nil {
		slog.Error("Failed to read JSON request", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Debug("JSON request body read successfully", "content", c)

	coll := chi.URLParam(r, "collection")
	id := chi.URLParam(r, "id")
	slog.Debug("Collection and ID parameters extracted", "collection", coll, "id", id)

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Updating content of id " + id)
	id, err = repo.UpdateContent(coll, id, &c)
	if err != nil {
		slog.Error("Failed to update content", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Content updated successfully", "id", id)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully updated content",
		Data:    id,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent for HandleUpdateContent", "status", http.StatusOK)
}

func (s *ContentService) HandleDeleteContent(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleDeleteContent called")
	coll := chi.URLParam(r, "collection")
	id := chi.URLParam(r, "id")
	slog.Debug("Collection and ID parameters extracted", "collection", coll, "id", id)

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Deleting content of id " + id)
	_, err := repo.DeleteContent(coll, id)
	if err != nil {
		slog.Error("Failed to delete content", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Content deleted successfully", "id", id)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully deleted content",
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent for HandleDeleteContent", "status", http.StatusOK)
}

func (s *ContentService) HandleDeleteClass(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleDeleteClass called")
	coll := chi.URLParam(r, "collection")
	class := chi.URLParam(r, "class")
	slog.Debug("Collection and Class parameters extracted", "collection", coll, "class", class)

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Deleting class...")
	_, err := repo.DeleteClass(coll, class)
	if err != nil {
		slog.Error("Failed to delete class", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Class deleted successfully", "class", class)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully deleted class",
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent for HandleDeleteClass", "status", http.StatusOK)
}

func (s *ContentService) HandleDeleteCollection(w http.ResponseWriter, r *http.Request) {
	slog.Debug("HandleDeleteCollection called")
	coll := chi.URLParam(r, "collection")
	slog.Debug("Collection parameter extracted", "collection", coll)

	repo := repositories.ContentRepository{DB: s.DB}
	slog.Debug("ContentRepository initialized")

	slog.Info("Deleting collection...")
	_, err := repo.DeleteCollection(coll)
	if err != nil {
		slog.Error("Failed to delete collection", "error", err)
		utils.ErrorJSON(w, err)
		return
	}
	slog.Info("Collection deleted successfully", "collection", coll)

	responsePayload := models.JsonResponse{
		Error:   false,
		Message: "Successfully deleted collection",
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)
	slog.Info("Response sent for HandleDeleteCollection", "status", http.StatusOK)
}
