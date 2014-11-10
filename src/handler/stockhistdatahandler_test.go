package handler_test

import (
    "testing"
    "handler"
    "fmt"
    "parser"
    "io/ioutil"
    "util"
)

func Test_StockHistDataHandler(t *testing.T) {
    filename := "../resource/stockhistdata.dat"
    //filename := "../resource/000002-2007-3.dat"
    chunks, err := ioutil.ReadFile(filename)
    util.CheckError(err)
    str := string(chunks)
    h := handler.NewStockHistDataHandler()
    h.Init()

    parser := parser.NewTextParser(h)
    parser.ParseStr(str)
    
    fmt.Println(h.Code)
    //fmt.Println(h.Data)
    fmt.Println(h.Years)
    PrintHistData(h.Data)
}

func PrintHistData(datas []handler.StockHistData) {
    format := "Date: %v, open: %v, close: %v, high: %v, low: %v, volume: %v, money: %v"
    for i, d := range datas {
        s := fmt.Sprintf(format, d.Date, d.Open, d.Close, d.Highest, d.Lowest, d.Volume, d.Money)
        fmt.Println(i, s)
    }
}

