package "main"

import (
    "listmanager"
)

func main() {
    manager := new(listmanager.StockListManager)
    manager.Init("config/stocklist.json")
    manager.Process()
}
