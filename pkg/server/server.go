package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/YanSystems/cms/pkg/services"
	utils "github.com/YanSystems/cms/pkg/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Port string
	DB   *mongo.Database
}

func (s *Server) NewRouter() http.Handler {
	slog.Info("Setting up new router")
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "https://localhost", "http://localhost:3000", "https://localhost:3000", "https://abyan.dev"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	slog.Info("CORS middleware configured")

	// Health check
	router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	slog.Info("Health check route configured")

	contentService := services.ContentService{DB: s.DB}

	// Content services
	router.Post("/contents/{collection}", contentService.HandleCreateContent)
	router.Get("/contents/{collection}", contentService.HandleGetCollection)
	router.Get("/contents/{collection}/id/{id}", contentService.HandleGetContent)
	router.Get("/contents/{collection}/class/{class}", contentService.HandleGetClass)
	router.Put("/contents/{collection}/id/{id}", contentService.HandleUpdateContent)
	router.Delete("/contents/{collection}", contentService.HandleDeleteCollection)
	router.Delete("/contents/{collection}/id/{id}", contentService.HandleDeleteContent)
	router.Delete("/contents/{collection}/class/{class}", contentService.HandleDeleteClass)
	slog.Info("Content service routes configured")

	return router
}

func (s *Server) NewServer() *http.Server {
	s.Port = "8000"

	router := s.NewRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.Port),
		Handler: router,
	}

	slog.Info("New server instance created", "port", s.Port)
	return server
}

func (s *Server) Run() {
	client, err := utils.ConnectToDB()
	slog.Info("Connecting to database")
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		log.Fatal(err)
	}
	defer func() {
		slog.Info("Disconnecting from database")
		if err := client.Disconnect(context.TODO()); err != nil {
			slog.Error("Failed to disconnect from database", "error", err)
			panic(err)
		}
		slog.Info("Disconnected from database successfully")
	}()

	s.DB = client.Database("content")
	slog.Info("Database connection established", "db", "content")

	slog.Info(fmt.Sprintf("The server is now live on port %s", s.Port))

	srv := s.NewServer()
	err = srv.ListenAndServe()
	if err != nil {
		slog.Error("Server encountered an error", "error", err)
		log.Panic(err)
		return
	}
}
