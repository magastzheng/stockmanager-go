package stockhandler

import (
    //"fmt"
    "strings"
    "util"
)

const ( 
    HistDivClass = "tagmain"
    HistTableId = "FundHoldSharesTable"
)

type StockHistData struct {
    Date string
    Open float32
    Highest float32
    Close float32
    Lowest float32
    Volume int
    Money int
}

type StockHistDataHandler struct {
    Code string
    Data [] StockHistData
    isTargetDiv bool
    isTargetTable bool
    isTargetHead bool
    isTargetHeadTr bool
    isTargetHeadTh bool
    isTargetTr bool
    isTargetTd bool
    isTargetTdDiv bool
    isTargetTdAnchor bool
    targetTrNum int
    targetTdNum int
    tempData StockHistData
}

func (h *StockHistDataHandler) Init() {
    h.Data = make([]StockHistData, 10, 92)
    h.isTargetDiv = false
    h.isTargetTable = false
    h.isTargetHead = false
    h.isTargetHeadTr = false
    h.isTargetHeadTh = false
    h.isTargetTr = false
    h.isTargetTd = false
    h.isTargetTdDiv = false
    h.isTargetTdAnchor = false
    h.targetTrNum = 0
    h.targetTdNum = 0
}

func (h *StockHistDataHandler) OnStartElement(tag string, attrs map[string]string){
    switch tag {
        case "div":
            var classname string
            var ok bool
            classname, ok = attrs["class"]
            if ok && classname == HistDivClass {
                h.isTargetDiv = true
            } else if h.isTargetTd {
                h.isTargetTdDiv = true
            }
        case "table":
            if h.isTargetDiv {
                var id string
                var ok bool
                id, ok = attrs["id"]
                if ok && id == HistTableId {
                    h.isTargetTable = true
                }
            }
        case "thead":
            if h.isTargetTable {
                h.isTargetHead = true
            }
        case "tr":
            if h.isTargetTable {
                if h.isTargetHead {
                    h.isTargetHeadTr = true
                } else {
                    h.targetTrNum++
                    h.isTargetTr = true
                    
                    if h.targetTrNum > 1 {
                        h.tempData = StockHistData{}
                    }
                }
            }
        case "th":
            if h.isTargetHeadTr {
                h.isTargetHeadTh = true
            } 
        case "td":
            if h.isTargetTr {
                h.isTargetTd = true
                h.targetTdNum++
            }
        case "a":
            if h.isTargetTdDiv {
                h.isTargetTdAnchor = true
            }
    }
}

func (h *StockHistDataHandler) OnEndElement(tag string) {
    switch tag {
        case "div":
            if h.isTargetDiv && !h.isTargetTable {
                h.isTargetDiv = false
            } else if h.isTargetTd {
                h.isTargetTdDiv = false
            }
        case "table":
            if h.isTargetDiv && h.isTargetTable {
                h.isTargetTable = false
                //reset the readed <TR> number
                h.targetTrNum = 0
            }
        case "thead":
            if h.isTargetTable && h.isTargetHead {
                h.isTargetHead = false
            }
        case "th":
            if h.isTargetHead && h.isTargetHeadTh {
                h.isTargetHeadTh = false
            }
        case "tr":
            if h.isTargetTable {
                if h.isTargetHead && h.isTargetHeadTr {
                    h.isTargetHeadTr = false
                } else if h.isTargetTr {
                    h.isTargetTr = false
                    //reset the readed <TD> number
                    h.targetTdNum = 0
                    if h.targetTrNum > 1 {
                        h.Data = append(h.Data, h.tempData)
                    }
                }
            }
        case "td":
            if h.isTargetTr && h.isTargetTd {
                h.isTargetTd = false
            }
        case "a": 
            if h.isTargetTdDiv && h.isTargetTdAnchor {
                h.isTargetTdAnchor = false    
            }
    }
}

func (h *StockHistDataHandler) OnText(text string) {
    if h.isTargetTd && h.isTargetTdDiv && h.isTargetTdAnchor {
        h.tempData.Date = strings.TrimSpace(text)
    } else if h.isTargetTd && h.isTargetTdDiv && h.targetTrNum > 1 {
        switch h.targetTdNum {
            case 1:
                //tempData.
            case 2:
                h.tempData.Open = util.ToFloat32(text)
            case 3:
                h.tempData.Highest = util.ToFloat32(text)
            case 4:
                h.tempData.Close = util.ToFloat32(text)
            case 5:
                h.tempData.Lowest = util.ToFloat32(text)
            case 6:
                h.tempData.Volume = util.ToInt(text)
            case 7:
                h.tempData.Money = util.ToInt(text)
        } 
    } else if h.isTargetHeadTh {
        code := h.GetCode(text)
        if util.IsStringNotEmpty(code){        
            h.Code = code
        }
    }
}

func (h *StockHistDataHandler) OnComment(text string) {
    //fmt.Println("Comment: ", text)
} 

func (h *StockHistDataHandler) OnPIElement(tag string, attrs map[string]string) {

}

func (h *StockHistDataHandler) OnCData(text string){
    //fmt.Println("CData: ", text)
}

func (h *StockHistDataHandler) OnError(line int, row int, message string) {
    //do nothing
}

func (h *StockHistDataHandler) GetCode(text string) string {
    start := strings.IndexRune(text, '(')
    end := strings.IndexRune(text, ')')

    var code string
    if end > start && start < len(text) && end < len(text) {
        code =  string(text[start+1: end])
    }

    return code
}

func NewStockHistDataHandler() *StockHistDataHandler {
    s := new(StockHistDataHandler)
    s.Init()

    return s
}
