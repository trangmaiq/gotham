package driver

import (
	"github.com/gorilla/mux"
	"github.com/trangmaiq/gotham/selfservice/flow/registration"
)

var _ Registry = new(RegistryDefault)

// RegistryDefault implements Registry interface
type RegistryDefault struct {
	selfserviceRegistrationHandler *registration.Handler
}

// NewRegistryDefault constructs a new RegistryDefault
func NewRegistryDefault() *RegistryDefault {
	return &RegistryDefault{}
}

// RegisterPublicRoutes registers all route
func (rd RegistryDefault) RegisterPublicRoutes(public *mux.Router) {
	rd.RegistrationHandler().RegisterPublicRoutes(public)
}
