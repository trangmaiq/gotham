package driver

import (
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/trangmaiq/gotham/driver/configuration"
	"github.com/trangmaiq/gotham/x/dbal"
)

type Registry interface {
	RegisterPublicRoutes(public *mux.Router)
}

// NewRegistry returns Registry with the provided DSN from configuration
func NewRegistry(c configuration.Provider) (Registry, error) {
	var (
		dsn         = c.DSN()
		driver, err = dbal.GetDriverFor(dsn)
	)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	var registry, ok = driver.(Registry)
	if !ok {
		return nil, errors.Errorf("driver of type %T does not implement interface Registry", driver)
	}

	return registry, nil
}
