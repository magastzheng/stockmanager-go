package stmanager

import(
    "download"
    "parser"
    "handler/szsehandler"
    //"fmt"
    "entity/stentity"
)

type SZSECompanyManager struct {
    CompanyManagerBase
}

func (m *SZSECompanyManager) Process() {
    companies := m.GetCompanies()
    m.InsertDB(companies)
}

func (m *SZSECompanyManager) GetCompanies() []stentity.Company {
    down := download.NewSZSEDownloader()
    stlist := down.GetList()
    h := szsehandler.NewStockListHandler()
    p := parser.NewTextParser(h)
    p.ParseStr(stlist)
    
    //fmt.Println(h.Companies)

    return h.Companies
}

func NewSZSECompanyManager() *SZSECompanyManager {
    m := new(SZSECompanyManager)
    m.Init()
    return m
}
