DROP DATABASE IF EXISTS synapsis_db;
CREATE DATABASE synapsis_db;

\c synapsis_db


CREATE TABLE accounts(
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    password VARCHAR NOT NULL,
    account_type VARCHAR(6) NOT NULL,
    profile_set BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

CREATE TABLE admins(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_account BIGINT NOT NULL REFERENCES accounts(id)
);

CREATE TABLE users(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    photo_url VARCHAR,
    date_of_birth DATE,
    gender VARCHAR(6),
    phone_number VARCHAR,
    main_address_id BIGINT,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_account BIGINT NOT NULL REFERENCES accounts(id)
);

CREATE TABLE user_addresses(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    phone_number VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    latitude NUMERIC NOT NULL,
    longitude NUMERIC NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_user BIGINT NOT NULL REFERENCES users(id)
);

CREATE TABLE sellers(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    date_of_birth DATE NOT NULL,
    gender VARCHAR(6),
    phone_number VARCHAR UNIQUE NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_account BIGINT NOT NULL REFERENCES accounts(id)
);

CREATE TABLE email_verify_tokens(
    id BIGSERIAL PRIMARY KEY,
    token VARCHAR NOT NULL,
    expired_at TIMESTAMP NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_account BIGINT NOT NULL REFERENCES accounts(id)
);

CREATE TABLE shops(
    id BIGSERIAL PRIMARY KEY,
    shop_name VARCHAR UNIQUE NOT NULL,
    logo_url VARCHAR,
    phone_number VARCHAR UNIQUE NOT NULL,
    description VARCHAR,
    address VARCHAR NOT NULL,
    latitude NUMERIC NOT NULL,
    longitude NUMERIC NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_seller BIGINT NOT NULL REFERENCES sellers(id)
);

CREATE TABLE shipment_methods(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

CREATE TABLE payment_methods(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

CREATE TABLE shop_shipment_methods(
    id BIGSERIAL PRIMARY KEY,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_shop BIGINT NOT NULL REFERENCES shops(id),
    id_shipment_method BIGINT NOT NULL REFERENCES shipment_methods(id)
);

CREATE TABLE shop_payment_methods(
    id BIGSERIAL PRIMARY KEY,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_shop BIGINT NOT NULL REFERENCES shops(id),
    id_payment_method BIGINT NOT NULL REFERENCES payment_methods(id)
);

CREATE TABLE categories(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    parent_id BIGINT,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP
);

CREATE TABLE products(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    photo_url VARCHAR,
    description VARCHAR,
    stock INT NOT NULL DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_category BIGINT NOT NULL REFERENCES categories(id),
    id_shop BIGINT NOT NULL REFERENCES shops(id)
);

CREATE TABLE transactions(
    id BIGSERIAL PRIMARY KEY,
    invoice VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    num_of_items INT NOT NULL,
    subtotal DECIMAL NOT NULL,
    shipment_fee DECIMAL NOT NULL,
    total_fee DECIMAL NOT NULL,
    address VARCHAR NOT NULL,
    latitude NUMERIC NOT NULL,
    longitude NUMERIC NOT NULL,
    
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_shop BIGINT NOT NULL REFERENCES shops(id),
    id_user BIGINT NOT NULL REFERENCES users(id),
    id_shipment_method BIGINT NOT NULL REFERENCES shipment_methods(id),
    id_payment_method BIGINT NOT NULL REFERENCES payment_methods(id)    
);

CREATE TABLE transaction_items(
    id BIGSERIAL PRIMARY KEY,
    amount INT NOT NULL,
    total_price DECIMAL NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_transaction BIGINT NOT NULL REFERENCES transactions(id),
    id_product BIGINT NOT NULL REFERENCES products(id)
);

CREATE TABLE payments(
    id BIGSERIAL PRIMARY KEY,
    file_url VARCHAR NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP,

    id_transaction BIGINT NOT NULL REFERENCES transactions(id)
);