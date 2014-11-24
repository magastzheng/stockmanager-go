package accmanager

import(
    "stockdb"
    "dbcreator"
    "parser"
    "excel"
    "handler/acchandler"
    "download"
    "download/accdownload"
    acc "entity/accountentity"
    "entity/dbentity"
    "util"
    "math"
    //"github.com/axgle/mahonia"
    "code.google.com/p/mahonia"
    "time"
    "fmt"
    "runtime"
    "path/filepath"
)

type FiManager struct {
    ep *excel.AccountColumnParser
    //dp *parser.TextParser
    //dh *acchandler.FiHandler
    db *stockdb.AccountFinancialIndexDB
    listdb *stockdb.StockListDB
    down *accdownload.FiDownloader
    generator *stockdb.SqlGenerator
    decoder mahonia.Decoder
    logger *util.StockLog
}

func (m *FiManager) Init() {
    m.ep = excel.NewAccountColumnParser()
    m.db = new(stockdb.AccountFinancialIndexDB)

    m.db.Init("chinastock")
    m.listdb = new(stockdb.StockListDB)
    m.listdb.Init("chinastock")
    m.down = accdownload.NewFiDownloader()
    m.generator = stockdb.NewSqlGenerator()
    m.decoder = mahonia.NewDecoder("gbk")
    m.logger = util.NewLog()
    
    pc, filename, line, ok := runtime.Caller(0)
    if pc < 0 || line < 0 || !ok {
        fmt.Println("Cannot read the fimanager.json")
        util.NewLog().Error("Cannot read the file fimanager.go")
    }
    filename = filepath.Dir(filename) + "/../../" + "resource/account/financialindexdb.xlsx"

    //m.ep.Parse("../../resource/account/financialindexdb.xlsx")
    //filename = baseDir + "/" + 
    m.ep.Parse(filename)
}

func (m *FiManager) Process() {
    
    dbTabMap := dbcreator.ConvertToDBColumn(m.ep.ColumnTableMap)
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
        data = m.decoder.ConvertString(data)

        if len(data) == 0{
            m.logger.Error("Cannot get data of: ", id)
            continue
        }
        
        m.WriteFile(id, year, data)
        dh := acchandler.NewFiHandler()
        dp := parser.NewTextParser(dh)
        dp.ParseStr(data)
        m.Insert(dh.DataMap, id, dbTabMap, columnMap)
        
        //handle historical financial index data
        m.ProcessHist(year, id, dh.DateMap, dbTabMap, columnMap)
    }   
}

func (m *FiManager) Insert(datedatamap map[string]map[string]float32, code string, tables []*dbentity.DBTable, columnMap map[string]*acc.Column) {
    //datedatamap := dh.DataMap
    //m.OutputDataMap(datedatamap)
    colIdNameMap := make(map[string]string)
    for k, col := range columnMap{
        colIdNameMap[col.Column] = k
    }
    
    //insert data by date
    for date, dataMap := range datedatamap {
        //insert to each data
        for _, table := range tables {
            cols := make([]string, 0)
            tabColNames := make([]string, 0)
            dbdata := dbentity.DBExecData{
                Rows: make([][]interface{}, 0),
            }
            row := make([]interface{}, 0)
            for _, col := range table.Columns{
                colName := col.Name
                cols = append(cols, colName)
                nm, ok := colIdNameMap[colName]
                if ok {
                    tabColNames = append(tabColNames, nm)
                } else {
                    tabColNames = append(tabColNames, colName)
                    m.logger.Error("Cannot find the column: ", colName, " while inserting table: ", table.TableName) 
                }
                
                if colName == "date"{
                    row = append(row, date)
                } else if colName == "code" {
                    row = append(row, code)
                } else {
                    val, ok := dataMap[nm]
                    if ok {
                        //fmt.Println("**********", col, val)
                        row = append(row, val)
                    } else {
                        row = append(row, math.NaN())
                    }
                }
            }
            
            sql := m.generator.GenerateInsert(table.TableName, cols)
            dbdata.Rows = append(dbdata.Rows, row)
            m.db.Exec(sql, dbdata)
        } 
    } 
}

func (m *FiManager) ProcessHist(currentYear, code string, dateUrlMap map[string]string, tables []*dbentity.DBTable, columnMap map[string]*acc.Column){
    for year, url := range dateUrlMap{
        if year != currentYear {
            data := download.HttpGet(url)
            data = m.decoder.ConvertString(data)
            
            if len(data) == 0{
                s := fmt.Sprintf("Cannot get data of: %s | year: %s | url: %s", code, year, url)
                m.logger.Error(s)
                continue
            }

            m.WriteFile(code, year, data)
            dh := acchandler.NewFiHandler()
            dp := parser.NewTextParser(dh)
            dp.ParseStr(data)
            m.Insert(dh.DataMap, code, tables, columnMap)
        }
    }
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
    //filename := fmt.Sprintf("../../data/fi/%s/%s-%s.dat", code, code, date)
    filename := fmt.Sprintf("../data/fi/%s/%s-%s.dat", code, code, date)
    util.WriteFile(filename, content)

    return filename
}

func NewFiManager() *FiManager{
    m := new(FiManager)
    m.Init()

    return m
}
