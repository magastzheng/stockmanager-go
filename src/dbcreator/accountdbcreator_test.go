package dbcreator_test

import(
    "testing"
    "dbcreator"
    //"fmt"
)

func Test_AccountDBCreator_Process(t *testing.T){
    m := dbcreator.NewAccountDBCreator()
    m.Process()
}
