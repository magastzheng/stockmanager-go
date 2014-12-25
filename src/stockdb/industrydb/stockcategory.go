package industrydb

import(
    "stockdb"
    "entity/xlsentity"
    "fmt"
)

const (
    SCInsert = "insert %s set id=?, code=?"
    SCDelete = "delete from %s where id=?"
    SCUpdate = "update %s set code=? where id=?"
    SCSelectCat = "select id, code from %s where id=?"
    SCSelectAll = "select id, code from %s"
    SCSelectStock = "select id from %s where code=?"
)

type StockCategoryDB struct {
    stockdb.DBBase
    dbtable string
}

func (s *StockCategoryDB) getSql(sql string) string {
    return fmt.Sprintf(sql, s.dbtable)
}

func (s *StockCategoryDB) Insert(sc xlsentity.StockCategory) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(SCInsert)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(sc.Id, sc.Code)
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

func (s *StockCategoryDB) Delete(id string) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(SCDelete)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(id)
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

func (s *StockCategoryDB) Update(sc xlsentity.StockCategory) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(SCUpdate)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(sc.Code, sc.Id)
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

func (s *StockCategoryDB) TranInsert(scs []xlsentity.StockCategory) int {
    db := s.Open()
    defer db.Close()
    
    tx, err := db.Begin()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    sql := s.getSql(SCInsert)
    for _, sc := range scs {
        stmt, err := tx.Prepare(sql)
        if err != nil {
            s.Logger.Error("Database error in transaction insert stockcategory: ", s.Dbtype, s.Dbcon, err, sql)
            continue
        }

        _, reserr := stmt.Exec(sc.Id, sc.Code)
        if reserr != nil {
            s.Logger.Error("Database error in transaction insert stockcategory: ", s.Dbtype, s.Dbcon, reserr, sc)
            continue
        }

        defer stmt.Close()
    }
    
    err = tx.Commit()
    if err != nil {
        s.Logger.Error("Database error - cannot commit in transaction insert stockcategory: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    return 0
} 


func (s *StockCategoryDB) Query(id string) xlsentity.StockCategory {
    db := s.Open()
    defer db.Close()

    sc := xlsentity.StockCategory{}
    
    sql := s.getSql(SCSelectCat)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return sc
    }

    err = stmt.QueryRow(id).Scan(&sc.Id, &sc.Code)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return sc
    }

    return sc
}

func (s *StockCategoryDB) QueryAll() []xlsentity.StockCategory {
    db := s.Open()
    defer db.Close()

    scs := make([]xlsentity.StockCategory, 0)
    
    sql := s.getSql(SCSelectAll)
    rows, err := db.Query(sql)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return scs
    }

    var id, code string
    for rows.Next() {
        err = rows.Scan(&id, &code)
        if err != nil {
            continue
        }

        sc := xlsentity.StockCategory{
            Id: id,
            Code: code,
        }

        scs = append(scs, sc)
    }

    return scs
}

func (s *StockCategoryDB) QueryStockIds(code string) []string {
    db := s.Open()
    defer db.Close()

    ids := make([]string, 0)
    
    sql := s.getSql(SCSelectStock)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return ids
    }

    rows, err := stmt.Query(code)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return ids
    }

    var id string
    for rows.Next() {
        err = rows.Scan(&id)
        if err != nil {
            continue
        }

        ids = append(ids, id)
    }

    return ids
}

func NewStockCategoryDB(dbname, dbtable string) *StockCategoryDB{
    db := new(StockCategoryDB)
    db.Init(dbname)
    db.dbtable = dbtable

    return db
}

