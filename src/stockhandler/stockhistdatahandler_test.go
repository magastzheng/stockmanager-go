package stockhandler_test

import (
    "testing"
    "stockhandler"
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
    handler := stockhandler.NewStockHistDataHandler()
    handler.Init()

    parser := parser.NewTextParser(handler)
    parser.ParseStr(str)
    
    fmt.Println(handler.Code)
    //fmt.Println(handler.Data)
    fmt.Println(handler.Years)
    PrintHistData(handler.Data)
}

func PrintHistData(datas []stockhandler.StockHistData) {
    format := "Date: %v, open: %v, close: %v, high: %v, low: %v, volume: %v, money: %v"
    for i, d := range datas {
        s := fmt.Sprintf(format, d.Date, d.Open, d.Close, d.Highest, d.Lowest, d.Volume, d.Money)
        fmt.Println(i, s)
    }
}

