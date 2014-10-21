package stockdb

import (
    "testing"
    "stockhandler"
    "fmt"
)

func Test_StockDatabaseInsert(t *testing.T) {
    stdb := NewStockDatabase("mysql", "root@/chinastock")
    
    stock := stockhandler.Stock{}
    stock.Id = "1234"
    stock.Name = "test"
    stock.Website = "http://www.1234.com"
    
    res := stdb.DeleteStock(stock)

    res = stdb.InsertStock("Sh", stock)
    fmt.Println(res)
}
