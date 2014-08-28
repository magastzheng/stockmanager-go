//Parse the stock code and name from the HMTL page
package main

import(
    "fmt"
    "os"
    "io"
    "bufio"
    "code.google.com/p/go.net/html"
)

type Stock struct{
    id string
    name string
    website string
}

type StockParser struct{
    Stocks map [string] Stock    
}

func (p *StockParser) parseHtml(r io.Reader){
    d := html.NewTokenizer(r)
    for{
        //token type
        tokenType := d.Next()
        if tokenType == html.ErrorToken{
            return
        }

        token := d.Token()
        switch tokenType{
            case html.StartTagToken:
                fmt.Println(d.Data)
            case html.TextToken:
                fmt.Println(d.Data)
            case html.EndTagToken:
                fmt.Println(d.Data)
            case html.SelfClosingTagToken:
                fmt.Println(d.Data)
        }
    }
}

func (p *StockParser) parseString(html string){
    r := strings.NewReader(html)
    p.parseHtml(r)
}

func main(){
    filename = "stock_a-ha.dat"
    f, err := os.Open(filename)
 
    if err != nil {
        panic(err)
    }
    defer f.Close()
    
    //content, err := ioutil.ReadFile(filename)
    //if err != nil {
    //    fmt.Println("Error to read file!")
    //}

    reader := bufio.NewReader(f)
    parser := new(StockParser)
    parser.parseString(reader)
}
