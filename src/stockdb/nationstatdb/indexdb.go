package nationstatdb

import(
    "stockdb"
    ns "entity/nsentity"
    //"util"
    "fmt"
)

const(
    IndexInsert = "insert %s set id=?, parent=?, name=?, ename=?, unit=?, eunit=?, note=?, enote=?, readid=?"
    IndexDelete = "delete from %s where id=?"
    IndexUpdate = "update %s set parent=?, name=?, ename=?, unit=?, eunit=?, note=?, enote=?, readid=?"
    IndexSelect = "select id, parent, name, ename, unit, eunit, note, enote, readid from %s where id=?"
    IndexSelectAll = "select id, parent, name, ename, unit, eunit, note, enote, readid from %s"
)

type IndexDB struct {
    stockdb.DBBase
    dbtable string
}

func (s *IndexDB) getSql(sql string) string {
    return fmt.Sprintf(sql, s.dbtable)
}

func (s *IndexDB) Insert(idx ns.NSDBIndex) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(IndexInsert)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        //fmt.Println(sql)
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(idx.Id, idx.Parent, idx.Name, idx.EName, idx.Unit, idx.Eunit, idx.Note, idx.Enote, idx.Readid)
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

func (s *IndexDB) Delete(id string) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(IndexDelete)
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

func (s *IndexDB) Update(idx ns.NSDBIndex) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(IndexUpdate)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(idx.Parent, idx.Name, idx.EName, idx.Unit, idx.Eunit, idx.Note, idx.Enote, idx.Readid)
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

func (s *IndexDB) Query(id string) ns.NSDBIndex {
    db := s.Open()
    defer db.Close()

    idx := ns.NSDBIndex{}
    
    sql := s.getSql(IndexSelect)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return idx
    }

    err = stmt.QueryRow(id).Scan(&idx.Id, &idx.Parent, &idx.Name, &idx.EName, &idx.Unit, &idx.Eunit, &idx.Note, &idx.Enote, &idx.Readid)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return idx
    }

    return idx
}

func (s *IndexDB) TranInsert(idxes []ns.NSDBIndex) int {
    db := s.Open()
    defer db.Close()
    
    tx, err := db.Begin()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    sql := s.getSql(IndexInsert)
    for _, idx := range idxes {
        stmt, err := tx.Prepare(sql)
        if err != nil {
            s.Logger.Error("Database error in transaction insert index: ", s.Dbtype, s.Dbcon, err, sql)
            continue
        }

        _, reserr := stmt.Exec(idx.Id, idx.Parent, idx.Name, idx.EName, idx.Unit, idx.Eunit, idx.Note, idx.Enote, idx.Readid)
        if reserr != nil {
            s.Logger.Error("Database error in transaction insert index: ", s.Dbtype, s.Dbcon, reserr, idx)
            continue
        }

        defer stmt.Close()
    }
    
    err = tx.Commit()
    if err != nil {
        s.Logger.Error("Database error - cannot commit in transaction insert index: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    return 0
} 

func NewIndexDB(dbname, dbtable string) *IndexDB {
    db := new(IndexDB)
    db.Init(dbname)
    db.dbtable = dbtable

    return db
}
