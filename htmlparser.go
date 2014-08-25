//Parse the stock code and name from the HMTL page
package main

import(
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

            case html.TextToken:

            case html.EndTagToken:

            case html.SelfClosingTagToken:

        }
    }
}
