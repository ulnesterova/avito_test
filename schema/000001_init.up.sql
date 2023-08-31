-- Создание таблицы "users"
CREATE TABLE users (
    id serial PRIMARY KEY,
    name VARCHAR(50),
    username VARCHAR(50),
    password_hash VARCHAR(100) NOT NULL
);

-- Создание таблицы "segments"
CREATE TABLE segments (
    id serial PRIMARY KEY,
    slug VARCHAR(255) UNIQUE
);

-- Создание таблицы для отношения "users_segments"
CREATE TABLE users_segments (
    user_id INT,
    segment_id INT,
    PRIMARY KEY (user_id, segment_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (segment_id) REFERENCES segments(id)
);