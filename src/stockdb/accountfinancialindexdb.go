package stockdb

//import(
//    "util"
//)

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

func NewAccountFinancialIndexDB(dbname string) *AccountFinancialIndexDB {
    db := new(AccountFinancialIndexDB)
    db.Init(dbname)

    return db
}
