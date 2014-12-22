/*--drop database if exists stocksummary;
--create database stocksummary;
--use stocksummary;
*/
use chinastock;

/*
--drop table if exists stocklist;
--create table stocklist(
--   id varchar(10) not null primary key,
--    name varchar(100) not null,
--    exchange varchar(10) not null,
--);
*/
drop table if exists newstocklist;
create table newstocklist(
	id varchar(10) not null primary key,
	name varchar(100) not null,
	exchange varchar(10) not null
);

drop table if exists stocksummary;
create table stocksummary(
	id varchar(10) not null primary key,
	name varchar(20) not null,
	fullname varchar(100) not null,
	fullname_en varchar(100),
	inceptdate date,
	regaddr varchar(120),
	addr varchar(120),
	legalrepresentative varchar(30),
	boardsecretary varchar(30),
	email varchar(50),
	phone varchar(15),
	website varchar(100),
	csrccategory varchar(30),
	csrcbigcategory varchar(30),
	csrcmidcategory varchar(30),
	sseindustory varchar(30),
	states varchar(10),
	city varchar(10),
	currentstate int,
	shsample int,
	inforeign int,
	foreignaddr varchar(100)
);