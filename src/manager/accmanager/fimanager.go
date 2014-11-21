package "accmanager"

import(
    "stockdb"
    "dbcreator"
    "parser"
    "excel"
    "handler/acchandler"
    "download/accdownload"
    acc "entity/accountentity"
    "entity/dbentity"
    "util"
)

type FiManager struct {
    ep *excel.AccountColumnParser
    //dp *parser.TextParser
    //dh *acchandler.FiHandler
    db *stockdb.AccountFinancialIndexDB
    listdb *stockdb.StockListDB
    down *accdownload.FiDownloader
    logger *util.StockLog
}

func (m *FiManager) Init() {
    m.ep = excel.NewAccountColumnParser()
    m.db = stockdb.NewAccountFinancialIndexDB()
    m.listdb = stockdb.NewStockListDB()
    m.down = accdownload.NewFiDownloader()
    m.logger = util.NewLog()

    m.ep.Parse("../../resource/account/financialindexdb.xlsx")
}

func (m *FiManager) Process() {
    
    dbTabMap := dbcreator.ConvertToDBColumn(m.ep.ColumnTableMap)
    columnMap := m.ep.ColumnMap

    ids := m.listdb.QueryIds()
    for _, id := range ids {
        data := m.down.GetData(id)
        
        if len(data) == 0{
            m.logger.Error("Cannot get data of: ", id)
            continue
        }

        dh := acchandler.NewFiHandler()
        dp := parser.NewTextParser(dh)
        dp.ParseStr(data)


    }   
}

func (m *FiManager) Insert(dh *acchandler.FiHandler, code string, tables []*dbentity.DBTable, columnMap map[string]*acc.Column) {
    datedatamap := dh.DataMap
    //insert data by date
    for date, dataMap := range datedatamap {
        //insert to each data
        for _, table := range tables {
            cols := make([]string, 0)
            //row := make
            for _, col := range table.Columns{
                cols = append(cols, col.Name)

            }
        } 
    } 
}

func NewFiManager() *FiManager{
    m :=  new(AccFiManager)
    m.Init()

    return m
}
