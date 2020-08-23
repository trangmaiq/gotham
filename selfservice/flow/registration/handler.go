package registration

import (
	"github.com/gorilla/mux"
	"net/http"
)

const (
	BrowserRegistrationPath = "/self-service/browser/flows/registration"
)

type Handler struct{}

// NewHandler returns pointer of Handler struct
func NewHandler() *Handler {
	return &Handler{}
}

// RegisterPublicRoutes registers all new routes related to registration
func (h *Handler) RegisterPublicRoutes(public *mux.Router) {
	public.HandleFunc(BrowserRegistrationPath, func(writer http.ResponseWriter, request *http.Request) {
	})
}
