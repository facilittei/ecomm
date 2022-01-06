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
	query := `INSERT INTO charges (
		id, sku, customer_name, customer_email, customer_document, 
        customer_address_street, customer_address_number, customer_address_complement,
        customer_address_city, customer_address_state, customer_address_postcode
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	args := []interface{}{
		charge.ID.String(),
		charge.SKU,
		charge.Customer.Name,
		charge.Customer.Email,
		charge.Customer.Document,
		charge.Customer.Address.Street,
		charge.Customer.Address.Number,
		charge.Customer.Address.Complement,
		charge.Customer.Address.City,
		charge.Customer.Address.State,
		charge.Customer.Address.PostCode,
	}

	err := c.Conn.QueryRowContext(ctx, query, args...)
	if err != nil {
		return err.Err()
	}

	return nil
}
