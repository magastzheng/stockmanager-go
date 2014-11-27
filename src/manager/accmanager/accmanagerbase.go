package accmanager

import(
    "entity/dbentity"
    acc "entity/accountentity"
    "stockdb"
    "math"
    "util"
    "runtime"
    "path/filepath"
    "fmt"
)

type AccManagerBase struct{
    generator *stockdb.SqlGenerator
    logger *util.StockLog
    baseDir string
}

func (m *AccManagerBase) Init() {
    m.generator = stockdb.NewSqlGenerator()
    m.logger = util.NewLog()
    pc, filename, line, ok := runtime.Caller(0)
    if pc < 0 || line < 0 || !ok {
        fmt.Println("Cannot read the fimanager.go")
        util.NewLog().Error("Cannot read the file fimanager.go")
    }

    m.baseDir = filepath.Dir(filename) + "/../../"
}

func (m *AccManagerBase) GetTableSql(tables []*dbentity.DBTable) map[string]string{
    tabSql := make(map[string]string)
    for _, table := range tables {
        cols := make([]string, 0)
        for _, col := range table.Columns{
            colName := col.Name
            cols = append(cols, colName)
        }

        sql := m.generator.GenerateInsert(table.TableName, cols)
        tabSql[table.TableName] = sql
    }

    return tabSql
}

func (m *AccManagerBase) GetTableData(datedatamap map[string]map[string]float32, code string, tables []*dbentity.DBTable, columnMap map[string]*acc.Column) map[string]dbentity.DBExecData {
    //there may be duplicated column: two Chinese index map to the same db colum
    colIdNameMap := make(map[string][]string)
    for k, col := range columnMap{
		cols, ok := colIdNameMap[col.Column]
		if ok {
			colIdNameMap[col.Column] = append(colIdNameMap[col.Column], k)
		} else {
			cols = make([]string, 0)
			cols = append(cols, k)
			colIdNameMap[col.Column] = cols
		}
    }
    
    tabsData := make(map[string]dbentity.DBExecData)
    //get data by date
    for date, dataMap := range datedatamap {
        //get data of each table
        for _, table := range tables {
            row := make([]interface{}, 0)
            for _, col := range table.Columns{
                colName := col.Name
                if colName == "date"{
                    row = append(row, date)
                } else if colName == "code" {
                    row = append(row, code)
                } else {
                    nmarr, ok := colIdNameMap[colName]
                    if ok {
                        var val float32
                        hasVal := false
                        for _, nm := range nmarr{
                            val, ok = dataMap[nm]
                            if ok {
                                hasVal = true
                                row = append(row, val)
                                break
                            }
                        }
                        if !hasVal {
                            row = append(row, math.NaN())
                        }
                    } else {
                        m.logger.Error("Cannot find the column: ", colName, " while inserting table: ", table.TableName)
						//skip of the case because it will fail to insert db even if it continue
						break
                    }
                }
            }
            
            tabData, ok := tabsData[table.TableName]
            if ok {
                tabData.Rows = append(tabData.Rows, row)
                tabsData[table.TableName] = tabData
            } else {
                tabData = dbentity.DBExecData{
                    Rows: make([][]interface{}, 0),
                }
                tabData.Rows = append(tabData.Rows, row)
                tabsData[table.TableName] = tabData
            }
        } 
    }

    return tabsData
}
