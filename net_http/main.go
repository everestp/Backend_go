package main

import (
	
	"log"
	"net/http"
)

// api represents our HTTP application.
type api struct {
	addr string
}

// ---- Handlers ----

// GET /users
func (a *api) getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User List >>"))
}

// POST /users
func (a *api) createUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Created User"))
}

// GET /
func (a *api) indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Index page"))
}

// GET /home
func (a *api) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is home page"))
}

// ---- Route Registration ----

func (a *api) registerRoutes(mux *http.ServeMux) {
	// New Go 1.22 pattern matching
	mux.HandleFunc("GET /", a.indexHandler)
	mux.HandleFunc("GET /home", a.homeHandler)

	// User routes
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
