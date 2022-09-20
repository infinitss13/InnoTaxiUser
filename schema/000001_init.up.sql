CREATE TABLE users
(
    id serial,
    name varchar(70) not null,
    phone varchar(15) not null unique,
    email varchar(255) not null unique,
    password_hash varchar(255) not null,
    rating real,
    created_at timestamp,
    updated_at timestamp,
    deleted boolean
);