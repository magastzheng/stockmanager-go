package util_test

import (
    "testing"
    "util"
)

func Test_IsEqual(t *testing.T) {
    res := util.IsEqual(3.1426, 3.142601)
    if res {
       t.Error("Fail") 
    }
}
