package stockdb_test

import (
    "testing"
    "stockhandler"
    "stockdb"
    "fmt"
)

func Test_StockDatabaseInsert(t *testing.T) {
    stdb := stockdb.NewStockDatabase("mysql", "root@/chinastock")
    
    stock := stockhandler.Stock{}
    stock.Id = "1234"
    stock.Name = "test"
    stock.Website = "http://www.1234.com"
    
    res := stdb.Delete(stock)

    res = stdb.Insert("Sh", stock)
    fmt.Println(res)

    res = stdb.Delete(stock)
}

func Test_StockDatabaseQuery(t *testing.T) {
    stdb := stockdb.NewStockDatabase("mysql", "root@/chinastock")
    stock := stdb.Query("601005")
    fmt.Println(stock)
}

func Test_StockDatabaseQueryIds(t *testing.T) {
    stdb := stockdb.NewStockDatabase("mysql", "root@/chinastock")
    ids := stdb.QueryIds()
    fmt.Println("ID num:", len(ids))
    //fmt.Println(ids)
}
