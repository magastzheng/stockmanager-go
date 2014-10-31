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
    chunks, err := ioutil.ReadFile(filename)
    util.CheckError(err)
    str := string(chunks)
    handler := stockhandler.NewStockHistDataHandler()
    handler.Init()

    parser := parser.NewTextParser(handler)
    parser.ParseStr(str)
    
    fmt.Println(handler.Code)
    fmt.Println(handler.Data)
}

