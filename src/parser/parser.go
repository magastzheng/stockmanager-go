package parser

import (
    //"fmt"
    "strings"
)

const (
    ElementOpen = iota
    ElementClose
    Text
    Comment
    CData
    Other
)

const (
    STAT_NONE = iota
    STAT_AFTER_LT
    STAT_START_TAG
    STAT_END_TAG
    STAT_TEXT
    STAT_PRE_COMMENT1
    STAT_PRE_COMMENT2
    STAT_COMMENT
    STAT_PROCESS_INSTRUCTION
    STAT_CDATA
    STAT_PRE_KEY
    STAT_KEY
    STAT_PRE_VALUE
    STAT_VALUE
    STAT_NAME
    STAT_ATTR
    STAT_END
    STAT_MINUS1
    STAT_MINUS2
    STAT_DOCTYPE
)

const (
    Lt rune = '<'
    Gt rune = '>'
    Slash rune = '/'
    And rune = '&'
    Apos rune = '\''
    Quot rune = '"'
    Blank rune = ' '
    Exclam = '!'
    Dash = '-'
    CR = '\n'
    RE = '\r'
    Tab = '\t'
    Question = '?'
    Underscore = '_'
    Eq = '='
    LeftBracket = '['
    RightBracket = ']'
    Semicolon = ';'
)

const (
    Script = "SCRIPT"
)

const MAX_ATTR_NR = 1024

const(
    End rune = rune(-1)
)

const (
    CDataStr = "CDATA"
)

type Handler interface {
    OnStartElement(tag string, attrs map[string]string)
    OnEndElement(tag string)
    OnText(text string)
    OnComment(text string)
    OnPIElement(tag string, attrs map[string]string)
    OnCData(text string)
    OnError(line int, row int, message string)
}

type Parser interface {
    SetHandler(handler Handler)
    Parse()
}

type TextReader interface {
    Read() rune 
    //ReadCharacter() rune
    ReadElement() string
    ReadText() string
    IsSpecialCh(r rune) bool
    ClearStatus()
}

type TextParser struct {
    handler Handler
    Data string
    buffer []rune
    status int
    length int
    current int
    isScript bool
}

func (p *TextParser) IsSpace(ch rune) bool {
    return ch == Blank || ch == Tab
}

func (p *TextParser) IsAlpha(ch rune) bool {
    return ('a' <= ch && ch <= 'z') || ( 'A' <= ch && ch <= 'Z')
}

func (p *TextParser) IsScriptElement(tag string) bool {
    return strings.ToUpper(tag) == Script
}

func (p *TextParser) IsNormalText() bool {
    if p.current + 1 < p.length && p.buffer[p.current + 1] == Slash && p.current + 2 + len(Script) < p.length {
        nextStr := string(p.buffer[p.current + 2: p.current + 2 + len(Script)])
        //fmt.Println("IsNormalText(): ", nextStr)
        return !p.IsScriptElement(nextStr)
    }

    return false
}

func (p *TextParser) IsDocType(tag string) bool {
    return strings.ToUpper(tag) == "DOCTYPE"
}

func (p *TextParser) SetData(data string){
    p.Data = data
    p.buffer = []rune(p.Data)
    p.length = len(p.buffer)
    p.current = 0
    p.isScript = false
} 

func (p *TextParser) SetHandler(handler Handler) {
    p.handler = handler
}

func (p *TextParser) ParseStr(data string) {
    p.SetData(data)
    p.Parse()
}

func (p *TextParser) Parse(){ 
    p.status = STAT_NONE
    for p.current = 0; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch p.status {
            case STAT_NONE:
                if ch == Lt {    
                    //reset_buffer
                    p.status = STAT_AFTER_LT
                } else if !p.IsSpace(ch) {
                    p.status = STAT_TEXT
                } else {
                    
                }

            case STAT_AFTER_LT:
                
                if ch == Question {
                    p.status = STAT_PROCESS_INSTRUCTION
                } else if ch == Slash {
                    p.status = STAT_END_TAG
                } else if ch == Exclam {
                    p.status = STAT_PRE_COMMENT1
                } else if p.IsAlpha(ch) || ch == Underscore {
                    p.status = STAT_START_TAG
                } else {
                    //do nothing
                }

                //fmt.Println("STAT_AFTER_LT: ", string(ch), " Status: ", p.status)
            case STAT_START_TAG:
                //parse start tag
				//fmt.Println("STAT_START_TAG: ", string(ch))
                p.ParseStartTag()
                p.status = STAT_NONE
            case STAT_END_TAG:
                //parse end tag
                p.ParseEndTag()
                p.status = STAT_NONE
            case STAT_PROCESS_INSTRUCTION:
                //parse process instruction
                p.ParsePI()
                p.status = STAT_NONE
            case STAT_TEXT:
                //parse text
                //fmt.Println("STAT_TEXT: ", string(ch))
                p.ParseText()
                p.status = STAT_NONE
				//fmt.Println("STAT_TEXT **END", string(p.buffer[p.current - 2: p.current + 2]))
            case STAT_PRE_COMMENT1:
                if ch == Dash {
                    p.status = STAT_PRE_COMMENT2
                } else if ch == LeftBracket {
                    p.status = STAT_CDATA
                } else {
                    //do nothing
                    var text string
                    if p.current + 7 < p.length {
                        text = string(p.buffer[p.current: p.current + 7])
                    }
                    //fmt.Println("Comment: ", text)
                    if p.IsDocType(text) {
                        p.status = STAT_DOCTYPE
                    } else {
                        //fmt.Println("do nothing in <!")
                    }
                }

				//fmt.Println("COMMENT: ", string(ch))
            case STAT_PRE_COMMENT2:
                if ch == Dash {
                    p.status = STAT_COMMENT
                } else {
                    //do nothing
                }
            case STAT_COMMENT:
                //parse comment
                p.ParseComment()
                p.status = STAT_NONE
            case STAT_DOCTYPE:
                p.ParseDoctype()
                p.status = STAT_NONE
            case STAT_CDATA:
                p.ParseCData()
                p.status = STAT_NONE
        }
    }
}

