package stmanager

import(
    "download"
    "stockdb"
    "entity/stentity"
    "parser/shseparser"
    "config"
    "util"
    "fmt"
)

type SHSECompanyManager struct {
    exchmanager *config.ExchangeConfigManager
    download *download.SHSEDownloader
    db *stockdb.StockListDB
    logger *util.StockLog
}

func (m *SHSECompanyManager) Init() {
    m.exchmanager = config.NewExchangeConfigManager()
    m.download = download.NewSHSEDownloader()
    m.db = stockdb.NewStockListDB("chinastock")
    m.logger = util.NewLog()
}

func (m *SHSECompanyManager) Process() {
    exchange, _ := m.exchmanager.GetExchange("CHS", "Shanghai")
    stockids := m.db.QueryIdsByExchange(exchange.Code)

    companies := make([]stentity.Company, 0)
    for _, code := range stockids {
        c := m.GetCompany(code)
        companies = append(companies, c)
    }

    fmt.Println(companies)
}

func (m *SHSECompanyManager) GetCompany(code string) stentity.Company {
    data := m.download.GetCompanyInfo(code)
    p := shseparser.NewCompanyParser()
    p.Parse(data)
    c := p.Company
    
    data = m.download.GetCompanyIncpt(code)
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
