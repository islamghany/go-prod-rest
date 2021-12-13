package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/islamghany/go-prod-rest/internals/comment"
	"net/http"
	"strconv"
)

// stores pointers to our comments services
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

//factory function
// return a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// setup our routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("setting Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/${id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, h *http.Request) {
		fmt.Fprint(w, "app-health")
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)

	if err != nil {
		fmt.Fprintf(w, "Unable to convert id")
	}
	comment, err := h.Service.GetComment(uint(id))

	if err != nil {
		fmt.Fprintf(w, "Can not retieve comment with that id")
	}
	fmt.Fprintf(w, "es %+v", comment)
}
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {

	comments, err := h.Service.GetAllComments()

	if err != nil {
		fmt.Fprintf(w, "Can not retieve all comments ")
	}
	fmt.Fprintf(w, "%+v", comments)
}
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		Slug: "/",
	})

	if err != nil {
		fmt.Fprintf(w, "falid to post a new comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		Slug: "/new",
	})

	if err != nil {
		fmt.Fprintf(w, "falid to update comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to convert id")
	}
	err = h.Service.DeleteComment(uint(id))

	if err != nil {
		fmt.Fprintf(w, "falid to delete comment")
	}
	fmt.Fprintf(w, "%+v", "deleted successfully!")
}
