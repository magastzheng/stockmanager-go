package util

import(
    "os"
)

func WriteFile(filename string, content string) {
    file, err := os.Create(filename)
    CheckError(err)
    defer file.Close()
    file.WriteString(content)
}
