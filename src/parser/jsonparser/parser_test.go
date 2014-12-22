package jsonparser_test

import(
    "testing"
    parser "parser/jsonparser"
    "io/ioutil"
    "fmt"
)

/*
type TestHandler struct {
    
}

func (h *TestHandler) OnObject(key string, keyValues map[string]string) {

}

func (h *TestHandler) OnArray(key string, elems []string) {
    
}
*/

func readFile(filename string) string {
    chunks, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    return string(chunks)
}

/*
func Test_JsonParser_ParseArray(t *testing.T){
    filename := "../../resource/json/testarr.json"
    str := readFile(filename)
    h := new(parser.JsonHandler)
    p := parser.NewJsonParser(h)
    p.ParseStr(str)
}

func Test_JsonParser_ParseArray_Str(t *testing.T){
    filename := "../../resource/json/teststrarr.json"
    str := readFile(filename)
    h := new(parser.JsonHandler)
    p := parser.NewJsonParser(h)
    p.ParseStr(str)
}

func Test_JsonParser_ParseObject(t *testing.T){
    filename := "../../resource/json/testobj.json"
    str := readFile(filename)
    h := new(parser.JsonHandler)
    p := parser.NewJsonParser(h)
    p.ParseStr(str)
}

func Test_JsonParser_ParseMix(t *testing.T){
    filename := "../../resource/json/test.json"
    str := readFile(filename)
    h := new(parser.JsonHandler)
    p := parser.NewJsonParser(h)
    p.ParseStr(str)
}*/

func Test_JsonParser_ParseShSeCompany(t *testing.T){
    filename := "../../resource/json/shsecompany.json"
    str := readFile(filename)
    h := new(parser.JsonHandler)
    p := parser.NewJsonParser(h)
    p.ParseStr(str)
}
