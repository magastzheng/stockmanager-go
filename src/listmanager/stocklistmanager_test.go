package listmanager

import(
    "testing"
    //"fmt"
)

func Test_Process(t *testing.T) {
    s := new(StockListManager)
    s.Init("../config/stocklist.json")
    s.Process()
}
