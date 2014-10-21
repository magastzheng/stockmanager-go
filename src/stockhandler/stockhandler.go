//Parse the stock code and name from the HMTL page
package stockhandler

import(
    "fmt"
    "strings"
)

type Stock struct{
    Id string
    Name string
    Website string
}

type StockHandler struct{
    Stocks map [string] Stock
    isTargetDiv bool
    isStockLi bool
    isStockLink bool
    tempStock Stock
}

func (h *StockHandler) Init() {
    h.Stocks = make(map [string] Stock)
    h.isTargetDiv = false
    h.isStockLi = false
    h.isStockLink = false
    //h.tempStock = nil
}

func (h *StockHandler) OnStartElement(tag string, attrs map[string]string){
    //fmt.Println("Start: ", tag)
    if tag == "div" {
        //fmt.Println("Start: ", tag)
        var classname string
        var ok bool
        classname, ok = attrs["class"]
        if ok && classname == "result" {
            h.isTargetDiv = true
        }

        //fmt.Println("attrs: ", len(attrs), ok, h.isTargetDiv)
    } else if h.isTargetDiv && tag == "li" {
        h.isStockLi = true
        //fmt.Println("tag: ", tag, h.isStockLi)
    } else if h.isStockLi && tag == "a" {
        h.isStockLink = true
        //fmt.Println("start: ", tag)
        h.tempStock = Stock{Website: attrs["href"]}
    } else {
        //do nothing
    }
}

func (h *StockHandler) OnEndElement(tag string) {
    if tag == "div" && h.isTargetDiv {
        h.isTargetDiv = false
    } else if h.isTargetDiv && tag == "li" && h.isStockLi {
        h.isStockLi = false
    } else if h.isStockLi && tag == "a" && h.isStockLink{
        h.isStockLink = false
        h.Stocks[h.tempStock.Id] = h.tempStock
    }
}

func (h *StockHandler) OnText(text string) {
    if h.isStockLink {
        h.tempStock.Name, h.tempStock.Id = h.Split(text)
    }
}
func (h *StockHandler) OnComment(text string) {
    //fmt.Println("Comment: ", text)
} 
func (h *StockHandler) OnPIElement(tag string, attrs map[string]string) {

}

func (h *StockHandler) OnCData(text string){
    //fmt.Println("CData: ", text)
}

func (h *StockHandler) OnError(line int, row int, message string) {
    //do nothing
}

func (h *StockHandler) Split(text string) (string, string) {
    text = strings.TrimSpace(text)
    st := strings.Split(text, "(")
    var name string
    var id string
    if len(st) == 2 {
        name = st[0]
        id = strings.TrimRight(st[1], ")")
    } else {
        //fmt.Println("Parse name and id wrong")
        name = ""
        id = ""
    }

    return name, id
}

func (h *StockHandler) PrintStocks(){
    for k, st := range h.Stocks {
        fmt.Println("key: ", k, "Id: ", st.Id, " name: ", st.Name, " website: ", st.Website)
    }
}

func NewStockHandler() *StockHandler {
    s := new(StockHandler)
    s.Init()

    return s
}

