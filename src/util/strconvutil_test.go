package util_test

import (
    "testing"
    "util"
)

func Test_ToInt(t *testing.T) {
    v := util.ToInt("1145")
    if v == 1145 {
        t.Log("success")
    } else {
        t.Error("Fail")
    }

    v = util.ToInt("-502")
    if v == -502 {
        t.Log("success")
    } else {
        t.Error("Fail")
    }
    
    v = util.ToInt("ds123")
}

func Test_ToFloat32(t *testing.T) {
    v := util.ToFloat32("3.26587")
    if v > 3.1658 && v < 3.1659 {
        t.Log("success")
    }
}
