package parser

import (
    "fmt"
)

type HtmlHandler struct {
    
}

func (h * HtmlHandler) OnStartElement(tag string, attrs map[string]string) {
    fmt.Println("Start el: ", tag)
    for k, v := range attrs {
        fmt.Println(k, v)
    }
}
    
func (h *HtmlHandler) OnEndElement(tag string) {
    fmt.Println("End el: ", tag)
}

func (h *HtmlHandler) OnText(text string) {
    fmt.Println("Text: ", text)
}
func (h *HtmlHandler) OnComment(text string) {
    fmt.Println("Comment: ", text)
} 
func (h *HtmlHandler) OnPIElement(tag string, attrs map[string]string) {
    fmt.Println("PI: ", tag)

    for k, v := range attrs {
        fmt.Println(k, v)
    }
}

func (h *HtmlHandler) OnCData(text string){
    fmt.Println("CData: ", text)
}

func (h *HtmlHandler) OnError(line int, row int, message string) {
    //do nothing
}
