package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/daparadoks/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
)

// Handler - stores pointer to our comments service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

type Response struct {
	Success bool
	Message string
	Code int
}

// NewHandler - returns a pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I'm alive!"}); err != nil {
			panic(err)
		}
	})
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	h.SetHeaders(w)
	comment, err := h.Service.GetAllComments()
	if err != nil {
		h.GetMessageAsJson(w, "Failed to retrieve all comments")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	h.SetHeaders(w)
	var comment comment.Comment
	if err:=json.NewDecoder(r.Body).Decode(&comment); err!=nil{
		h.GetMessageAsJson(w,"Failed to decode json from body")
		return
	}
	comment, err := h.Service.PostComment(comment)

	if err != nil {
		h.GetMessageAsJson(w,"Failed to post new comment")
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// GetComment - retrieve a comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	h.SetHeaders(w)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "Unable to parse int from id")
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		h.GetMessageAsJson(w,"Error retrieving comment by id")
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// UpdateComment - at gibi update ediyor mk
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	h.SetHeaders(w)
	var comment comment.Comment
	if err:=json.NewDecoder(r.Body).Decode(&comment); err!=nil{
		h.GetMessageAsJson(w,"Failed to decode json")
		return
	}
	vars :=mux.Vars(r)
	id:=vars["id"]
	commentId, err:=strconv.ParseUint(id, 10, 64)
	if err!=nil{
		h.GetMessageAsJson(w,"Failed to parse id")
		return
	}
	comment, err = h.Service.UpdateComment(uint(commentId), comment)
	if err != nil {
		h.GetMessageAsJson(w,"Failed to update comment")
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// DeleteComment - siliyor işte, baya temiz siliyor hemde
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	h.SetHeaders(w)
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		h.GetMessageAsJson(w, "Failed to parse int from id")
		return
	}

	err = h.Service.DeleteComment(uint(commentId))
	if err != nil {
		h.GetMessageAsJson(w,"Failed to delete comment")
		return
	}

	h.GetMessageAsJson(w,"Comment deleted")
}

func (h *Handler) SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetMessageAsJson(w http.ResponseWriter, message string) error {
	return json.NewEncoder(w).Encode(Response{Success: false, Message: message, Code: 400})
}