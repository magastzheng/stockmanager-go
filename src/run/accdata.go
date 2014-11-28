package main

import(
    "manager/accmanager"
)

func main() {
    m := accmanager.NewAccountManager()
    m.Process()
}
