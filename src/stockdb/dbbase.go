package stockdb

import (
    "database/sql"
    "util"
    _ "github.com/go-sql-driver/mysql"
    "config"
    "entity/dbentity"
    //"fmt"
)

type DBBase struct {
    Dbtype string
    Dbcon string
    logger *util.StockLog
}

func (s *DBBase) Init(name string) {
    //dbconfig := config.NewDBConfig("../config/dbconfig.json")
    dbconfig := config.NewDBConfig()
    config := dbconfig.GetConfig(name)
    s.Dbtype = config.Dbtype
    s.Dbcon = config.Dbcon

    s.logger = util.NewLog()
}

//func (s *DBBase) InitDB(filename, name string) {
//    dbconfig := config.NewDBConfig(filename)
//    config := dbconfig.GetConfig(name)
//    s.Dbtype = config.Dbtype
//    s.Dbcon = config.Dbcon

//    s.logger = util.NewLog()
//}

func (s *DBBase) Open() *sql.DB {
    db, err := sql.Open(s.Dbtype, s.Dbcon)
    if err != nil {
        s.logger.Error("Cannot open database: ", s.Dbtype, s.Dbcon, err)
    }
    
    return db
}

func (s *DBBase)ExecOnce(query string, args ... interface{}) int {
    db := s.Open()
    defer db.Close()
    
    s.logger.Info(query, args)

    stmt, err := db.Prepare(query)
    defer stmt.Close()
    if err != nil{
        s.logger.Error("Cannot prepare the sql:", query, s.Dbtype, s.Dbcon)
        return -1
    }
    
    res, err := stmt.Exec(args ...)
    if err != nil {
        s.logger.Error("Fail to execute the sql:", query, s.Dbtype, s.Dbcon, err)
        return -1
    }

    _, reserr := res.RowsAffected()
    if reserr != nil{
        s.logger.Error("Fail to get the affected row:", query, s.Dbtype, s.Dbcon, err)
        return -1
    }

    return 0
}

func (s *DBBase)Exec(query string, data dbentity.DBExecData) int {
    db := s.Open()
    defer db.Close()
    
    tx, err := db.Begin()
    if err != nil {
        s.logger.Error("Fail to begin the transaction", s.Dbtype, s.Dbcon)
        return -1
    }
    for i, row := range data.Rows {
        stmt, err := tx.Prepare(query)
        defer stmt.Close()
        if err != nil{
            s.logger.Error("Fail to prepare the sql:", query, i, s.Dbtype, s.Dbcon, err)
        }
        
        //fmt.Println(row)
        _, reserr := stmt.Exec(row ... )
        if reserr != nil {
            s.logger.Error("Fail to execute the row in:", query, i, s.Dbtype, s.Dbcon, row, reserr)
        }
    }
    
    err = tx.Commit()
    if err != nil {
        s.logger.Error("Cannot commit the transaction: ", query, s.Dbtype, s.Dbcon)
        return -1
    }

    return 0
}   

func (s *DBBase)Query(query string, args ... interface{}) dbentity.DBData {
    db := s.Open()
    defer db.Close()
    
    data := dbentity.DBData{}
    var rows *sql.Rows
	var err error
    
    rows, err = db.Query(query, args ...)
    defer rows.Close()
    
    if err != nil {
        s.logger.Error("Cannot execute the sql:", query, s.Dbtype, s.Dbcon, err)
        return data
    }
    
    columns, err := rows.Columns()
    if err != nil {
        s.logger.Error("Cannot get the columns of rows:", s.Dbtype, s.Dbcon, err)
        return data
    }
    
    clen := len(columns)
    values := make([]sql.RawBytes, clen)
    scanArgs := make([]interface{}, clen)

    for i := range values {
        scanArgs[i] = &values[i]
    }
    
    data.Columns = columns
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
