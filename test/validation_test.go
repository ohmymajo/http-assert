package test

import (
	"encoding/json"
	"testing"

	"github.com/ohmymajo/http-assert/pkg/validation"
)

func TestGetValueTypeArrayString(t *testing.T) {
	val := validation.GetValueType([]string{"a", "b"})
	if val != "array-string" {
		t.Fail()
	}
}

func TestGetValueTypeArrayInt(t *testing.T) {
	val := validation.GetValueType([]int{1, 2})
	if val != "array-int" {
		t.Fail()
	}
}

func TestGetValueTypeArrayFloat(t *testing.T) {
	val := validation.GetValueType([]float64{1.1, 1.2})
	if val != "array-float" {
		t.Fail()
	}
}

func TestGetValueTypeString(t *testing.T) {
	val := validation.GetValueType("Hello World")
	if val != "string" {
		t.Fail()
	}
}

func TestGetValueTypeInt(t *testing.T) {
	val := validation.GetValueType(1000)
	if val != "int" {
		t.Fail()
	}
}

func TestGetValueTypeFloat(t *testing.T) {
	val := validation.GetValueType(3.14)
	if val != "float" {
		t.Fail()
	}
}

func TestGetValueTypeBool(t *testing.T) {
	val := validation.GetValueType(true)
	if val != "bool" {
		t.Fail()
	}
}

func TestGetValueTypeObject(t *testing.T) {
	b := []byte(`{"message": "Hello World"}`)

	var data interface{}
	json.Unmarshal(b, &data)

	val := validation.GetValueType(data)
	if val != "object" {
		t.Fail()
	}
}

func TestGetValueTypeArrayObject(t *testing.T) {
	b := []byte(`[{"message": "Hello World"}]`)

	var data interface{}
	json.Unmarshal(b, &data)

	val := validation.GetValueType(data)
	if val != "array-object" {
		t.Fail()
	}
}
