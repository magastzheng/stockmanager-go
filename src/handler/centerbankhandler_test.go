package handler_test

import(
    "parser"
	"handler"
	"testing"
	"io/ioutil"
	"fmt"
	"util"
)

func Test_CenterBankHandler(t *testing.T) {
	filename := "../resource/afre-2014.html"
	chunks, err := ioutil.ReadFile(filename)
	util.CheckError(err)

	str := string(chunks)
	h := handler.NewCenterBankHandler()

	parser := parser.NewTextParser(h)
	parser.ParseStr(str)
	//fmt.Println(h.Data)
	fmt.Println(len(h.Data))
	h.Output();
}
