package main

import (
	"fmt"
	"net/http"
	"userlogin/config"
	"userlogin/database"
	"userlogin/httphandler"
	"userlogin/services"
	"userlogin/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.StandardLogger()
	cfg := config.New()
	// db, err := database

	client, ctx, cancel, err := database.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	// Release resource when main function is retured
	defer database.Close(client, ctx, cancel)

	// store
	userStore := store.NewUserStore(client, ctx, logger)

	// service
	userService := services.NewUserService(client, ctx, userStore, logger)

	// handler
	userHandler := httphandler.NewUserHandler(userService, logger)

	httpCorsHandler := cors.New(cors.Options{
		// AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"LINK"},
		AllowCredentials: false,
		MaxAge:           300,
		Debug:            true,
	})

	router := chi.NewRouter()
	// router.Use(httpCorsHandler.Handler, middleware.Logger())
	router.Use(httpCorsHandler.Handler)
	router.Route("/api/v1", func(r chi.Router) {
		r.Post("/user", userHandler.CreateUser)
		r.Get("/test", userHandler.TestAPI)
	})

	Serve(router, cfg, logger)
}

func Serve(router chi.Router, cfg config.Config, logger *logrus.Logger) {
	addr := fmt.Sprintf(":%s", cfg.Port)
	logger.Infof("Serving API at %s", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		logger.WithError(err).Fatal("failed to start the server")
	}
}
