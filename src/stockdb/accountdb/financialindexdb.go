package accountdb

import(
    "stockdb"
//  "util"
    "fmt"
)

const (
    Truncate = "truncate table %s"
)
type FinancialIndexDB struct{
    stockdb.DBBase
}

func (s *FinancialIndexDB) Create(sqls []string) int {
    //db := s.Open()
    
    for _, sql := range sqls {
        s.ExecOnce(sql)
    }

    return 0
}

func (s *FinancialIndexDB) Clear(tables []string) int {
    sqls := make([]string, 0)

    for _, table := range tables {
        sql := fmt.Sprintf(Truncate, table)
        sqls = append(sqls, sql)
    }

    for _, sql := range sqls {
        s.ExecOnce(sql)
    }

    return 0
}

func NewFinancialIndexDB(dbname string) *FinancialIndexDB {
    db := new(FinancialIndexDB)
    db.Init(dbname)

    return db
}
