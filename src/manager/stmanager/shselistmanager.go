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
    //"encoding/json"
)

type SHSEListManager struct {
    //config config.StockListConfig
    exchmanager *config.ExchangeConfigManager
    download *download.SHSEDownloader
    db *stockdb.StockListDB
    newdb *stsummary.StockListDB
    logger *util.StockLog
}

func (m *SHSEListManager) Init() {
    //m.config = config.Parse(filename)
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
    
    //for _, c := range categories {
        //fmt.Println(i,c)
        //pageStr := s.download.GetPage(baseUrl, c.Type, c.Class)
        
        //s.parser.ParseStr(pageStr)
        //h := handler.NewStockHandler()
        //parser := parser.NewTextParser(h)
        //parser.ParseStr(pageStr)
        //fmt.Println(len(h.Stocks))
        //stockstr := h.ToJson()
        //s.WriteFile(c.Type + c.Class, stockstr)
        //s.db.TranInsert(c.Exchange, h.Stocks)
        //exchange := c.Exchange
        //for _, st := range s.h.Stocks {
        //    fmt.Println(id)
        //    s.db.DeleteStock(st)
        //    s.db.InsertStock(exchange, st)
        //}
    //}
}

func (m *SHSEListManager) ProcessList(stockids []string) []entity.Stock {
    exchange, _ := m.exchmanager.GetExchange("CHS", "Shanghai")

    //get stocklist
    stlist := m.download.GetList()
    p := shseparser.NewListParser()
    p.Parse(stlist)
    //fmt.Println(shparser.Stocks)
    
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
