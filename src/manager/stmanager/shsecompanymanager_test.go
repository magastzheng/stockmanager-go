package stmanager_test

import(
    "testing"
    "manager/stmanager"
)

func Test_SHSECompanyManager_Process(t *testing.T){
    m := stmanager.NewSHSECompanyManager()
    m.Process()
}

