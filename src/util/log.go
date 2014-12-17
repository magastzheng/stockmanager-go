package util

import(
    "log"
    "fmt"
    "time"
    "runtime"
    "path/filepath"
)

type StockLog struct {
    filename string
    logger *log.Logger
}

func (l *StockLog) Init(){
    format := "../log/log-%v-%v-%v-%v-%v.log"
    lfilename := ""
    pc, filename, line, ok := runtime.Caller(0)
    if pc < 0 || line < 0 || !ok {
        fmt.Println("Cannot read the file log.go")
        lfilename = format
    } else {
        lfilename = filepath.Dir(filename) + "/" + format
    }
    
    now := time.Now()
    l.filename = fmt.Sprintf(lfilename, now.Year(), int(now.Month()), now.Day(), now.Minute(), now.Second())
    logfile, err := Create(l.filename)
    //defer logfile.Close()
    if err != nil {
        log.Fatalln("Cannot open log file: ", l.filename, err)
    }

    //l.logger = log.New(logfile, "[Info]", log.Llongfile | log.LstdFlags)
    l.logger = log.New(logfile, "[Info]", log.LstdFlags)
    l.logger.Println("Success to initialize the logging!")
}

func (l *StockLog) Error(v ... interface{}) {
    l.logger.SetPrefix("[Error]")
    l.logger.Println(v)
}

func (l *StockLog) Info(v ...interface{}){
    l.logger.SetPrefix("[Info]")
    l.logger.Println(v)
}

var stlog *StockLog
func NewLog() *StockLog{
    if stlog == nil {
        l := new(StockLog)
        l.Init()

        stlog = l
    }

    return stlog
}



