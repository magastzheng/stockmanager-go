package industrymanager

import (
    "manager"
    xls "excel/industry"
    "stockdb/industrydb"
    //"fmt"
)

type IndustryManager struct {
    manager.ManagerBase
    parser *xls.IndustryParser
    db *industrydb.IndustryDB
}

func (m *IndustryManager) Init() {
    m.ManagerBase.Init()
    filename := m.BaseDir + "resource/industry/hyflbz.xlsx"
    //fmt.Println("filename: ", filename)
    m.parser = xls.NewIndustryParser(filename)
    m.db = industrydb.NewIndustryDB("chinastock", "csrcindustry")
}

func (m *IndustryManager) Process() {
    
    //fmt.Println(m.parser.BigMap)
    //fmt.Println(m.parser.MinorMap)

    m.db.TranInsert(m.parser.BigMap)
    m.db.TranInsert(m.parser.MinorMap)
}

func NewIndustryManager()*IndustryManager{
    m := new(IndustryManager)
    m.Init()

    return m
}


