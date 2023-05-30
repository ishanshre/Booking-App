CREATE TABLE "reservations" (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email VARCHAR(255),
    phone VARCHAR(10),
    start_date DATE,
    end_date DATE,
    room_id INTEGER,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    CONSTRAINT fk_rooms FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE
)