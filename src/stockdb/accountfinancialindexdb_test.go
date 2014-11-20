package stockdb_test

import(
    "testing"
    "stockdb"
    //"fmt"
)

func Test_AccountFinancialIndexDB_Create(t *testing.T){
    db := stockdb.NewAccountFinancialIndexDB("stocktest")
    sqls := []string{"drop table if exists accountfinancialindex_test",
    "create table accountfinancialindex_test (id int not null, name varchar(10) not null, age int)"}
    db.Create(sqls) 
}
