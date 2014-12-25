package industry_test

import(
    "testing"
    "excel/industry"
    "fmt"
)

func Test_ParseIndustry(t *testing.T) {
    parser := industry.NewIndustryParser("../../resource/industry/hyflbz.xlsx")
    fmt.Println(len(parser.NewRows))
    fmt.Println(len(parser.OldRows))
    fmt.Println("BigMap", len(parser.BigMap))
    fmt.Println("MinorMap", len(parser.MinorMap))
    fmt.Println("===========New Rows===========")
    fmt.Println(parser.NewRows)
    //PrintRows(parser.NewRows)
    fmt.Println("===========Old Rows===========")
    //PrintRows(parser.OldRows)
    fmt.Println(parser.OldRows)
    
    fmt.Println(parser.BigMap)
    fmt.Println(parser.MinorMap)
}

func PrintRows(rows []*industry.IndustryRow) {
    for _, row := range rows {
        fmt.Println(row)
        fmt.Println(row.Column1, row.Column2, row.Column3)
    }
}

