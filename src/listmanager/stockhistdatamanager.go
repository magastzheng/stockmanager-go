package listmanager

import (
    "download"
    "stockhandler"
    "parser"
    "stockdb"
    "config"
)

type StockHistDataManager struct {
    downloader *download.StockHistDownloader
    parser *parser.TextParser
    handler *stockhandler.StockHistDataHandler
    stocklistdb *stockdb.StockDatabase
    db *stockdb.StockHistDataDB
}

func (m *StockHistDataManager) Init() {
    m.downloader = download.NewStockHistDownloader()
    m.handler = stockhandler.NewStockHistDataHandler()
    m.parser = parser.NewTextParser(m.handler)

    dbconfig := config.NewDBConfig("../config/dbconfig.json")
    config := dbconfig.GetConfig("chinastock")
    m.stocklistdb = stockdb.NewStockDatabase(config.Dbtype, config.Dbcon)
    m.db = stockdb.NewStockHistDataDB(config.Dbtype, config.Dbcon)
}

func (m *StockHistDataManager) Process() {
   //fmt.Println("text") 
}

func NewStockHistDataManager() *StockHistDataManager{
    manager := new(StockHistDataManager)
    manager.Init()

    return manager
}
