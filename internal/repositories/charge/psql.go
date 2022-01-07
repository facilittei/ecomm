package repository

import (
	"context"
	"database/sql"
	"fmt"
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
	var err error

	tx, err := c.Conn.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("charge begin transation failed: %s", err.Error())
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	query := `INSERT INTO charges (id, sku, amount, description, customer_name, customer_email, customer_document, 
        customer_address_street, customer_address_number, customer_address_complement,
        customer_address_city, customer_address_state, customer_address_postcode
	) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`

	args := []interface{}{
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
	}

	if _, err := tx.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("charge insert failed: %s", err)
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("charge commit transaction failed: %s", err.Error())
	}

	return err
}
