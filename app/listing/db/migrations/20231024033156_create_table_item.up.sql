CREATE TABLE items(
    id              varchar(100) not null primary key,
    name            varchar(100) not null,
    description     text,
    quantity        int not null default 0,
    category        int not null,
    int             timestamp not null default current_timestamp,
    out             timestamp,
    constraint fk_items_categories Foreign Key (category) REFERENCES categories(id),
    created_at      timestamp not null default current_timestamp
)