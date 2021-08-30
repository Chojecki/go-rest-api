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
	fmt.Println("Setupping routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Am alive")
	})
}
