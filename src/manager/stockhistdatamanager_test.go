package manager_test

import (
    "testing"
    "manager"
)

func Test_StockHistDataManager(t *testing.T) {
    m := manager.NewStockHistDataManager()
    m.Process()
}
