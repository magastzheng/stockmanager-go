package accmanager

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
    "math"
    //"github.com/axgle/mahonia"
    "code.google.com/p/mahonia"
    "fmt"
)

type FiManager struct {
    ep *excel.AccountColumnParser
    //dp *parser.TextParser
    //dh *acchandler.FiHandler
    db *stockdb.AccountFinancialIndexDB
    listdb *stockdb.StockListDB
    down *accdownload.FiDownloader
    generator *stockdb.SqlGenerator
    logger *util.StockLog
}

func (m *FiManager) Init() {
    m.ep = excel.NewAccountColumnParser()
    m.db = new(stockdb.AccountFinancialIndexDB)
    m.db.InitDB("../../config/dbconfig.json", "chinastock")
    m.listdb = new(stockdb.StockListDB)
    m.listdb.InitDB("../../config/dbconfig.json", "chinastock")
    m.down = accdownload.NewFiDownloader()
    m.generator = stockdb.NewSqlGenerator()
    m.logger = util.NewLog()

    m.ep.Parse("../../resource/account/financialindexdb.xlsx")
}

func (m *FiManager) Process() {
    
    dbTabMap := dbcreator.ConvertToDBColumn(m.ep.ColumnTableMap)
    columnMap := m.ep.ColumnMap

    ids := m.listdb.QueryIds()
    //fmt.Println(ids)
    
    //enc := mahonia.NewEncoder("UTF-8")
    enc := mahonia.NewDecoder("gbk")
    ids = ids[1:2]
    for _, id := range ids {
        data := m.down.GetData(id)
        //filename := m.WriteFile(id, data)
        //data = util.ReadFile(filename)
        //fmt.Println(data)
        //fmt.Println("=========convert===========")
        data = enc.ConvertString(data)
        //fmt.Println(data)
        if len(data) == 0{
            m.logger.Error("Cannot get data of: ", id)
            continue
        }

        dh := acchandler.NewFiHandler()
        dp := parser.NewTextParser(dh)
        dp.ParseStr(data)
        m.Insert(dh, id, dbTabMap, columnMap)
    }   
}

func (m *FiManager) Insert(dh *acchandler.FiHandler, code string, tables []*dbentity.DBTable, columnMap map[string]*acc.Column) {
    datedatamap := dh.DataMap
    m.OutputDataMap(datedatamap)
    //fmt.Println(datedatamap)
    colIdNameMap := make(map[string]string)
    for k, col := range columnMap{
        //fmt.Println(col.Name, col.Column)
        colIdNameMap[col.Column] = k
    }
    fmt.Println(colIdNameMap)
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
                        fmt.Println("**********", col, val)
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

func (m *FiManager)WriteFile(code, content string) string {
    filename := fmt.Sprintf("../../data/%s.dat", code)
    util.WriteFile(filename, content)

    return filename
}

func NewFiManager() *FiManager{
    m := new(FiManager)
    m.Init()

    return m
}
