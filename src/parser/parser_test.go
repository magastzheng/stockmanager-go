package parser

import (
    "testing"
    "os"
    "io"
    "bytes"
)

func Test_ParseStr(t *testing.T){
    t.Log("test start")
    filename := "../resource/stocklist_sa.dat"
    file, err := os.Open(filename)
 
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    //content, err := ioutil.ReadFile(filename)
    //if err != nil {
    //    fmt.Println("Error to read file!")
    //}
    
    chunks := bytes.NewBuffer(nil)
    io.Copy(chunks, file)
    str := string(chunks.Bytes())
    handler := new(HtmlHandler)
    //parser := new(TextParser)
    //parser.SetHandler(handler)
    parser := NewTextParser(handler)
    parser.ParseStr(str)
    
    //fmt.Println(len(handler.Stocks))
}
