package manager

import(
    "download"
    "parser"
    "stockdb"
    "util"
   // "fmt"
)

type StockRtDataManager struct{
    downloader *download.StockRtDownloader
    listdb *stockdb.StockListDB
    datadb *stockdb.StockHistDataDB
    logger *util.StockLog
}

func (m *StockRtDataManager) Init() {
    m.downloader = download.NewStockRtDownloader()
    m.listdb = stockdb.NewStockListDB("chinastock")
    m.datadb = stockdb.NewStockHistDataDB("chinastock")
    m.logger = util.NewLog()
}

func (m *StockRtDataManager) Process() {
    idexchs := m.listdb.GetIdExchange()
    if len(idexchs) == 0 {
        m.logger.Error("Cannot get stock list id, exchange from database")
    }

    for _, idexch := range idexchs {
        str := m.downloader.GetData(idexch.Id, idexch.Exchange)
		if len(str) == 0 {
			m.logger.Error("Cannot get realtime date", idexch)
		} else {
			p := parser.NewStockRtParser()
			p.ParseStr(str)
			data := p.Data
        
			//fmt.Println(data)
			m.datadb.Insert(idexch.Id, data)
		}
    }
}

func NewStockRtDataManager() *StockRtDataManager{
    m := new(StockRtDataManager)
    m.Init()

    return m
}
