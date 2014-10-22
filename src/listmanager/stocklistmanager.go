package listmanager

import(
    "download"
    "config"
    "stockdb"
    "parser"
    "stockhandler"
    "fmt"
)

type StockListManager struct {
    config config.StockListConfig
    download *download.StockDownloader
    parser *parser.TextParser
    handler *stockhandler.StockHandler
    db *stockdb.StockDatabase
}

func (s *StockListManager) Init(filename string) {
    s.config = config.Parse(filename)
    s.download = download.NewDownloader()
    s.handler = stockhandler.NewStockHandler()
    s.parser = parser.NewTextParser(s.handler)
    s.db = stockdb.NewStockDatabase("mysql", "root@/chinastock?charset=utf8")
}

func (s *StockListManager) Process() {
    baseUrl := s.config.Sites.BaseUrl
    categories := s.config.Sites.Categories
    //fmt.Println(baseUrl)
    //fmt.Println(len(categories))
    for _, c := range categories {
        //fmt.Println(i,c)
        pageStr := s.download.GetPage(baseUrl, c.Type, c.Class)
        s.parser.ParseStr(pageStr)
        fmt.Println(len(s.handler.Stocks))
        s.db.TranInsertStock(c.Exchange, s.handler.Stocks)
        //exchange := c.Exchange
        //for _, st := range s.handler.Stocks {
        //    fmt.Println(id)
        //    s.db.DeleteStock(st)
        //    s.db.InsertStock(exchange, st)
        //}
    }
}
