package repository

// sqlChargeInsert creates a record with current information about charge
// it's also used for associating charge states to get its history
var sqlChargeInsert = `INSERT INTO charges (
	id, sku, amount, description, customer_name, customer_email, customer_document, 
	customer_address_street, customer_address_number, customer_address_complement,
	customer_address_city, customer_address_state, customer_address_postcode
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`
