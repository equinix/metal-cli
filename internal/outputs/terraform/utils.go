package terraform

import (
	"fmt"
	"reflect"
	"time"
)

func addQuotesToString(v interface{}) {
	val := reflect.ValueOf(v)

	switch val.Kind() {
	case reflect.Ptr:
		val = val.Elem()
		if val.Kind() != reflect.Struct {
			return
		}

		if val.Type() == reflect.TypeOf(new(string)) {
			oldValue := val.Elem().String()
			newValue := fmt.Sprintf(`"%s"`, oldValue)
			val.Elem().SetString(newValue)
			return
		}
	case reflect.String:
		oldValue := val.String()
		newValue := fmt.Sprintf(`"%s"`, oldValue)
		val.SetString(newValue)
		return
	case reflect.Slice:
		for i := 0; i < val.Len(); i++ {
			elem := val.Index(i)
			if elem.Kind() == reflect.Struct || (elem.Kind() == reflect.Ptr && elem.Elem().Kind() == reflect.Struct) {
				addQuotesToString(elem.Interface())
			}
		}
		return
	case reflect.Map:
		for _, key := range val.MapKeys() {
			elem := val.MapIndex(key)
			if elem.Kind() == reflect.Struct || (elem.Kind() == reflect.Ptr && elem.Elem().Kind() == reflect.Struct) {
				addQuotesToString(elem.Interface())
			}
		}
		return
	default:
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		switch field.Kind() {
		case reflect.String:
			oldValue := field.String()
			newValue := fmt.Sprintf(`"%s"`, oldValue)
			field.SetString(newValue)
		case reflect.Ptr:
			if field.IsNil() {
				continue
			}
			// Check if the pointer is to a string
			if field.Type().Elem() == reflect.TypeOf("") {
				oldValue := field.Elem().String()
				newValue := fmt.Sprintf(`"%s"`, oldValue)
				field.Elem().SetString(newValue)
			} else {
				// Exclude *time.Time from recursion
				if field.Type() != reflect.TypeOf(&time.Time{}) {
					addQuotesToString(field.Interface())
				}
			}
		case reflect.Struct:
			// Exclude time.Time from recursion
			if field.Type() != reflect.TypeOf(time.Time{}) {
				addQuotesToString(field.Interface())
			}
		}
	}
}

func nullIfNilOrEmpty(v interface{}) interface{} {
    if v == nil {
        return "null"
    }

    // Use reflection to check if the value is an empty value (e.g., empty string, empty slice, or empty map)
    val := reflect.ValueOf(v)
    switch val.Kind() {
    case reflect.String:
        if val.String() == "\"\"" {
            return "null"
        }
    case reflect.Array, reflect.Slice, reflect.Map:
        if val.Len() == 0 {
            return "null"
        }
    case reflect.Ptr:
        if val.IsNil() || val.IsZero() {
            return "null"
        }

        elem := val.Elem()

        // Check if it's a pointer to a string
        if elem.Kind() == reflect.String && elem.String() == "\"\"" {
            return "null"
        }

        switch elem.Kind() {
        case reflect.Struct:
			if isPointerStructEmpty(elem) {
                return "null"
            }
        case reflect.Array, reflect.Slice:
            if elem.Len() == 0 {
                return "null"
            }
        case reflect.Map:
            if elem.Len() == 0 {
                return "null"
            }
        }
    }

    return v
}

func isPointerStructEmpty(structVal reflect.Value) bool {
    // Iterate through the struct fields
    for i := 0; i < structVal.NumField(); i++ {
        field := structVal.Field(i)

        // You can define custom logic to determine if a field is empty
        // For example, check if a string field is empty, or if a slice/map field is empty
        switch field.Kind() {
        case reflect.String:
            if field.String() != "" {
                return false
            }
        case reflect.Slice, reflect.Map:
            if field.Len() > 0 {
                return false
            }
        // Add more cases for other field types as needed
        }
    }

    // All fields are empty
    return true
}
