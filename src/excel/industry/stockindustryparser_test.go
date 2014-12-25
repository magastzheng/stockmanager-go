package industry_test

import(
    "testing"
    "excel/industry"
    "fmt"
)

func Test_Parse(t *testing.T) {
    parser := industry.NewStockIndustryParser("../../resource/industry/csrcindustry.xlsx")
    fmt.Println(len(parser.Rows))
    fmt.Println("BigMap", len(parser.BigMap))
    fmt.Println("MinorMap", len(parser.MinorMap))
    fmt.Println(parser.BigMap)
    fmt.Println(parser.MinorMap)
}

