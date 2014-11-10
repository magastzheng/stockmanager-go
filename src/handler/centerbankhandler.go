package handler

import(
    //"parser"
	"strings"
	"fmt"
)

type CenterBankHandler struct {
    Name string
	Data []map[int]string
    //Months [] string
    //Data map [string][]float
	row map[int] string
    isTargetTable bool
    isTargetTr bool
    isTargetTd bool
	isTargetFont bool
	trNum int
	tdNum int
}

func (h *CenterBankHandler) Init() {
    h.Data = make([]map[int]string, 0, 20)
	//h.months = make([]string, 0, 12)
    //h.Data = make(map[string][]float)
	//h.row = make([]string, 0, 12)
    h.isTargetTable = false
    h.isTargetTr = false
    h.isTargetTd = false
	h.isTargetFont = false
	h.trNum = 0
	h.tdNum = 0
}

func (h *CenterBankHandler) OnStartElement(tag string, attrs map[string]string){
    switch tag {
        case "table":
            h.isTargetTable = true
        case "tr":
            h.isTargetTr = true
			h.row = make(map[int]string)
			h.trNum++
		case "td":
			h.isTargetTd = true
			var xnum string
			var ok bool
			xnum, ok = attrs["x:num"]
			if ok {
				h.row[h.tdNum] = xnum
			}
			h.tdNum++
		case "font":
			if h.isTargetTd {
				h.isTargetFont = true
			}
    }
}

func (h *CenterBankHandler) OnEndElement(tag string) {
    switch tag {
        case "table":
            if h.isTargetTable {
                h.isTargetTable = false
				//h.trNum = 0
            }
        case "tr":
            if h.isTargetTr{
                h.isTargetTr = false
				fmt.Println("Total td: ", h.tdNum)
				h.tdNum = 0
				h.Data = append(h.Data, h.row)
				//var str string
				//for k, v := range h.row {
				//	str += fmt.Sprintf("%v:%v ", k, v)
				//}
				//fmt.Println(h.trNum, str)
            }
        case "td":
            if h.isTargetTd {
                h.isTargetTd = false
            }
    }

}

func (h *CenterBankHandler) OnText(text string) {
	value := strings.TrimSpace(text)
	if h.isTargetTr && h.isTargetTd {
		if len(value) > 0 {
			h.row[h.tdNum] = value
		} else if h.isTargetFont {
			h.row[h.tdNum] = h.row[h.tdNum] + value
		}
	}
}

func (h *CenterBankHandler) OnComment(text string) {
    //fmt.Println("Comment: ", text)
} 
func (h *CenterBankHandler) OnPIElement(tag string, attrs map[string]string) {

}

func (h *CenterBankHandler) OnCData(text string){
    //fmt.Println("CData: ", text)
}

func (h *CenterBankHandler) OnError(line int, row int, message string) {
    //do nothing
}

func (h *CenterBankHandler) Output() {
	fmt.Println("rows: ", h.trNum)
	for _, row := range h.Data {
		//var str string
        str := fmt.Sprintf("# %v #", len(row))
		for k, v := range row {
			str += fmt.Sprintf("%v: %v    ", k, v)
		}
		fmt.Println("*******************************")
		fmt.Println(str)
	}
}

func NewCenterBankHandler() *CenterBankHandler {
    s := new(CenterBankHandler)
    s.Init()

    return s
}
