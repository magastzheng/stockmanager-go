use chinastock;

drop table if exists csrcindustry;
create table csrcindustry(
    code varchar(10) not null primary key,
    parent varchar(10),
    name varchar(80) not null,
    name_en varchar(80) not null
);

drop table if exists csrcstockcategory;
create table csrcstockcategory(
	id varchar(10) not null primary key,
	code varchar(10) not null
);