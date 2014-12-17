package stsummary

import (
    //"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "stockdb"
    "entity"
    "util"
    //"strconv"
    "fmt"
)

const(
    ListInsert = "insert %s set id=?, name=?, exchange=?"
    ListDelete = "delete from %s where id=?"
    ListUpdate = "update %s set name=? where id=?"
    ListSelect = "select id, name, exchange from %s where id = ?"
    ListQueryCount = "select count(id) from %s"
    ListQueryId = "select id from %s"
    ListQueryIdExchange = "select id, exchange from %s"
)

type StockListDB struct {
   stockdb.DBBase
   dbtable string
}

func (s *StockListDB) getSql(sql string) string{
    return fmt.Sprintf(sql, s.dbtable)
}

func (s *StockListDB) Insert(stock entity.Stock) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(ListInsert)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }
    
    res, err := stmt.Exec(stock.Id, stock.Name, stock.Exchange)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    _, reserr := res.LastInsertId()
    if reserr != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, reserr)
        return -1
    }
    
    return 0
}

func (s *StockListDB) Delete(stock entity.Stock) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(ListDelete)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(stock.Id)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }
    
    _, reserr := res.RowsAffected()
    if reserr != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, reserr)
        return -1
    }

    return 0
}

func (s *StockListDB) Update(stock entity.Stock) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(ListUpdate)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(stock.Name, stock.Id)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }
    
    _, reserr := res.RowsAffected()
    if reserr != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, reserr)
        return -1
    }

    return 0
}

func (s *StockListDB) TranInsert(stocks []entity.Stock) int {
    db := s.Open()
    defer db.Close()
	
    sql := s.getSql(ListInsert)
	tx, err := db.Begin()
	if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }
	
	for _, stock := range stocks {
		stmt, err := tx.Prepare(sql)
        defer stmt.Close()
        if err != nil {
            s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
            continue
        }
        
		_, reserr := stmt.Exec(stock.Id, stock.Name, stock.Exchange)
        if reserr != nil {
            s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, reserr)
            return -1
        }
	}
	
	err = tx.Commit()
	if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }
	
    return 0
}

func (s *StockListDB) Query(id string) entity.Stock {
    stock := entity.Stock{}
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(ListSelect)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return stock
    }
    
    var stockid, stockname, exchange string
    err = stmt.QueryRow(id).Scan(&stockid, &stockname, &exchange)
    if err != nil{
        s.Logger.Error("Cannot query the stock with id: ", id, err)
    }
    
    stock.Id = stockid
    stock.Name = stockname
    stock.Exchange = exchange

    return stock
}

func (s *StockListDB) QueryIds() []string {
    db := s.Open()
    defer db.Close()
    
    ids := make([]string, 0)
    
    sql := s.getSql(ListQueryCount)
    rows, err := db.Query(sql)
    if err != nil{
        s.Logger.Error("Cannot query the stock count.", err)
        return ids
    }

    var count int
    for rows.Next() {
        err = rows.Scan(&count)
    }
    
    sql = s.getSql(ListQueryId)
    rows, err = db.Query(sql)
    if err != nil {
        s.Logger.Error("Cannot query the stock list id.", err)
        return ids
    }
    
    var id string
    for rows.Next() {
         err = rows.Scan(&id)
         util.CheckError(err)
         ids = append(ids, id)
    }
    
    return ids
}

func (s *StockListDB) GetIdExchange() []entity.Stock {
    idexchs := make([]entity.Stock, 0)

    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(ListQueryCount)
    rows, err := db.Query(sql)
    if err != nil {
        s.Logger.Error("Cannot get stock count.", err)
        return idexchs
    }

    var count int
    for rows.Next() {
        err = rows.Scan(&count)
    }
    
    sql = s.getSql(ListQueryIdExchange)
    rows, err = db.Query(sql)
    if err != nil {
        s.Logger.Error("Cannot query the stock list id.", err)
        return idexchs
    }
    

    var id, exch string
    for rows.Next() {
         err = rows.Scan(&id, &exch)
         util.CheckError(err)
         
         idexch := entity.Stock{
            Id: id,
            Exchange: exch,
         }

         idexchs = append(idexchs, idexch)
    }
    
    return idexchs
}

func NewStockListDB(dbname, dbtable string) *StockListDB {
    stdb := new(StockListDB)
    stdb.Init(dbname)
    stdb.dbtable = dbtable

    return stdb
}
