use chinastock;

create procedure getcsrccategorystock(
	in p_categoryid varchar(10)
	)
begin
	select id from csrcstockcategory where code=p_categoryid;
end