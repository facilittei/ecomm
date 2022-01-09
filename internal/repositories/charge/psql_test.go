package repository

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/facilittei/ecomm/internal/domains/customer"
	"github.com/facilittei/ecomm/internal/domains/payment"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"

	_ "github.com/lib/pq"
)

func TestChargePsql_Store(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.Nil(t, err)
	defer db.Close()

	chargeRepository := NewChargePsql(db)
	charge := payment.Charge{
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
	}

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(sqlChargeInsert)).WithArgs(
		charge.ID.String(),
		charge.SKU,
		charge.Amount,
		charge.Description,
		charge.Customer.Name,
		charge.Customer.Email,
		charge.Customer.Document,
		charge.Customer.Address.Street,
		charge.Customer.Address.Number,
		charge.Customer.Address.Complement,
		charge.Customer.Address.City,
		charge.Customer.Address.State,
		charge.Customer.Address.PostCode,
	).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	ctx := context.Background()
	err = chargeRepository.Store(ctx, charge)
	require.Nil(t, err)
}
