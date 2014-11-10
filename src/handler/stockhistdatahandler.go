package handler

import (
    //"fmt"
    "strings"
    "util"
    "entity"
)

const ( 
    HistDivClass = "tagmain"
    HistTableId = "FundHoldSharesTable"
    HistFormName = "daily"
    HistSelectName = "year"
)

type StockHistDataHandler struct {
    Code string
    Data [] entity.StockHistData
    Years [] int
    isTargetForm bool
    isTargetSelectYear bool
    isTargetOptionYear bool
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
    tempData entity.StockHistData
}

func (h *StockHistDataHandler) Init() {
    h.Data = make([]entity.StockHistData, 0, 92)
    h.Years = make([]int, 0, 30)
    h.isTargetForm = false
    h.isTargetSelectYear = false
    h.isTargetOptionYear = false
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
        case "form":
            if h.isTargetDiv {
                var name string
                var ok bool
                name, ok = attrs["name"]
                if ok && name == HistFormName {
                    h.isTargetForm = true
                }
            }
        case "select":
            if h.isTargetForm {
                var name string
                var ok bool
                name, ok = attrs["name"]
                if ok && name == HistSelectName {
                    h.isTargetSelectYear = true
                }
            }
        case "option":
            if h.isTargetSelectYear {
                h.isTargetOptionYear = true
                var value string
                var ok bool
                value, ok = attrs["value"]
                if ok && util.IsStringNotEmpty(value) {
                    year := util.ToInt(value)
                    h.Years = append(h.Years, year)
                }
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
                        h.tempData = entity.StockHistData{}
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
        case "form":
            if h.isTargetDiv && h.isTargetForm {
                h.isTargetForm = false
            }
        case "select":
            if h.isTargetForm && h.isTargetSelectYear {
                h.isTargetSelectYear = false
            }
        case "option":
            if h.isTargetSelectYear && h.isTargetOptionYear {
                h.isTargetOptionYear = false
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
                        //if len(h.tempData.Date) == 0 {
                        //    fmt.Println(h.tempData)
                        //}
                        //o := fmt.Sprintf("Each row: %#v", h.tempData)
                        //fmt.Println(o)
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
    //if h.isTargetTd && h.isTargetTdDiv {
    //    h.tempData.Date = strings.TrimSpace(text)
    //    if len(h.tempData.Date) == 0 && h.isTargetTdAnchor {
    //        h.tempData.Date = strings.TrimSpace(text)
    //    }
    //} else 
    if h.isTargetTd && h.isTargetTdDiv && h.targetTrNum > 1 {
        switch h.targetTdNum {
            case 1:
                date := strings.TrimSpace(text)
                if len(h.tempData.Date) == 0 {
                    if len(date) > 0 {
                        h.tempData.Date = date
                    }
                }
                //if len(h.tempData.Date) == 0 && h.isTargetTdAnchor {
                    //fmt.Println("Before", h.tempData.Date)
                //    h.tempData.Date = strings.TrimSpace(text)
                    //fmt.Println("After", h.tempData.Date)
                //}
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
