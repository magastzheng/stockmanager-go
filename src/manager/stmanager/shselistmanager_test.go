package manager_test

import(
    "testing"
    manager "manager/stmanager"
    //"fmt"
    //"os"
    //"encoding/json"
)

func Test_SHSEListManager_Process(t *testing.T){
    m := manager.NewSHSEListManager()
    m.Process()
}
