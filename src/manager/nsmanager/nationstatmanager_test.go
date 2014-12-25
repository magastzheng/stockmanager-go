package nsmanager_test

import(
    "testing"
    "manager/nsmanager"
)

func Test_NationStatManager_GetData(t *testing.T){
    m := nsmanager.NewNationStatManager()
    m.Process()
}
