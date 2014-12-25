package accgenerator

import(
    "stockdb"
    "stockdb/accountdb"
    "excel/account"
    "util"
	"entity/dbentity"
    "fmt"
)

type AccDBCreatorBase struct {
    parser *account.AccountColumnParser
    generator *stockdb.SqlGenerator
    db *accountdb.FinancialIndexDB
    logger *util.StockLog
}

func (m *AccDBCreatorBase) Init(){
    m.parser = account.NewAccountColumnParser()
    m.generator = stockdb.NewSqlGenerator()
    m.db = accountdb.NewFinancialIndexDB("chinastock")
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

func (m *AccDBCreatorBase) DropDB(tables []*dbentity.DBTable){
	sqls := make([]string, 0)
	for _, dbTab := range tables {
		sql := fmt.Sprintf("drop table %s", dbTab.TableName)
		sqls = append(sqls, sql)
	}
	
	m.db.Create(sqls)
}

