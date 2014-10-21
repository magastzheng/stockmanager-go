package stockdb

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "stockhandler"
    "strconv"
    "fmt"
)

type StockDatabase struct {
   dbtype string
   dbcon string
}

func (s *StockDatabase) CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func (s *StockDatabase) InsertStock(exch string, stock stockhandler.Stock) int {
    db, err := sql.Open(s.dbtype, s.dbcon)
    s.CheckError(err)

    stmt, err := db.Prepare("insert stocklist set id=?, name=?, exchange=?, website=?")
    s.CheckError(err)
    
    id, err := strconv.Atoi(stock.Id)
    s.CheckError(err)

    res, err := stmt.Exec(id, stock.Name, exch, stock.Website)
    s.CheckError(err)

    newid, err := res.LastInsertId()
    s.CheckError(err)
    fmt.Println(newid)
    
    db.Close()
    return 0
}

func (s *StockDatabase) DeleteStock(stock stockhandler.Stock) int {
    db, err := sql.Open(s.dbtype, s.dbcon)
    s.CheckError(err)

    stmt, err := db.Prepare("delete from stocklist where id=?")
    s.CheckError(err)
    
    id, err := strconv.Atoi(stock.Id)
    s.CheckError(err)
    
    res, err := stmt.Exec(id)
    s.CheckError(err)
    
    affect, err := res.RowsAffected()
    s.CheckError(err)

    fmt.Println(affect)

    db.Close()
    return 0
}

func (s *StockDatabase) UpdateStock(stock stockhandler.Stock) int {
    db, err := sql.Open(s.dbtype, s.dbcon)
    s.CheckError(err)

    stmt, err := db.Prepare("update stocklist set name=? where id=?")
    s.CheckError(err)
    
    id, err := strconv.Atoi(stock.Id)
    s.CheckError(err)
    
    res, err := stmt.Exec(stock.Name, id)
    s.CheckError(err)
    
    affect, err := res.RowsAffected()
    s.CheckError(err)

    fmt.Println(affect)

    db.Close()
    return 0
}

func NewStockDatabase(dbtype string, dbcon string) *StockDatabase {
    return &StockDatabase{
        dbtype: dbtype,
        dbcon: dbcon,
    }
}
