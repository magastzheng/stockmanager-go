package main

import (
    "manager"
    "fmt"
)

func main() {
    m := manager.NewIndustryManager()
    m.Process()
    fmt.Println("Industry complete!")
}


