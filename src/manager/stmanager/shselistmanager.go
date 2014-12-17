package stmanager

import(
    "download"
    "stockdb"
    "stockdb/stsummary"
    "parser/shseparser"
    "config"
    //"fmt"
    "os"
    "entity"
    "util"
)

type SHSEListManager struct {
    exchmanager *config.ExchangeConfigManager
    download *download.SHSEDownloader
    db *stockdb.StockListDB
    newdb *stsummary.StockListDB
    logger *util.StockLog
}

func (m *SHSEListManager) Init() {
    m.exchmanager = config.NewExchangeConfigManager()
    m.download = download.NewSHSEDownloader()
    m.db = stockdb.NewStockListDB("chinastock")
    m.newdb = stsummary.NewStockListDB("chinastock", "newstocklist")
    m.logger = util.NewLog()
}

func (m *SHSEListManager) Process() {
    stockids := m.db.QueryIds()
    newstocks := m.ProcessList(stockids)
    if len(newstocks) > 0 {
        m.logger.Info("New stock added from SE:", len(newstocks))
        m.newdb.TranInsert(newstocks)
    }
}

func (m *SHSEListManager) ProcessList(stockids []string) []entity.Stock {
    exchange, _ := m.exchmanager.GetExchange("CHS", "Shanghai")

    //get stocklist
    stlist := m.download.GetList()
    p := shseparser.NewListParser()
    p.Parse(stlist)
    m.logger.Info("SH SE new stock count: ", len(p.Stocks))
    
    newstocks := make([]entity.Stock, 0)
    for _, stockitem := range p.Stocks {
        isExisted := m.Contains(stockids, stockitem.Id)
        if !isExisted {
            stockitem.Exchange = exchange.Code

            newstocks = append(newstocks, stockitem)
        }
    }

    return newstocks
}

func (m *SHSEListManager) Contains(list []string, key string) (ok bool) {
    for _, v := range list {
        if v == key {
            ok = true
            break
        }
    }

    return
}

func (m *SHSEListManager) WriteFile(filename string, content string) {
    file, err := os.Create(filename)
    defer file.Close()
    if err != nil {
        panic(err)
    }

    file.WriteString(content)
}

func NewSHSEListManager() *SHSEListManager {
    m := new(SHSEListManager)
    m.Init()
    return m
}
