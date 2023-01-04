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
	b := []byte(`{"message": "Hello World"}`)

	resp := http.Response{
		Body: io.NopCloser(bytes.NewReader(b)),
	}

	http := assert.New(&resp)
	httpBody := http.AssertBody()
	val := httpBody.Has("message").Check()
	if !val {
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