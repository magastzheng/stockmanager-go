package util

import (
    "strconv"
    "math"
    "fmt"
    "strings"
)

func ToInt(s string) int {
    s = strings.TrimSpace(s)
    i, err := strconv.Atoi(s)
    if err != nil {
        fmt.Println(err)
    }

    return i
}

func ToInt64(s string) int64 {
    s = strings.TrimSpace(s)
    i, err := strconv.ParseInt(s, 10, 64)
    if err != nil {
        fmt.Println(err)
    }

    return i
}

func ToFloat32(s string) float32 {
    s = strings.TrimSpace(s)
    f, err := strconv.ParseFloat(s, 32)
    
    if err != nil {
        fmt.Println(err)
        f = math.NaN()
    }

    return float32(f)
}

func ToFloat64(s string) float64 {
    s = strings.TrimSpace(s)
    f, err := strconv.ParseFloat(s, 64)
    
    if err != nil {
        fmt.Println(err)
        f = math.NaN()
    }

    return f
}
