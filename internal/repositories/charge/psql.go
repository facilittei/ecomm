package repository

import (
	"context"
	"database/sql"
	"github.com/facilittei/ecomm/internal/domains/payment"
)

// ChargePsql holds payment charge history
type ChargePsql struct {
	Conn *sql.DB
}

// NewChargePsql creates an instance of Juno
func NewChargePsql(conn *sql.DB) Charge {
	return &ChargePsql{Conn: conn}
}

// Store adds charge state by each interaction made with payment provider
func (c ChargePsql) Store(ctx context.Context, charge payment.Charge) error {
	panic("implement me")
}
