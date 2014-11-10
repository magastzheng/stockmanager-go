package parser_test

import(
    "testing"
    "io/ioutil"
    "util"
    "parser"
    "fmt"
)

func Test_NSParseData(t *testing.T){
    filename := "../resource/nationstat-data.dat"
    bytes, err := ioutil.ReadFile(filename)
    util.CheckError(err)
    data := string(bytes)

    p := parser.NewNSParser()
    result := p.ParseData(data)
    
    if len(result.TableData) == 0 {
        t.Error("Parse json wrong")
    }
    fmt.Println(result)
}

func Test_NSParseIndexData(t *testing.T){
    filename := "../resource/nationstat-index.dat"
    bytes, err := ioutil.ReadFile(filename)
    util.CheckError(err)
    data := string(bytes)

    p := parser.NewNSParser()
    result := p.ParseIndex(data)
    
    if len(result) == 0 {
        t.Error("Parse index json data wrong")
    }
    
    for _, v := range result {
        str := fmt.Sprintf("Id: %v, Name: %v, PId: %v, EName: %v, IfData: %v, IsParent: %v", v.Id, v.Name, v.PId, v.EName, v.IfData, v.IsParent)
        fmt.Println(str)
    }

    //fmt.Println(result)
}

func Test_NSParsePeriodData(t *testing.T){
    filename := "../resource/nationstat-period.dat"
    bytes, err := ioutil.ReadFile(filename)
    util.CheckError(err)
    data := string(bytes)

    p := parser.NewNSParser()
    result := p.ParsePeriod(data)
    
    if len(result) == 0 {
        t.Error("Parse period json data wrong")
    }
    fmt.Println(result)
}
