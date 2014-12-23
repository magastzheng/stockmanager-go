package stmanager

import(
    "download"
    "entity/stentity"
    "parser/shseparser"
    //"fmt"
)

type SHSECompanyManager struct {
    CompanyManagerBase
}

func (m *SHSECompanyManager) Process() {
    companies := m.GetCompanies()
    m.InsertDB(companies)
}

func (m *SHSECompanyManager) GetCompanies() []stentity.Company {
    exchange, _ := m.exchmanager.GetExchange("CHS", "Shanghai")
    stockids := m.db.QueryIdsByExchange(exchange.Code)
    //fmt.Println(stockids)
    //stockids = stockids[0: 2]

    companies := make([]stentity.Company, 0)
    for _, code := range stockids {
        c := m.GetCompany(code)
        companies = append(companies, c)
    }

    return companies
}

func (m *SHSECompanyManager) GetCompany(code string) stentity.Company {
    down := download.NewSHSEDownloader()
    data := down.GetCompanyInfo(code)
    p := shseparser.NewCompanyParser()
    p.Parse(data)
    c := p.Company
    
    data = down.GetCompanyIncpt(code)
    p = shseparser.NewCompanyParser()
    p.Parse(data)
    c.InceptDate = p.Company.InceptDate

    return c
}

func NewSHSECompanyManager() *SHSECompanyManager {
    m := new(SHSECompanyManager)
    m.Init()
    return m
}
