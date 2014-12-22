package jsonparser

import(
    //"io/ioutil"
    "fmt"
)

const (
    STAT_NONE = iota
    STAT_START_OBJ  //Left bracket
    STAT_START_ARR //left square bracket
    STAT_END_OBJ //end of bracket
    STAT_END_ARR //end of square bracket
    STAT_TEXT
    STAT_PRE_KEY
    STAT_KEY
    STAT_KEY_VALUE_DELIMITER
    STAT_PRE_VALUE
    STAT_VALUE
    STAT_MEMBER_DELIMITER
    STAT_PRE_ELEMENT
    STAT_PRE_CHILD_ELEMENT
    STAT_ELEMENT
    STAT_ELEMENT_DELIMITER
    STAT_ARR_OBJ_START
)

const (
    LeftBrace rune = '{'
    RightBrace rune = '}'
    Apos rune = '\''
    Quot rune = '"'
    Blank rune = ' '
    CR = '\n'
    RE = '\r'
    Tab = '\t'
    LeftBracket = '['
    RightBracket = ']'
    Colon = ':'
    Comma = ','
)

const (
    Object = "object"
    Array = "array"
)

type Handler interface {
    OnObject(key string, keyValues map[string]string)
    OnArray(key string, elems []string)
}

type JsonHandler struct {

}

func (h *JsonHandler) OnObject(key string, keyValues map[string]string){
    fmt.Println("***Parse object: ", key)
    for k, v := range keyValues {
        fmt.Println(k, v)
    }
}

func (h *JsonHandler) OnArray(key string, elems []string) {
    fmt.Println("***Parse array: ", key)
    for _, e := range elems {
        fmt.Println(e)
    }
}

type JsonParser struct {
    handler Handler
    Data string
    buffer []rune
    status int
    length int
    current int
}

func (p *JsonParser) IsSpace(ch rune) bool {
    return ch == Blank || ch == Tab
}

func (p *JsonParser) IsAlpha(ch rune) bool {
    return ('a' <= ch && ch <= 'z') || ( 'A' <= ch && ch <= 'Z')
}

func (p *JsonParser) Skip(ch rune) bool {
    return ch == Blank || ch == Tab || ch == CR || ch == RE
}

func (p *JsonParser) SetData(data string) {
    p.Data = data
    p.buffer = []rune(p.Data)
    p.length = len(p.buffer)
    p.current = 0
}

func (p *JsonParser) SetHandler(handler Handler) {
    p.handler = handler
}

func (p *JsonParser) ParseStr(data string) {
    p.SetData(data)
    p.Parse()
}

func (p *JsonParser) Parse() {
    p.status = STAT_NONE
    for p.current = 0; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch p.status {
            case STAT_NONE:
                if ch == LeftBrace {
                    p.status = STAT_START_OBJ
                } else if ch == LeftBracket {
                    p.status = STAT_START_ARR 
                } else if !p.Skip(ch) {
                    p.status = STAT_TEXT
                } else {
                    
                }
            case STAT_START_OBJ:
                p.ParseStartObject("")
                p.status = STAT_NONE
            case STAT_START_ARR:
                p.ParseStartArray("")
                p.status = STAT_NONE
            case STAT_TEXT:
                p.ParseText()
                p.status = STAT_NONE
        }
    }
}

