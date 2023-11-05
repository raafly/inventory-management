CREATE TABLE users (
    id varchar(100) primary key not null,
    username varchar(100) not null,
    email varchar(100) not null,
    password varchar(100) not null,
    cpassword varchar(100) not null,
    constraint username_unique unique(username),
    constraint email_unique unique(email),
    created_at timestamp default current_timestamp
)