CREATE TABLE "restrictions" (
    id SERIAL PRIMARY KEY,
    restriction_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ
)