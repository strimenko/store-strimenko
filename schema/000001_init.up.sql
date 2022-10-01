CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE backpacks
(
    itemId serial not null unique,
    itemName varchar(255) not null,
    price serial not null
);

CREATE TABLE bodyarmors
(
    itemId serial not null unique,
    itemName varchar(255) not null,
    price serial not null
);

CREATE TABLE helmets
(
    itemId serial not null unique,
    itemName varchar(255) not null,
    price serial not null
);