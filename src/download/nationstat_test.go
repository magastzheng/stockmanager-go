package download_test

import (
    "testing"
    "download"
    "fmt"
    "net/url"
)

func Test_NationStatGetRoot(t *testing.T) {
    res := download.GetRoot()
    fmt.Println(len(res))

    query := "code=A0B&level=1&dbcode=hgyd&dimension=zb"
    values, err := url.ParseQuery(query)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(values)
}

