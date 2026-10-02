-- +goose Up
create table if not exists users (

);

create table if not exists services (

);

create table if not exists gallery (

);

-- +goose Down
drop table if exists gallery;
drop table if exists services;
drop table if exists users;