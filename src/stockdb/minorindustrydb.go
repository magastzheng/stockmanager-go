package stockdb

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "excel"
    "util"
)

type MinorIndustryDatabase struct {
    dbtype string
    dbcon string
}

func (s* MinorIndustryDatabase) Open() *sql.DB {
    db, err := sql.Open(s.dbtype, s.dbcon)
    util.CheckError(err)
    
    return db
}

func (s *MinorIndustryDatabase) InsertIndustry(industry excel.MinorIndustry) int {
    db := s.Open()
    
    stmt, err := db.Prepare("insert csrcminorindustry set code=?, name=?, name_en=?, bigcode=?")
    util.CheckError(err)

    res, err := stmt.Exec(industry.MinorCode, industry.Name, industry.Name_en, industry.BigCode)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.LastInsertId()
    util.CheckError(reserr)
    
    db.Close()
    return 0
}

func (s *MinorIndustryDatabase) DeleteIndustry(code int) int {
    db := s.Open()

    stmt, err := db.Prepare("delete from csrcminorindustry where code=?")
    util.CheckError(err)

    res, err := stmt.Exec(code)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    db.Close()
    return 0
}

func (s *MinorIndustryDatabase) UpdateIndustry(industry excel.MinorIndustry) int {
    db := s.Open()
    
    stmt, err := db.Prepare("update csrcminorindustry set name=?, name_en=?, bigcode=? where code=?")
    util.CheckError(err)

    res, err := stmt.Exec(industry.Name, industry.Name_en, industry.BigCode, industry.MinorCode)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    db.Close()
    return 0
}

func (s *MinorIndustryDatabase) QueryIndustry(code int) excel.MinorIndustry {
    db := s.Open()
    
    stmt, err := db.Prepare("select code, name, name_en, bigcode from csrcminorindustry where code = ?")
    util.CheckError(err)
    defer stmt.Close()
    
    var name, name_en, bigcode string
    var minorcode int

    err = stmt.QueryRow(code).Scan(&minorcode, &name, &name_en, &bigcode)
    util.CheckError(err)

    db.Close()

    return excel.MinorIndustry{
        MinorCode: minorcode,
        BigCode: bigcode,
        Name: name,
        Name_en: name_en,
    }
}

func (s *MinorIndustryDatabase) TranInsertIndustry(industries map[int] excel.MinorIndustry) int {
    db := s.Open()
    
    tx, err := db.Begin()
    util.CheckError(err)

    for key, industry := range industries {
        stmt, err := tx.Prepare("insert csrcminorindustry set code=?, name=?, name_en=?, bigcode=?")
        util.CheckError(err)

        _, reserr := stmt.Exec(key, industry.Name, industry.Name_en, industry.BigCode)
        util.CheckError(reserr)
        defer stmt.Close()
    }
    
    err = tx.Commit()
    util.CheckError(err)

    db.Close()
    return 0
} 

func NewMinorIndustryDB(dbtype string, dbcon string) *MinorIndustryDatabase {
    return &MinorIndustryDatabase{
        dbtype: dbtype,
        dbcon: dbcon,
    }
}
