package test

import (
	"encoding/json"
	"testing"

	"github.com/ohmymajo/http-assert/pkg/filter"
	"github.com/ohmymajo/http-assert/pkg/validation"
)

func TestFindObject(t *testing.T) {
	j := []byte(`{"data": { "message": "Hello World" }}`)

	var data interface{}
	json.Unmarshal(j, &data)

	val := filter.Find("data.message", data)

	if val != "Hello World" {
		t.Fail()
	}
}

func TestFindObjectWithIndex(t *testing.T) {
	j := []byte(`{"data": { "options": [1, 2, 3] }}`)

	var data interface{}
	json.Unmarshal(j, &data)

	val := filter.Find("data.options.1", data)
	valid := validation.EqualValue(val, 2, "int")

	if !valid {
		t.Fail()
	}
}

func TestFindObjectWithArrayObject(t *testing.T) {
	j := []byte(`{"data": { "options": [{ "value": 1.5 }] }}`)

	var data interface{}
	json.Unmarshal(j, &data)

	val := filter.Find("data.options.0.value", data)
	valid := validation.EqualValue(val, 1.5, "float")

	if !valid {
		t.Fail()
	}
}
