package manager_test

import(
    "testing"
    "manager"
    //"util"
)

func Test_ProcessRealtimeDataManager(t *testing.T) {
    m := manager.NewStockRtDataManager()
    m.Process()
}

