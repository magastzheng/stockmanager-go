package nsmanager

import(
    //"parser"
    //"stockdb"
    "util"
    "os"
    "bufio"
    "strings"
    //"fmt"
)

type NSManagerBase struct{
    logger *util.StockLog
}

func (m *NSManagerBase) Init(){
    m.logger = util.NewLog()
}

func (m *NSManagerBase) Process() {

}

func (m *NSManagerBase) LoadData(filename string) map[string] string {
    dataMap := make(map[string] string)

    //filename := "../data/actualdata-A0B0101-198301--1.dat"
    f, err := os.Open(filename)
    if err != nil {
        m.logger.Error("Cannot open file:", filename, err)
        return dataMap
    }
    
    r := bufio.NewReader(f)
    err = nil
    count := 0
    for err == nil {
        line, err := r.ReadString('\n'); 
        if err != nil {
            m.logger.Error(line)
            m.logger.Error("Fail or end of the file", err)
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
    
    return dataMap
}

