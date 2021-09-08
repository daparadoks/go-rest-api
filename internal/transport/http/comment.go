package http

import (
	"encoding/json"
	"github.com/daparadoks/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	comment, err := h.Service.GetAllComments()
	if err != nil {
		GetErrorResponse(w, "Failed to retrieve all comments")
	}

	SendOkResponse(w, comment)
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	var comment comment.Comment
	if err:=json.NewDecoder(r.Body).Decode(&comment); err!=nil{
		GetErrorResponse(w,"Failed to decode json from body")
		return
	}
	comment, err := h.Service.PostComment(comment)

	if err != nil {
		GetErrorResponse(w,"Failed to post new comment")
		return
	}

	SendOkResponse(w, comment)
}

// GetComment - retrieve a comment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		GetErrorResponse(w, "Unable to parse int from id")
		return
	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		GetErrorResponse(w,"Error retrieving comment by id")
		return
	}

	SendOkResponse(w, comment)
}

// UpdateComment - at gibi update ediyor mk
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	var comment comment.Comment
	if err:=json.NewDecoder(r.Body).Decode(&comment); err!=nil{
		GetErrorResponse(w,"Failed to decode json")
		return
	}
	vars :=mux.Vars(r)
	id:=vars["id"]
	commentId, err:=strconv.ParseUint(id, 10, 64)
	if err!=nil{
		GetErrorResponse(w,"Failed to parse id")
		return
	}
	comment, err = h.Service.UpdateComment(uint(commentId), comment)
	if err != nil {
		GetErrorResponse(w,"Failed to update comment")
		return
	}

	SendOkResponse(w, comment)
}

// DeleteComment - siliyor işte, baya temiz siliyor hemde
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	SetHeaders(w)
	vars := mux.Vars(r)
	id := vars["id"]
	commentId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		GetErrorResponse(w, "Failed to parse int from id")
		return
	}

	err = h.Service.DeleteComment(uint(commentId))
	if err != nil {
		GetErrorResponse(w,"Failed to delete comment")
		return
	}

	GetErrorResponse(w,"Comment deleted")
}
