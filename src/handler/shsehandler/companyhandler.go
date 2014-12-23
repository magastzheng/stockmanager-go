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
                CSRCCategory: keyValues["CSRC_CODE_DESC"],
                CSRCBigCategory: keyValues["CSRC_GREAT_CODE_DESC"],
                CSRCMidCategory: keyValues["CSRC_MIDDLE_CODE_DESC"],
                SSEIndustry: keyValues["SSE_CODE_DESC"],
                Website: h.EscapeUrl(keyValues["WWW_ADDRESS"]),
                LegalRepresent: keyValues["LEGAL_REPRESENTATIVE"],
                Phone: keyValues["REPR_PHONE"],
                Email: keyValues["E_MAIL_ADDRESS"],
                OfficeAddr: keyValues["OFFICE_ADDRESS"],
                OfficeZip: keyValues["OFFICE_ZIP"],
                Status: h.getCurrentStatus(keyValues["STATE_CODE_A_DESC"]),
                IsSample: h.getIsSample(keyValues["SECURITY_30_DESC"]),
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

func (h *CompanyHandler) getCurrentStatus(v string) int {
    if v == "上市" {
        return 1
    }

    return 0
}

func (h *CompanyHandler) getIsSample(v string) int {
    if v == "是" {
        return 1
    }

    return 0
}

func NewCompanyHandler() *CompanyHandler{
    return &CompanyHandler{}
}
