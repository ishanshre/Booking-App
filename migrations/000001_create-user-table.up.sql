CREATE TABLE "users" (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ,
    update_at TIMESTAMPTZ,
    last_login TIMESTAMPTZ,
    access_level INTEGER DEFAULT 1
);