package accmanager

import(
    "stockdb"
    "stockdb/accountdb"
    "dbcreator/accgenerator"
    //"parser"
    "excel/account"
    //"handler/acchandler"
    "download/accdownload"
    acc "entity/accountentity"
    //"entity/dbentity"
    "util"
    //"time"
    "fmt"
)

const(
    Balance = "balancesheet"
    Income = "incomestatement"
    Cashflow = "cashflowstatement"
)

type AccountManager struct{
    AccManagerBase
    ep *account.AccountColumnParser
    dp *account.AccountParser
    db *accountdb.FinancialIndexDB
    listdb *stockdb.StockListDB
    down *accdownload.AccountDownloader

    sheets []string
}

func (m *AccountManager) Init() {
	m.AccManagerBase.Init()
    m.ep = account.NewAccountColumnParser()
    m.dp = account.NewAccountParser()
    m.db = accountdb.NewFinancialIndexDB("chinastock")

    m.listdb = stockdb.NewStockListDB("chinastock")
    m.down = accdownload.NewAccountDownloader()

    filename := m.baseDir + "resource/account/accountdb.xlsx"

    m.ep.Parse(filename)

    m.sheets = make([]string, 0)
    m.sheets = append(m.sheets, Balance, Income, Cashflow)
}

func (m *AccountManager) Process() {
    ids := m.listdb.QueryIds()
    if len(ids) == 0 {
        m.logger.Error("Cannot get stocklist from database")
    } else {
        m.logger.Info("Get the stocklist from database: ", len(ids)) 
    }
    ids = ids[1:2]
    
    for _, sheet := range m.sheets {
        m.ProcessSheet(sheet, ids, m.ep.ColumnMap)
    }
}

func (m *AccountManager) ProcessSheet(sheet string, ids []string, columnMap map[string]*acc.Column) {
    categoryMap := m.ep.GetSheetMap(sheet)
    tables := accgenerator.ConvertToDBTable(categoryMap)
    tableSqlMap := m.GetTableSql(tables)
    
    for _, id := range ids {
        var data string 
        switch sheet{
            case Balance:
                data = m.down.GetBalanceData(id)
            case Income:
                data = m.down.GetIncomeData(id)
            case Cashflow:
                data = m.down.GetCashFlowData(id)
        }
        

        if len(data) == 0{
            m.logger.Error("Cannot get data of: ", sheet, " code: ", id)
            continue
        }
        
        m.WriteFile(id, sheet, data)
        dateDataMap := m.dp.Parse(data)
        m.Output(dateDataMap)
        
        tabdbData := m.GetTableData(dateDataMap, id, tables, columnMap)
        for name, sql := range tableSqlMap{
            tabData, ok := tabdbData[name]
            if ok {
                m.db.Exec(sql, tabData)
            } else {
                m.logger.Error("There is no data to insert:", name)
            }
        }
    }   
}

func (m *AccountManager) ClearDB() {
    
    for _, sheet := range m.sheets {
        categoryMap := m.ep.GetSheetMap(sheet)
        tables := accgenerator.ConvertToDBTable(categoryMap)

        tabNames := make([]string, 0)
        for _, table := range tables {
            tabNames = append(tabNames, table.TableName)
        }

        m.db.Clear(tabNames)
    }
}

func (m *AccountManager)WriteFile(code, stype, content string) string {
    filename := fmt.Sprintf("%sdata/account/%s/%s-%s.dat", m.baseDir, code, stype, code)
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

