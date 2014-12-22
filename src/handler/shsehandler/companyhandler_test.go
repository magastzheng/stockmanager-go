package shsehandler_test

import(
    "testing"
    "handler/shsehandler"
    "parser/jsonparser"
    "io/ioutil"
    "strings"
    "fmt"
)

func Test_CompanyHandler_Parse(t *testing.T){
    filename := "../../resource/shsecompany.dat"
    chunks, err := ioutil.ReadFile(filename)
    if err != nil{
        fmt.Println(err)
        panic(err)
    }

    data := string(chunks)
    start := strings.Index(data, "({")
    end := strings.LastIndex(data, "})")
    str := string(data[start+1: end+1])

    handler := shsehandler.NewCompanyHandler()
    parser := jsonparser.NewJsonParser(handler)
    parser.ParseStr(str)

    fmt.Println(handler.Company)
}


