package main

import(
    "manager/accmanager"
)

func main() {
    m := accmanager.NewFiManager()
    m.Process()
}
