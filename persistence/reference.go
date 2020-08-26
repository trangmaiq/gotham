package persistence

import "github.com/gobuffalo/pop/v5"

type Persister interface {
	GetConnection() *pop.Connection
}
