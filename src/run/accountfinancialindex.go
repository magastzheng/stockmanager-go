package main

import(
    "dbcreator"
)

func main() {
    m := dbcreator.NewAccountDBCreator()
    m.Process()
}
