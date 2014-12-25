package nsmanager

import(
    "download"
    "parser"
    "entity"
    "util"
    "fmt"
    "strings"
)

const(
    NSStart = "198301"
    NSEnd = "-1"
)

type NationStatManager struct {
    downloader *download.NationStatDownloader
    parser *parser.NSParser
    logger *util.StockLog
    idxmap map[string] []entity.NSDataIndex
}

func (m *NationStatManager) Init() {
    m.downloader = download.NewNationStatDownloader()
    m.parser = parser.NewNSParser()
    m.logger = util.NewLog()
    m.idxmap = make(map[string] []entity.NSDataIndex)
}

func (m *NationStatManager) Process() {
    datastr := m.downloader.GetRoot()
    if len(datastr) == 0 {
        m.logger.Error("Cannot get the children of root")
    }

    m.WriteIndexData("root", 0, datastr)
    rootData := m.parser.ParseIndex(datastr)
    m.OutputIndex(rootData)
    for _, root := range rootData {
        if root.IsParent {
            m.GetIndex(root, 1)
        }
    }

    m.WriteIndex()
}

func (m *NationStatManager) GetIndex(idxdata entity.NSIndex, level int) {
    datastr := m.downloader.GetChild(idxdata.Id, level)
    if len(datastr) == 0 {
        m.logger.Error("Cannot get the children of:", idxdata.Id, " in level: ", level)
    }
    m.WriteIndexData(idxdata.Id, level, datastr)
    idxchild := m.parser.ParseIndex(datastr)
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
        m.GetData(idxdata.Id, dataid, NSStart, NSEnd)
    }
}

func (m *NationStatManager) GetData(pid string, ids []string, start string, end string) {
    datastr := m.downloader.GetData(ids, start, end)
    if len(datastr) == 0 {
        m.logger.Error("Cannot get data of: ", ids, start, end)
    }
    m.WriteData(ids, start, end, datastr)
    nsvalue := m.parser.ParseData(datastr)
    m.WriteValue(ids, start, end, nsvalue.TableData)

    //indexes := nsvalue.Value.Index
    //for _, idx := range indexes {
    //    m.idxmap[idx.Id] = idx.Name
    //}
    
    values, ok := m.idxmap[pid]
    if !ok {
        values = nsvalue.Value.Index
        m.idxmap[pid] = values
    } else {
        values = append(values, nsvalue.Value.Index ...)
        m.idxmap[pid] = values
    }
}

func (m *NationStatManager) WriteIndexData(id string, level int, content string) {
    format := "../data/index-%v-%v.dat"
    filename := fmt.Sprintf(format, level, id)
    util.WriteFile(filename, content)
}

func (m *NationStatManager) WriteData(ids []string, start, end, content string) {
    idstr := strings.Join(ids, "-")
    format := "../data/data-%v-%v-%v.dat"
    filename := fmt.Sprintf(format, ids[0], start, end)
    content += idstr + "\n" + content
    util.WriteFile(filename, content)
}

func (m *NationStatManager) WriteValue(ids []string, start string, end string, data map[string]string) {
    idstr := strings.Join(ids, "-")

    format := "../data/actualdata-%v-%v-%v.dat"
    filename := fmt.Sprintf(format, ids[0], start, end)
    content := idstr + "\n"
    for k, v := range data{
        content += k + ": " + v + "\n"
    }
    util.WriteFile(filename, content)
}

func (m *NationStatManager) WriteIndex() {
    filename := "../data/macroindex.dat"

    var content string
    for k, vs := range m.idxmap {
        content += fmt.Sprintf("=============Parent: %v", k)
        for _, v := range vs {
            content += fmt.Sprintf("%v: %v\n", v.Id, v.Name)
        }
    }

    util.WriteFile(filename, content)
}

func (m *NationStatManager) OutputIndex(idxdata []entity.NSIndex){
    var str string
    for _, v := range idxdata {
        str += fmt.Sprintf("Id: %v, Name: %v, PId: %v, EName: %v, IfData: %v, IsParent: %v\n", v.Id, v.Name, v.PId, v.EName, v.IfData, v.IsParent)
        //fmt.Println(str)
        
    }
    m.logger.Info(str)
}

func NewNationStatManager() *NationStatManager{
    m := new(NationStatManager)
    m.Init()

    return m
}
