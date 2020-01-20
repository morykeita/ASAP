CREATE TABLE  IF NOT EXISTS  "uuid-ossp"

Create TABLE users(
    user_id BINARY(16) PRIMARY KEY AUTO_INCREMENT,
    email TEXT NOT NULL,
    password_hash binary(64),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);
CREATE UNIQUE INDEX user_email
    ON users(email);
