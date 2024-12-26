package main

import (
	"fmt"
	"gotraining/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	// Add middlewares
	r.Use(middleware.Logger)    // Logs each request
	r.Use(middleware.Recoverer) // Recovers from panics

	// Define routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!!"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/create_user", handlers.CreateUser)   // User creation API
		r.Get("/get_all_users", handlers.GetAllUsers) // Get all users with all the details
		r.Get("/{id}", handlers.GetUser)              // Get user details based on the user ID
		r.Put("/update_user", handlers.UpdateUser)    // Updates user details
		r.Delete("/delete_user", handlers.DeleteUser) // Deletes the user data
	})

	// Starting the server on Localhost
	fmt.Println("Serving on localhost:8080 ...")
	http.ListenAndServe(":8080", r)
}
