package manager

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
}

func (m *NationStatManager) Init() {
    m.downloader = download.NewNationStatDownloader()
    m.parser = parser.NewNSParser()
    m.logger = util.NewLog()
}

func (m *NationStatManager) Process() {
    datastr := m.downloader.GetRoot()
    if len(datastr) == 0 {
        m.logger.Error("Cannot get the children of root")
    }

    m.WriteIndexData("root", 0, datastr)
    rootData := m.parser.ParseIndex(datastr)
    for _, root := range rootData {
        if root.IsParent {
            m.GetIndex(root, 1)
        }
    }
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
        m.GetData(dataid, NSStart, NSEnd)
    }
}

func (m *NationStatManager) GetData(ids []string, start string, end string) {
    datastr := m.downloader.GetData(ids, start, end)
    if len(datastr) == 0 {
        m.logger.Error("Cannot get data of: ", ids, start, end)
    }
    m.WriteData(ids, start, end, datastr)
    nsvalue := m.parser.ParseData(datastr)
    m.WriteValue(ids, start, end, nsvalue.TableData)
}

func (m *NationStatManager) WriteIndexData(id string, level int, content string) {
    format := "../data/index-%v-%v.dat"
    filename := fmt.Sprintf(format, level, id)
    util.WriteFile(filename, content)
}

func (m *NationStatManager) WriteData(ids []string, start, end, content string) {
    idstr := strings.Join(ids, "-")
    format := "../data/data-%v-%v-%v.dat"
    filename := fmt.Sprintf(format, idstr, start, end)
    util.WriteFile(filename, content)
}

func (m *NationStatManager) WriteValue(ids []string, start string, end string, data map[string]string) {
    idstr := strings.Join(ids, "-")
    format := "../data/actualdata-%v-%v-%v.dat"
    filename := fmt.Sprintf(format, idstr, start, end)
    var content string
    for k, v := range data{
        content += "key: " + k + "|value: " + v + ","
    }
    util.WriteFile(filename, content)
}

func NewNationStatManager() *NationStatManager{
    m := new(NationStatManager)
    m.Init()

    return m
}
