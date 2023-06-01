CREATE TABLE "rooms" (
    id SERIAL PRIMARY KEY,
    room_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ,
    update_at TIMESTAMPTZ
);