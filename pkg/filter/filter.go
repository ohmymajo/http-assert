package filter

import (
	"strconv"
	"strings"

	"github.com/ohmymajo/http-assert/pkg/validation"
)

func Find(cursor string, data interface{}) interface{} {
	var value interface{}
	var lf string

	fields := strings.Split(cursor, ".")

	d := data
	for idx, field := range fields {
		numeric, ndx := isInt(field)

		if !numeric {
			dtmp := filter(field, d)
			t := validation.GetBodyType(dtmp)

			if t != "object" {
				value = dtmp
			} else {
				d = dtmp
			}
		} else {
			tmp := d.(map[string]interface{})[lf]
			d = tmp.([]interface{})[ndx]
			if idx == len(fields)-1 {
				value = tmp.([]interface{})[ndx]
			}
		}

		lf = field
	}

	return value
}

func isInt(field string) (bool, int) {
	v, err := strconv.Atoi(field)
	if err == nil {
		return true, v
	}

	return false, 0
}

func filter(field string, data interface{}) interface{} {
	btype := validation.GetBodyType(data)
	if btype == "object" {
		d := data.(map[string]interface{})
		for key := range d {
			if key == field {
				return d[key]
			}
		}
	}

	return nil
}
