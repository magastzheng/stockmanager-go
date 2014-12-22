package shsehandler

import(
    "entity/stentity"
    "strings"
    //"fmt"
)

type CompanyHandler struct{
    Company stentity.Company
}

func (h *CompanyHandler) OnObject(key string, keyValues map[string]string) {    
    switch key {
        case "result":
            h.Company = stentity.Company{
                Code: keyValues["COMPANY_CODE"],
                AbbrName: keyValues["COMPANY_ABBR"],
                Name: keyValues["FULLNAME"],
                Name_en: keyValues["FULL_NAME_IN_ENGLISH"],
                RegAddr: keyValues["COMPANY_ADDRESS"],
                InceptDate: keyValues["LISTINGDATEA"],
                Region: "",
                State: keyValues["AREA_NAME_DESC"],
                City: "",
                Industry: keyValues["CSRC_CODE_DESC"],
                Website: h.EscapeUrl(keyValues["WWW_ADDRESS"]),
            }
    }
}

func (h *CompanyHandler) OnArray(key string, elems []string){
    //fmt.Println("Array:", key)
    //for _, v := range elems {
    //    fmt.Println(v)
    //}
}

func (h *CompanyHandler) EscapeUrl(url string) string {
    return strings.Replace(url, "\\", "", -1)
}

func NewCompanyHandler() *CompanyHandler{
    return &CompanyHandler{}
}
