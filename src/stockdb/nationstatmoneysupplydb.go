package stockdb

import(
    _ "github.com/go-sql-driver/mysql"
    "entity"
    "util"
)

type NSMoneySupplyDB struct{
    DBBase
}

func (s *NSMoneySupplyDB) Insert(ms entity.MoneySupply) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("insert moneysupply set date=?, m0=?, m0pct=?, m1=?, m1pct=?, m2=?, m2pct=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(ms.Date, ms.M0, ms.M0pct, ms.M1, ms.M1pct, ms.M2, ms.M2pct)
    util.CheckError(err)

    _, reserr := res.LastInsertId()
    util.CheckError(reserr)
    
    return 0
}

func (s *NSMoneySupplyDB) Delete(date string) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("delete from moneysupply where date=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(date)
    util.CheckError(err)

    _,reserr := res.RowsAffected()
    util.CheckError(reserr)

    return 0
}

func (s *NSMoneySupplyDB)Update(ms entity.MoneySupply) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("update moneysupply set m0=?, m0pct=?, m1=?, m1pct=?, m2=?, m2pct=? where date=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(ms.M0, ms.M0pct, ms.M1, ms.M1pct, ms.M2, ms.M2pct, ms.Date)
    util.CheckError(err)

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    return 0
}

func (s *NSMoneySupplyDB)Query(date string) entity.MoneySupply {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("select m0, m0pct, m1, m1pct, m2, m2pct from moneysupply where date=?")
    defer stmt.Close()
    util.CheckError(err)
    
    ms := entity.MoneySupply{
        Date: date,
    }

    err = stmt.QueryRow(date).Scan(&ms.M0, &ms.M0pct, &ms.M1, &ms.M1pct, &ms.M2, &ms.M2pct)
    util.CheckError(err)

    return ms
}

func (s *NSMoneySupplyDB) TranInsert(mses []entity.MoneySupply) int {
    db := s.Open()
    defer db.Close()
    
    tx, err := db.Begin()
    util.CheckError(err)

    for _, ms := range mses {
        stmt, err := tx.Prepare("insert moneysupply set date=?, m0=?, m0pct=?, m1=?, m1pct=?, m2=?, m2pct=?")
        defer stmt.Close()
        util.CheckError(err)
    
        _, reserr := stmt.Exec(ms.Date, ms.M0, ms.M0pct, ms.M1, ms.M1pct, ms.M2, ms.M2pct)
        util.CheckError(reserr)
    }
    
    err = tx.Commit()
    util.CheckError(err)
    
    return 0
}


func NewMSMoneySupplyDB(dbname string) *NSMoneySupplyDB{
    nsdb := new(NSMoneySupplyDB)
    nsdb.Init(dbname)

    return nsdb
}
