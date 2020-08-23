package driver

import "github.com/trangmaiq/gotham/selfservice/flow/registration"

func (rd *RegistryDefault) RegistrationHandler() *registration.Handler {
	if rd.selfserviceRegistrationHandler == nil {
		rd.selfserviceRegistrationHandler = registration.NewHandler()
	}

	return rd.selfserviceRegistrationHandler

}
