package main

import (
    "manager"
    "fmt"
)

func main() {
    m := manager.NewStockListManager()
    m.Process()
    fmt.Println("Stock list run complete!")
}
