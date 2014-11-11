package util

import (
    "strconv"
    "math"
    //"fmt"
    "strings"
)

func processToValid(s string) string {
    s = strings.TrimSpace(s)
    s = strings.Replace(s, ",", "", -1)
    return s
}

func GetIntStr(s string) string {
    s = processToValid(s)
    if strings.Contains(s, ".") {
        s = strings.Split(s, ".")[0]
    }

    return s
}

func ToInt(s string) int {
    s = GetIntStr(s)
    i, err := strconv.Atoi(s)
    if err != nil {
        //fmt.Println(err)
        NewLog().Error("Cannot convert to int type:", s)
    }

    return i
}

func ToInt64(s string) int64 {
    s = GetIntStr(s)
    i, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        //fmt.Println(err)
        NewLog().Error("Cannot convert to int64 type:", s)
    }

    return i
}

func ToFloat32(s string) float32 {
    s = processToValid(s)
    f, err := strconv.ParseFloat(s, 32)
    
    if err != nil {
        //fmt.Println(err)
        NewLog().Error("Cannot convert to float32 type:", s)
        f = math.NaN()
    }

    return float32(f)
}

func ToFloat64(s string) float64 {
    s = processToValid(s)
    f, err := strconv.ParseFloat(s, 64)
    
    if err != nil {
        //fmt.Println(err)
        NewLog().Error("Cannot convert to float64 type:", s)
        f = math.NaN()
    }

    return f
}
