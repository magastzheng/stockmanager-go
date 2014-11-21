package manager_test

import(
    "testing"
    "manager"
    //"fmt"
)

func Test_AccountDBManager_Process(t *testing.T){
    m := manager.NewAccountDBManager()
    m.Process()
}
