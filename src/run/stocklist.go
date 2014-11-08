package main

import (
    "listmanager"
    "fmt"
)

func main() {
    manager := listmanager.NewStockListManager()
    manager.Process()
    fmt.Println("Stock list run complete!")
}
