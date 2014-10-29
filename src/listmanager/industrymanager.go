package listmanager

import (
    "excel"
    "stockdb"
    "config"
)

type IndustryManager struct {
    config config.DBItem
    parser *excel.IndustryParser
}

func (m *IndustryManager) Init() {
    dbconfig := config.NewDBConfig("../config/dbconfig.json")
    m.config = dbconfig.GetConfig("chinastock")

    m.parser = excel.NewIndustryParser("../resource/hyflbz.xlsx")
}

func (m *IndustryManager) Process() {
    bigdb := stockdb.NewIndustryDB(m.config.Dbtype, m.config.Dbcon)
    bigdb.TranInsertIndustry(m.parser.BigMap)

    minordb := stockdb.NewMinorIndustryDB(m.config.Dbtype, m.config.Dbcon)
    minordb.TranInsertIndustry(m.parser.MinorMap)
}


