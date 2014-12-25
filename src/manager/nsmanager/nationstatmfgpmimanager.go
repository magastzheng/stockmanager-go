package nsmanager

import(
    "parser"
    nsdb "stockdb/nationstatdb"
    //"fmt"
)

type NSMfgPmiManager struct{
    NSManagerBase
    parser *parser.NSMfgPmiParser
    db *nsdb.NSMfgPmiDB
}

func (m *NSMfgPmiManager) Init(){
    m.parser = parser.NewNSMfgPmiParser()
    m.db = nsdb.NewNSMfgPmiDB("macroindecis")
    m.NSManagerBase.Init()
}

func (m *NSMfgPmiManager) Process() {
    filename := "../data/actualdata-A090101-198301--1.dat"
    dataMap := m.LoadData(filename)
    if len(dataMap) == 0 {
        m.logger.Error("Cannot load data of money supply")
    }
    m.parser.Parse(dataMap)
    
    count := len(m.parser.Data)
    if count > 0 {
        m.logger.Info("Insert the manufacturing PMI into database")
        //fmt.Println(m.parser.Data)
        m.db.TranInsert(m.parser.Data)
        m.logger.Info("Complete to insert manufacturing PMI into database: ", count)
    } else {
        m.logger.Error("No data can be inserted!")
    }
}

func NewNSMfgPmiManager() *NSMfgPmiManager{
    m := new(NSMfgPmiManager)
    m.Init()

    return m
}
