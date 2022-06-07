package goutils

import "reflect"

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

func ContainsValue(m map[string]string, v string) (bool, string) {
	for k, x := range m {
		if x == v {
			return true, k
		}
	}
	return false, ""
}

func ContainsStructFieldValue(slice interface{}, fields ...KeyValue) int {

	rangeOnMe := reflect.ValueOf(slice)

	for i := 0; i < rangeOnMe.Len(); i++ {
		s := rangeOnMe.Index(i)
		matched := true
		for j := 0; j < len(fields); j++ {
			fieldName := fields[j].Key
			fieldValueToCheck := fields[j].Value
			f := s.FieldByName(fieldName)
			if f.Kind() == reflect.Ptr {
				f = f.Elem()
			}
			if f.IsValid() {
				fval := f.Interface()
				if fval != fieldValueToCheck {
					matched = false
					break
				}
			}
		}
		if matched {
			return i + 1
		}
	}
	return 0
}
