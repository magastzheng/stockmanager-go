package stsummary_test

import (
    "testing"
    db "stockdb/stsummary"
    "entity"
    "fmt"
)

func Test_StockListDBInsert(t *testing.T) {
    //stdb := db.NewStockListDB("mysql", "root@/stocktest")
    stdb := db.NewStockListDB("stocktest", "stlistunittest")
    stock := entity.Stock{}
    stock.Id = "1234"
    stock.Name = "test"
    stock.Exchange = "Shenzhen"
    
    res := stdb.Delete(stock)

    res = stdb.Insert(stock)
    fmt.Println(res)

    //res = stdb.Delete(stock)
}

func Test_StockListDBQuery(t *testing.T) {
    //stdb := stockdb.NewStockListDB("mysql", "root@/stocktest")
    stdb := db.NewStockListDB("stocktest", "stlistunittest")
    stock := stdb.Query("601005")
    fmt.Println(stock)
}

func Test_StockListDBQueryIds(t *testing.T) {
    //stdb := db.NewStockListDB("mysql", "root@/stocktest")
    stdb := db.NewStockListDB("stocktest", "stlistunittest")
    ids := stdb.QueryIds()
    fmt.Println("ID num:", len(ids))
    //fmt.Println(ids)
}

func Test_StockListDBIdExchange(t *testing.T){
    stdb := db.NewStockListDB("stocktest", "stlistunittest")
    idexchs := stdb.GetIdExchange()
    fmt.Println("IDEXCH num:", len(idexchs))
    //fmt.Println(idexchs)
}
