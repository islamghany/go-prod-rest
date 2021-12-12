package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// stores pointers to our comments services
type Handler struct {
	Router *mux.Router
}

//factory function
// return a pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// setup our routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("setting Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, h *http.Request) {
		fmt.Fprint(w, "app-health")
	})
}
