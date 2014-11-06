package stockhandler_test

import(
    "parser"
	"stockhandler"
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
	handler := stockhandler.NewCenterBankHandler()

	parser := parser.NewTextParser(handler)
	parser.ParseStr(str)
	//fmt.Println(handler.Data)
	fmt.Println(len(handler.Data))
	handler.Output();
}
