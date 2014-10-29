package stockdb

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "stockhandler"
    //"strconv"
    //"fmt"
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
    
    //id, err := strconv.Atoi(stock.Id)
    //s.CheckError(err)

    res, err := stmt.Exec(stock.Id, stock.Name, exch, stock.Website)
    s.CheckError(err)

    _, reserr := res.LastInsertId()
    s.CheckError(reserr)
    //fmt.Println(newid)
    
    db.Close()
    return 0
}

func (s *StockDatabase) DeleteStock(stock stockhandler.Stock) int {
    db, err := sql.Open(s.dbtype, s.dbcon)
    s.CheckError(err)

    stmt, err := db.Prepare("delete from stocklist where id=?")
    s.CheckError(err)
    
    //id, err := strconv.Atoi(stock.Id)
    //s.CheckError(err)
    
    res, err := stmt.Exec(stock.Id)
    s.CheckError(err)
    
    _, reserr := res.RowsAffected()
    s.CheckError(reserr)

    //fmt.Println(affect)

    db.Close()
    return 0
}

func (s *StockDatabase) UpdateStock(stock stockhandler.Stock) int {
    db, err := sql.Open(s.dbtype, s.dbcon)
    s.CheckError(err)

    stmt, err := db.Prepare("update stocklist set name=? where id=?")
    s.CheckError(err)
    
    //id, err := strconv.Atoi(stock.Id)
    //s.CheckError(err)
    
    res, err := stmt.Exec(stock.Name, stock.Id)
    s.CheckError(err)
    
    _, reserr := res.RowsAffected()
    s.CheckError(reserr)

    //fmt.Println(affect)

    db.Close()
    return 0
}

func (s *StockDatabase) TranInsertStock(exch string, stocks map[string] stockhandler.Stock) int {
    db, err := sql.Open(s.dbtype, s.dbcon)
    s.CheckError(err)
	
	tx, err := db.Begin()
	s.CheckError(err)
	
	for key, stock := range stocks {
		stmt, err := tx.Prepare("insert stocklist set id=?, name=?, exchange=?, website=?")
		s.CheckError(err)
		
		//id, err := strconv.Atoi(key)
		//s.CheckError(err)

		_, reserr := stmt.Exec(key, stock.Name, exch, stock.Website)
		s.CheckError(reserr)
		//fmt.Println(res)
	}
	
	err = tx.Commit()
	s.CheckError(err)
	
    db.Close()
    return 0
}

func (s *StockDatabase) QueryStock(id string) stockhandler.Stock {
    db, err := sql.Open(s.dbtype, s.dbcon)
    s.CheckError(err)
    defer db.Close()

    stmt, err := db.Prepare("select id, name, exchange, website from stocklist where id = ?")
    s.CheckError(err)
    defer stmt.Close()
    
    var stockid, stockname, exchange, website string
    err = stmt.QueryRow(id).Scan(&stockid, &stockname, &exchange, &website)
    s.CheckError(err)

    return stockhandler.Stock{
        Id: stockid,
        Name: stockname,
        Website: website,
    }
}

func NewStockDatabase(dbtype string, dbcon string) *StockDatabase {
    return &StockDatabase{
        dbtype: dbtype,
        dbcon: dbcon,
    }
}
