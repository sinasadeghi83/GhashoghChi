-- +goose Up
-- +goose StatementBegin
create table users (
    id integer not null primary key AUTO_INCREMENT,
    full_name varchar(255) not null,
    phone varchar(15) not null,
    email varchar(70) unique,
    password varchar(255) not null,
    role enum('buyer', 'seller', 'courier') not null,
    address varchar(255),
    profileImageBase64 varchar(255),
    bank_name varchar(255),
    account_number varchar(255),
    created_at datetime default CURRENT_TIMESTAMP,
    updated_at datetime default CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd
