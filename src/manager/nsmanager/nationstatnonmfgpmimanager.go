package nsmanager

import(
    "parser/nsparser"
    nsdb "stockdb/nationstatdb"
    //"fmt"
)

type NSNonMfgPmiManager struct{
    NSManagerBase
    parser *nsparser.NSNonMfgPmiParser
    db *nsdb.NSNonMfgPmiDB
}

func (m *NSNonMfgPmiManager) Init(){
    m.parser = nsparser.NewNSNonMfgPmiParser()
    m.db = nsdb.NewNSNonMfgPmiDB("macroindecis")
    m.NSManagerBase.Init()
}

func (m *NSNonMfgPmiManager) Process() {
    filename := "../data/actualdata-A090201-198301--1.dat"
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

func NewNSNonMfgPmiManager() *NSNonMfgPmiManager{
    m := new(NSNonMfgPmiManager)
    m.Init()

    return m
}
