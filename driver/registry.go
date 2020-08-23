package driver

import "github.com/gorilla/mux"

type Registry interface {
	RegisterPublicRoutes(public *mux.Router)
}
