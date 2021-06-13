package mr

type KV struct {
	Key   string
	Value string
}

type MapF func(fileName string, content string) []*KV

// ReduceF accumulate the key's value together
type ReduceF func(key string, values []string) string

type MapReduce interface {
	Run() bool
}

type mapReduceImpl struct {
	mf         MapF
	rf         ReduceF
	splitFiles []string
}

func NewMapReduce(mf MapF, rf ReduceF, files []string) MapReduce {
	return &mapReduceImpl{mf: mf, rf: rf, splitFiles: files}
}

func (impl *mapReduceImpl) Run() bool {
	return false
}
