package parser

import(
    //"parser"
	"strings"
	"fmt"
)

type CBHandler struct {
	trNum int
	tdNum int
}

func (h *CBHandler) Init() {
	h.trNum = 0
	h.tdNum = 0
}

func (h *CBHandler) OnStartElement(tag string, attrs map[string]string){
	if tag == "tr" {
		h.trNum++
	}
	str := tag
	for k, v := range attrs {
		str += fmt.Sprintf(" %v:%v", k, v)
	}
	fmt.Println("start el:", str)
}

func (h *CBHandler) OnEndElement(tag string) {
	fmt.Println("End el:", tag)
}

func (h *CBHandler) OnText(text string) {
	value := strings.TrimSpace(text)
	fmt.Println("text:", value)
}

func (h *CBHandler) OnComment(text string) {
	//fmt.Println("===================Comment-start===============")
    //fmt.Println("Comment: ", text)
	//fmt.Println("===================Comment-end==================")
} 
func (h *CBHandler) OnPIElement(tag string, attrs map[string]string) {

}

func (h *CBHandler) OnCData(text string){
    //fmt.Println("CData: ", text)
}

func (h *CBHandler) OnError(line int, row int, message string) {
    //do nothing
}

func (h *CBHandler) Output() {
	fmt.Println("rows: ", h.trNum)
	//for _, row := range h.Data {
	//	var str string
	//	for k, v := range row {
	//		str += fmt.Sprintf("%v: %v    ", k, v)
	//	}
	//	fmt.Println("*******************************")
	//	fmt.Println(str)
	//}
}

func NewCBHandler() *CBHandler {
    s := new(CBHandler)
    s.Init()

    return s
}
