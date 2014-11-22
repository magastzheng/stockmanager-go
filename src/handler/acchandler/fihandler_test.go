package acchandler_test

import(
    "testing"
    "parser"
    "handler/acchandler"
    "util"
    "fmt"
)

func Test_FiHanlder(t *testing.T) {
    //filename := "../../resource/account/financialindex-600001.html"
    //filename := "../../data/000002.dat"
    filename := "../../data/000002.html"
    str := util.ReadFile(filename)
    //fmt.Println(str) 
    h := acchandler.NewFiHandler()
    p := parser.NewTextParser(h)
    p.ParseStr(str)

    fmt.Println(len(h.DateMap))
    fmt.Println(len(h.DataMap))

    Output_FiHandler_DateMap(h.DateMap)
    Output_FiHandler_DataMap(h.DataMap)
}

func Output_FiHandler_DateMap(dateMap map[string]string) {
    for k, v := range dateMap {
        s := fmt.Sprintf("%s: %s", k, v)
        fmt.Println(s)
    }
}


func Output_FiHandler_DataMap(dataMap map[string]map[string]float32) {
    str := ""
    for date, dm := range dataMap {
        s := fmt.Sprintf("=============Date: %s ============", date)
        fmt.Println(s)
        for k, v := range dm {
            s1 := fmt.Sprintf("%s: %f\n", k, v)
            str += s1
            //fmt.Print(s1)
        }

        str += s +"\n"
    }
    //fmt.Println(str)
    util.WriteFile("writedata.dat", str)
}
