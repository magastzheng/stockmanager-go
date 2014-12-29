package nationstatdb_test

import(
    "testing"
    nsdb "stockdb/nationstatdb"
    ns "entity/nsentity"
    "fmt"
)

func getIndexDB() *nsdb.IndexDB {
    return nsdb.NewIndexDB("macroindecis", "nsindex")
}

func Test_IndexDB_Insert(t *testing.T) {
    db := getIndexDB()
    idx := ns.NSDBIndex{}
    idx.Id = "A00001"
    idx.Name = "TestA"
    idx.Parent = "A00"
    
    res := db.Delete(idx.Id)
    res = db.Insert(idx)
    fmt.Println(res)

    idx = db.Query(idx.Id)
    fmt.Println(idx)
}

func Test_IndexDB_Delete(t *testing.T) {
    db := getIndexDB()
    res := db.Delete("A00001")

    fmt.Println(res)
}

func Test_IndexDB_Update(t *testing.T) {
    db := getIndexDB()
    idx := ns.NSDBIndex{}
    idx.Id = "A00001"
    idx.Name = "TestA"
    idx.Parent = "A00"
    
    res := db.Insert(idx)
    fmt.Println(res)

    idx.Readid = "onlytest"
    res = db.Update(idx)

    idx = db.Query(idx.Id)
    fmt.Println("After: ", idx)
}
