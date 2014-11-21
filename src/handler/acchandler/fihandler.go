package acchandler

import(
    "strings"
    "util"
    //"fmt"
)

const(
    FIDivClass = "tagmain"
    FITabId = "BalanceSheetNewTable0"
)

type FiHandler struct {
    DateMap map[string]string
    DataMap map[string]map[string]float32
    isTargetDiv bool
    isTargetTab bool
    isTargetDateTab bool
    isTargetTBody bool
    isTargetDataTr bool
    isTargetDateTr bool
    isTargetDataTd bool
    isTargetDateTd bool
    isTargetStrong bool
    isTargetAnch bool
    ignoreTr bool
    //tabNum int
    trNum int
    tdNum int
    tempUrl string

    reportDates []string
    tempRowData []float32
    tempRowId string
}

func (h *FiHandler) Init() {
    h.DateMap = make(map[string]string)
    h.DataMap = make(map[string]map[string]float32)
    h.isTargetDiv = false
    h.isTargetTab = false
    h.isTargetDateTab = false
    h.isTargetTBody = false
    h.isTargetDataTr = false
    h.isTargetDateTr = false
    h.isTargetDataTd = false
    h.isTargetDateTd = false
    h.isTargetStrong = false
    h.isTargetAnch = false
    h.ignoreTr = false
    
    //h.tabNum = 0
    h.trNum = 0
    h.tdNum = 0
    h.tempUrl = ""
    h.reportDates = make([]string, 0)
    //h.tempRowData = make([]float32, 0)
    h.tempRowId = ""
}

func (h * FiHandler) OnStartElement(tag string, attrs map[string]string) {
    //fmt.Println(tag, attrs)
    switch tag {
        case "div":
            clsname, ok := attrs["class"]
            if ok && clsname == FIDivClass {
                h.isTargetDiv = true
                //fmt.Println(tag, attrs)
            }
        case "table":
            if h.isTargetDiv {
                tabId, ok := attrs["id"]
                _, ok2 := attrs["class"]

                if ok && tabId == FITabId {
                    h.isTargetTab = true
                } else if !ok2{
                    h.isTargetDateTab = true
                }
            }
        case "tbody":
            if h.isTargetTab {
                h.isTargetTBody = true
            }
        case "tr":
            if h.isTargetTBody {
                h.isTargetDataTr = true
                h.trNum++
                h.tempRowData = make([]float32, len(h.reportDates))
            } else if h.isTargetDateTab {
                h.isTargetDateTr = true
                h.trNum++
            } else {
                //do nothing
            }
        case "td":
            if h.isTargetDataTr {
                h.isTargetDataTd = true
                h.tdNum++

                cspan, ok := attrs["colspan"]
                if ok && len(cspan) > 0 {
                    h.ignoreTr = true
                }
            } else if h.isTargetDateTab || h.isTargetDateTr {
                h.isTargetDateTd = true
                h.tdNum++
            } else {
                //do nothing
            }
        case "strong":
            if h.isTargetDataTd || h.isTargetDateTd {
                h.isTargetStrong = true
            }
        case "a":
            if h.isTargetDataTd || h.isTargetDateTd {
                h.isTargetAnch = true
                h.tempUrl = attrs["href"]
            }
    }
}
    
func (h *FiHandler) OnEndElement(tag string) {
    switch tag {
        case "div":
            if h.isTargetDiv {
               h.isTargetDiv = false 
            }
        case "table":
            if h.isTargetDiv {
                if h.isTargetDateTab {
                    h.isTargetDateTab = false
                } else if h.isTargetTab {
                    h.isTargetTab = false
                    h.trNum = 0
                } else {
                    //do nothing
                }
            }
        case "tbody":
            if h.isTargetTab && h.isTargetTBody{
                h.trNum = 0
            }
        case "tr":
            if h.isTargetTBody && h.isTargetDataTr {
                h.isTargetDataTr = false
                h.tdNum = 0
                if !h.ignoreTr && h.trNum > 1 {
                    h.AddData()
                }
                
                if h.ignoreTr {
                    h.ignoreTr = false
                }
            } else if h.isTargetDateTab && h.isTargetDateTr {
                h.isTargetDateTr = false
                h.tdNum = 0
            }
        case "td":
            if h.isTargetDataTr && h.isTargetDataTd {
                h.isTargetDataTd = false
            } else if (h.isTargetDateTab || h.isTargetDateTr) && h.isTargetDateTd {
                h.isTargetDateTd = false
            }
        case "strong":
            if (h.isTargetDataTd || h.isTargetDateTd) && h.isTargetStrong {
                h.isTargetStrong = false
            }
        case "a":
            if (h.isTargetDataTd || h.isTargetDateTd) && h.isTargetAnch {
                h.isTargetAnch = false
                h.tempUrl = ""
            }
    }
}

func (h *FiHandler) OnText(text string) {
    if h.isTargetDateTd && h.isTargetAnch {
        h.DateMap[text] = h.tempUrl
    } else if h.isTargetDataTd {
        text = strings.TrimSpace(text)
        if h.trNum == 1 {
            //fmt.Println(text)
            h.reportDates = append(h.reportDates, text)
        } else if !h.ignoreTr {
            if h.tdNum == 1 {
                h.tempRowId = text
            } else if h.tdNum > 1{
                value := util.ToFloat32(text)
                h.tempRowData[h.tdNum - 1] = value
            } else {
                //do nothing
            }
        }
    }
}

func (h *FiHandler) OnComment(text string) {
    //fmt.Println("Comment: ", text)
}

func (h *FiHandler) OnPIElement(tag string, attrs map[string]string) {
    //fmt.Println("PI: ", tag)

    //for k, v := range attrs {
     //   fmt.Println(k, v)
    //}
}

func (h *FiHandler) OnCData(text string){
    //fmt.Println("CData: ", text)
}

func (h *FiHandler) OnError(line int, row int, message string) {
    //do nothing
}

func (h *FiHandler) AddData() {
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

func NewFiHandler() *FiHandler{
    h := new(FiHandler)
    h.Init()

    return h
}
