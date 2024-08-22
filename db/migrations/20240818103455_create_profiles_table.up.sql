CREATE TABLE profiles (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id),
  bio VARCHAR,
  profile_picture VARCHAR,
  location VARCHAR,
  birthdate TIMESTAMP,
  status VARCHAR
)
