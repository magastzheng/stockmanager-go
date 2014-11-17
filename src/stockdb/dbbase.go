package stockdb

import (
    "database/sql"
    "util"
    _ "github.com/go-sql-driver/mysql"
    "config"
)

type DBData struct {
    Columns []string
    Rows [][]string
}

type DBExecData struct {
    Rows [][]interface{}
}

type DBBase struct {
    Dbtype string
    Dbcon string
    logger *util.StockLog
}

func (s *DBBase) Init(name string) {
    dbconfig := config.NewDBConfig("../config/dbconfig.json")
    config := dbconfig.GetConfig(name)
    s.Dbtype = config.Dbtype
    s.Dbcon = config.Dbcon

    s.logger = util.NewLog()
}

func (s *DBBase) Open() *sql.DB {
    db, err := sql.Open(s.Dbtype, s.Dbcon)
    if err != nil {
        s.logger.Error("Cannot open database: ", s.Dbtype, s.Dbcon, err)
    }
    
    return db
}


func (s *DBBase)Insert(query string, data DBExecData) int {
    db := s.Open()
    defer db.Close()
    //rlen := len(data)
    
    //stmt, err := db.Prepare(query)
    //defer stmt.Close()
    //if err != nil{
    //    s.logger.Error("Cannot prepare the sql:", query, s.Dbtype, s.Dbcon)
    //}
    
    //res, err := stmt.Exec
    tx, err := db.Begin()
    if err != nil {
        s.logger.Error("Fail to begin the transaction", s.Dbtype, s.Dbcon)
    }
    for i, row := range data.Rows {
        stmt, err := tx.Prepare(query)
        defer stmt.Close()
        if err != nil{
            s.logger.Error("Fail to prepare the sql:", query, i, s.Dbtype, s.Dbcon)
        }

        _, reserr := stmt.Exec(row)
        if reserr != nil {
            s.logger.Error("Fail to execute the row in:", i, row)
        }
    }
    
    err = tx.Commit()
    if err != nil {
        s.logger.Error("Cannot commit the transaction", s.Dbtype, s.Dbcon)
    }

    return 0
}   

func (s *DBBase)Query(query string, args ... interface{}) DBData {
    db := s.Open()
    defer db.Close()
    
    var rows *sql.Rows
	var err error
    //if len(args) > 0 {
        rows, err = db.Query(query, args ...)
    //} else {
    //    rows, err = db.Query(query)
    //}
    defer rows.Close()
    
    if err != nil {
        s.logger.Error("Cannot execute the sql:", query, s.Dbtype, s.Dbcon, err)
    }
    
    columns, err := rows.Columns()
    if err != nil {
        s.logger.Error("Cannot get the columns of rows:", s.Dbtype, s.Dbcon, err)
    }
    
    clen := len(columns)
    values := make([]sql.RawBytes, clen)
    scanArgs := make([]interface{}, clen)

    for i := range values {
        scanArgs[i] = &values[i]
    }
    
    data := DBData{Columns: columns}
    data.Rows = make([][]string, 0)
    for rows.Next() {
        err = rows.Scan(scanArgs ...)
        if err != nil {
            s.logger.Error("Cannot read the rows:", s.Dbtype, s.Dbcon, err)
        }

        var value string
        rvalues := make([]string, clen)
        for _, col := range values {
            if col == nil {
                value = "NULL"
            } else {
                value = string(col)
            }

            rvalues = append(rvalues, value)
        }
        data.Rows = append(data.Rows, rvalues)
    }

    return data
} 
