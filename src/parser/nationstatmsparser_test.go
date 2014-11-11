package parser_test

import(
    "testing"
    "parser"
    //"entity"
    "strings"
    "os"
    "bufio"
    "fmt"
)

func Test_NSMSParser_Parse(t *testing.T) {
    filename := "../data/actualdata-A0B0101-198301--1.dat"
    f, err := os.Open(filename)
    if err != nil {
        fmt.Println("Cannot open file:", filename, err)
        return
    }
    
    dataMap := make(map[string] string)
    r := bufio.NewReader(f)
    err = nil
    count := 0
    for err == nil {
        line, err := r.ReadString('\n'); 
        if err != nil {
            fmt.Println(line)
            fmt.Println("Fail or end of the file", err)
            break
        }
        count++
        if strings.Contains(line, ":") {
            arr := strings.Split(line, ":")
            key := strings.TrimSpace(arr[0])
            value := strings.TrimSpace(arr[1])
            dataMap[key] = value
        }
    }
    
    p := parser.NewNSMSParser()
    p.Parse(dataMap)

    fmt.Println(len(p.Data))
    //fmt.Println(p.Data)
    fmt.Println(count)
}
