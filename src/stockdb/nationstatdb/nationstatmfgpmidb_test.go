package nationstatdb_test

import(
    "testing"
    "stockdb/nationstatdb"
    ns "entity/nsentity"
)

func Test_NSMfgPmiDB_Insert(t *testing.T){
    ms := ns.NSMfgPmi{
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
        Production: 45.78,
        PurchasingVolume: 47.32,
        Import: 51.23,
        MainRawMaterial: 49.12,
        RawMaterialInventory: 48.52,
        PbizActExpectation: 44.51,
    }

    db := nationstatdb.NewNSMfgPmiDB("macroindecis")
    db.Delete(ms.Date)
    db.Insert(ms)
}
