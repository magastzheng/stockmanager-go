package nationstatdb_test

import(
    "testing"
    "stockdb/nationstatdb"
    ns "entity/nsentity"
)

func Test_NSNonMfgPmiDB_Insert(t *testing.T){
    ms := ns.NSNonMfgPmi{
        NSPmiCommon:ns.NSPmiCommon{
            "2014-10-31",
            52.12,
            51.23,
            49.42,
            54.45,
            40.2,
            48.21,
            62.1,
        },
        ImInputPrice: 45.78,
        SubscriptionPrice: 47.32,
        BizActExpectation: 49.12,
    }

    db := nationstatdb.NewNSNonMfgPmiDB("macroindecis")
    db.Delete(ms.Date)
    db.Insert(ms)
}
