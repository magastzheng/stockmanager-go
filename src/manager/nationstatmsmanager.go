package manager

import(
    "parser"
    "stockdb"
    "util"
    "os"
    "bufio"
    "strings"
    //"fmt"
)

type NSMSManager struct{
    parser *parser.NSMSParser
    db *stockdb.NSMoneySupplyDB
    logger *util.StockLog
}

func (m *NSMSManager) Init(){
    m.parser = parser.NewNSMSParser()
    m.db = stockdb.NewMSMoneySupplyDB("macroindecis")
    m.logger = util.NewLog()
}

func (m *NSMSManager) Process() {
    dataMap := m.LoadData()
    if len(dataMap) == 0 {
        m.logger.Error("Cannot load data of money supply")
    }
    m.parser.Parse(dataMap)
    
    count := len(m.parser.Data)
    if count > 0 {
        m.logger.Info("Insert the money supply into database")
        m.db.TranInsert(m.parser.Data)
        m.logger.Info("Complete to insert money supply into database: ", count)
    } else {
        m.logger.Error("No data can be inserted!")
    }
}

func (m *NSMSManager) LoadData() map[string] string {
    dataMap := make(map[string] string)

    filename := "../data/actualdata-A0B0101-198301--1.dat"
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

func NewNSMSManager() *NSMSManager{
    m := new(NSMSManager)
    m.Init()

    return m
}
