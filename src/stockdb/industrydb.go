package stockdb

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "excel"
    "util"
)

type IndustryDatabase struct {
    dbtype string
    dbcon string
}

func (s* IndustryDatabase) Open() *sql.DB {
    db, err := sql.Open(s.dbtype, s.dbcon)
    util.CheckError(err)
    
    return db
}

func (s *IndustryDatabase) InsertIndustry(industry excel.Industry) int {
    db := s.Open()
    
    stmt, err := db.Prepare("insert csrcbigindustry set code=?, name=?, name_en=?")
    util.CheckError(err)

    res, err := stmt.Exec(industry.BigCode, industry.Name, industry.Name_en)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.LastInsertId()
    util.CheckError(reserr)
    
    db.Close()
    return 0
}

func (s *IndustryDatabase) DeleteIndustry(code string) int {
    db := s.Open()

    stmt, err := db.Prepare("delete from csrcbigindustry where code=?")
    util.CheckError(err)

    res, err := stmt.Exec(code)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    db.Close()
    return 0
}

func (s *IndustryDatabase) UpdateIndustry(industry excel.Industry) int {
    db := s.Open()
    
    stmt, err := db.Prepare("update csrcbigindustry set name=?, name_en=? where code=?")
    util.CheckError(err)

    res, err := stmt.Exec(industry.Name, industry.Name_en, industry.BigCode)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    db.Close()
    return 0
}

func (s *IndustryDatabase) QueryIndustry(code string) excel.Industry {
    db := s.Open()
    
    stmt, err := db.Prepare("select code, name, name_en from csrcbigindustry where code = ?")
    util.CheckError(err)
    defer stmt.Close()

    var newcode, name, name_en string
    err = stmt.QueryRow(code).Scan(&newcode, &name, &name_en)
    util.CheckError(err)

    db.Close()

    return excel.Industry{
        BigCode: newcode,
        Name: name,
        Name_en: name_en,
    }
}

func (s *IndustryDatabase) TranInsertIndustry(industries map[string] excel.Industry) int {
    db := s.Open()
    
    tx, err := db.Begin()
    util.CheckError(err)

    for key, industry := range industries {
        stmt, err := tx.Prepare("insert csrcbigindustry set code=?, name=?, name_en=?")
        util.CheckError(err)

        _, reserr := stmt.Exec(key, industry.Name, industry.Name_en)
        util.CheckError(reserr)
        defer stmt.Close()
    }
    
    err = tx.Commit()
    util.CheckError(err)

    db.Close()
    return 0
} 

func NewIndustryDB(dbtype string, dbcon string) *IndustryDatabase {
    return &IndustryDatabase{
        dbtype: dbtype,
        dbcon: dbcon,
    }
}
