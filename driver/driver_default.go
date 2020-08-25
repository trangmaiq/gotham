package driver

import (
	"github.com/pkg/errors"
	"github.com/trangmaiq/gotham/driver/configuration"
)

// DefaultDriver implements Driver interface
type DefaultDriver struct {
	c configuration.Provider
	r Registry
}

// NewDefaultDriver returns DefaultDriver and error if any
func NewDefaultDriver() (Driver, error) {
	var (
		c      = configuration.NewViperProvider()
		r, err = NewRegistry(c)
	)

	if err != nil {
		return nil, errors.Wrap(err, "unable to instantiate service registry")
	}

	return &DefaultDriver{c: c, r: r}, nil
}

func (d *DefaultDriver) Configuration() configuration.Provider {
	return d.c
}

func (d *DefaultDriver) Registry() Registry {
	return d.r
}
