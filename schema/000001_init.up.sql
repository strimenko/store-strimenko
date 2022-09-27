CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE backpacks
(
    itemid serial not null unique,
    itemname varchar(255) not null,
    price serial not null
);

CREATE TABLE bodyarmors
(
    itemid serial not null unique,
    itemname varchar(255) not null,
    price serial not null
);

CREATE TABLE helmets
(
    itemid serial not null unique,
    itemname varchar(255) not null,
    price serial not null
);