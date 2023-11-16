CREATE TABLE items(
    id serial primary key,
    name varchar(100) not null,
    description text default 'description not set',
    quantity int not null default 0,
    category varchar(100) not null,
    status boolean default 'true',
    constraint fk_items_categories Foreign Key (category) REFERENCES categories(id),
    created_at timestamp not null default current_timestamp
)