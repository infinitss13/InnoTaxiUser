CREATE TABLE users
(
    id serial not null unique,
    name varchar(70) not null,
    phone varchar(15) not null unique,
    email varchar(255) not null unique,
    password_hash varchar(255) not null,
    rating real
);