func (p *TextParser) ParseAttributes(endch rune) map[string]string {
    status := STAT_PRE_KEY
    valueEnd := Quot
    start := 0
    attrNR := 0
    attrs := make(map[string]string)
    name := ""
    value := ""
    for ; p.current < p.length && attrNR < MAX_ATTR_NR; p.current++ {
        ch := p.buffer[p.current]
        switch status {
            case STAT_PRE_KEY:
                if ch == endch || ch == Gt {
                    //read '/' or '>' then go to end status
                    
                    status = STAT_END
                   //if ch == Gt {
                   //     p.handler.OnEndElement(name)
                    //}
                } else if !p.IsSpace(ch) {
                    status = STAT_KEY
                    start = p.current
                }
            case STAT_KEY:
                if ch == Eq {
                    //read the name (p.current - start)
                    names := p.buffer[start: p.current]
                    name = string(names)
                    status = STAT_PRE_VALUE
                } else if ch == Gt {
                    // the case as <script src="" lang="" defer></script>
                    //fmt.Println("======", string(p.buffer[p.current]))
                    status = STAT_END
                } else {
                    // do nothing to get more key name
                }
            case STAT_PRE_VALUE:
                if ch == Quot || ch == Apos {
                    //read " or '
                    status = STAT_VALUE
                    valueEnd = ch
                    start = p.current + 1
                } else {
					//the case as <tr class=xss font="test"></tr>
					status = STAT_VALUE
					valueEnd = Blank
					start = p.current
				}
            case STAT_VALUE:
                if ch == valueEnd {
                    values := p.buffer[start:p.current]
                    value = string(values)
                    attrs[name] = value
                    status = STAT_PRE_KEY
                } else if ch == Gt {
					values := p.buffer[start: p.current]
					value = string(values)
                    attrs[name] = value
					status = STAT_END
                } else {
					//do nothing
				}
        }

        if status == STAT_END {
            break
        }
    }

    return attrs
}

func (p *TextParser) ParseStartTag() {
    status := STAT_NAME
    start := p.current - 1 
    end := p.current
    isTag := true
    attrs := make(map[string]string)

    for ; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch status {
            case STAT_NAME:
                if p.IsSpace(ch) || ch == Gt || ch == Slash {
                    if ch != Gt && ch != Slash { 
                        if isTag && p.IsSpace(ch) {
                            end = p.current
                            isTag = false
                        }
                        status = STAT_ATTR 
                    } else {
                        if ch == Slash {
                            if isTag {
                                end = p.current
                                isTag = false
                            }
                        } else {
                            //those node only as node name
                            end = p.current
                        }

                        status = STAT_END
                    }
                }
            case STAT_ATTR:
                attrs = p.ParseAttributes('/')
                status = STAT_END
        }

        if status == STAT_END {
            break
        }
    }

    ch := p.buffer[p.current]
    names := p.buffer[start: end]
    tag := string(names)
    p.isScript = p.IsScriptElement(tag)
    p.handler.OnStartElement(tag, attrs)
	//fmt.Println("start-parser: ", tag)
    if ch == Slash {
        //if it is a self-close element, read more a char '>' to end it
        p.current += 1
        p.handler.OnEndElement(tag)
        if p.isScript {
            p.isScript = false
        }
    }
}

func (p *TextParser) ParsePI() {
    status := STAT_NAME
    start := p.current
    end := p.current
    firstSpace := true

    attrs := make(map[string]string)
    for ; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch status {
            case STAT_NAME:
                if p.IsSpace(ch) || ch == Gt {
                    if ch != Gt {
                        if firstSpace && p.IsSpace(ch) {
                            end = p.current
                        }
                        status = STAT_ATTR
                    } else {
                        status = STAT_END
                    }
                }
            case STAT_ATTR:
                attrs = p.ParseAttributes('?')
                status = STAT_END
        }

        if status == STAT_END {
            break
        }
    }

    tag := string(p.buffer[start:end])
    p.handler.OnPIElement(tag, attrs)
    //fmt.Println("PI: ", tag, " current ", p.current)
    for ; p.buffer[p.current] != Gt && p.current < p.length; p.current++ {
        //read continue to end the element
    }
}

