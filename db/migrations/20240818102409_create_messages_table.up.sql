CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  chat_id INT REFERENCES chats(id),
  user_id INT REFERENCES users(id),
  content VARCHAR(250),
  sent_at TIMESTAMP,
  is_read BOOLEAN DEFAULT FALSE
)
