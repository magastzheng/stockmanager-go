package stockdb

import(
    "fmt"
)

type SPDB struct {
    DBBase
}

func (d *SPDB) CallStoreProc(code string) {
    db := d.Open()
    
    stmt, err := db.Prepare("CALL getcsrccategorystock('?')")
    if err != nil {
        panic(err)
    }
    
    rows, err := stmt.Query(code)

    fmt.Println(code)
    //rows, err := db.Query("CALL getcsrccategorystock(?)", code)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }
    
    ids := make([]string, 0)
    id := ""
    for rows.Next() {
        if err := rows.Scan(&id); err != nil {
            panic(err)
        }
        ids = append(ids, id)
    }

    fmt.Println(ids)
}

func NewSPDB() *SPDB {
    db := new(SPDB)
    db.Init("csrcstockcategory")

    return db
}