func (p *JsonParser) ParseStartObject(key string) {
    //end while get {,},[,]
    status := STAT_PRE_KEY
    valueQuot := Apos
    start := p.current - 1
    end := p.current
    name := ""
    value := ""
    isStr := false

    keyValues := make(map[string]string)
    for ; p.current < p.length; p.current++ {
        //p.Print("Object", "start")
        ch := p.buffer[p.current]
        switch status {
            case STAT_PRE_KEY:
                //p.Print("STAT_PRE_KEY", "start")
                if ch == Apos || ch == Quot {
                    valueQuot = ch
                    start = p.current
                    status = STAT_KEY
                } else if ch == RightBrace {
                    status = STAT_END_OBJ   
                } else {
                    //do nothing
                }
            case STAT_KEY:
                if ch == valueQuot {
                    end = p.current
                    name = string(p.buffer[start+1: end])
                } else if ch == Colon {
                    status = STAT_KEY_VALUE_DELIMITER
                } else {
                    //do nothing
                }
            case STAT_KEY_VALUE_DELIMITER:
                if ch == valueQuot {
                    start = p.current
                    status = STAT_PRE_VALUE
                } else if ch == LeftBracket {
                    p.current = p.current + 1
                    p.ParseStartArray(name)
                    keyValues[name] = Array
                    status = STAT_MEMBER_DELIMITER
                } else if ch == LeftBrace {
                    p.current = p.current + 1
                    p.ParseStartObject(name)
                    keyValues[name] = Object
                    status = STAT_MEMBER_DELIMITER
                } else if ch == Comma {
                    value = ""
                    status = STAT_MEMBER_DELIMITER
                }else if !p.Skip(ch){
                    start = p.current
                    status = STAT_VALUE
                } else {
                    //fmt.Println("Do nothing in key-value delimiter")
                }
            case STAT_PRE_VALUE:
                isStr = true
                status = STAT_VALUE
            case STAT_VALUE:
                if ch == valueQuot {
                    isStr = false
                    end = p.current
                    value = string(p.buffer[start+1: end])
                    keyValues[name] = value
                    status = STAT_MEMBER_DELIMITER
                } else if ch == Comma {
                    end = p.current 
                    value = string(p.buffer[start: end])
                    keyValues[name] = value
                    status = STAT_PRE_KEY
                } else if !isStr && p.Skip(ch) {
                    _, ok := keyValues[name]
                    if !ok {
                        keyValues[name] = string(p.buffer[start: p.current])
                    }
                }else if ch == RightBrace {
                    _, ok := keyValues[name]
                    if !ok {
                        keyValues[name] = string(p.buffer[start: p.current])
                    }
                    status = STAT_END_OBJ
                } else {
                    //do nothing
                }
            case STAT_MEMBER_DELIMITER: 
                if ch == Comma {
                    status = STAT_PRE_KEY
                } else if ch == RightBrace {
                    _, ok := keyValues[name]
                    if !ok {
                        keyValues[name] = string(p.buffer[start: p.current])
                    }
                    status = STAT_END_OBJ
                } else {
                    //do nothing
                }
            }

        if status == STAT_END_OBJ {
            break
        }
    }

    p.handler.OnObject(key, keyValues)
}

func (p *JsonParser) ParseStartArray(key string) {
    //end while get {,},[,]
    status := STAT_PRE_ELEMENT
    valueQuot := Apos
    start := p.current - 1
    end := p.current
    value := ""
    elems := make([]string, 0)
    isNoStr := false

    for ; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch status {
            case STAT_PRE_ELEMENT:
                if ch == Apos || ch == Quot {
                    start = p.current + 1
                    valueQuot = ch
                    status = STAT_ELEMENT
                } else if ch == LeftBracket {
                    status = STAT_PRE_CHILD_ELEMENT
                } else if ch == LeftBrace {
                    status = STAT_ARR_OBJ_START
                } else if !p.Skip(ch) {
                    isNoStr = true
                    start = p.current
                    status = STAT_ELEMENT
                }
            case STAT_ELEMENT:
                //p.Print("STAT_ELEMENT", "start")
                if ch == valueQuot {
                    end = p.current 
                    value = string(p.buffer[start: end])
                    elems = append(elems, value)
                    status = STAT_ELEMENT_DELIMITER
                } else if ch == Comma {
                    isNoStr = false
                    end = p.current
                    value = string(p.buffer[start: end])
                    elems = append(elems, value)
                    status = STAT_ELEMENT_DELIMITER
                } else if ch == RightBracket {
                    if isNoStr {
                        value = string(p.buffer[start: p.current])
                        elems = append(elems, value)
                    }

                    status = STAT_END_ARR
                } else {
                    //do nothing
                }
            case STAT_PRE_CHILD_ELEMENT:
                //p.Print("STAT_PRE_CHILD_ELEMENT", "start")
                p.ParseStartArray(key)
                status = STAT_ELEMENT_DELIMITER
            case STAT_ARR_OBJ_START:
                p.ParseStartObject(key)
                status = STAT_ELEMENT_DELIMITER
            case STAT_ELEMENT_DELIMITER:
                //p.Print("STAT_ELEMENT_DELIMITER", "start")
                if ch == valueQuot {
                    start = p.current + 1
                    status = STAT_ELEMENT
                } else if ch == LeftBracket {
                    status = STAT_PRE_CHILD_ELEMENT
                } else if ch == LeftBrace {
                    status = STAT_ARR_OBJ_START
                } else if ch == RightBracket {
                    status = STAT_END_ARR
                } else if ch == Comma {
                    status = STAT_PRE_ELEMENT
                } else if !p.Skip(ch) {
                    isNoStr = true
                    start = p.current
                    status = STAT_ELEMENT
                } else {
                    //do nothing
                }
        }

        if status == STAT_END_ARR {
            break
        }
    }

    p.handler.OnArray(key, elems)
}

func (p *JsonParser) Print(stat, message string) {
    start := p.current
    end := p.current + 3
    if start < 0 {
        start = 0
    }
    if end > p.length {
        end = p.length
    }

    fmt.Println(stat, message, "current: ", p.current, string(p.buffer[p.current]), "context: ", string(p.buffer[start: end]))
}

func (p *JsonParser) ParseText() {
    
}

func NewJsonParser(handler Handler) *JsonParser{
    p := new(JsonParser)
    p.SetHandler(handler)
    return p
}
