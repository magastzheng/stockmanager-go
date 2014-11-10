package manager_test

import (
    "testing"
    "manager"
)

func Test_Industry(t *testing.T) {
    m := manager.NewIndustryManager()
    m.Process()
}


