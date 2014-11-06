package listmanager

import (
    "download"
    "stockhandler"
    "parser"
    "stockdb"
    "config"
    "time"
    "util"
    "fmt"
)

type StockHistDataManager struct {
    downloader *download.StockHistDownloader
    //parser *parser.TextParser
    //handler *stockhandler.StockHistDataHandler
    stocklistdb *stockdb.StockDatabase
    db *stockdb.StockHistDataDB
}

func (m *StockHistDataManager) Init() {
    m.downloader = download.NewStockHistDownloader()

    dbconfig := config.NewDBConfig("../config/dbconfig.json")
    config := dbconfig.GetConfig("chinastock")
    m.stocklistdb = stockdb.NewStockDatabase(config.Dbtype, config.Dbcon)
    m.db = stockdb.NewStockHistDataDB(config.Dbtype, config.Dbcon)
}

func (m *StockHistDataManager) Process() {
    ids := m.stocklistdb.QueryIds()
    //ids := []string{"000002"}
    for _, id := range ids {
        if util.IsStringNotEmpty(id) {
            data := m.GetStockData(id)
            //m.WriteData(id, data)
            //fmt.Println(data)
            m.db.TranInsert(id, data)
        }
    }
}

func (m *StockHistDataManager) GetStockData(code string) []stockhandler.StockHistData {
    mainPage := m.downloader.GetMainPage(code)
    handler := stockhandler.NewStockHistDataHandler()
    mparser := parser.NewTextParser(handler)
    mparser.ParseStr(mainPage)
    //m.WriteMainPage(code, mainPage)
    
    mainData := handler.Data
    now := time.Now()
    nowYear := now.Year()
    nowMonth := int(now.Month())
    maxSeason := 1 + int((nowMonth - 1) / 3)
    for _, year := range handler.Years {
        if year == nowYear {
            for i := 1; i < maxSeason; i ++ {
                seasonPage := m.downloader.GetSeasonPage(code, year, i)
                //m.WriteSeasonPage(code, seasonPage, year, i)
                shandler := stockhandler.NewStockHistDataHandler()
                sparser := parser.NewTextParser(shandler)
                sparser.ParseStr(seasonPage)
                //m.WriteSeasonData(code, year, i, shandler.Data)
                mainData = append(mainData, shandler.Data...)
            }
        } else if year > 0 {
            for i := 1; i <= 4; i ++ {
                seasonPage := m.downloader.GetSeasonPage(code, year, i)
                //m.WriteSeasonPage(code, seasonPage, year, i)
                shandler := stockhandler.NewStockHistDataHandler()
                sparser := parser.NewTextParser(shandler)
                sparser.ParseStr(seasonPage)
                //m.WriteSeasonData(code, year, i, shandler.Data)
                mainData = append(mainData, shandler.Data...)
            }
        }
    }

    return mainData
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

func (m *StockHistDataManager) WriteSeasonData(code string, year int, season int, data []stockhandler.StockHistData) {
    format := "../data/%v/%v-%v-%v-data.dat"
    filename := fmt.Sprintf(format, code, code, year, season)
    dataStr := fmt.Sprintf("%#v", data)
    util.WriteFile(filename, dataStr)
}

func (m *StockHistDataManager) WriteData(code string, data []stockhandler.StockHistData) {
    format := "../data/%v/%v-data.dat"
    filename := fmt.Sprintf(format, code, code)
    dataStr := fmt.Sprintf("%#v", data)
    util.WriteFile(filename, dataStr)
}

func NewStockHistDataManager() *StockHistDataManager{
    manager := new(StockHistDataManager)
    manager.Init()

    return manager
}
