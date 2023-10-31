CREATE TABLE users (
    id varchar(100) not null,
    username varchar(100) not null,
    email varchar(100) not null,
    password text not null,
    primary key(id),
    constraint username_unique unique(username),
    constraint email_unique unique(email),
    created_at timestamp default current_timestamp
)