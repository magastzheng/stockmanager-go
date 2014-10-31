package stockdb

import (
    "database/sql"
    "util"
    _ "github.com/go-sql-driver/mysql"
)

type DBBase struct {
    Dbtype string
    Dbcon string
}

func (s *DBBase) Open() *sql.DB {
    db, err := sql.Open(s.Dbtype, s.Dbcon)
    util.CheckError(err)
    
    return db
}
