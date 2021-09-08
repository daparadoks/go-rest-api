package http

import (
	"encoding/json"
	"fmt"
	"github.com/daparadoks/go-rest-api/internal/comment"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
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

// LoggingMiddleware - a handy middleware function that logs out incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"Method": r.Method,
			"Path": r.URL.Path,
		})
		log.Info("Endpoint hit!")
		next.ServeHTTP(w, r)
	})
}

// SetupRoutes - sets up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.Use(LoggingMiddleware)

	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")

	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Success: true, Message: "I'm alive!", Code: http.StatusOK}); err != nil {
			panic(err)
		}
	})
}

func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
}

func SendOkResponse(w http.ResponseWriter, resp interface{}){
	w.WriteHeader(http.StatusOK)
	if err:= json.NewEncoder(w).Encode(resp); err!=nil{
		panic(err)
	}
}

func GetErrorResponse(w http.ResponseWriter, message string) {
	if err:= json.NewEncoder(w).Encode(Response{Success: false, Message: message, Code: http.StatusBadRequest});err!=nil {
		panic(err)
	}
}