package parser_test

import(
    "parser"
	"testing"
	"io/ioutil"
	//"fmt"
	"util"
)

func Test_CenterBankHandler(t *testing.T) {
	filename := "../resource/afre-2014.html"
	chunks, err := ioutil.ReadFile(filename)
	util.CheckError(err)

	str := string(chunks)
	handler := parser.NewCBHandler()

	parser := parser.NewTextParser(handler)
	parser.ParseStr(str)
	//fmt.Println(handler.Data)
	handler.Output();
}
