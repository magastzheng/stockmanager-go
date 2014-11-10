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

    v = util.ToInt("123.456")
    if v == 123 {
        t.Log("success")
    } else {
        t.Error("Fail")
    }
}

func Test_ToFloat32(t *testing.T) {
    v := util.ToFloat32("3.26587")
    if v > 3.1658 && v < 3.1659 {
        t.Log("success")
    }

    v = util.ToFloat32("312")
    if v > 311.0 && v < 313.0 {
        t.Log("Parse the str 312 to float success")
    } else {
        t.Error("Fail to parse 312")
    }
}
