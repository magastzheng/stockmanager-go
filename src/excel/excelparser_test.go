package excel_test

import(
    "testing"
    "excel"
    "fmt"
)

func Test_Parse(t *testing.T) {
    parser := new(excel.IndustryParser)
    parser.Parse("../resource/csrcindustry.xlsx")
    fmt.Println(len(parser.Rows))
    fmt.Println("BigMap", len(parser.BigMap))
    fmt.Println("MinorMap", len(parser.MinorMap))
    fmt.Println(parser.BigMap)
    fmt.Println(parser.MinorMap)
}

