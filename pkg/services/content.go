package services

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type ContentService struct {
	DB *mongo.Database
}

func (s *ContentService) HandleCreateContent(w http.ResponseWriter, r *http.Request)    {}
func (s *ContentService) HandleGetContent(w http.ResponseWriter, r *http.Request)       {}
func (s *ContentService) HandleGetCollection(w http.ResponseWriter, r *http.Request)    {}
func (s *ContentService) HandleGetClass(w http.ResponseWriter, r *http.Request)         {}
func (s *ContentService) HandleUpdateContent(w http.ResponseWriter, r *http.Request)    {}
func (s *ContentService) HandleDeleteContent(w http.ResponseWriter, r *http.Request)    {}
func (s *ContentService) HandleDeleteClass(w http.ResponseWriter, r *http.Request)      {}
func (s *ContentService) HandleDeleteCollection(w http.ResponseWriter, r *http.Request) {}
