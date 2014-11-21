package dbcreator

import(
    "stockdb"
    "excel"
    "util"
    //"fmt"
)

type AccountDBCreator struct {
    parser *excel.AccountColumnParser
    generator *stockdb.SqlGenerator
    db *stockdb.AccountFinancialIndexDB
    logger *util.StockLog
}

func (m *AccountDBCreator) Init(){
    m.parser = excel.NewAccountColumnParser()
    m.generator = stockdb.NewSqlGenerator()
    m.db = stockdb.NewAccountFinancialIndexDB("chinastock")
    m.logger = util.NewLog()
}

func (m *AccountDBCreator) Process() {
    m.parser.Parse("../resource/account/financialindexdb.xlsx")
    dbTabs := ConvertToDBColumn(m.parser.ColumnTableMap)
    
    sqls := make([]string, 0)
    for _, dbTab := range dbTabs {
        sql := m.generator.GenerateCreate(*dbTab)
        //fmt.Println(sql)
        sqls = append(sqls, sql)
    }

    m.db.Create(sqls)
}

func NewAccountDBCreator() *AccountDBCreator {
    m := new(AccountDBCreator)
    m.Init()

    return m
}
