package stmanager

type ListManager struct {
    ch chan int
}

func (m *ListManager) Init() {
    m.ch = make(chan int)
}

func (m *ListManager) Process() {
    szm := NewSZSEListManager()
    go func() {
        szm.Process()
        m.ch <- 1
    }()

    shm := NewSHSEListManager()
    go func() {
        shm.Process()
        m.ch <- 2
    }()

    <-m.ch
    <-m.ch
}

func NewListManager() *ListManager{
    m := new(ListManager)
    m.Init()

    return m
}
