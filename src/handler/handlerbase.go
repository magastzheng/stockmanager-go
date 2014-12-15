package handler

import(
    "fmt"
)

type HandlerBase struct {
    IsOutput bool    
}

func (h *HandlerBase) OnStartElement(tag string, attrs map[string]string) {
    if h.IsOutput {
        fmt.Println("Start-el: ", tag)
        for k, v := range attrs {
            fmt.Println(k, v) 
        }
    }
}

func (h *HandlerBase) OnEndElement(tag string) {
    if h.IsOutput {
        fmt.Println("End-el: ", tag)
    }
}

func (h *HandlerBase) OnText(text string) {
    if h.IsOutput {
        fmt.Println("Text: ", text)
    }
}

func (h *HandlerBase) OnComment(text string) {
    if h.IsOutput {
        fmt.Println("Comment: ", text)
    }
}

func (h *HandlerBase) OnPIElement(tag string, attrs map[string]string) {
    if h.IsOutput {
        fmt.Println("PI-el: ", tag)
        for k, v := range attrs {
            fmt.Println(k, v) 
        }
    }
}

func (h *HandlerBase) OnCData(text string) {
    if h.IsOutput {
        fmt.Println("CData: ", text)
    }
}

func (h *HandlerBase) OnError(line int, row int, message string) {
    if h.IsOutput {
        fmt.Println("Error: ", line, row, message)
    }
}
