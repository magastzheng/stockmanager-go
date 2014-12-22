package stsummary

import (
    //"database/sql"
    _ "github.com/go-sql-driver/mysql"
    "stockdb"
    "entity/stentity"
    //"util"
    //"strconv"
    "fmt"
)

const(
    CompanyInsert = "insert %s set id=?, name=?, fullname=?, fullname_en=?, inceptdate=?, regaddr=?, website=?, states=?, city=?"
    CompanyDelete = "delete from %s where id=?"
    CompanyUpdate = "update %s set name=?, fullname=?, fullname_en=?, inceptdate=?, regaddr=?, website=?, states=?, city=? where id=?"
    CompanySelect = "select id, name, fullname, fullname_en, inceptdate, regaddr, website, states, city from %s where id = ?"
    //ListQueryCount = "select count(id) from %s"
    //ListQueryId = "select id from %s"
    //ListQueryIdExchange = "select id, exchange from %s"
)

type CompanyDB struct {
   stockdb.DBBase
   dbtable string
}

func (s *CompanyDB) getSql(sql string) string{
    return fmt.Sprintf(sql, s.dbtable)
}

func (s *CompanyDB) Insert(c stentity.Company) int {
    db := s.Open()
    defer db.Close()

    sql := s.getSql(CompanyInsert)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }
    
    res, err := stmt.Exec(c.Code, c.AbbrName, c.Name, c.Name_en, c.InceptDate, c.RegAddr, c.Website, c.State, c.City)
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

func (s *CompanyDB) Delete(c stentity.Company) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(CompanyDelete)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(c.Code)
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

func (s *CompanyDB) Update(c stentity.Company) int {
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(CompanyUpdate)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }

    res, err := stmt.Exec(c.AbbrName, c.Name, c.Name_en, c.InceptDate, c.RegAddr, c.Website, c.State, c.City)
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

func (s *CompanyDB) TranInsert(companies []stentity.Company) int {
    db := s.Open()
    defer db.Close()
	
    sql := s.getSql(CompanyInsert)
	tx, err := db.Begin()
	if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return -1
    }
	
	for _, c := range companies {
		stmt, err := tx.Prepare(sql)
        defer stmt.Close()
        if err != nil {
            s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
            continue
        }
        
		_, reserr := stmt.Exec(c.Code, c.AbbrName, c.Name, c.Name_en, c.InceptDate, c.RegAddr, c.Website, c.State, c.City)
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

func (s *CompanyDB) Query(id string) stentity.Company {
    c := stentity.Company{}
    db := s.Open()
    defer db.Close()
    
    sql := s.getSql(CompanySelect)
    stmt, err := db.Prepare(sql)
    defer stmt.Close()
    if err != nil {
        s.Logger.Error("Database error: ", s.Dbtype, s.Dbcon, err)
        return c
    }
    
    err = stmt.QueryRow(id).Scan(&c.Code, &c.AbbrName, &c.Name, &c.Name_en, &c.InceptDate, &c.RegAddr, &c.Website, &c.State, &c.City)
    if err != nil{
        s.Logger.Error("Cannot query the stock with id: ", id, err)
    }

    return c
}

func NewCompanyDB(dbname, dbtable string) *CompanyDB {
    stdb := new(CompanyDB)
    stdb.Init(dbname)
    stdb.dbtable = dbtable

    return stdb
}
