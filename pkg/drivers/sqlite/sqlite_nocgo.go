//go:build !cgo
// +build !cgo

package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tashima42/kine/pkg/drivers/generic"
	"github.com/tashima42/kine/pkg/server"
)

var errNoCgo = errors.New("this binary is built without CGO, sqlite is disabled")

func New(ctx context.Context, dataSourceName string, connPoolConfig generic.ConnectionPoolConfig, metricsRegisterer prometheus.Registerer) (server.Backend, error) {
	return nil, errNoCgo
}

func NewVariant(driverName, dataSourceName string, connPoolConfig generic.ConnectionPoolConfig, metricsRegisterer prometheus.Registerer) (server.Backend, *generic.Generic, error) {
	return nil, nil, errNoCgo
}

func setup(db *sql.DB) error {
	return errNoCgo
}
