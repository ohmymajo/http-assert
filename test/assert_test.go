package test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	assert "github.com/ohmymajo/http-assert"
)

func TestAssertStatus(t *testing.T) {
	resp := http.Response{
		StatusCode: 200,
	}

	http := assert.New(&resp)
	val := http.AssertStatus(200)
	if !val {
		t.Fail()
	}
}

func TestAssertStatusFail(t *testing.T) {
	resp := http.Response{
		StatusCode: 201,
	}

	http := assert.New(&resp)
	val := http.AssertStatus(200)
	if val {
		t.Fail()
	}
}

func TestAssertHeader(t *testing.T) {
	header := http.Header{}
	header.Add("x-test-value", "test")

	resp := http.Response{
		Header: header,
	}

	http := assert.New(&resp)
	httpHeader := http.AssertHeader()
	val := httpHeader.
		Has("x-test-value").
		Where("x-test-value", "test").
		Check()

	if !val {
		t.Fail()
	}
}

func TestAssertHeaderFail(t *testing.T) {
	header := http.Header{}
	header.Add("x-test-value", "test")

	resp := http.Response{
		Header: header,
	}

	http := assert.New(&resp)
	httpHeader := http.AssertHeader()
	val := httpHeader.
		Has("x-test-value").
		Where("x-test-value", "123").
		Check()

	if val {
		t.Fail()
	}
}

func TestAssertBodyObjectHas(t *testing.T) {
	b := []byte(`{"message": "Hello World", "obj": {"str": "Hellow"}}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		Has("message").
		Has("obj.str").
		Check()
	if !val {
		t.Fail()
	}
}

func TestAssertBodyObjectHasAll(t *testing.T) {
	b := []byte(`{"message": "Hello World", "obj": {"str": "Hellow"}}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		HasAll([]string{"message", "obj", "obj.str"}).
		Check()
	if !val {
		t.Fail()
	}
}

func TestAssertBodyObjectHasAllFail(t *testing.T) {
	b := []byte(`{"message": "Hello World", "obj": {"str": "Hellow"}}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		HasAll([]string{"message", "obj", "obj.strs"}).
		Check()
	if val {
		t.Fail()
	}
}

func TestAssertBodyObjectWhere(t *testing.T) {
	b := []byte(`{"int": 123, "float": 1.5, "bool": false, "str": "Hello World", "object": {"str": "Hello World"}, "arr": [{"str": "array"}]}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		Where("int", 123).
		Where("float", 1.5).
		Where("bool", false).
		Where("str", "Hello World").
		Where("object.str", "Hello World").
		Where("arr.0.str", "array").
		Check()
	if !val {
		t.Fail()
	}
}

func TestAssertBodyObjectWhereNot(t *testing.T) {
	b := []byte(`{"int": 123, "float": 1.5, "bool": false, "str": "Hello World", "object": {"str": "Hello World"}, "arr": [{"str": "array"}]}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereNot("int", 1234).
		WhereNot("object.str", "Hello Worlds").
		Check()
	if !val {
		t.Fail()
	}
}

func TestAssertBodyWhereType(t *testing.T) {
	b := []byte(`{"int": 1, "str": "Hello", "obj": {"str": false}, "arr":[{"idx": 1.5}]}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereType("int", "int").
		WhereType("str", "string").
		WhereType("obj.str", "bool").
		WhereType("arr.0.idx", "float").
		Check()

	if !val {
		t.Fail()
	}
}

func TestAssertHasLength(t *testing.T) {
	b := []byte(`{"arr": [1, 2, 3]}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		HasLength("arr", 3).
		Check()

	if !val {
		t.Fail()
	}
}

func TestAssertWhereGte(t *testing.T) {
	b := []byte(`{"int": 2}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereGte("int", 3).
		WhereGte("obj.int", 3).
		Check()

	if !val {
		t.Fail()
	}
}

func TestAssertWhereGteFail(t *testing.T) {
	b := []byte(`{"int": 2, "obj": {"int": 2}}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereGte("int", 1).
		WhereGte("int", 3).
		Check()

	if val {
		t.Fail()
	}
}

func TestAssertWhereGt(t *testing.T) {
	b := []byte(`{"int": 2, "obj":{"int":2}}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereGt("int", 3).
		WhereGt("obj.int", 3).
		Check()

	if !val {
		t.Fail()
	}
}

func TestAssertWhereGtFail(t *testing.T) {
	b := []byte(`{"int": 2}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereGt("int", 2).
		WhereGt("int", 3).
		Check()

	if val {
		t.Fail()
	}
}

func TestAssertWhereLte(t *testing.T) {
	b := []byte(`{"int": 2, "obj":{"int":2}}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereLte("int", 1).
		WhereLte("obj.int", 1).
		Check()

	if !val {
		t.Fail()
	}
}

func TestAssertWhereLteFail(t *testing.T) {
	b := []byte(`{"int": 2}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereLte("int", 3).
		Check()

	if val {
		t.Fail()
	}
}

func TestAssertWhereLt(t *testing.T) {
	b := []byte(`{"int": 2, "obj":{"int":2}}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereLt("int", 1).
		WhereLt("obj.int", 1).
		Check()

	if !val {
		t.Fail()
	}
}

func TestAssertWhereLtFail(t *testing.T) {
	b := []byte(`{"int": 2}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.
		WhereLt("int", 2).
		Check()

	if val {
		t.Fail()
	}
}
