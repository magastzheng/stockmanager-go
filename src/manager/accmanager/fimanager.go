package accmanager

import(
    "stockdb"
    "stockdb/accountdb"
    "dbcreator/accgenerator"
    "parser"
    "excel/account"
    "handler/acchandler"
    "download/accdownload"
    acc "entity/accountentity"
    "entity/dbentity"
    "util"
    "time"
    "fmt"
)

type FiManager struct {
	AccManagerBase
    ep *account.AccountColumnParser
    db *accountdb.FinancialIndexDB
    listdb *stockdb.StockListDB
    down *accdownload.FiDownloader
	
	categoryMap map[string][]*acc.Column
}

func (m *FiManager) Init() {
	m.AccManagerBase.Init()
    m.ep = account.NewAccountColumnParser()
    m.db = accountdb.NewFinancialIndexDB("chinastock")

    m.listdb = stockdb.NewStockListDB("chinastock")
    m.down = accdownload.NewFiDownloader()

    filename := m.baseDir + "resource/account/financialindexdb.xlsx"

    m.ep.Parse(filename)
	m.categoryMap = m.ep.GetSheetMap("findex")
}

func (m *FiManager) Process() {
    categoryMap := m.ep.GetSheetMap("findex")
    tables := accgenerator.ConvertToDBTable(categoryMap)
    columnMap := m.ep.ColumnMap
    
    now := time.Now()
    year := fmt.Sprintf("%d", now.Year())
    
    //enc := mahonia.NewEncoder("UTF-8")
    
    ids := m.listdb.QueryIds()
    if len(ids) == 0 {
        m.logger.Error("Cannot get stocklist from database")
    } else {
        m.logger.Info("Get the stocklist from database: ", len(ids)) 
    }
    ids = ids[1:2]
    for _, id := range ids {
        data := m.down.GetData(id)

        if len(data) == 0{
            m.logger.Error("Cannot get data of: ", id)
            continue
        }
        
        m.WriteFile(id, year, data)
        dh := acchandler.NewFiHandler()
        dp := parser.NewTextParser(dh)
        dp.ParseStr(data)

        //m.OutputDataMap(dh.DataMap)
        m.Insert(dh.DataMap, id, tables, columnMap)
        
        //handle historical financial index data
        m.ProcessHist(year, id, dh.DateMap, tables, columnMap)
    }   
}

func (m *FiManager) ProcessHist(currentYear, code string, dateUrlMap map[string]string, tables []*dbentity.DBTable, columnMap map[string]*acc.Column){
    for year, url := range dateUrlMap{
        if year != currentYear {
            data := m.down.GetHistData(url)
            
            if len(data) == 0{
                s := fmt.Sprintf("Cannot get data of: %s | year: %s | url: %s", code, year, url)
                m.logger.Error(s)
                continue
            }

            m.WriteFile(code, year, data)
            dh := acchandler.NewFiHandler()
            dp := parser.NewTextParser(dh)
            dp.ParseStr(data)

            //m.OutputDataMap(dh.DataMap)
            m.Insert(dh.DataMap, code, tables, columnMap)
        }
    }
}

func (m *FiManager) Insert(datedatamap map[string]map[string]float32, code string, tables []*dbentity.DBTable, columnMap map[string]*acc.Column) {
    tabSqlMap := m.GetTableSql(tables)
    tabsData := m.GetTableData(datedatamap, code, tables, columnMap)
    for name, sql := range tabSqlMap{
        dbdata, ok := tabsData[name]
        if ok {
            m.db.Exec(sql, dbdata)
        }
    }
}

func (m *FiManager) ClearDB() {
    tables := accgenerator.ConvertToDBTable(m.categoryMap)
    tabNames := make([]string, 0)
    for _, table := range tables {
        tabNames = append(tabNames, table.TableName)
    }

    m.db.Clear(tabNames)
}

func (m *FiManager)OutputDataMap(dataMap map[string]map[string]float32) {
    for date, dm := range dataMap {
        s := fmt.Sprintf("=============Date: %s ============", date)
        fmt.Println(s)
        for k, v := range dm {
            s = fmt.Sprintf("%v: %f", k, v)
            fmt.Println(s)
        }
    }
}

func (m *FiManager)WriteFile(code, date, content string) string {
    filename := fmt.Sprintf("%sdata/fi/%s/%s-%s.dat", m.baseDir, code, code, date)
    util.WriteFile(filename, content)

    return filename
}

func NewFiManager() *FiManager{
    m := new(FiManager)
    m.Init()

    return m
}
