package httpc

import (
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

func Object2UrlValues(obj any) url.Values {
	result := url.Values{}

	v := reflect.Indirect(reflect.ValueOf(obj))
	fields := reflect.VisibleFields(v.Type())

	var fieldTag string
	var tagName string
	var fValue reflect.Value
	var fType reflect.Type

	for _, field := range fields {
		if field.Anonymous || !field.IsExported() {
			continue
		}

		fieldTag = field.Tag.Get("form")
		if fieldTag == "" || fieldTag == "-" {
			continue
		}

		tagName = strings.SplitN(fieldTag, ",", 2)[0]
		fValue = v.FieldByIndex(field.Index)
		fType = field.Type

		if fType.Kind() == reflect.Pointer {
			if fValue.IsNil() {
				continue
			}

			fValue = fValue.Elem()
			fType = fType.Elem()
		}

		switch fType.Kind() {
		case reflect.Slice, reflect.Array:
			strSlice := make([]string, fValue.Len())
			for i := 0; i < len(strSlice); i++ {
				strSlice[i] = fmt.Sprintf("%v", fValue.Index(i).Interface())
			}
			result[tagName] = strSlice
		default:
			result.Set(tagName, fmt.Sprintf("%v", fValue.Interface()))
		}
	}

	return result
}
