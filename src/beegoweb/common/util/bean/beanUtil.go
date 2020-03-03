package beanutil

import (
	"fmt"
	"reflect"
)

//利用反射
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("non-struct type")
		return nil
	}
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructToMapRmNil(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("non-struct type")
		return nil
	}
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		val := v.Field(i).Interface()
		if val == nil || val == "" || val == 0 {
			continue
		}
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

//
func MapRmNil(data map[string]interface{}) {
	for k, v := range data {
		if v == nil {
			delete(data, k)
		} else if v == float64(0) || v == "" {
			delete(data, k)
		}
	}
}
