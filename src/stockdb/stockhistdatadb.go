package stockdb

import(
    _ "github.com/go-sql-driver/mysql"
    "stockhandler"
    "util"
)

type StockHistDataDB struct {
    DBBase
}

func (s *StockHistDataDB) Insert(code string, d stockhandler.StockHistData) int {
    db := s.Open()
    stmt, err := db.Prepare("insert stockhistdata set code=?, date=?, open=?, close=?, highest=?, lowest=?, volume=?, money=?")
    util.CheckError(err)
    defer stmt.Close()

    res, err := stmt.Exec(code, d.Date, d.Open, d.Close, d.Highest, d.Lowest, d.Volume, d.Money)
    util.CheckError(err)
    
    _, reserr := res.LastInsertId()
    util.CheckError(reserr)
    
    db.Close()
    return 0
}

func (s *StockHistDataDB) Delete(code string, date string) int {
    db := s.Open()

    stmt, err := db.Prepare("delete from stockhistdata where code=? and date=?")
    util.CheckError(err)

    res, err := stmt.Exec(code, date)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    db.Close()
    return 0
}

func (s *StockHistDataDB) Update(code string, d stockhandler.StockHistData) int {
    db := s.Open()
    
    stmt, err := db.Prepare("update stockhistdata set open=?, close=?, highest=?, lowest=?, volume=?, money=? where code=? and date=?")
    util.CheckError(err)

    res, err := stmt.Exec(d.Date, d.Open, d.Close, d.Highest, d.Lowest, d.Volume, d.Money, code, d.Date)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    db.Close()
    return 0
}

func (s *StockHistDataDB) Query(code string, date string) stockhandler.StockHistData {
    db := s.Open()
    
    stmt, err := db.Prepare("select code, date, open, close, highest, lowest, volume, money from stockhistdata where code = ? and date = ?")
    util.CheckError(err)
    defer stmt.Close()

    var newcode, newdate string
    var open, close, highest, lowest float32
    var volume, money int
    err = stmt.QueryRow(code, date).Scan(&newcode, &newdate, &open, &close, &highest, &lowest, &volume, &money)
    util.CheckError(err)

    db.Close()

    return stockhandler.StockHistData{
        Date: newdate,
        Open: open,
        Close: close,
        Highest: highest,
        Lowest: lowest,
        Volume: volume,
        Money: money,
    }
}

func (s *StockHistDataDB) TranInsert(code string, datas [] stockhandler.StockHistData) int {
    db := s.Open()
    
    tx, err := db.Begin()
    util.CheckError(err)

    for _, d := range datas {
        stmt, err := tx.Prepare("insert stockhistdata set code=?, date=?, open=?, close=?, highest=?, lowest=?, volume=?, money=?")
        util.CheckError(err)

        _, reserr := stmt.Exec(code, d.Date, d.Open, d.Close, d.Highest, d.Lowest, d.Volume, d.Money)
        util.CheckError(reserr)
        defer stmt.Close()
    }
    
    err = tx.Commit()
    util.CheckError(err)

    db.Close()
    return 0
} 

func NewStockHistDataDB(dbtype string, dbcon string) *StockHistDataDB {
    return &StockHistDataDB{
        DBBase: DBBase{
            Dbtype: dbtype,
            Dbcon: dbcon,
        },
    }
}
