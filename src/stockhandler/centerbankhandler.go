package stockhandler

import(
    "parser"
)

type CenterBankHandler struct {
    Name string
    Months [] string
    Data map [string][]float
    isTargetTable
    isTargetTr
    isTargetTd
}

func (h *CenterBankHandler) Init() {
    h.months := make([]string, 0, 12)
    h.Data := make(map[string][]float)
    h.isTargetTable := false
    h.isTargetTr := false
    h.isTargetTd := false
}

func (h *CenterBankHandler) OnStartElement(tag string, attrs map[string]string){
    switch tag {
        case "table":
            h.isTargetTable = true
        case "tr":
            
        case "td"
    }
}

func (h *CenterBankHandler) OnEndElement(tag string) {
    switch tag {
        case "table":
            if h.isTargetTable {
                h.isTargetTable = false
            }
        case "tr":
            if h.isTargetTr{
                h.isTargetTr = false
            }
        case "td":
            if h.isTargetTd {
                h.isTargetTd = false
            }
    }

}

func (h *CenterBankHandler) OnText(text string) {
 
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

func NewCenterBankHandler() *CenterBankHandler {
    s := new(CenterBankHandler)
    s.Init()

    return s
}
