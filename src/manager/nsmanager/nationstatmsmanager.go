package nsmanager

import(
    "parser"
    nsdb "stockdb/nationstatdb"
    //"fmt"
)

type NSMSManager struct{
    NSManagerBase
    parser *parser.NSMSParser
    db *nsdb.NSMoneySupplyDB
}

func (m *NSMSManager) Init(){
    m.parser = parser.NewNSMSParser()
    m.db = nsdb.NewMSMoneySupplyDB("macroindecis")
    m.NSManagerBase.Init()
}

func (m *NSMSManager) Process() {
    filename := "../data/actualdata-A0B0101-198301--1.dat"
    dataMap := m.LoadData(filename)
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

func NewNSMSManager() *NSMSManager{
    m := new(NSMSManager)
    m.Init()

    return m
}
