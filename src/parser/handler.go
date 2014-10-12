package parser

import (
    "fmt"
)

type HtmlHandler struct {
    
}

func (h * HtmlHandler) OnStartElement(tag string, attrs map[string]string) {
    str := "StartEl: " + tag
    for k, v := range attrs {
        str += " " + k + "=" + v
    }
    fmt.Println(str)
}
    
func (h *HtmlHandler) OnEndElement(tag string) {
    fmt.Println("EndEl: ", tag)
}

func (h *HtmlHandler) OnText(text string) {
    fmt.Println("Text: ", text)
}
func (h *HtmlHandler) OnComment(text string) {
    fmt.Println("Comment: ", text)
} 
func (h *HtmlHandler) OnPIElement(tag string, attrs map[string]string) {
    //fmt.Println("PI: ", tag)

    //for k, v := range attrs {
     //   fmt.Println(k, v)
    //}
}

func (h *HtmlHandler) OnCData(text string){
    //fmt.Println("CData: ", text)
}

func (h *HtmlHandler) OnError(line int, row int, message string) {
    //do nothing
}
