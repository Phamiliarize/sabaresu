package data

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

/*
 * This is a collection of validators used when validating incoming data against a schema.
 * Currently, only Numbers (float64), Strings, Booleans, and Slices should be supported
 * keeping in line with JSON and our target runtime of Lua.
 *
 * Try to use type inferences where possible and reflection minimally- but also don't be afraid
 * to use reflection. It'll still be faster than other options including waiting till LuaRuntime.
 */

var lookup = map[string]func(rule string, value any) error{
	"required": Required,
	"string":   Type,
	"number":   Type,
	"boolean":  Type,
	"min":      Min,
	"max":      Max,
	"int":      Int,
	"uuid":     UUID,
}

func GetValidator(rule string) func(rule string, value any) error {
	r := strings.Split(rule, "=")[0]

	return lookup[r]
}

type ValidatorError struct {
	Rule       string
	Validator  string
	Constraint string
	Value      any
	Detail     string
}

func (v ValidatorError) Error() string {
	if v.Detail != "" {
		return v.Detail
	}

	return fmt.Sprintf("Error validating %v for %s", v.Value, v.Rule)
}

func Required(rule string, value any) error {
	err := ValidatorError{Rule: rule, Validator: rule, Constraint: "", Value: value}
	if value != nil {
		return nil
	}

	return &err
}

func Type(rule string, value any) error {
	err := ValidatorError{Rule: rule, Validator: rule, Constraint: "", Value: value}
	switch v := value.(type) {
	case string:
		if rule == "string" {
			return nil
		}
	case float64:
		if rule == "number" {
			return nil
		}
	case bool:
		if rule == "boolean" {
			return nil
		}
	default:
		err.Detail = fmt.Sprintf("%T not supported for %s validator", v, rule)
	}

	return &err
}

func Min(rule string, value any) error {
	validator := rule[:3]
	constraint := rule[4:]
	err := ValidatorError{Rule: rule, Validator: validator, Constraint: constraint, Value: value}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Float64:
		min, _ := strconv.ParseFloat(constraint, 64)
		if v.Float() > min {
			return nil
		}
	case reflect.String:
		min, _ := strconv.Atoi(constraint)
		if len(v.String()) > min {
			return nil
		}
	case reflect.Slice:
		min, _ := strconv.Atoi(constraint)
		if v.Len() > min {
			return nil
		}
	default:
		err.Detail = fmt.Sprintf("%T not supported for %s validator", value, validator)
	}

	return &err
}

func Max(rule string, value any) error {
	validator := rule[:3]
	constraint := rule[4:]
	err := ValidatorError{Rule: rule, Validator: validator, Constraint: constraint, Value: value}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Float64:
		max, _ := strconv.ParseFloat(constraint, 64)
		if v.Float() < max {
			return nil
		}
	case reflect.String:
		max, _ := strconv.Atoi(constraint)
		if len(v.String()) < max {
			return nil
		}
	case reflect.Slice:
		max, _ := strconv.Atoi(constraint)
		if v.Len() < max {
			return nil
		}
	default:
		err.Detail = fmt.Sprintf("%T not supported for %s validator", value, validator)
	}

	return &err
}

func Int(rule string, value any) error {
	validator := rule[:3]
	err := ValidatorError{Rule: rule, Validator: validator, Constraint: "", Value: value}

	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Float64:
		float := v.Float()
		intVal := int64(float)
		if float64(intVal) == float {
			return nil
		}
	case reflect.String:
		_, err := strconv.Atoi(v.String())
		if err == nil {
			return nil
		}
	default:
		err.Detail = fmt.Sprintf("%T not supported for %s validator", value, validator)
	}

	return &err
}

func UUID(rule string, value any) error {
	err := ValidatorError{Rule: rule, Validator: rule, Constraint: "", Value: value}

	v, ok := value.(string)
	if ok {
		_, err := uuid.Parse(v)
		if err == nil {
			return nil
		}
	} else {
		err.Detail = fmt.Sprintf("%T not supported for %s validator", value, rule)
	}

	return &err
}
