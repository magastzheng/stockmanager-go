package nsdownload_test

import (
    "testing"
    "download/nsdownload"
    "fmt"
)

func Test_NationStatGetRoot(t *testing.T) {
    d := nsdownload.NewNationStatDownloader()
    res := d.GetRoot()
    fmt.Println(len(res))
}

func Test_NationStatGetChild(t *testing.T){
    d := nsdownload.NewNationStatDownloader()
    code := "A01"
    result := d.GetChild(code, 1)
    if len(result) == 0 {
        t.Error("Cannot get level 1")
    }

    code = "A0101"
    result = d.GetChild(code, 2)
    if len(result) == 0 {
        t.Error("Cannot get level 2")
    }

    code = "A010101"
    result = d.GetChild(code, 3)
    if len(result) == 0 {
        t.Error("Cannot get level 3")
    }
}

func Test_NationStatGetPeriod(t *testing.T){
    d := nsdownload.NewNationStatDownloader()
    result := d.GetPeriod()
    if len(result) == 0 {
        t.Error("Cannot get period")
    }
}

func Test_NationStatGetData(t *testing.T){
    d := nsdownload.NewNationStatDownloader()
    codes := []string{"A01010101"}
    result := d.GetData(codes, "200101", "-1")
    if len(result) == 0{
        t.Error(result)
    }
}
