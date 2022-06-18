package goutils

import "reflect"

func GetField(v interface{}, field string) interface{} {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Map:
		return InterfaceMap(v)[field]
	case reflect.Struct:
		r := reflect.ValueOf(v)
		f := reflect.Indirect(r).FieldByName(field)
		return f.Interface()
	default:
		return v
	}
	return nil
}

func GetFields(v interface{}, fields []string) interface{} {
	if len(fields) < 2 {
		return GetField(v, fields[0])
	}
	myV := GetField(v, fields[0])
	return GetFields(myV, fields[1:])
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice && s.Kind() != reflect.Array {
		panic("InterfaceSlice() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	//if s.IsNil() {
	//	return nil
	//}

	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func InterfaceMap(mp interface{}) map[string]interface{} {
	s := reflect.ValueOf(mp)
	if s.Kind() != reflect.Map {
		panic("InterfaceMap() given a non-slice type")
	}

	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}

	ret := make(map[string]interface{})
	mpRange := s.MapRange()
	for mpRange.Next() {
		ret[mpRange.Key().String()] = mpRange.Value().Interface()
	}

	return ret
}

func Search(sov interface{}, key string, value interface{}) int {
	s := InterfaceSlice(sov)
	for ix, v := range s {
		vv := GetField(v, key)
		if reflect.ValueOf(vv).Kind() == reflect.Ptr {
			vv = reflect.Indirect(reflect.ValueOf(vv)).Interface()
		}
		if vv == value {
			return ix
		}
	}
	return -1
}

func MultiSearch(pObj interface{}, pFields ...KeyValue) int {
	if pObj == nil {
		return -1
	}
	s := InterfaceSlice(pObj)
	for ix, v := range s {
		matched := true
		for _, field := range pFields {
			vv := GetField(v, field.Key)
			if reflect.ValueOf(vv).Kind() == reflect.Ptr {
				vv = reflect.Indirect(reflect.ValueOf(vv)).Interface()
			}
			if vv != field.Value {
				matched = false
			}
		}

		if matched {
			return ix
		}
	}
	return -1
}

func Searchs(sov interface{}, key []string, value interface{}) int {
	s := InterfaceSlice(sov)
	for ix, v := range s {
		vv := GetFields(v, key)
		if vv == value {
			return ix
		}
	}
	return -1
}

func Contains(sov interface{}, key string, value interface{}) bool {
	return Search(sov, key, value) > -1
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
