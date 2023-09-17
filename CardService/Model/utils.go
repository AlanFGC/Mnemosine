package Model

import "reflect"

func ToInterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("ToInterfaceSlice: not a slice")
	}

	result := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		result[i] = s.Index(i).Interface()
	}
	return result
}
