package listmanager

import (
    "excel"
    "stockdb"
)

type IndustryManager struct {
    parser *excel.IndustryParser
}

func (m *IndustryManager) Init() {
    m.parser = excel.NewIndustryParser("../resource/hyflbz.xlsx")
}

func (m *IndustryManager) Process() {
    bigdb := stockdb.NewIndustryDB("chinastock")
    bigdb.TranInsertIndustry(m.parser.BigMap)

    minordb := stockdb.NewMinorIndustryDB("chinastock")
    minordb.TranInsertIndustry(m.parser.MinorMap)
}

func NewIndustryManager()*IndustryManager{
    m := new(IndustryManager)
    m.Init()

    return m
}


