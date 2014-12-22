package shseparser_test

import(
    "testing"
    "parser/shseparser"
    "io/ioutil"
    "fmt"
)

func Test_CompanyParser_Parse(t *testing.T){
    filename := "../../resource/shsecompany.dat"
    chunks, err := ioutil.ReadFile(filename)
    if err != nil{
        fmt.Println(err)
        panic(err)
    }

    str := string(chunks)
    p := shseparser.NewCompanyParser()
    p.Parse(str)

    fmt.Println(p.Companies)
}
