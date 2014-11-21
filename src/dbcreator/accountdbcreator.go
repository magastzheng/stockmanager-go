package dbcreator

import(
    "stockdb"
    "excel"
    acc "entity/accountentity"
    "entity/dbentity"
    "util"
    "fmt"
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
    dbTabs := m.ConvertDB(m.parser.ColumnTableMap)
    
    sqls := make([]string, 0)
    for _, dbTab := range dbTabs {
        sql := m.generator.GenerateCreate(*dbTab)
        fmt.Println(sql)
        sqls = append(sqls, sql)
    }

    m.db.Create(sqls)
}

func (m *AccountDBCreator) ConvertDB(tabMap map[string][]*acc.Column) []*dbentity.DBTable {
    comDBCols := make([]*dbentity.DBColumn, 0)
    pcols, ok := tabMap[acc.Common]
    if ok {
        for _, col := range pcols {
            dbcol := dbentity.DBColumn{
                Name: col.Column,
                Type: col.Type,
                Maxsize: col.Maxsize,
                IsNotNull: true,
            }

            comDBCols = append(comDBCols, &dbcol)
        }
    }
    
    dbTabs := make([]*dbentity.DBTable, 0)
    for k, cols := range tabMap{
        if k == acc.Common {
            continue
        }

        dbTab := dbentity.DBTable{
            TableName: k,
        }
        dbTab.Columns = make([]*dbentity.DBColumn, 0)
        dbTab.Keys = make([]*dbentity.DBColumn, 0)
        dbTab.Columns = append(dbTab.Columns, comDBCols ... )

        for _, col := range cols {
            dbcol := dbentity.DBColumn{
                Name: col.Column,
                Type: col.Type,
                Maxsize: col.Maxsize,
            }

            dbTab.Columns = append(dbTab.Columns, &dbcol)
        }

        dbTabs = append(dbTabs, &dbTab)
    }

    return dbTabs
}

func NewAccountDBCreator() *AccountDBCreator {
    m := new(AccountDBCreator)
    m.Init()

    return m
}
