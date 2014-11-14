use chinastock;

/*drop table if exists balancesheet;*/
drop table if exists bs_currentasset;

create table bs_currentasset(
    code varchar(10) not null
    date date not null
    monetaryfund decimal(15, 2)
    
);