func (p *TextParser) ParseCData() {
    //status := STAT_CDATA
    start := p.current - 3

    for ; p.current + 2 < p.length && !(p.buffer[p.current] == RightBracket && p.buffer[p.current + 1] == RightBracket && p.buffer[p.current + 2] == Gt); p.current++ {
        //do nothing
    }
    
    p.current += 2
    //fmt.Println("CData: ", string(p.buffer[start: p.current+1]))
    cdata := string(p.buffer[start:p.current+1])
    p.handler.OnCData(cdata)
}

func (p *TextParser) ParseComment() {
    status := STAT_COMMENT
    start := p.current
    completed := false 
    for ; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        switch status {
            case STAT_COMMENT:
                if ch == Dash {
                    status = STAT_MINUS1
                } else if ch == Gt {
					//fmt.Println("STAT_COMMET: ", string(ch), "more: ", string(p.buffer[p.current - 2: p.current + 2]))
				}

            case STAT_MINUS1:
                if ch == Dash {
                    status = STAT_MINUS2
                //} else if ch == Gt{
					//there are some case as <!---------> 
				//	completed = true
				//	break
				} else {
                    status = STAT_COMMENT
                }
            case STAT_MINUS2:
                if ch == Gt {
                    completed = true
                    break
                } else if ch == Dash {
					//there are some case as <!--------->
					status = STAT_MINUS2
				} else {
                    status = STAT_COMMENT
                }
        }

        if completed {
            break
        }
    }
   
    comment := p.buffer[start:p.current - 2]
    p.handler.OnComment(string(comment))
    return
}

func (p *TextParser) ParseDoctype() {
    //start := p.current
    for ; p.buffer[p.current] != Gt && p.current < p.length; p.current++ {
        //read continue to end the element
    }
    //fmt.Println("DocType: ", string(p.buffer[start: p.current]))
}

func (p *TextParser) ParseEndTag() {
    start := p.current
    for ; p.current < p.length; p.current++ {
        if p.buffer[p.current] == Gt {
            break
        }
    }
    
    name := p.buffer[start: p.current]
    //fmt.Println("End tag name: ", string(name))
    p.handler.OnEndElement(string(name))
    if p.isScript {
        p.isScript = false
    }
    return
}

func (p *TextParser) ParseEntity() []rune {
    start := p.current 
    
    isSemicolon := false
    isSpace := false
    for ; p.current < p.length; p.current ++ {
        ch := p.buffer[p.current]
        if ch == Lt {
            p.current = p.current - 1
            break
        }else if ch == Semicolon {
            isSemicolon = true
            break
        } else if p.IsSpace(ch) {
            isSpace = true
            break
        } else {
            //do nothing
        }
    }
    
    if isSpace {
        return p.buffer[start: p.current]
    } else if isSemicolon {
        en := p.buffer[start: p.current+1]
        entity := string(en)
        entity = strings.Trim(entity, " ")
    
        //fmt.Println("start", string(p.buffer[start]), "end", string(p.buffer[p.current]), "entity: ", entity)

        if entity == "&lt;" {
            return [] rune {Lt}
        }else if entity == "&gt;" {
            return [] rune {Gt}
        }else if entity == "&amp;"{
            return [] rune {And}
        }else if entity == "&apos;"{
            return [] rune {Apos}
        }else if entity == "&quot;" {
            return [] rune {Quot}  
        }else{
            return en
        }
    }else{
        return p.buffer[start: p.current]
    }
    
    //fmt.Println("current ", p.current, " length ", p.length)
    return nil
}

func (p *TextParser) ParseText() {
    //read back a char to get back the one in Parse()
    p.current = p.current - 1
    start := p.current
    end := p.current
    stext := []rune{p.buffer[start]}

    //str := ""
    for ; p.current < p.length; p.current++ {
        ch := p.buffer[p.current]
        //str += string(ch)
        //TODO: handle the script text
        //read < and end of the parsing
        if ch == Lt {
            if p.isScript && p.IsNormalText() {
            
            } else {

                if p.current > start {
            
                }
            
                end = p.current
            
                //return back a char
                p.current = p.current - 1
                break;
            }
        } else if ch == And {
            //read & and parse the entity
            entity := p.ParseEntity()
            for i := 0; i < len(entity); i++{
                stext = append(stext, entity[i])
            }
        } else {
            stext = append(stext, ch)
        }
    }
    
    //fmt.Println("Text: ", str)
    if p.current > start {
        text := p.buffer[start: end]
        p.handler.OnText(string(text))
    }
}

func NewTextParser(handler Handler) *TextParser {
    p := new(TextParser)
    p.SetHandler(handler)

    return p
}
