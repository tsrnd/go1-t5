drop table posts;
drop table threads;
drop table users;
drop table sessions;

create table users (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL
)
create table sessions (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL
)
create table threads (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    topic TEXT,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP NOT NULL
)
create table posts (
    id SERIAL PRIMARY KEY,
    uuid VARCHAR(64) NOT NULL UNIQUE,
    body TEXT,
    user_id INTEGER REFERENCES users(id),
    thread_id INTEGER REFERENCES threads(id),
    created_at TIMESTAMP NOT NULL
)