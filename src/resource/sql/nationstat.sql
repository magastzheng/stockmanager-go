use macroindecis;

drop table if exists nsindex;
create table nsindex(
	id varchar(15) not null primary key,
	parent varchar(15),
	name varchar(100) not null,
	ename varchar(100),
	unit varchar(15),
	eunit varchar(15),
	note varchar(2048),
	enote varchar(2048),
	readid varchar(50)
);