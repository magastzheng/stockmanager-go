package stockdb

import (
    "database/sql"
    "util"
    _ "github.com/go-sql-driver/mysql"
    "config"
)

type DBBase struct {
    Dbtype string
    Dbcon string
}

func (s *DBBase) Init(name string) {
    dbconfig := config.NewDBConfig("../config/dbconfig.json")
    config := dbconfig.GetConfig(name)
    s.Dbtype = config.Dbtype
    s.Dbcon = config.Dbcon
}

func (s *DBBase) Open() *sql.DB {
    db, err := sql.Open(s.Dbtype, s.Dbcon)
    util.CheckError(err)
    
    return db
}
