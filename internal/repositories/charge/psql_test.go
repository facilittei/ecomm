package repository

import (
	"context"
	"database/sql"
	"github.com/facilittei/ecomm/internal/domains/customer"
	"github.com/facilittei/ecomm/internal/domains/payment"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/lib/pq"
)

func TestChargePsql_Store(t *testing.T) {
	dsn := "postgres://facilittei:4321@localhost/facilittei?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	require.Nil(t, err)

	chargeRepository := NewChargePsql(db)

	ctx := context.Background()
	err = chargeRepository.Store(ctx, payment.Charge{
		ID:          uuid.New(),
		SKU:         "abc1234",
		Amount:      100,
		Description: "Go Development",
		Customer: customer.Customer{
			Name:     "Jeff Bezos",
			Email:    "jeff@amazon.com",
			Document: "11144740452",
			Address: customer.Address{
				Street:   "Rua Guedes Perreira",
				Number:   "90",
				City:     "Recife",
				State:    "PE",
				PostCode: "52060150",
			},
		},
	})

	require.Nil(t, err)
}
