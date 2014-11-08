package util

import(
    "log"
    "fmt"
    "time"
)

type StockLog struct {
    filename string
    logger *log.Logger
}

func (l *StockLog) Init(){
    format := "../log/log-%v-%v-%v-%v-%v.log"
    now := time.Now()
    l.filename = fmt.Sprintf(format, now.Year(), int(now.Month()), now.Day(), now.Minute(), now.Second())
    logfile, err := Create(l.filename)
    //defer logfile.Close()
    if err != nil {
        log.Fatalln("Cannot open log file: ", l.filename, err)
    }

    l.logger = log.New(logfile, "[Info]", log.Llongfile | log.LstdFlags)
    l.logger.Println("Success to initialize the logging!")
}

func (l *StockLog) Error(v ... interface{}) {
    l.logger.Println("Error start")
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



