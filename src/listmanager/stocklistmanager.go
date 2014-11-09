package listmanager

import(
    "download"
    "stockdb"
    "parser"
    "stockhandler"
    "config"
    //"fmt"
    "os"
    //"encoding/json"
)

type StockListManager struct {
    config config.StockListConfig
    download *download.StockDownloader
    db *stockdb.StockDatabase
}

func (s *StockListManager) Init() {
    const filename = "../config/stocklist.json"
    s.config = config.Parse(filename)
    s.download = download.NewDownloader()
    s.db = stockdb.NewStockDatabase("chinastock")
}

func (s *StockListManager) Process() {
    baseUrl := s.config.Sites.BaseUrl
    categories := s.config.Sites.Categories
    //fmt.Println(baseUrl)
    //fmt.Println(len(categories))
    for _, c := range categories {
        //fmt.Println(i,c)
        pageStr := s.download.GetPage(baseUrl, c.Type, c.Class)
        
        //s.parser.ParseStr(pageStr)
        handler := stockhandler.NewStockHandler()
        parser := parser.NewTextParser(handler)
        parser.ParseStr(pageStr)
        //fmt.Println(len(handler.Stocks))
        //stockstr := handler.ToJson()
        //s.WriteFile(c.Type + c.Class, stockstr)
        s.db.TranInsert(c.Exchange, handler.Stocks)
        //exchange := c.Exchange
        //for _, st := range s.handler.Stocks {
        //    fmt.Println(id)
        //    s.db.DeleteStock(st)
        //    s.db.InsertStock(exchange, st)
        //}
    }
}

func (s *StockListManager) WriteFile(filename string, content string) {
    file, err := os.Create(filename)
    defer file.Close()
    if err != nil {
        panic(err)
    }

    file.WriteString(content)
}

func NewStockListManager() *StockListManager {
    m := new(StockListManager)
    m.Init()
    return m
}
