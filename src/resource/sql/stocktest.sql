--drop database if exists stocktest;
--create database stocktest;
use stocktest;

drop table if exists stunittest;
create table stunittest(
    id int not null primary key,
    name varchar(10) not null,
    age int not null
);

drop table if exists stlistunittest;
create table stlistunittest(
	id varchar(10) not null primary key,
	name varchar(100) not null,
	exchange varchar(10),
	website varchar(100)
);