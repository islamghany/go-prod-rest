package http

import (
	"encoding/json"
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

type Response struct {
	Message string
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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am a Alive"}); err != nil {
			panic(err)
		}
	})
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err != nil {

		sendErrorResponse(w, "Unable to convert id", err)
		return
	}
	comment, err := h.Service.GetComment(uint(id))

	if err != nil {

		sendErrorResponse(w, "Can not retieve comment with that id", err)
		return
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	comments, err := h.Service.GetAllComments()

	if err != nil {

		sendErrorResponse(w, "Can not retieve all comments ", err)
		return
	}
	if err := json.NewEncoder(w).Encode(comments); err != nil {
		panic(err)
	}
}
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		panic(err)
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {

		sendErrorResponse(w, "falid to post a new comment", err)
		return
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {

		sendErrorResponse(w, "Unable to convert id", err)
		return
	}
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		panic(err)
	}

	comment, err = h.Service.UpdateComment(uint(id), comment)

	if err != nil {

		sendErrorResponse(w, "falid to update comment", err)
		return
	}
	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to convert id", err)
		return
	}

	err = h.Service.DeleteComment(uint(id))

	if err != nil {
		sendErrorResponse(w, "falid to delete comment", err)
		return
	}
	if err := json.NewEncoder(w).Encode(Response{Message: "deleted: succeefully!"}); err != nil {
		panic(err)
	}
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
