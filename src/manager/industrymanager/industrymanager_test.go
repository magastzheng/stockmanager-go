package industrymanager_test

import (
    "testing"
    "manager/industrymanager"
)

func Test_IndustryManager(t *testing.T) {
    m := industrymanager.NewIndustryManager()
    m.Process()
}


