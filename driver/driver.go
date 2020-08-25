package driver

import "github.com/trangmaiq/gotham/driver/configuration"

type Driver interface {
	Configuration() configuration.Provider
	Registry() Registry
}
