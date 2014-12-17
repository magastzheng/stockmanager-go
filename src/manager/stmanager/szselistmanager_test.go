package stmanager_test

import(
    "testing"
    "manager/stmanager"
    //"fmt"
    //"os"
    //"encoding/json"
)

func Test_SZSEListManager_Process(t *testing.T){
    m := stmanager.NewSZSEListManager()
    m.Process()
}
