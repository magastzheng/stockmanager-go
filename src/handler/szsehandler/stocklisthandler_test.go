//Parse the stock code and name from the HMTL page
package szsehandler_test

import(
    "testing"
    "handler/szsehandler"
    "code.google.com/p/mahonia"
    "fmt"
    "io/ioutil"
    "parser"
)

func Test_StockListHandler_Parse(t *testing.T){
    filename := "../../resource/szcompanylist.xls"
    chunks, err := ioutil.ReadFile(filename)
    if err != nil{
        fmt.Println(err)
        t.Error("Cannot read the file:", filename)
    }
    str := string(chunks)
    decoder := mahonia.NewDecoder("gbk")
    str = decoder.ConvertString(str)
    h := szsehandler.NewStockListHandler()
    parser := parser.NewTextParser(h)
    parser.ParseStr(str)
    
    fmt.Println(len(h.Companies))
    h.Output()
}
