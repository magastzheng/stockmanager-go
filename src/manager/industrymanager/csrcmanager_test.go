package industrymanager_test

import(
    "testing"
    "manager/industrymanager"
)

func Test_CSRCManager_Process(t *testing.T) {
    m := industrymanager.NewCSRCManager()
    m.Process()
}
