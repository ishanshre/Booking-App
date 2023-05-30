CREATE TABLE "room_restrictions" (
    id SERIAL PRIMARY KEY,
    start_date TIMESTAMPTZ,
    end_date TIMESTAMPTZ,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    restriction_id INTEGER,
    reservation_id INTEGER,
    room_id INTEGER,
    CONSTRAINT fk_restrictions FOREIGN KEY (restriction_id) REFERENCES restrictions(id) ON DELETE CASCADE,
    CONSTRAINT fk_reservations FOREIGN KEY (reservation_id) REFERENCES reservations(id) ON DELETE CASCADE,
    CONSTRAINT fk_rooms_restrictions FOREIGN KEY (room_id) REFERENCES rooms(id) ON DELETE CASCADE
);