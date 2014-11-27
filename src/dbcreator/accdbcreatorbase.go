package dbcreator

import(
    "stockdb"
    "excel"
    "util"
	"entity/dbentity"
    //"fmt"
)

type AccDBCreatorBase struct {
    parser *excel.AccountColumnParser
    generator *stockdb.SqlGenerator
    db *stockdb.AccountFinancialIndexDB
    logger *util.StockLog
}

func (m *AccDBCreatorBase) Init(){
    m.parser = excel.NewAccountColumnParser()
    m.generator = stockdb.NewSqlGenerator()
    m.db = stockdb.NewAccountFinancialIndexDB("chinastock")
    m.logger = util.NewLog()
}

func (m *AccDBCreatorBase) CreateDB(tables []*dbentity.DBTable) {
	sqls := make([]string, 0)
    for _, dbTab := range tables {
        sql := m.generator.GenerateCreate(*dbTab)
        //fmt.Println(sql)
        sqls = append(sqls, sql)
    }

    m.db.Create(sqls)
}

