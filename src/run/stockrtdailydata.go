package main

import(
    "manager"
)

func main() {
    m := manager.NewStockRtDataManager()
    m.Process()
}
