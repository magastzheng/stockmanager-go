package shseparser_test

import(
    "testing"
    "shseparser"
    "io/ioutil"
    "fmt"
)

func Test_LisParser_Parse(t *testing.T){
    filename := "../data/shse_list.dat"
    buf, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Cannot read file: ", filename)
    }
    src := string(buf)
    p := shseparser.NewListParser()
    row := p.Parse(src)
    fmt.Println(row)
    fmt.Println(len(p.Stocks))
    fmt.Println(p.Stocks)
}

