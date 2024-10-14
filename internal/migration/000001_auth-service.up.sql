create type role as ENUM('admin','super-admin','user');

CREATE TABLE IF NOT EXISTS users(
    id uuid primary key not null,
    name varchar(100) not null,
    lastname varchar(100) default 'empty',
    email varchar(150) unique not null,
    password varchar(350) not null,
    role role default 'user' not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at bigint default 0 not null
);