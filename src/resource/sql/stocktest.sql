drop database if exists stocktest;
create database stocktest;
use stocktest;

drop table if exists stunittest;
create table stunittest(
    id int not null primary key,
    name varchar(10) not null,
    age int not null
);
