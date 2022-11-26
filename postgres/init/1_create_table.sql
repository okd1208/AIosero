CREATE TABLE combat_experience (
    id BIGINT PRIMARY KEY,
    Is_user_won boolean NOT NULL,
    ts TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);