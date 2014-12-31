package stockdb_test

import(
    "testing"
    "stockdb"
    //"fmt"
)

func Test_SPDB_CallSP(t *testing.T){
    db := stockdb.NewSPDB()
    db.CallStoreProc("21")
}
