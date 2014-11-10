package download_test

import (
    "testing"
    "download"
    "fmt"
)

func Test_NationStatGetRoot(t *testing.T) {
    d := download.NewNationStatDownloader()
    res := d.GetRoot()
    fmt.Println(len(res))
}

func Test_NationStatGetChild(t *testing.T){
    d := download.NewNationStatDownloader()
    codes := []string{"A01"}
    result := d.GetChild(codes, 1)
    if len(result) == 0 {
        t.Error("Cannot get level 1")
    }

    codes = []string{"A0101"}
    result = d.GetChild(codes, 2)
    if len(result) == 0 {
        t.Error("Cannot get level 2")
    }

    codes = []string{"A010101"}
    result = d.GetChild(codes, 3)
    if len(result) == 0 {
        t.Error("Cannot get level 3")
    }
}

func Test_NationStatGetPeriod(t *testing.T){
    d := download.NewNationStatDownloader()
    result := d.GetPeriod()
    if len(result) == 0 {
        t.Error("Cannot get period")
    }
}

func Test_NationStatGetData(t *testing.T){
    d := download.NewNationStatDownloader()
    codes := []string{"A01010101"}
    result := d.GetData(codes, "200101", "-1")
    if len(result) == 0{
        t.Error(result)
    }
}
