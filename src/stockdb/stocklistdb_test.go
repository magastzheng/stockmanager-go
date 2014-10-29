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
    
    res := stdb.DeleteStock(stock)

    res = stdb.InsertStock("Sh", stock)
    fmt.Println(res)

    res = stdb.DeleteStock(stock)
}

func Test_StockDatabaseQuery(t *testing.T) {
    stdb := stockdb.NewStockDatabase("mysql", "root@/chinastock")
    stock := stdb.QueryStock("601005")
    fmt.Println(stock)
}
