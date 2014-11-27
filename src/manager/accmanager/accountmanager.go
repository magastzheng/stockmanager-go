package accmanager

import(
    "stockdb"
    //"dbcreator"
    //"parser"
    "excel"
    //"handler/acchandler"
    "download/accdownload"
    //acc "entity/accountentity"
    //"entity/dbentity"
    "util"
    //"time"
    "fmt"
)

type AccountManager struct{
    AccManagerBase
    ep *excel.AccountColumnParser
    dp *excel.AccountParser
    db *stockdb.AccountFinancialIndexDB
    listdb *stockdb.StockListDB
    down *accdownload.AccountDownloader
}

func (m *AccountManager) Init() {
	m.AccManagerBase.Init()
    m.ep = excel.NewAccountColumnParser()
    m.dp = excel.NewAccountParser()
    m.db = new(stockdb.AccountFinancialIndexDB)

    m.db.Init("chinastock")
    m.listdb = new(stockdb.StockListDB)
    m.listdb.Init("chinastock")
    m.down = accdownload.NewAccountDownloader()

    filename := m.baseDir + "resource/account/accountdb.xlsx"

    m.ep.Parse(filename)
}

func (m *AccountManager) Process() {
    tables := dbcreator.ConvertToDBTable(m.ep.CategoryColumnMap)
    columnMap := m.ep.ColumnMap
    tableSqlMap := m.GetTableSql(tables)

    ids := m.listdb.QueryIds()
    if len(ids) == 0 {
        m.logger.Error("Cannot get stocklist from database")
    } else {
        m.logger.Info("Get the stocklist from database: ", len(ids)) 
    }
    ids = ids[1:2]
    for _, id := range ids {
        data := m.down.GetBalanceData(id)

        if len(data) == 0{
            m.logger.Error("Cannot get data of: ", id)
            continue
        }
        
        m.WriteFile(id, "bs", data)
        dateDataMap := m.dp.Parse(data)
        m.Output(dateDataMap)
        
        dbData := m.GetTableData(dateDataMap, id, tables, columnMap)
    }   
}

func (m *AccountManager)WriteFile(code, type, content string) string {
    filename := fmt.Sprintf("%sdata/account/%s/%s-%s.dat", m.baseDir, code, type, code)
    util.WriteFile(filename, content)

    return filename
}

func (m *AccountManager)Output(dateDataMap map[string]map[string]float32){
	for date, datakeyval := range dateDataMap{
		fmt.Println("=======", date, "=========")
		for k, v := range datakeyval{
			fmt.Println(k, "\t:", v)
		}
	}
}

func NewAccountManager() *AccountManager{
    m := new(AccountManager)
    m.Init()

    return m
}

