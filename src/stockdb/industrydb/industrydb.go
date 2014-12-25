package industrydb

import (
    _ "github.com/go-sql-driver/mysql"
    "stockdb"
    "entity/xlsentity"
    "fmt"
    //"excel"
    //"util"
)

const(
    IndustryInsert = "insert %s set code=?, parent=?, name=?, name_en=?"
    IndustryDelete = "delete from %s where code=?"
    IndustryUpdate = "update %s set parent=?, name=?, name_en=? where code=?"
    IndustrySelect = "select code, parent, name, name_en from %s where code=?"
)

type IndustryDB struct {
    stockdb.DBBase
    dbtable string
}

func (s *IndustryDB) getSql(sql string) string{
    return fmt.Sprintf(sql, s.dbtable)
}

func (s *IndustryDB) Insert(industry xlsentity.Industry) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(IndustryInsert)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(industry.Code, industry.Parent, industry.Name, industry.Name_en)
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

func (s *IndustryDB) Delete(code string) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(IndustryDelete)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(code)
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

func (s *IndustryDB) Update(industry xlsentity.Industry) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(IndustryDelete)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(industry.Parent, industry.Name, industry.Name_en)
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

func (s *IndustryDB) Query(code string) xlsentity.Industry {
    db := s.Open()
    defer db.Close()

    industry := xlsentity.Industry{}
    
    sql := s.getSql(IndustrySelect)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return industry
    }

    err = stmt.QueryRow(code).Scan(&industry.Code, &industry.Parent, &industry.Name, &industry.Name_en)
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return industry
    }

    return industry
}

func (s *IndustryDB) TranInsert(industries map[string] xlsentity.Industry) int {
    db := s.Open()
    defer db.Close()
    
    tx, err := db.Begin()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    sql := s.getSql(IndustryInsert)
    for key, industry := range industries {
        stmt, err := tx.Prepare(sql)
        if err != nil {
            s.Logger.Error("Database error in transaction insert industry: ", s.Dbtype, s.Dbcon, err, sql)
            continue
        }

        _, reserr := stmt.Exec(key, industry.Parent, industry.Name, industry.Name_en)
        if reserr != nil {
            s.Logger.Error("Database error in transaction insert industry: ", s.Dbtype, s.Dbcon, reserr, industry)
            continue
        }

        defer stmt.Close()
    }
    
    err = tx.Commit()
    if err != nil {
        s.Logger.Error("Database error - cannot commit in transaction insert industry: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    return 0
} 

func NewIndustryDB(dbname, dbtable string) *IndustryDB {
    stdb := new(IndustryDB)
    stdb.Init(dbname)
    stdb.dbtable = dbtable

    return stdb
}
