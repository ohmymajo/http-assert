package assert

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ohmymajo/http-assert/pkg/filter"
	"github.com/ohmymajo/http-assert/pkg/validation"
)

type Http struct {
	Resp *http.Response
}

type HttpJson struct {
	Type          string
	Header        *http.Header
	Body          interface{}
	AssertCorrect bool
}

func New(resp *http.Response) Http {
	return Http{
		Resp: resp,
	}
}

func (h Http) AssertStatus(statusCode int) bool {
	return h.Resp.StatusCode == statusCode
}

func (h Http) AssertHeader() HttpJson {
	return HttpJson{
		Type:          "header",
		Header:        &h.Resp.Header,
		AssertCorrect: true,
	}
}

func (h Http) AssertBody() HttpJson {
	var body interface{}

	d := json.NewDecoder(h.Resp.Body)
	d.UseNumber()
	err := d.Decode(&body)
	if err != nil {
		panic("cannot decode json data")
	}

	return HttpJson{
		Type:          "body",
		Body:          body,
		AssertCorrect: true,
	}
}

func (h HttpJson) Has(fieldName string) HttpJson {
	var correct bool
	if h.Type == "header" && h.AssertCorrect {
		correct = h.Header.Get(fieldName) != ""
	} else if h.Type == "body" && h.AssertCorrect {
		f := strings.Split(fieldName, ".")
		t := validation.GetBodyType(h.Body)

		if t == "" {
			panic("cannot read the response body")
		} else if t == "object" {
			if len(f) == 1 {
				b := h.Body.(map[string]interface{})
				for key := range b {
					if key == fieldName {
						correct = true
						break
					}
				}
			} else {
				v := filter.Find(fieldName, h.Body)
				if v != nil {
					correct = true
				}
			}
		} else {
			panic("body should be JSON object")
		}
	}

	return HttpJson{
		Type:          h.Type,
		Header:        h.Header,
		Body:          h.Body,
		AssertCorrect: correct,
	}
}

func (h HttpJson) HasLength(fieldName string, length int) HttpJson {
	var correct bool
	if h.Type == "body" && h.AssertCorrect {
		f := strings.Split(fieldName, ".")
		t := validation.GetBodyType(h.Body)

		if t == "" {
			panic("cannot read the response body")
		} else if t == "object" {
			if len(f) == 1 {
				b := h.Body.(map[string]interface{})
				for key, val := range b {
					if key == fieldName {
						correct = len(val.([]interface{})) == length
						break
					}
				}
			} else {
				v := filter.Find(fieldName, h.Body)
				if v != nil {
					correct = len(v.([]interface{})) == length
				}
			}
		} else {
			panic("body should be JSON object")
		}
	}

	return HttpJson{
		Type:          h.Type,
		Header:        h.Header,
		Body:          h.Body,
		AssertCorrect: correct,
	}
}

func (h HttpJson) Where(fieldName string, value interface{}) HttpJson {
	var correct bool
	if h.Type == "header" && h.AssertCorrect {
		hVal := h.Header.Get(fieldName)
		correct = hVal == value.(string)
	} else if h.Type == "body" && h.AssertCorrect {
		f := strings.Split(fieldName, ".")
		t := validation.GetBodyType(h.Body)

		if t == "" {
			panic("cannot read the response body")
		} else if t == "object" {
			if len(f) == 1 {
				b := h.Body.(map[string]interface{})
				eq := false
				for key, val := range b {
					if key == fieldName {
						vType := validation.GetValueType(value)
						eq = validation.EqualValue(value, val, vType)

						break
					}
				}

				correct = eq
			} else {
				v := filter.Find(fieldName, h.Body)

				t = validation.GetValueType(value)
				correct = validation.EqualValue(v, value, t)
			}
		} else {
			panic("body should be JSON object")
		}
	}

	return HttpJson{
		Type:          h.Type,
		Header:        h.Header,
		Body:          h.Body,
		AssertCorrect: correct,
	}
}

func (h HttpJson) WhereType(fieldName, valueType string) HttpJson {
	var correct bool
	if h.Type == "header" && h.AssertCorrect {
		hVal := h.Header.Get(fieldName)
		t := validation.GetValueType(hVal)

		correct = t == valueType
	} else if h.Type == "body" && h.AssertCorrect {
		f := strings.Split(fieldName, ".")
		t := validation.GetBodyType(h.Body)

		if t == "" {
			panic("cannot read the response body")
		} else if t == "object" {
			if len(f) == 1 {
				b := h.Body.(map[string]interface{})
				for key, val := range b {
					if key == fieldName {
						t := validation.GetValueType(val)
						correct = t == valueType

						break
					}
				}
			} else {
				v := filter.Find(fieldName, h.Body)
				t = validation.GetValueType(v)

				correct = t == valueType
			}
		} else {
			panic("body should be JSON object")
		}
	}

	return HttpJson{
		Type:          h.Type,
		Header:        h.Header,
		Body:          h.Body,
		AssertCorrect: correct,
	}
}

func (h HttpJson) Check() bool {
	return h.AssertCorrect
}
