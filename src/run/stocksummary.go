package main

import(
    "manager/stmanager"
    "fmt"
)

func main(){
    var c chan int
    c = make(chan int)
    shm := stmanager.NewSHSECompanyManager()
    go func(){
        shm.Process()
        c <- 1
    }()

    szm := stmanager.NewSZSECompanyManager()
    go func() {
        szm.Process()
        c <- 2
    }()

    <-c
    <-c

    fmt.Println("Success to get company information")
}
