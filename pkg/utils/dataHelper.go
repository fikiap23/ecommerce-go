package utils

import (
	"reflect"
)

// StructToMap mengubah struct dengan pointer field jadi map[string]interface{}
func StructToMap(input interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	v := reflect.ValueOf(input)
	if v.Kind() != reflect.Struct {
		// Handle pointer to struct
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := t.Field(i)
		jsonTag := fieldType.Tag.Get("json")

		if jsonTag == "" || jsonTag == "-" {
			continue
		}

		jsonKey := jsonTag
		if commaIdx := len(jsonTag); commaIdx > 0 {
			if idx := reflect.StructTag(jsonTag).Get("json"); idx != "" {
				jsonKey = fieldType.Tag.Get("json")
			}
		}
		jsonKey = parseJSONTag(jsonTag)

		if field.Kind() == reflect.Ptr && !field.IsNil() {
			result[jsonKey] = field.Elem().Interface()
		}
	}
	return result
}

func parseJSONTag(tag string) string {
	if tag == "" {
		return ""
	}
	if commaIdx := len(tag); commaIdx > 0 {
		for i, c := range tag {
			if c == ',' {
				return tag[:i]
			}
		}
	}
	return tag
}
