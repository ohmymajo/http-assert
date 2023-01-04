package assert

import (
	"encoding/json"
	"io"
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
	data, _ := io.ReadAll(h.Resp.Body)
	json.Unmarshal(data, &body)

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
		t := validation.GetBodyType(h.Body)
		if t == "" {
			panic("cannot read the response body")
		} else if t == "object" {
			b := h.Body.(map[string]interface{})
			has := false
			for key := range b {
				if key == fieldName {
					has = true
					break
				}
			}

			correct = has
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

func (h HttpJson) Check() bool {
	return h.AssertCorrect
}
