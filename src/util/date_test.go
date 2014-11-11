package util_test

import(
    "testing"
    "util"
    "fmt"
)

func Test_ParseDate(t *testing.T){
    l := "201301"
    d := util.ParseDate(l)
    fmt.Println(d)
    if d.Year() == 2013 && int(d.Month()) == 1 && d.Day() == 31 {
        t.Log("Success to parse: ", l)
    } else {
        t.Error("Cannot parse: ", l)
    }
    
    l = "190002"
    d = util.ParseDate(l)
    fmt.Println(d)
    if d.Year() == 1900 && int(d.Month()) == 2 && d.Day() == 28 {
        t.Log("Success to parse: ", l)
    } else {
        t.Error("Cannot parse: ", l)
    }

    l = "200002"
    d = util.ParseDate(l)
    fmt.Println(d)
    if d.Year() == 2000 && int(d.Month()) == 2 && d.Day() == 29 {
        t.Log("Success to parse: ", l)
    } else {
        t.Error("Cannot parse: ", l)
    }

    l = "200402"
    d = util.ParseDate(l)
    fmt.Println(d)
    if d.Year() == 2004 && int(d.Month()) == 2 && d.Day() == 29 {
        t.Log("Success to parse: ", l)
    } else {
        t.Error("Cannot parse: ", l)
    }
}

func Test_IsLeapYear(t *testing.T){
    year := 1900
    ret := util.IsLeapYear(year)
    if !ret {
        t.Log("Success to judge the leap year: ", year)
    } else {
        t.Error("Cannot judge the leap year: ", year)
    }

    year = 2000
    ret = util.IsLeapYear(year)
    if ret {
        t.Log("Success to judge the leap year: ", year)
    } else {
        t.Error("Cannot judge the leap year: ", year)
    }

    year = 2008
    ret = util.IsLeapYear(year)
    if ret {
        t.Log("Success to judge the leap year: ", year)
    } else {
        t.Error("Cannot judge the leap year: ", year)
    }

    year = 1993
    ret = util.IsLeapYear(year)
    if !ret {
        t.Log("Success to judge the leap year: ", year)
    } else {
        t.Error("Cannot judge the leap year: ", year)
    }
}

func Test_LastDay(t *testing.T){
    y := 2001
    m := 2
    ret := util.LastDay(y, m)
    
    if ret == 28 {
        t.Log("Success to get the right last day: ", y, m, ret)
    } else {
        t.Error("Wrong to get the last day: ", y, m, ret)
    }

    y = 2000
    m = 2
    ret = util.LastDay(y, m)
    
    if ret == 29 {
        t.Log("Success to get the right last day: ", y, m, ret)
    } else {
        t.Error("Wrong to get the last day: ", y, m, ret)
    }
}

