package nationstatdb

import(
    _ "github.com/go-sql-driver/mysql"
    "stockdb"
    "entity"
    "util"
)

type NSNonMfgPmiDB struct{
    stockdb.DBBase
}

func (s *NSNonMfgPmiDB) Insert(ms entity.NSNonMfgPmi) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("insert nonmfgpmi set date=?, pmi=?, neworder=?, newexportorder=?, inhandorder=?, inventory=?, employees=?, supplierdeliverytime=?, iminputprice=?, subscriptionprice=?, bizactexpectation=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(ms.Date, ms.Pmi, ms.NewOrder, ms.NewExportOrder, ms.InHandOrder, ms.Inventory, ms.Employees, ms.SupplierDeliveryTime, ms.ImInputPrice, ms.SubscriptionPrice, ms.BizActExpectation)
    util.CheckError(err)

    _, reserr := res.LastInsertId()
    util.CheckError(reserr)
    
    return 0
}

func (s *NSNonMfgPmiDB) Delete(date string) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("delete from nonmfgpmi where date=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(date)
    util.CheckError(err)

    _,reserr := res.RowsAffected()
    util.CheckError(reserr)

    return 0
}

func (s *NSNonMfgPmiDB)Update(ms entity.NSNonMfgPmi) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("update nonmfgpmi set pmi=?, neworder=?, newexportorder=?, inhandorder=?, inventory=?, employees=?, supplierdeliverytime=?, iminputprice=?, subscriptionprice=?, bizactexpectation=? where date=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(ms.Pmi, ms.NewOrder, ms.NewExportOrder, ms.InHandOrder, ms.Inventory, ms.Employees, ms.SupplierDeliveryTime, ms.ImInputPrice, ms.SubscriptionPrice, ms.BizActExpectation, ms.Date)
    util.CheckError(err)

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    return 0
}

func (s *NSNonMfgPmiDB)Query(date string) entity.NSNonMfgPmi {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("select pmi,neworder,newexportorder,inhandorder,inventory,employees,supplierdeliverytime,production,purchasingvolume,iminputprice, subscriptionprice, bizactexpectation from nonmfgpmi where date=?")
    defer stmt.Close()
    util.CheckError(err)
    
    //ms := entity.NSMfgPmi{
    //    entity.NSPmiCommon{
    //        Date: date,
    //        Pmi: 0.0,
    //    },
    //    Production: 0.0,
    //}
    
    var ms entity.NSNonMfgPmi
    ms.Date = date

    err = stmt.QueryRow(date).Scan(&ms.Pmi, &ms.NewOrder, &ms.NewExportOrder, &ms.InHandOrder, &ms.Inventory, &ms.Employees, &ms.SupplierDeliveryTime, &ms.ImInputPrice, &ms.SubscriptionPrice, &ms.BizActExpectation)
    util.CheckError(err)

    return ms
}

func (s *NSNonMfgPmiDB) TranInsert(mses []entity.NSNonMfgPmi) int {
    db := s.Open()
    defer db.Close()
    
    tx, err := db.Begin()
    util.CheckError(err)

    for _, ms := range mses {
        stmt, err := tx.Prepare("insert nonmfgpmi set date=?, pmi=?, neworder=?, newexportorder=?, inhandorder=?, inventory=?, employees=?, supplierdeliverytime=?, iminputprice=?, subscriptionprice=?, bizactexpectation=?")
        defer stmt.Close()
        util.CheckError(err)
    
        _, reserr := stmt.Exec(ms.Date, ms.Pmi, ms.NewOrder, ms.NewExportOrder, ms.InHandOrder, ms.Inventory, ms.Employees, ms.SupplierDeliveryTime, ms.ImInputPrice, ms.SubscriptionPrice, ms.BizActExpectation)
        util.CheckError(reserr)
    }
    
    err = tx.Commit()
    util.CheckError(err)
    
    return 0
}


func NewNSNonMfgPmiDB(dbname string) *NSNonMfgPmiDB{
    nsdb := new(NSNonMfgPmiDB)
    nsdb.Init(dbname)

    return nsdb
}
