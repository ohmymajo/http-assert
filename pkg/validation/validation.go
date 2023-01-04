package validation

import (
	"fmt"
	"reflect"
	"strconv"
)

func GetBodyType(b interface{}) string {
	rv := reflect.ValueOf(b)

	switch rv.Kind() {
	case reflect.Slice:
		return "array-object"
	case reflect.Map:
		return "object"
	default:
		return ""
	}
}

func GetValueType(v interface{}) string {
	rv := reflect.ValueOf(v)

	switch rv.Type().String() {
	case "[]string":
		return "array-string"
	case "[]int", "[]int8", "[]int16", "[]int32", "[]int64":
		return "array-int"
	case "[]float64", "[]float32":
		return "array-float"
	case "[]interface {}":
		return "array-object"
	case "bool":
		return "bool"
	case "string":
		return "string"
	case "int", "int8", "int16", "int32", "int64":
		return "int"
	case "float32", "float64":
		return "float"
	case "map[string]interface {}":
		return "object"
	default:
		return ""
	}
}

func EqualValue(a, b interface{}, aType string) bool {
	if aType == "string" {
		return a.(string) == b.(string)
	} else if aType == "int" {
		valA, _ := strconv.Atoi(fmt.Sprintf("%v", a))
		valB, _ := strconv.Atoi(fmt.Sprintf("%v", b))
		return valA == valB
	} else if aType == "float" {
		b, _ := strconv.ParseFloat(fmt.Sprintf("%v", b), 64)
		return a.(float64) == b
	} else if aType == "bool" {
		return a.(bool) == b.(bool)
	}

	return false
}
