CREATE  EXTENSION IF NOT EXIST "uuid-ossp";


CREATE TABLE users(
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    email TEXT NOT NULL,
    password_hash bytea,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE UNIQUE INDEX users_email
    on users(email)