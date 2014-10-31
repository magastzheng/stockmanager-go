package util

import (
    "math"
)

const (
    FloatPrecise = 0.000000001
)

func IsEqual(f1, f2 float64) bool {
    if f1 > f2 {
        return math.Dim(f1, f2) < FloatPrecise
    } else {
        return math.Dim(f2, f1) < FloatPrecise
    }
}
