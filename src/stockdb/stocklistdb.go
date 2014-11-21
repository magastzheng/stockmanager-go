package stockdb

import (
    //"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "entity"
    "util"
    //"strconv"
    //"fmt"
)

type StockIdExchange struct {
    Id string
    Exchange string
}

type StockListDB struct {
   DBBase
}

func (s *StockListDB) Insert(exch string, stock entity.Stock) int {
    db := s.Open()

    stmt, err := db.Prepare("insert stocklist set id=?, name=?, exchange=?, website=?")
    util.CheckError(err)
    defer stmt.Close()
    
    //id, err := strconv.Atoi(stock.Id)
    //s.CheckError(err)

    res, err := stmt.Exec(stock.Id, stock.Name, exch, stock.Website)
    util.CheckError(err)

    _, reserr := res.LastInsertId()
    util.CheckError(reserr)
    //fmt.Println(newid)
    
    db.Close()
    return 0
}

func (s *StockListDB) Delete(stock entity.Stock) int {
    db := s.Open()

    stmt, err := db.Prepare("delete from stocklist where id=?")
    util.CheckError(err)
    defer stmt.Close()
    
    //id, err := strconv.Atoi(stock.Id)
    //s.CheckError(err)
    
    res, err := stmt.Exec(stock.Id)
    util.CheckError(err)
    
    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    //fmt.Println(affect)

    db.Close()
    return 0
}

func (s *StockListDB) Update(stock entity.Stock) int {
    db := s.Open()

    stmt, err := db.Prepare("update stocklist set name=? where id=?")
    util.CheckError(err)
    
    //id, err := strconv.Atoi(stock.Id)
    //s.CheckError(err)
    
    res, err := stmt.Exec(stock.Name, stock.Id)
    util.CheckError(err)
    
    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    //fmt.Println(affect)

    db.Close()
    return 0
}

func (s *StockListDB) TranInsert(exch string, stocks map[string] entity.Stock) int {
    db := s.Open()
	
	tx, err := db.Begin()
	util.CheckError(err)
	
	for key, stock := range stocks {
		stmt, err := tx.Prepare("insert stocklist set id=?, name=?, exchange=?, website=?")
		util.CheckError(err)
		
		//id, err := strconv.Atoi(key)
		//s.CheckError(err)

		_, reserr := stmt.Exec(key, stock.Name, exch, stock.Website)
		util.CheckError(reserr)
		//fmt.Println(res)
	}
	
	err = tx.Commit()
	util.CheckError(err)
	
    db.Close()
    return 0
}

func (s *StockListDB) Query(id string) entity.Stock {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("select id, name, exchange, website from stocklist where id = ?")
    util.CheckError(err)
    defer stmt.Close()
    
    var stockid, stockname, exchange, website string
    err = stmt.QueryRow(id).Scan(&stockid, &stockname, &exchange, &website)
    util.CheckError(err)

    return entity.Stock{
        Id: stockid,
        Name: stockname,
        Website: website,
    }
}

func (s *StockListDB) QueryIds() []string {
    db := s.Open()
    defer db.Close()
   
    //rows, err := db.Query("select count(id) from stocklist")
    rows, err := db.Query("select count(id) from stocklist where id not in (select distinct code from stockhistdata)")
    util.CheckError(err)

    var count int
    for rows.Next() {
        err = rows.Scan(&count)
    }
    //fmt.Println("Total:", count)

    //rows, err = db.Query("select id from stocklist")
    rows, err = db.Query("select id from stocklist where id not in (select distinct code from stockhistdata)")
    util.CheckError(err)
    
    //Get column names
    //columns, err := rows.Columns()
    //util.CheckError(err)
    //make a slice for the values
    //values := make([]sql.RawBytes, len(columns))
    
    //scanArgs := make([]interface{}, len(values))
    //for i := range values {
    //    scanArgs[i] = &values[i]
    //}
    
    ids := make([]string, 0, count + 1)

    //total := 0
    var id string
    for rows.Next() {
         //err = rows.Scan(scanArgs...)
         err = rows.Scan(&id)
         util.CheckError(err)
           
         //total++
         //var value string
         //for _, col := range values {
         //   if col == nil {
         //       value = "NULL"
         //   } else {
         //       value = string(col)
         //   }
         //   
            //fmt.Println(i, value)
         //   ids = append(ids, value)
         //}

         ids = append(ids, id)
    }
    
    //fmt.Println(total)
    return ids
}

func (s *StockListDB) GetIdExchange() []StockIdExchange {
    db := s.Open()
    defer db.Close()
   
    rows, err := db.Query("select count(id) from stocklist")
    util.CheckError(err)

    var count int
    for rows.Next() {
        err = rows.Scan(&count)
    }
    //fmt.Println("Total:", count)

    rows, err = db.Query("select id, exchange from stocklist")
    util.CheckError(err)
    
    //Get column names
    //columns, err := rows.Columns()
    //util.CheckError(err)
    //make a slice for the values
    //values := make([]sql.RawBytes, len(columns))
    
    //scanArgs := make([]interface{}, len(values))
    //for i := range values {
    //    scanArgs[i] = &values[i]
    //}
    
    idexchs := make([]StockIdExchange, 0, count + 1)

    //total := 0
    var id, exch string
    for rows.Next() {
         //err = rows.Scan(scanArgs...)
         err = rows.Scan(&id, &exch)
         util.CheckError(err)
         
         idexch := StockIdExchange{
            Id: id,
            Exchange: exch,
         }

         idexchs = append(idexchs, idexch)
    }
    
    return idexchs
}

func NewStockListDB(dbname string) *StockListDB {
    stdb := new(StockListDB)
    stdb.Init(dbname)
    return stdb
    //return &StockListDB{
    //    DBBase: DBBase{
    //        Dbtype: dbtype,
    //        Dbcon: dbcon,
    //    },
    //}
}
