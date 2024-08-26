package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"basedantoni/habits-be/internal/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load environment variables")
	}

	port := os.Getenv("PORT")

	// Database
	db, err := sql.Open("sqlite3", "./habits.db")
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	// Router
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/health", healthHandler)

	// Habits Handlers
	v1Router.Route("/habits", func(r chi.Router) {
		r.Post("/", apiCfg.createHabitHandler)
		r.Get("/", apiCfg.indexHabitHandler)
		r.Get("/{id}", apiCfg.showHabitHandler)
		r.Put("/{id}", apiCfg.updateHabitHandler)
		r.Delete("/{id}", apiCfg.deleteHabitHandler)
	})

	router.Mount("/v1", v1Router)

	server := &http.Server{Handler: router, Addr: ":" + port}

	log.Fatal(server.ListenAndServe())
}
