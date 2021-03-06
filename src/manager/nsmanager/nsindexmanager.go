package nsmanager

import(
    "download/nsdownload"
    "parser/nsparser"
    nsdb "stockdb/nationstatdb"
    ns "entity/nsentity"
    "manager"
    "util"
    "fmt"
    "strings"
    "time"
)

//const(
    //NSStart = "198301"
    //NSStart = "201401"
    //NSEnd = "-1"
    //NSDateFormat = "2006-01"
    //NSInputDateFormat = "200601"
//)

type NSIndexManager struct {
    manager.ManagerBase
    downloader *nsdownload.NationStatDownloader
    db *nsdb.IndexDB
    parser *nsparser.NSParser
    indexes []ns.NSDBIndex
    nsstart string //"yyyyMM" or "-1" as to present
    nsend string
}

func (m *NSIndexManager) Init() {
    m.ManagerBase.Init()
    m.downloader = nsdownload.NewNationStatDownloader()
    m.db = nsdb.NewIndexDB("macroindecis", "nsindex")
    m.parser = nsparser.NewNSParser()
    m.indexes = make([]ns.NSDBIndex, 0)
}

func (m *NSIndexManager) Process() {
    datastr := m.downloader.GetRoot()
    if len(datastr) == 0 {
        m.Logger.Error("Cannot get the children of root")
    }

    //m.WriteIndexData("root", 0, datastr)
    rootData := m.parser.ParseIndex(datastr)
    m.AddIndex("root", rootData)
    //m.OutputIndex(rootData)
    for _, root := range rootData {
        if root.IsParent {
            m.GetIndex(root, 1)
        }
    }
    
    m.db.TranInsert(m.indexes)
    m.WriteIndex()
}

func (m *NSIndexManager) GetIndex(idxdata ns.NSIndex, level int) {
    datastr := m.downloader.GetChild(idxdata.Id, level)
    if len(datastr) == 0 {
        m.Logger.Error("Cannot get the children of:", idxdata.Id, " in level: ", level)
    }
    m.WriteIndexData(idxdata.Id, level, datastr)
    idxchild := m.parser.ParseIndex(datastr)
    //m.AddIndex(idxdata.Id, idxchild)

    dataid := make([]string, 0)
    for _, idx := range idxchild {
        if idx.IfData == "1" {
            dataid = append(dataid, idx.Id)
        } else if idx.IsParent || idx.IfData == "4" {
            m.GetIndex(idx, level + 1)
        } else {
            //do nothing
        }
    }
    
    if len(dataid) > 0 {
        m.GetData(idxdata.Id, dataid, m.nsstart, m.nsend)
    }
}

func (m *NSIndexManager) GetData(pid string, ids []string, start string, end string) {
    datastr := m.downloader.GetData(ids, start, end)
    if len(datastr) == 0 {
        m.Logger.Error("Cannot get data of: ", ids, start, end)
    }
    m.WriteData(ids, start, end, datastr)
    nsvalue := m.parser.ParseData(datastr)
    m.WriteValue(ids, start, end, nsvalue.TableData)

    m.AddDataIndex(pid, nsvalue.Value.Index)
}

func (m *NSIndexManager) AddIndex(pid string, idxes []ns.NSIndex) {
    for _, idx := range idxes {
        dbidx := ns.NSDBIndex{}
        dbidx.Id = idx.Id
        dbidx.Name = idx.Name
        dbidx.EName = idx.EName

        if len(idx.PId) > 0 {
            dbidx.Parent = idx.PId
        } else {
            dbidx.Parent = pid
        }

        m.indexes = append(m.indexes, dbidx)
    }
}

func (m *NSIndexManager) AddDataIndex(pid string, dataIdxes []ns.NSDataIndex) {
    for _, idx := range dataIdxes {
        dbidx := ns.NSDBIndex {}
        dbidx.NSDataIndex = idx
        dbidx.Parent = pid

        m.indexes = append(m.indexes, dbidx)
    }
}

func (m *NSIndexManager) WriteIndexData(id string, level int, content string) {
    format := "%sdata/ns/index-%v-%v.dat"
    filename := fmt.Sprintf(format, m.BaseDir, level, id)
    fmt.Println(filename)
    util.WriteFile(filename, content)
}

func (m *NSIndexManager) WriteData(ids []string, start, end, content string) {
    idstr := strings.Join(ids, "-")
    format := "%sdata/ns/data-%v-%v-%v.dat"
    filename := fmt.Sprintf(format, m.BaseDir, ids[0], start, end)
    content += idstr + "\n" + content
    util.WriteFile(filename, content)
}

func (m *NSIndexManager) WriteValue(ids []string, start string, end string, data map[string]string) {
    idstr := strings.Join(ids, "-")

    format := "%sdata/ns/actualdata-%v-%v-%v.dat"
    filename := fmt.Sprintf(format, m.BaseDir, ids[0], start, end)
    content := idstr + "\n"
    for k, v := range data{
        content += k + ": " + v + "\n"
    }
    util.WriteFile(filename, content)
}

func (m *NSIndexManager) WriteIndex() {
    format := "%sdata/ns/macroindex.dat"
    filename := fmt.Sprintf(format, m.BaseDir)

    var content string
    for _, idx := range m.indexes {
        content += fmt.Sprintf("=============Parent: %v\n", idx.Parent)
        content += fmt.Sprintf("%v: %v\n", idx.Id, idx.Name)
    }

    util.WriteFile(filename, content)
}

func (m *NSIndexManager) OutputIndex(idxdata []ns.NSIndex){
    var str string
    for _, v := range idxdata {
        str += fmt.Sprintf("Id: %v, Name: %v, PId: %v, EName: %v, IfData: %v, IsParent: %v\n", v.Id, v.Name, v.PId, v.EName, v.IfData, v.IsParent)
        //fmt.Println(str)
        
    }
    m.Logger.Info(str)
}

func NewNSIndexManager(start, end string) *NSIndexManager{
    m := new(NSIndexManager)
    m.Init()
    
    t, err := time.Parse(NSDateFormat, start)
    if err != nil {
        m.nsstart = "-1"
    } else {
        m.nsstart = t.Format(NSInputDateFormat)
    }

    t, err = time.Parse(NSDateFormat, end)
    if err != nil {
        m.nsend = "-1"
    } else {
        m.nsend = t.Format(NSInputDateFormat)
    }

    fmt.Println("start: ", m.nsstart, " end: ", m.nsend)

    return m
}
