package util

import(
    "os"
    "fmt"
    "strings"
)

func WriteFile(filename string, content string) {
    //file, err := os.Create(filename)
    pos := strings.LastIndex(filename, "/")
    if pos > 0 {
        filepath := filename[0: pos]
        if filepath != "." || filepath != ".." {
            os.MkdirAll(filepath, 0777)
        }
    }

    file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0660)
    defer file.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

    file.WriteString(content)
}
