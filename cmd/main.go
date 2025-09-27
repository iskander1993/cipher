package main

import (
	http2 "ave_project/internal/handlers"
	"ave_project/internal/infrastructure/postgres"
	"ave_project/internal/infrastructure/repositories"
	"ave_project/internal/middleware"
	"ave_project/internal/usecase/cipher"
	"ave_project/internal/usecase/user"
	// остальные импорты...
	"fmt"
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func main() {
	// Подключаем PostgreSQL
	db := postgres.ConnectPostgres()

	// Репозиторий
	userRepo := repositories.NewUserRepository(db)

	// Юзкейсы
	userUC := &user.UserUsecase{Repo: userRepo}
	cipherUC := &cipher.CipherUsecase{}

	// Хендлеры
	userHandler := &http2.UserHandler{Usecase: userUC}
	cipherHandler := &http2.CipherHandler{Usecase: cipherUC}

	// Chi router
	r := chi.NewRouter()

	// === MIDDLEWARE ===
	r.Use(middleware.LoggingMiddleware) // кастомный логгер
	r.Use(chiMiddleware.Logger)         // встроенный логгер chi
	r.Use(chiMiddleware.Recoverer)      // отлов паник

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// API группа
	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/register", userHandler.RegisterHandler)
		r.Post("/login", userHandler.LoginHandler)
		r.Post("/encrypt", cipherHandler.EncryptHandler)
		r.Post("/decrypt", cipherHandler.DecryptHandler)

		// Тестовый эндпоинт
		r.Get("/test", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Test OK"))
		})
	})

	// Запуск сервера
	srvAddr := ":8080"
	fmt.Println("Server starting on", srvAddr)
	fmt.Println("Logging middleware activated!")
	http.ListenAndServe(srvAddr, r)
}
