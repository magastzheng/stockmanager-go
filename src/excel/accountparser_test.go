package excel_test

import(
    "testing"
    "excel"
    "io/ioutil"
	"fmt"
	"code.google.com/p/mahonia"
)

func Test_TestAccountParser(t *testing.T){
    //filename := "../resource/account/balancesheet-300028.xls"
    filename := "../resource/account/cashflowstatement-300028.xls"
    buf, err := ioutil.ReadFile(filename)
    if err != nil{
        t.Error("Cannot load file:", filename)
    }
    
    data := string(buf)
	decoder := mahonia.NewDecoder("gbk")
	data = decoder.ConvertString(data)
    p := excel.NewAccountParser()
    dataMap := p.Parse(data)
	Output_DataMap(dataMap)
}

func Output_DataMap(dataMap map[string]map[string]float32){
	for date, datakeyval := range dataMap{
		fmt.Println("=======", date, "=========")
		for k, v := range datakeyval{
			fmt.Println(k, "\t:", v)
		}
	}
}
