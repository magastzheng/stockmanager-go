package szsehandler

import(
    "handler"
    "strings"
    "util"
    //"fmt"
)

const(
    STTabClass = "cls-data-table"
    STTrHead = "cls-data-tr-head"
    STTrData = "cls-data-tr"
)

type Company struct {
    Code string
    AbbrName string
    Name string
    Name_en string
    RegAddr string
    InceptDate string
    TotalShares_A int
    FlowShares_A int
    Code_B string
    InceptDate_B string
    TotalShares_B int
    FlowShares_B int
    Region string
    State string
    City string
    Industry string
    Website string
}

type StockListHandler struct {
    handler.HandlerBase
    Companies []Company
    tempComp Company

    isTargetTab bool
    isTargetTr bool
    isTargetTd bool
    tdNum int
}

func (h *StockListHandler) Init() {
    h.Companies = make([]Company, 0)
    
    h.isTargetTab = false
    h.isTargetTr = false
    h.isTargetTd = false
    h.tdNum = 0
}

func (h * StockListHandler) OnStartElement(tag string, attrs map[string]string) {
    clsname, ok := attrs["class"]

    switch tag {
        case "table":
            if ok && clsname == STTabClass {
                h.isTargetTab = true
            }
        case "tr":
            if h.isTargetTab {
                if ok && clsname == STTrData {
                    h.isTargetTr = true
                    h.tempComp = Company{}
                }
            }
        case "td":
            if h.isTargetTr {
                h.isTargetTd = true
                h.tdNum++
            }
    }
}
    
func (h *StockListHandler) OnEndElement(tag string) {
    switch tag {
        case "table":
            if h.isTargetTable {
                h.isTargetTable = false
            }
        case "tr":
            if h.isTargetTable && h.isTargetTr {
                h.isTargetTr = false
                h.Companies = append(h.Companies, h.tempComp)
                h.tdNum = 0
            }
        case "td":
            if h.isTargetTr && h.isTargetTd {
                h.isTargetTd = false
            }
    }
}

func (h *StockListHandler) OnText(text string) {
    if h.isTargetTd {
        
    } 
}

func (h *StockListHandler) AddData() {
    //fmt.Println(len(h.tempRowData), h.tempRowData)
    //fmt.Println(len(h.reportDates))
    for i, col := range h.reportDates {
        if i > 0 {
            val := h.tempRowData[i]
            dataMap, ok := h.DataMap[col]
            if ok {
                if dataMap == nil {
                    dataMap = make(map[string]float32)
                }

                dataMap[h.tempRowId] = val
            } else {
                dataMap := make(map[string]float32)
                dataMap[h.tempRowId] = val
            }

            h.DataMap[col] = dataMap
        } else {
            //fmt.Println(i, col)
        }
    }
}

func NewStockListHandler() *StockListHandler{
    h := new(StockListHandler)
    h.Init()

    return h
}
