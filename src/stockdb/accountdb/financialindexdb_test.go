package accountdb_test

import(
    "testing"
    "stockdb/accountdb"
    //"fmt"
)

func Test_FinancialIndexDB_Create(t *testing.T){
    db := accountdb.NewFinancialIndexDB("stocktest")
    sqls := []string{"drop table if exists accountfinancialindex_test",
    "create table accountfinancialindex_test (id int not null, name varchar(10) not null, age int)"}
    db.Create(sqls) 
}
