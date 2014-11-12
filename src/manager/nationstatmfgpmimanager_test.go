package manager_test

import(
    "testing"
    "manager"
    //"stockdb"
    //"fmt"
)

func Test_NSMfgPmiManager_Process(t *testing.T){
    m := manager.NewNSMfgPmiManager()
    m.Process()
}
