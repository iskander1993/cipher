package main

import (
	http2 "ave_project/internal/handlers"
	"fmt"
	"net/http"

	"ave_project/internal/infrastructure/postgres"
	"ave_project/internal/infrastructure/repositories"
	"ave_project/internal/usecase"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	// Подключаем PostgreSQL
	db := postgres.ConnectPostgres()

	// Репозиторий
	//passwordRepo := repositories.NewPasswordRepository(db)
	userRepo := repositories.NewUserRepository(db)

	// Юзкейсы
	userUC := &usecase.UserUsecase{Repo: userRepo}
	cipherUC := &usecase.CipherUsecase{}

	// Хендлеры
	userHandler := &http2.UserHandler{Usecase: userUC}
	cipherHandler := &http2.CipherHandler{Usecase: cipherUC}

	// Chi router
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS (для фронта)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"}, // можно указать конкретные фронт-домены
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // 5 минут
	}))

	// API группа
	r.Route("/api/v1", func(r chi.Router) {
		// User routes
		r.Post("/register", userHandler.RegisterHandler)
		r.Post("/login", userHandler.LoginHandler)

		// Cipher routes
		r.Post("/encrypt", cipherHandler.EncryptHandler)
		r.Post("/decrypt", cipherHandler.DecryptHandler)
	})

	// Запуск сервера
	srvAddr := ":8080"
	fmt.Println("Listening on", srvAddr)
	http.ListenAndServe(srvAddr, r)
}
