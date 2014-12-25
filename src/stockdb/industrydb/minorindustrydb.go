package industrydb

import (
    _ "github.com/go-sql-driver/mysql"
    "stockdb"
    "entity/xlsentity"
    "util"
)

type MinorIndustryDatabase struct {
    stockdb.DBBase
}

func (s *MinorIndustryDatabase) InsertIndustry(industry xlsentity.Industry) int {
    db := s.Open()
    
    stmt, err := db.Prepare("insert csrcminorindustry set code=?, name=?, name_en=?, bigcode=?")
    util.CheckError(err)

    res, err := stmt.Exec(industry.Code, industry.Name, industry.Name_en, industry.Parent)
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

func (s *MinorIndustryDatabase) UpdateIndustry(industry xlsentity.Industry) int {
    db := s.Open()
    
    stmt, err := db.Prepare("update csrcminorindustry set name=?, name_en=?, bigcode=? where code=?")
    util.CheckError(err)

    res, err := stmt.Exec(industry.Name, industry.Name_en, industry.Parent, industry.Code)
    util.CheckError(err)
    defer stmt.Close()

    _, reserr := res.RowsAffected()
    util.CheckError(reserr)

    db.Close()
    return 0
}

func (s *MinorIndustryDatabase) QueryIndustry(code int) xlsentity.Industry {
    db := s.Open()
    
    stmt, err := db.Prepare("select code, name, name_en, bigcode from csrcminorindustry where code = ?")
    util.CheckError(err)
    defer stmt.Close()
    
    var name, name_en, bigcode string
    var minorcode string

    err = stmt.QueryRow(code).Scan(&minorcode, &name, &name_en, &bigcode)
    util.CheckError(err)

    db.Close()
    
    return xlsentity.Industry{
        Code: minorcode,
        Parent: bigcode,
        Name: name,
        Name_en: name_en,
    }
}

func (s *MinorIndustryDatabase) TranInsertIndustry(industries map[string] xlsentity.Industry) int {
    db := s.Open()
    
    tx, err := db.Begin()
    util.CheckError(err)

    for key, industry := range industries {
        stmt, err := tx.Prepare("insert csrcminorindustry set code=?, name=?, name_en=?, bigcode=?")
        util.CheckError(err)

        _, reserr := stmt.Exec(key, industry.Name, industry.Name_en, industry.Parent)
        util.CheckError(reserr)
        defer stmt.Close()
    }
    
    err = tx.Commit()
    util.CheckError(err)

    db.Close()
    return 0
} 

func NewMinorIndustryDB(dbname string) *MinorIndustryDatabase {
    stdb := new(MinorIndustryDatabase)
    stdb.Init(dbname)
    return stdb
}
