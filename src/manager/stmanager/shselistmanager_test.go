package stmanager_test

import(
    "testing"
    "manager/stmanager"
    //"fmt"
    //"os"
    //"encoding/json"
)

func Test_SHSEListManager_Process(t *testing.T){
    m := stmanager.NewSHSEListManager()
    m.Process()
}
