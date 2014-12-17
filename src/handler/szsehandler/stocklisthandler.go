package szsehandler

import(
    "handler"
    "strings"
	"entity/stentity"
    "util"
    "time"
    //"fmt"
)

const(
    STTabClass = "cls-data-table"
    STTrHead = "cls-data-tr-head"
    STTrData = "cls-data-tr"
    STDateFormat = "2006-01-02"
)

type StockListHandler struct {
    handler.HandlerBase
    Companies []stentity.Company
    tempComp stentity.Company

    isTargetTab bool
    isTargetTrHead bool
    isTargetTr bool
    isTargetTdHead bool
    isTargetTd bool
    tdHeadNum int
    tdNum int
}

func (h *StockListHandler) Init() {
    h.Companies = make([]stentity.Company, 0)
    
    h.isTargetTab = false
    h.isTargetTrHead = false
    h.isTargetTr = false
    h.isTargetTdHead = false
    h.isTargetTd = false
    h.tdHeadNum = 0
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
                if ok && clsname == STTrHead {
                    h.isTargetTrHead = true
                } else {   
                    h.isTargetTr = true
                    h.tempComp = stentity.Company{}
                }
            }
        case "td":
            if h.isTargetTr {
                h.isTargetTd = true
                h.tdNum++
            } else if h.isTargetTrHead {
                h.isTargetTdHead = true
                h.tdHeadNum++
            }
    }
}
    
func (h *StockListHandler) OnEndElement(tag string) {
    switch tag {
        case "table":
            if h.isTargetTab {
                h.isTargetTab = false
            }
        case "tr":
            if h.isTargetTab {
                if h.isTargetTr {
                    h.isTargetTr = false
                    h.Companies = append(h.Companies, h.tempComp)
                    h.tdNum = 0
                } else if h.isTargetTrHead {
                    h.isTargetTrHead = false
                    h.tdHeadNum = 0
                }
            }
        case "td":
            if h.isTargetTr && h.isTargetTd {
                h.isTargetTd = false
            } else if h.isTargetTrHead && h.isTargetTdHead {
                h.isTargetTdHead = false
            }
    }
}

func (h *StockListHandler) OnText(text string) {
    if h.isTargetTd {
        text = strings.TrimSpace(text)
        switch h.tdNum {
            case 1:
                h.tempComp.Code = text
            case 2:
                h.tempComp.AbbrName = text
            case 3:
                h.tempComp.Name = text
            case 4:
                h.tempComp.Name_en = text
            case 5:
                h.tempComp.RegAddr = text
            case 6:
                //do nothing
            case 7:
                //do nothing
            case 8:
                h.tempComp.InceptDate = h.GetDate(text)
            case 9:
                h.tempComp.TotalShares_A = h.GetInt(text)
            case 10:
                h.tempComp.FlowShares_A = h.GetInt(text)
            case 11:
                h.tempComp.Code_B = text
            case 12:
                //B share abbr name
            case 13:
                h.tempComp.InceptDate_B = h.GetDate(text)
            case 14:
                h.tempComp.TotalShares_B = h.GetInt(text)
            case 15:
                h.tempComp.FlowShares_B = h.GetInt(text)
            case 16:
                h.tempComp.Region = text
            case 17:
                h.tempComp.State = text
            case 18:
                h.tempComp.City = text
            case 19:
                h.tempComp.Industry = text
            case 20:
                h.tempComp.Website = text
        }
    } 
}

func (h *StockListHandler) GetDate(text string) string {
    if len(text) == 0 {
        return ""
    }
    
    t, err := time.Parse(STDateFormat, text)
    if err != nil {
        util.NewLog().Error("Cannot parse the text to date: ", text)
        return ""
    }

    return util.FormatDate(t)
}

func (h *StockListHandler) Output() {
    content := ""
    for _, c := range h.Companies {
        //fmt.Println(c.ToString())
        content += c.ToString() + "\n"
    }

    util.WriteFile("Output.txt", content)
}

func (h *StockListHandler) GetInt(text string) int {
    if len(text) == 0{
        return 0
    }

    s := strings.Replace(text, ",", "", -1)
    return util.ToInt(s)
}

func NewStockListHandler() *StockListHandler{
    h := new(StockListHandler)
    h.Init()

    return h
}
