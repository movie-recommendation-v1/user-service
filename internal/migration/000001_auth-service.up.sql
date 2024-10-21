create type role as ENUM('admin','super-admin','user');

CREATE TABLE IF NOT EXISTS users(
    id uuid primary key not null,
    name varchar(100) not null,
    email varchar(150) unique not null,
    img_url varchar(250) not null,
    password varchar(350) not null,
    role role default 'user' not null,
    created_at timestamp default now() not null,
    updated_at timestamp default now() not null,
    deleted_at bigint default 0 not null
);

-- Create a partial unique index for img_url
CREATE UNIQUE INDEX unique_img_url ON users(img_url) WHERE img_url != 'empty';
