package manager

import(
    "download"
    "stockdb"
    "shseparser"
    "config"
    //"fmt"
    "os"
    //"encoding/json"
)

type SHSEListManager struct {
    config config.StockListConfig
    exchmanager config.ExchangeConfigManager
    download *download.SHSEDownloader
    db *stockdb.StockListDB
}

func (m *SHSEListManager) Init() {
    m.config = config.Parse(filename)
    m.exchmanager = config.NewExchangeConfigManager()
    m.download = download.NewSHSEDownloader()
    m.db = stockdb.NewStockListDB("chinastock")
}

func (m *SHSEListManager) Process() {
    


    for _, c := range categories {
        //fmt.Println(i,c)
        pageStr := s.download.GetPage(baseUrl, c.Type, c.Class)
        
        //s.parser.ParseStr(pageStr)
        h := handler.NewStockHandler()
        parser := parser.NewTextParser(h)
        parser.ParseStr(pageStr)
        //fmt.Println(len(h.Stocks))
        //stockstr := h.ToJson()
        //s.WriteFile(c.Type + c.Class, stockstr)
        s.db.TranInsert(c.Exchange, h.Stocks)
        //exchange := c.Exchange
        //for _, st := range s.h.Stocks {
        //    fmt.Println(id)
        //    s.db.DeleteStock(st)
        //    s.db.InsertStock(exchange, st)
        //}
    }
}

func (m *SHSEListManager) ProcessShanghai(){
    //get Shanghai stocklist
    stlist := m.download.GetList()
    shparser := shseparser.NewListParser()
    shparser.Parse(stlist)
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
