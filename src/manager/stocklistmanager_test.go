package manager_test

import(
    "testing"
    "manager"
    //"fmt"
)

func Test_Process(t *testing.T) {
    //s := new(StockListManager)
    //s.Init("../config/stocklist.json")
    s := manager.NewStockListManager()
    s.Process()
}
