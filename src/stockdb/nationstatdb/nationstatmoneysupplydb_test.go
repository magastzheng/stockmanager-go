package nationstatdb_test

import(
    "testing"
    "stockdb/nationstatdb"
    ns "entity/nsentity"
    "fmt"
)

func Test_NSMoneySupportInsert(t *testing.T) {
    db := nationstatdb.NewMSMoneySupplyDB("macroindecis")
    db.Delete("2013-01-31")
    ms := ns.MoneySupply{
        Date: "2013-01-31",
        M0: 12548778,
        M0pct: 2.5,
        M1: 657892185,
        M1pct: 4.5,
        M2: 745678778,
        M2pct: 4.6,
    }

    db.Insert(ms)

    newms := db.Query("2013-01-31")
    fmt.Println(newms)
    db.Delete("2013-01-31")
}
