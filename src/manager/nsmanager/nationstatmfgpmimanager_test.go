package nsmanager_test

import(
    "testing"
    "manager/nsmanager"
    //"stockdb"
    //"fmt"
)

func Test_NSMfgPmiManager_Process(t *testing.T){
    m := nsmanager.NewNSMfgPmiManager()
    m.Process()
}
