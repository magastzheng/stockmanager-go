//Parse the stock code and name from the HMTL page
package stockhandler

import(
    "testing"
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
    handler := new(StockHandler)
    handler.Init()
    parser := new(parser.TextParser)
    parser.SetHandler(handler)
    parser.ParseStr(str)
    
    fmt.Println(len(handler.Stocks))
    //handler.PrintStocks()
    //fmt.Println("600725", handler.Stocks["600725"].id)
    
    if len(handler.Stocks) == 982 {
        t.Log("Pass")
    } else {
        t.Error("The total stock number is wrong!")
    } 
}
