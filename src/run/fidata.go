package main

import(
    "manager/accmanager"
    "os"
    //"path/filepath"
    "fmt"
)

func main() {
    //fmt.Println(filepath.Dir("."))
    fmt.Println(os.Getwd())
    m := accmanager.NewFiManager()
    m.Process()
}
