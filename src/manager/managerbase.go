package manager

import(
    "util"
    "runtime"
    "path/filepath"
    "fmt"
)

type ManagerBase struct {
    Logger *util.StockLog
    BaseDir string
}

func (m *ManagerBase) Init() {
    m.Logger = util.NewLog()
    pc, filename, line, ok := runtime.Caller(0)
    if pc < 0 || line < 0 || !ok {
        fmt.Println("Cannot read the managerbase.go")
        util.NewLog().Error("Cannot read the file managerbase.go")
    }

    m.BaseDir = filepath.Dir(filename) + "/../"
}

func (m *ManagerBase) Process() {
    
}
