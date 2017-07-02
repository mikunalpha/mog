package utils

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var validationMethods = map[string]func(interface{}, string) error{
	"required":  IsRequired,
	"size":      IsSize,
	"max":       IsMax,
	"min":       IsMin,
	"sometimes": IsSet,
	"regex":     IsRegex,
}

// IsNonzero validates wheather v is nonzero value
func IsRequired(v interface{}, param string) error {
	value := reflect.ValueOf(v)
	valid := true

	switch value.Kind() {
	case reflect.String:
		valid = len(value.String()) != 0
	case reflect.Ptr, reflect.Interface:
		valid = !value.IsNil()
	case reflect.Slice, reflect.Map, reflect.Array:
		valid = value.Len() != 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valid = value.Int() != 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		valid = value.Uint() != 0
	case reflect.Float32, reflect.Float64:
		valid = value.Float() != 0
	case reflect.Bool:
		valid = value.Bool()
	case reflect.Invalid:
		valid = false // always invalid
	case reflect.Struct:
		valid = true // always valid since only nil pointers are empty
	default:
		return fmt.Errorf("Unsupported field type.")
	}

	if !valid {
		return fmt.Errorf("This filed is required.")
	}

	return nil
}

// IsSize validates wheather v is equal to param
func IsSize(v interface{}, param string) error {
	value := reflect.Indirect(reflect.ValueOf(v))
	valid := true

	invalidParamError := fmt.Errorf("The type of param value is invalid.")

	switch value.Kind() {
	case reflect.String:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		valid = int64(len(value.String())) == p
	case reflect.Slice, reflect.Map, reflect.Array:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		valid = int64(value.Len()) == p
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		valid = value.Int() == p
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p, err := strconv.ParseUint(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		valid = value.Uint() == p
	case reflect.Float32, reflect.Float64:
		p, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return invalidParamError
		}
		valid = value.Float() == p
	case reflect.Invalid:
		valid = false
	default:
		return fmt.Errorf("Unsupported field type.")
	}

	if !valid {
		return fmt.Errorf("This filed size should match %s.", param)
	}

	return nil
}

// IsMin validates wheather v is equal to or larger than param
func IsMin(v interface{}, param string) error {
	value := reflect.Indirect(reflect.ValueOf(v))
	invalid := false

	invalidParamError := fmt.Errorf("The type of param value is invalid.")

	switch value.Kind() {
	case reflect.String:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = int64(len(value.String())) < p
	case reflect.Slice, reflect.Map, reflect.Array:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = int64(value.Len()) < p
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = value.Int() < p
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p, err := strconv.ParseUint(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = value.Uint() < p
	case reflect.Float32, reflect.Float64:
		p, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = value.Float() < p
	case reflect.Invalid:
		return invalidParamError
	default:
		return fmt.Errorf("Unsupported field type.")
	}

	if invalid {
		return fmt.Errorf("This filed size should be not lower than %s.", param)
	}

	return nil
}

// IsMax validates wheather v is equal to or lower than param
func IsMax(v interface{}, param string) error {
	value := reflect.Indirect(reflect.ValueOf(v))
	invalid := false

	invalidParamError := fmt.Errorf("The type of param value is invalid.")

	switch value.Kind() {
	case reflect.String:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = int64(len(value.String())) > p
	case reflect.Slice, reflect.Map, reflect.Array:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = int64(value.Len()) > p
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		p, err := strconv.ParseInt(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = value.Int() > p
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		p, err := strconv.ParseUint(param, 0, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = value.Uint() > p
	case reflect.Float32, reflect.Float64:
		p, err := strconv.ParseFloat(param, 64)
		if err != nil {
			return invalidParamError
		}
		invalid = value.Float() > p
	case reflect.Invalid:
		return invalidParamError
	default:
		return fmt.Errorf("Unsupported field type.")
	}

	if invalid {
		return fmt.Errorf("This filed size should be not larger than %s.", param)
	}

	return nil
}

// IsRegex validates wheather v is matched to param pattern
func IsRegex(v interface{}, param string) error {
	value := reflect.Indirect(reflect.ValueOf(v))

	if value.Kind() != reflect.String {
		return fmt.Errorf("The type of filed value is invalid.")
	}

	re, err := regexp.Compile(param)
	if err != nil {
		return fmt.Errorf("The type of param value is invalid.")
	}

	if !re.MatchString(value.Interface().(string)) {
		return fmt.Errorf("This filed pattern should is invalid.")
	}

	return nil
}

func IsSet(v interface{}, param string) error {
	return IsRequired(v, param)
}

// Validate uses fieldsRules to Validate fields of struct v
func Validate(v interface{}, fieldsRules map[string]string) error {
	value := reflect.ValueOf(v)
	valueType := reflect.TypeOf(v)

	if value.Kind() == reflect.Ptr && !value.IsNil() {
		return Validate(value.Elem().Interface(), fieldsRules)
	}
	if value.Kind() != reflect.Struct {
		return fmt.Errorf("Unsupported field type.")
	}

	fieldsNum := value.NumField()
	for i := 0; i < fieldsNum; i++ {
		fieldName := valueType.Field(i).Name
		if _, ok := fieldsRules[fieldName]; !ok {
			continue
		}

		field := value.Field(i)
		rules := strings.Split(fieldsRules[fieldName], "|")
		for _, rule := range rules {
			ruleName := strings.Trim(rule, " ")
			param := ""

			sep := strings.Index(rule, ":")
			if sep != -1 {
				ruleName = rule[:sep]
				param = rule[sep+1:]
			}

			if _, ok := validationMethods[ruleName]; !ok {
				continue
			}

			if err := validationMethods[ruleName](field.Interface(), param); err != nil {
				if ruleName == "sometimes" {
					break
				}
				return fmt.Errorf("Field '%s' validation is failed. %s", fieldName, err.Error())
			}
		}
	}

	return nil
}
