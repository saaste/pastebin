package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/saaste/pastebin/pkg/auth"
	"github.com/saaste/pastebin/pkg/config"
	"github.com/saaste/pastebin/pkg/handlers"
)

func main() {

	appConfig, err := config.GetConfig()
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	handler := handlers.NewHandler(appConfig)
	authMiddleware := auth.NewAuthMiddleware(appConfig)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)
	r.Use(authMiddleware.Authenticate)

	r.Use(middleware.Timeout(60 * time.Second))

	r.Group(func(r chi.Router) {
		r.Use(authMiddleware.RequiresAuthentication)
		r.Get("/", handler.IndexHandler)
		r.Get("/new", handler.NewHandler)
		r.Post("/new", handler.NewHandler)
		r.Get("/edit/{id}", handler.EditHandler)
		r.Post("/edit/{id}", handler.EditHandler)
	})

	r.Group(func(r chi.Router) {
		r.Get("/login", handler.LoginHandler)
		r.Post("/login", handler.LoginHandler)
		r.Get("/logout", handler.LogoutHandler)
		r.Get("/paste/{public_path}", handler.PasteHandler)
	})

	handlers.FileServer(r, "/static/", http.Dir("static"))
	handlers.FileServer(r, "/assets/", http.Dir(fmt.Sprintf("ui/%s/assets", appConfig.Theme)))

	fmt.Println("Listening: http://localhost:8000")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("starting the server failed: %v", err)
		}
	}

}
