package stockdb

import(
    "fmt"
)

type SPDB struct {
    DBBase
}

func (d *SPDB) CallStoreProc(code string) {
    db := d.Open()
	
	//rows, err := db.Query("select id from csrcstockcategory")
    
    stmt, err := db.Prepare("CALL getcsrccategorystock(?)")
    if err != nil {
        panic(err)
    }
    
    rows, err := stmt.Query(code)
	//res, err := db.Exec("call getcsrccategorystock('?')", code)
	//if err != nil {
	//	fmt.Println(err)
	//}
    //fmt.Println(code)
    //rows, err := db.Query("call getcsrccategorystock(?)", code)
    if err != nil {
        fmt.Println(err)
        //panic(err)
		//return
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
    //db.Init("csrcstockcategory")
	db.Init("chinastock")

    return db
}
