package excel_test

import(
    "testing"
    "excel"
    "io/ioutil"
)

func Test_TestAccountParser(t *testing.T){
    //filename := "../resource/account/balancesheet-300028.xls"
    filename := "../resource/account/cashflowstatement-300028.xls"
    buf, err := ioutil.ReadFile(filename)
    if err != nil{
        t.Error("Cannot load file:", filename)
    }
    
    data := string(buf)

    p := excel.NewAccountParser()
    p.Parse(data)
}
