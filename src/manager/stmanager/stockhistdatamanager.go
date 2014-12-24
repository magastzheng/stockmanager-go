package stmanager

import (
    "download"
    "handler"
    "entity"
    "parser"
    "stockdb/stsummary"
    "stockdb"
    "config"
    "time"
    "util"
    "fmt"
)

type StockHistDataManager struct {
    exchmanager *config.ExchangeConfigManager
    listdb *stsummary.StockListDB
    newlistdb *stsummary.StockListDB
    downloader *download.StockHistDownloader
    db *stockdb.StockHistDataDB
    logger *util.StockLog
}

func (m *StockHistDataManager) Init() {
    m.downloader = download.NewStockHistDownloader()
    m.listdb = stsummary.NewStockListDB("chinastock", "stocklist")
    m.newlistdb = stsummary.NewStockListDB("chinastock", "newstocklist")
    m.db = stockdb.NewStockHistDataDB("chinastock")
    m.logger = util.NewLog()
}

func (m *StockHistDataManager) Process() {
    stocks := m.newlistdb.GetStocks()
    //ids := m.stocklistdb.QueryIds()
    //ids := []string{"000002"}
    for _, stock := range stocks {
        id := stock.Id
        if util.IsStringNotEmpty(id) {
            data := m.GetStockData(id)
            m.logger.Info("Complete to handle:", id)
            //m.WriteData(id, data)
            //fmt.Println(data)
            res := m.db.TranInsert(id, data)
            if res == 0 {
                m.listdb.Insert(stock)
                m.newlistdb.Delete(stock)
                //stock := m.stocklistdb.Query(id)
                //insert into stocklist
                //m.stocklistdb.Delete(stock)
                m.logger.Info("Delete the id from new stocklist after data updating: ", id) 
            }
        }
    }
}

func (m *StockHistDataManager) GetStockData(code string) []entity.StockHistData {
    mainPage := m.downloader.GetMainPage(code)
    mhandler := handler.NewStockHistDataHandler()
    mparser := parser.NewTextParser(mhandler)
    mparser.ParseStr(mainPage)
    //m.WriteMainPage(code, mainPage)
    
    mainData := mhandler.Data
    if len(mainData) == 0 {
        m.log(code, -1, -1)
    }
    now := time.Now()
    nowYear := now.Year()
    nowMonth := int(now.Month())
    maxSeason := 1 + int((nowMonth - 1) / 3)
    for _, year := range mhandler.Years {
        if year == nowYear {
            for i := 1; i < maxSeason; i ++ {
                seasonPage := m.downloader.GetSeasonPage(code, year, i)
                //m.WriteSeasonPage(code, seasonPage, year, i)
                shandler := handler.NewStockHistDataHandler()
                sparser := parser.NewTextParser(shandler)
                sparser.ParseStr(seasonPage)
                //m.WriteSeasonData(code, year, i, shandler.Data)
                if len(shandler.Data) == 0 {
                    m.log(code, year, i)
                }
                mainData = append(mainData, shandler.Data...)
            }
        } else if year > 0 {
            for i := 1; i <= 4; i ++ {
                seasonPage := m.downloader.GetSeasonPage(code, year, i)
                //m.WriteSeasonPage(code, seasonPage, year, i)
                shandler := handler.NewStockHistDataHandler()
                sparser := parser.NewTextParser(shandler)
                sparser.ParseStr(seasonPage)
                //m.WriteSeasonData(code, year, i, shandler.Data)
                if len(shandler.Data) == 0 {
                    m.log(code, year, i)
                }
                mainData = append(mainData, shandler.Data...)
            }
        }
    }

    return mainData
}

func (m *StockHistDataManager) log(code string, year, season int) {
    m.logger.Error("[historical-data]: miss: ", code, year, season)
} 

func (m *StockHistDataManager) WriteSeasonPage(code, content string, year, season int) {
    format := "../data/%v/%v-%v-%v.dat"
    filename := fmt.Sprintf(format, code, code, year, season)
    util.WriteFile(filename, content)
}

func (m *StockHistDataManager) WriteMainPage(code, content string) {
    format := "../data/%v/%v.dat"
    filename := fmt.Sprintf(format, code, code)
    util.WriteFile(filename, content)
}

func (m *StockHistDataManager) WriteSeasonData(code string, year int, season int, data []entity.StockHistData) {
    format := "../data/%v/%v-%v-%v-data.dat"
    filename := fmt.Sprintf(format, code, code, year, season)
    dataStr := fmt.Sprintf("%#v", data)
    util.WriteFile(filename, dataStr)
}

func (m *StockHistDataManager) WriteData(code string, data []entity.StockHistData) {
    format := "../data/%v/%v-data.dat"
    filename := fmt.Sprintf(format, code, code)
    dataStr := fmt.Sprintf("%#v", data)
    util.WriteFile(filename, dataStr)
}

func NewStockHistDataManager() *StockHistDataManager{
    m := new(StockHistDataManager)
    m.Init()

    return m
}
