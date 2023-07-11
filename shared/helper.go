package shared

import (
	"fmt"
	"reflect"
)

// a function to create a *[]*string
func StringSlice(s string) *[]*string {
	return &[]*string{&s}
}

func StructToMap(item interface{}) map[string]string {
	result := map[string]string{}
	elem := reflect.ValueOf(item).Elem()

	for i := 0; i < elem.NumField(); i++ {
		result[elem.Type().Field(i).Name] = fmt.Sprintf("%v", elem.Field(i).Interface())
	}

	return result
}
