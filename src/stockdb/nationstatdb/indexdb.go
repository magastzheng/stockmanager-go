package nationstatdb

import(
    "stockdb"
    ns "entity/nsentity"
    //"util"
    "fmt"
)

const(

)

type IndexDB struct {
    stockdb.DBBase
    dbtable string
}

func (s *IndexDB) getSql(sql string) string {
    return fmt.Sprintf(sql, s.dbtable)
}

func (s *IndexDB) Insert(idx ns.NSDataIndex) int {
    return 0
}

func NewIndexDB(dbname, dbtable string) *IndexDB {
    db := new(IndexDB)
    db.Init(dbname)
    db.dbtable = dbtable

    return db
}
