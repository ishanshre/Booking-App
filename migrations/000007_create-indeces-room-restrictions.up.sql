CREATE INDEX idx_room_restrictions ON room_restrictions (start_date, end_date);
CREATE INDEX idx_room_restrictions_room_id ON room_restrictions (room_id);
CREATE INDEX idx_room_restrictions_reservation_id ON room_restrictions (reservation_id);