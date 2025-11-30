package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// api represents our HTTP application.
type api struct {
	addr string
}

// User represents a user in the system.
type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

var users = []User{}

// ---- Handlers ----

// GET /users
func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// POST /users
func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var payload User
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert user with validation
	if err := insertUser(payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Respond with the created user
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// insertUser validates and inserts a new user into memory
func insertUser(u User) error {
	// Input validation
	if u.FirstName == "" {
		return errors.New("first name is required")
	}
	if u.LastName == "" {
		return errors.New("last name is required")
	}

	// Check for duplicate user
	for _, existing := range users {
		if existing.FirstName == u.FirstName && existing.LastName == u.LastName {
			return errors.New("user already exists")
		}
	}

	// Append to global slice
	users = append(users, u)
	return nil
}

// GET /
func (a *api) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Index page"))
}

// GET /home
func (a *api) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This is home page"))
}

// ---- Route Registration ----
func (a *api) registerRoutes(mux *http.ServeMux) {
	// Go 1.22 method + path routing
	mux.HandleFunc("GET /", a.indexHandler)
	mux.HandleFunc("GET /home", a.homeHandler)

	mux.HandleFunc("GET /users", a.getUserHandler)
	mux.HandleFunc("POST /users", a.createUserHandler)

	// 404 fallback
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
}

// ---- Main ----
func main() {
	a := &api{addr: ":8081"}

	mux := http.NewServeMux()
	a.registerRoutes(mux)

	srv := &http.Server{
		Addr:    a.addr,
		Handler: mux,
	}

	log.Printf("Server running at http://localhost%s\n", a.addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
