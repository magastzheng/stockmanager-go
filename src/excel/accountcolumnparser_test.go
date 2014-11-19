package excel_test

import(
    "testing"
    "excel"
)

func Test_AccountColumnParser_Parse(t *testing.T){
    filename := "../resource/financialindex.xlsx"
    p := excel.NewAccountColumnParser()
    p.Parse(filename)
}
