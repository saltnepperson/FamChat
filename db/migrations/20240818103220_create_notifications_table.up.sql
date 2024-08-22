CREATE TABLE notifications (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id),
  type VARCHAR(100),
  message VARCHAR(50),
  created_at TIMESTAMP,
  is_read BOOLEAN
)
