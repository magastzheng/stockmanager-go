package nationstatdb

import(
    _ "github.com/go-sql-driver/mysql"
    "stockdb"
    ns "entity/nsentity"
    "util"
)

type NSMfgPmiDB struct{
    stockdb.DBBase
}

func (s *NSMfgPmiDB) Insert(ms ns.NSMfgPmi) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("insert mfgpmi set date=?, pmi=?, neworder=?, newexportorder=?, inhandorder=?, inventory=?, employees=?, supplierdeliverytime=?, production=?, purchasingvolume=?, import=?, mainrawmaterialpurchaseprice=?, rawmaterialsinventory=?, pbizactexpectation=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(ms.Date, ms.Pmi, ms.NewOrder, ms.NewExportOrder, ms.InHandOrder, ms.Inventory, ms.Employees, ms.SupplierDeliveryTime, ms.Production, ms.PurchasingVolume, ms.Import, ms.MainRawMaterial, ms.RawMaterialInventory, ms.PbizActExpectation)
    util.CheckError(err)

    _, reserr := res.LastInsertId()
    util.CheckError(reserr)
    
    return 0
}

func (s *NSMfgPmiDB) Delete(date string) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("delete from mfgpmi where date=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(date)
    util.CheckError(err)

    _,reserr := res.RowsAffected()
    util.CheckError(reserr)

    return 0
}

func (s *NSMfgPmiDB)Update(ms ns.NSMfgPmi) int {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("update mfgpmi set pmi=?, neworder=?, newexportorder=?, inhandorder=?, inventory=?, employees=?, supplierdeliverytime=?, production=?, purchasingvolume=?, import=?, mainrawmaterialpurchaseprice=?, rawmaterialsinventory=?, pbizactexpectation=? where date=?")
    defer stmt.Close()
    util.CheckError(err)

    res, err := stmt.Exec(ms.Pmi, ms.NewOrder, ms.NewExportOrder, ms.InHandOrder, ms.Inventory, ms.Employees, ms.SupplierDeliveryTime, ms.Production, ms.PurchasingVolume, ms.Import, ms.MainRawMaterial, ms.RawMaterialInventory, ms.PbizActExpectation, ms.Date)
    util.CheckError(err)

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    return 0
}

func (s *NSMfgPmiDB)Query(date string) ns.NSMfgPmi {
    db := s.Open()
    defer db.Close()

    stmt, err := db.Prepare("select pmi,neworder,newexportorder,inhandorder,inventory,employees,supplierdeliverytime,production,purchasingvolume,import,mainrawmaterialpurchaseprice,rawmaterialsinventory,pbizactexpectation from mfgpmi where date=?")
    defer stmt.Close()
    util.CheckError(err)
    
    //ms := ns.NSMfgPmi{
    //    ns.NSPmiCommon{
    //        Date: date,
    //        Pmi: 0.0,
    //    },
    //    Production: 0.0,
    //}
    
    var ms ns.NSMfgPmi
    ms.Date = date

    err = stmt.QueryRow(date).Scan(&ms.Pmi, &ms.NewOrder, &ms.NewExportOrder, &ms.InHandOrder, &ms.Inventory, &ms.Employees, &ms.SupplierDeliveryTime, &ms.Production, &ms.PurchasingVolume, &ms.Import, &ms.MainRawMaterial, &ms.RawMaterialInventory, &ms.PbizActExpectation)
    util.CheckError(err)

    return ms
}

func (s *NSMfgPmiDB) TranInsert(mses []ns.NSMfgPmi) int {
    db := s.Open()
    defer db.Close()
    
    tx, err := db.Begin()
    util.CheckError(err)

    for _, ms := range mses {
        stmt, err := tx.Prepare("insert mfgpmi set date=?, pmi=?, neworder=?, newexportorder=?, inhandorder=?, inventory=?, employees=?, supplierdeliverytime=?, production=?, purchasingvolume=?, import=?, mainrawmaterialpurchaseprice=?, rawmaterialsinventory=?, pbizactexpectation=?")
        defer stmt.Close()
        util.CheckError(err)
    
        _, reserr := stmt.Exec(ms.Date, ms.Pmi, ms.NewOrder, ms.NewExportOrder, ms.InHandOrder, ms.Inventory, ms.Employees, ms.SupplierDeliveryTime, ms.Production, ms.PurchasingVolume, ms.Import, ms.MainRawMaterial, ms.RawMaterialInventory, ms.PbizActExpectation)
        util.CheckError(reserr)
    }
    
    err = tx.Commit()
    util.CheckError(err)
    
    return 0
}


func NewNSMfgPmiDB(dbname string) *NSMfgPmiDB{
    nsdb := new(NSMfgPmiDB)
    nsdb.Init(dbname)

    return nsdb
}
