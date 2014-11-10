//Parse the stock code and name from the HMTL page
package handler_test

import(
    "testing"
    "handler"
    "fmt"
    "os"
    "io"
    "bytes"
    "parser"
)

func Test_StockHandler(t *testing.T){
    filename := "../stock_a-ha.dat"
    file, err := os.Open(filename)
 
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    chunks := bytes.NewBuffer(nil)
    io.Copy(chunks, file)
    str := string(chunks.Bytes())
    h := new(handler.StockHandler)
    h.Init()
    parser := new(parser.TextParser)
    parser.SetHandler(h)
    parser.ParseStr(str)
    
    fmt.Println(len(h.Stocks))
    //h.PrintStocks()
    //fmt.Println("600725", h.Stocks["600725"].id)
    
    if len(h.Stocks) == 982 {
        t.Log("Pass")
    } else {
        t.Error("The total stock number is wrong!")
    } 
}
