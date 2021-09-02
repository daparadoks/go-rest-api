package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


type Handler struct {
	Router *mux.Router
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("api/health", func(w http.ResponseWriter, r.http.http.Request){
		fmt.Fprintf(w, "I'm alive")
	})
}
