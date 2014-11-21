package stockdb_test

import (
    "testing"
    "entity"
    "stockdb"
    "fmt"
)

func Test_StockListDBInsert(t *testing.T) {
    //stdb := stockdb.NewStockListDB("mysql", "root@/chinastock")
    stdb := stockdb.NewStockListDB("chinastock")
    stock := entity.Stock{}
    stock.Id = "1234"
    stock.Name = "test"
    stock.Website = "http://www.1234.com"
    
    res := stdb.Delete(stock)

    res = stdb.Insert("Sh", stock)
    fmt.Println(res)

    res = stdb.Delete(stock)
}

func Test_StockListDBQuery(t *testing.T) {
    //stdb := stockdb.NewStockListDB("mysql", "root@/chinastock")
    stdb := stockdb.NewStockListDB("chinastock")
    stock := stdb.Query("601005")
    fmt.Println(stock)
}

func Test_StockListDBQueryIds(t *testing.T) {
    //stdb := stockdb.NewStockListDB("mysql", "root@/chinastock")
    stdb := stockdb.NewStockListDB("chinastock")
    ids := stdb.QueryIds()
    fmt.Println("ID num:", len(ids))
    //fmt.Println(ids)
}

func Test_StockListDBIdExchange(t *testing.T){
    stdb := stockdb.NewStockListDB("chinastock")
    idexchs := stdb.GetIdExchange()
    fmt.Println("IDEXCH num:", len(idexchs))
    //fmt.Println(idexchs)
}
