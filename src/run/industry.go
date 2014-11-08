package main

import (
    "listmanager"
    "fmt"
)

func main() {
    manager := listmanager.NewIndustryManager()
    manager.Process()
    fmt.Println("Industry complete!")
}


