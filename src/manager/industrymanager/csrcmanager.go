package industrymanager

import(
    "manager"
    xls "excel/industry"
    "stockdb/industrydb"
)

type CSRCManager struct {
    manager.ManagerBase
    parser *xls.StockIndustryParser
    db *industrydb.StockCategoryDB
}

func (m *CSRCManager) Init() {
    m.ManagerBase.Init()
    filename := m.BaseDir + "resource/industry/csrcindustry.xlsx"
    m.parser = xls.NewStockIndustryParser(filename)
    m.db = industrydb.NewStockCategoryDB("chinastock", "csrcstockcategory")
}

func (m *CSRCManager) Process() {
    scs := m.parser.Scs
    m.db.TranInsert(scs)
}

func NewCSRCManager() *CSRCManager {
    m := new(CSRCManager)
    m.Init()

    return m
}
