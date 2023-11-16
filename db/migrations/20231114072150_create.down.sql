CREATE TABLE history(
    id serial primary key not null,
    item_id int not null,
    action  bool not null,
    quantity int not null,
    update_at  timestamp default current_timestamp
)