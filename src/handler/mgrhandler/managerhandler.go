package mgrhandler

import(
    "handler"
)

type ManagerHandler struct {
    handler.HandlerBase
}

func (h *ManagerHandler) Init() {

}

func (h *ManagerHandler) OnStartElement(tag string, attrs map[string]string) {
    
}

func (h *ManagerHandler) OnEndElement(tag string) {

}

func (h *ManagerHandler) OnText(text string) {

}

func NewManagerHandler() *ManagerHandler{
    h := new(ManagerHandler)
    h.Init()

    return h
}
