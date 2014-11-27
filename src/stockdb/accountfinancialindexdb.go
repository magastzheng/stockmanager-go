package stockdb

import(
//  "util"
    "fmt"
)

const (
    Truncate = "truncate table %s"
)
type AccountFinancialIndexDB struct{
    DBBase
}

func (s *AccountFinancialIndexDB) Create(sqls []string) int {
    //db := s.Open()
    
    for _, sql := range sqls {
        s.ExecOnce(sql)
    }

    return 0
}

func (s *AccountFinancialIndexDB) Clear(tables []string) int {
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

func NewAccountFinancialIndexDB(dbname string) *AccountFinancialIndexDB {
    db := new(AccountFinancialIndexDB)
    db.Init(dbname)

    return db
}
