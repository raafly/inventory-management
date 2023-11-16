CREATE TABLE categories (
    id varchar(100) not null primary key,
    name varchar(50) not null,
    description text default 'description not set',
    created_at timestamp default current_timestamp
)