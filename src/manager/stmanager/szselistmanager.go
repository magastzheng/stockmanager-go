package stmanager

import(
    "download"
    "stockdb"
    "stockdb/stsummary"
    "parser"
    "handler/szsehandler"
    "config"
    //"fmt"
    "os"
    "entity"
    "util"
)

type SZSEListManager struct {
    exchmanager *config.ExchangeConfigManager
    download *download.SZSEDownloader
    db *stockdb.StockListDB
    newdb *stsummary.StockListDB
    logger *util.StockLog
}

func (m *SZSEListManager) Init() {
    m.exchmanager = config.NewExchangeConfigManager()
    m.download = download.NewSZSEDownloader()
    m.db = stockdb.NewStockListDB("chinastock")
    m.newdb = stsummary.NewStockListDB("chinastock", "newstocklist")
    m.logger = util.NewLog()
}

func (m *SZSEListManager) Process() {
    stockids := m.db.QueryIds()
    newstocks := m.ProcessList(stockids)
    if len(newstocks) > 0 {
        m.logger.Info("New stock added from SE:", len(newstocks))
        m.newdb.TranInsert(newstocks)
    }
}

func (m *SZSEListManager) ProcessList(stockids []string) []entity.Stock {
    exchange, _ := m.exchmanager.GetExchange("CHS", "Shenzhen")

    //get stocklist
    stlist := m.download.GetList()
    h := szsehandler.NewStockListHandler()
    p := parser.NewTextParser(h)
    p.ParseStr(stlist)
    
    m.logger.Info("SZ SE new stock count: ", len(h.Companies))

    newstocks := make([]entity.Stock, 0)
    for _, company := range h.Companies {
        isExisted := m.Contains(stockids, company.Code)
        if !isExisted {
            stockitem := entity.Stock{
                Id: company.Code,
                Name: company.AbbrName,
                Exchange: exchange.Code,
            }

            newstocks = append(newstocks, stockitem)
        }
    }

    return newstocks
}

func (m *SZSEListManager) Contains(list []string, key string) (ok bool) {
    for _, v := range list {
        if v == key {
            ok = true
            break
        }
    }

    return
}

func (m *SZSEListManager) WriteFile(filename string, content string) {
    file, err := os.Create(filename)
    defer file.Close()
    if err != nil {
        panic(err)
    }

    file.WriteString(content)
}

func NewSZSEListManager() *SZSEListManager {
    m := new(SZSEListManager)
    m.Init()
    return m
}
