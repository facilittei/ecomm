CREATE TABLE IF NOT EXISTS charges
(
    id                          CHAR(16)                    NOT NULL PRIMARY KEY,
    sku                         VARCHAR(255)                NOT NULL,
    customer_name               VARCHAR(255)                NOT NULL,
    customer_email              VARCHAR(255)                NOT NULL,
    customer_document           VARCHAR(50)                 NOT NULL,
    customer_address_street     VARCHAR(255)                NOT NULL,
    customer_address_number     VARCHAR(50)                 NOT NULL,
    customer_address_complement VARCHAR(255),
    customer_address_city       VARCHAR(50)                 NOT NULL,
    customer_address_state      VARCHAR(50)                 NOT NULL,
    customer_address_postcode   VARCHAR(50)                 NOT NULL,
    created_at                  TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at                  TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS charges_history
(
    charge_id   CHAR(16)                    NOT NULL,
    status      INTEGER                     NOT NULL,
    description VARCHAR(255),
    created_at  TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_charge_history FOREIGN KEY (charge_id) REFERENCES charges (id)
)