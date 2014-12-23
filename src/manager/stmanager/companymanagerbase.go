package stmanager

import(
    //"download"
    "stockdb"
    "stockdb/stsummary"
    "entity/stentity"
    "config"
    "util"
    //"fmt"
)

type CompanyManagerBase struct {
    exchmanager *config.ExchangeConfigManager
    db *stockdb.StockListDB
    compdb *stsummary.CompanyDB
    logger *util.StockLog
}

func (m *CompanyManagerBase) Init() {
    m.exchmanager = config.NewExchangeConfigManager()
    //m.download = download.NewSHSEDownloader()
    m.db = stockdb.NewStockListDB("chinastock")
    m.compdb = stsummary.NewCompanyDB("chinastock", "stocksummary")
    m.logger = util.NewLog()
}

func (m *CompanyManagerBase) InsertDB(companies []stentity.Company) int {
    if len(companies) > 0 {
        return m.compdb.TranInsert(companies)
    }

    return -1
}
