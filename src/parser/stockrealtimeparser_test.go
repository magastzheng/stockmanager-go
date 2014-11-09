package parser_test

import(
    "testing"
    "parser"
    "io/ioutil"
    "util"
    "fmt"
)

func Test_ParseRtData(t *testing.T){
    filename := "../resource/realtime.dat"
    bytes, err := ioutil.ReadFile(filename)
    if err != nil {
        util.CheckError(err)
    }

    data := string(bytes)
    p := parser.NewStockRtParser()
    p.ParseStr(data)

    fmt.Println(p.Data)
}
