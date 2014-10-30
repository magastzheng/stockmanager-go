package util

import (
    "os"
)

func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}

func IsStringNotEmpty(s string) bool {
    // s != ""
    return len(s) > 0
}

func WriteFile(filename string, content string) {
    file, err := os.Create(filename)
    CheckError(err)
    defer file.Close()
    file.WriteString(content)
}
