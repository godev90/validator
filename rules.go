package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/godev90/validator/faults"
	"github.com/godev90/validator/typedef"
)

var (
	digitRe    = regexp.MustCompile(`^\d+$`)
	alphanumRe = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	alphaRe    = regexp.MustCompile(`^[a-zA-Z]+$`)
	emailRe    = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	nameRe     = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_ '. ,]*$`)
	textRe     = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_ '. ,!@#$%^&*()\-+=\[\]{}:;<>?/|~]*$`)
)

func requiredRule(value any, _ string) error {
	if value == nil {
		return faults.ErrRequired
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return faults.ErrRequired
	}
	if isZero(v) {
		return faults.ErrRequired
	}
	return nil
}

func minRule(value any, param string) error {
	minVal, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return faults.ErrInvalidParameter.Render(param)
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		if f, err := strconv.ParseFloat(val.String(), 64); err == nil {
			if f < minVal {
				return faults.ErrBelowMinimum.Render(int64(minVal))
			}
		} else {
			return faults.ErrInvalidNumericFormat
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(val.Int()) < minVal {
			return faults.ErrBelowMinimum.Render(int64(minVal))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(val.Uint()) < minVal {
			return faults.ErrBelowMinimum.Render(int64(minVal))
		}

	case reflect.Float32, reflect.Float64:
		if val.Float() < minVal {
			return faults.ErrBelowMinimum.Render(int64(minVal))
		}

	default:
		if v, ok := value.(typedef.Validatable); ok {
			if err := v.Err(); err != nil {
				return err
			}
		}

		s := fmt.Sprintf("%v", value)
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			if f < minVal {
				return faults.ErrBelowMinimum.Render(int64(minVal))
			}
		} else {
			return faults.ErrInvalidNumericFormat
		}
	}

	return nil
}

func maxRule(value any, param string) error {
	maxVal, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return faults.ErrInvalidParameter.Render(param)
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		if f, err := strconv.ParseFloat(val.String(), 64); err == nil {
			if f > maxVal {
				return faults.ErrAboveMaximum.Render(int64(maxVal))
			}
		} else {
			return faults.ErrInvalidNumericFormat
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(val.Int()) > maxVal {
			return faults.ErrAboveMaximum.Render(int64(maxVal))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(val.Uint()) > maxVal {
			return faults.ErrAboveMaximum.Render(int64(maxVal))
		}

	case reflect.Float32, reflect.Float64:
		if val.Float() > maxVal {
			return faults.ErrAboveMaximum.Render(int64(maxVal))
		}

	default:
		if errVal, ok := value.(typedef.Validatable); ok {
			if err := errVal.Err(); err != nil {
				return err
			}
		}

		s := fmt.Sprintf("%v", value)
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			if f > maxVal {
				return faults.ErrAboveMaximum.Render(int64(maxVal))
			}
		} else {
			return faults.ErrInvalidNumericFormat
		}
	}

	return nil
}

func minlenRule(value any, param string) error {
	min, _ := strconv.Atoi(param)
	if s, ok := value.(string); ok && len(s) < min {
		return faults.ErrLengthBelowMinimum.Render(min)
	}
	return nil
}

func maxlenRule(value any, param string) error {
	max, _ := strconv.Atoi(param)
	if s, ok := value.(string); ok && len(s) > max {
		return faults.ErrLengthAboveMaximum.Render(max)
	}
	return nil
}

func emailRule(value any, _ string) error {
	if s, ok := value.(string); ok && !emailRe.MatchString(s) {
		return faults.ErrMustBeEmail
	}
	return nil
}

func digitRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !digitRe.MatchString(s) {
		return faults.ErrMustBeDigit
	}
	return nil

}

func alphanumRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !alphanumRe.MatchString(s) {
		return faults.ErrMustBeAlphanum
	}
	return nil
}

func alphabetRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !alphaRe.MatchString(s) {
		return faults.ErrMustBeAlphabet
	}
	return nil
}

func dateRule(value any, _ string) error {
	var s string

	switch v := value.(type) {
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return faults.ErrInvalidDateFormat
	}

	const format = "2006-01-02"

	if _, err := time.Parse(format, s); err != nil {
		return faults.ErrInvalidDateFormat
	}

	return nil
}

func datetimeRule(value any, _ string) error {
	var s string

	switch v := value.(type) {
	case string:
		s = v
	case fmt.Stringer:
		s = v.String()
	default:
		return faults.ErrInvalidDatetimeFormat
	}

	layout := "2006-01-02 15:04:05"

	_, err := time.Parse(layout, s)
	if err != nil {
		return faults.ErrInvalidDatetimeFormat
	}
	return nil

}

func nameRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !nameRe.MatchString(s) {
		return faults.ErrMustBeName
	}
	return nil
}

func textRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !textRe.MatchString(s) {
		return faults.ErrMustBeText
	}
	return nil
}

func oneOfRule(value any, param string) error {
	allowed := splitByPipe(param)

	valStr := fmt.Sprintf("%v", value)

	for _, v := range allowed {
		if valStr == v {
			return nil
		}
	}

	return faults.ErrMustBeOneOf.Render(param)
}

func splitByPipe(param string) []string {
	raw := regexp.MustCompile(`\s*\|\s*`).Split(param, -1)
	out := make([]string, 0, len(raw))
	for _, v := range raw {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
