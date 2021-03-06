package goutils

type KeyValue struct {
	Key   string
	Value interface{}
}

func NewPointer[T any](v T) *T {
	return &v
}

func Bool2Int(pVar bool) int64 {
	if pVar {
		return 1
	} else {
		return 0
	}
}
