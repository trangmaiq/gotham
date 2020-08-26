package sql

import (
	"context"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/pop/v5"
	"github.com/pkg/errors"
)

var migrations = packr.New("migrations", "migrations/sql")

// Persister implements Persister interface
type Persister struct {
	c  *pop.Connection
	mb pop.MigrationBox
}

// NewPersister creates a new instance of Persister and return error if cannot create new migration box
func NewPersister(c *pop.Connection) (*Persister, error) {
	var mb, err = pop.NewMigrationBox(migrations, c)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &Persister{c: c, mb: mb}, nil
}

// GetConnection returns *pop.Connection it is kept in context or persister connection otherwise
func (p *Persister) GetConnection(ctx context.Context) *pop.Connection {
	var c = ctx.Value(transactionKey)
	if c != nil {
		if conn, ok := c.(*pop.Connection); ok {
			return conn
		}
	}

	return p.c
}