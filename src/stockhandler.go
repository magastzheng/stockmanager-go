//Parse the stock code and name from the HMTL page
package main

import(
    "fmt"
    "os"
    "io"
    "bytes"
    "parser"
    "strings"
)

type Stock struct{
    id string
    name string
    website string
}

type StockHandler struct{
    Stocks map [string] Stock
    isTargetDiv bool
    isStockLi bool
    isStockLink bool
    tempStock Stock
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
         fmt.Println("tag: ", tag, h.isStockLi)
    } else if h.isStockLi && tag == "a" {
        h.isStockLink = true
        fmt.Println("start: ", tag)
        h.tempStock = Stock{website: attrs["href"]}
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
        h.Stocks[h.tempStock.id] = h.tempStock
    }
}

func (h *StockHandler) OnText(text string) {
    if h.isStockLink {
        h.tempStock.name, h.tempStock.id = h.Split(text)
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
        fmt.Println("Parse name and id wrong")
        name = ""
        id = ""
    }

    return name, id
}

func main(){
    filename := "stock_a-ha.dat"
    file, err := os.Open(filename)
 
    if err != nil {
        panic(err)
    }
    defer file.Close()
    
    //content, err := ioutil.ReadFile(filename)
    //if err != nil {
    //    fmt.Println("Error to read file!")
    //}
    
    chunks := bytes.NewBuffer(nil)
    io.Copy(chunks, file)
    str := string(chunks.Bytes())
    handler := new(StockHandler)
    parser := new(parser.TextParser)
    parser.SetHandler(handler)
    parser.ParseStr(str)
    
    fmt.Println(len(handler.Stocks))

    //reader := bufio.NewReader(f)
    //parser := new(StockHandler)
    //parser.parseString(reader)
}
