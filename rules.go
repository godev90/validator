package validator

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/godev90/validator/errors"
	"github.com/godev90/validator/types"
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
		return errors.ErrFieldRequired
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return errors.ErrFieldRequired
	}
	if isZero(v) {
		return errors.ErrFieldRequired
	}
	return nil
}

func minRule(value any, param string) error {
	minVal, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return errors.ErrFieldInvalidParam(param)
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		if f, err := strconv.ParseFloat(val.String(), 64); err == nil {
			if f < minVal {
				return errors.ErrFieldBelowMinimum(int64(minVal))
			}
		} else {
			return errors.ErrInvalidNumericFormat
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(val.Int()) < minVal {
			return errors.ErrFieldBelowMinimum(int64(minVal))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(val.Uint()) < minVal {
			return errors.ErrFieldBelowMinimum(int64(minVal))
		}

	case reflect.Float32, reflect.Float64:
		if val.Float() < minVal {
			return errors.ErrFieldBelowMinimum(int64(minVal))
		}

	default:
		// Cek jika implement Validatable
		if v, ok := value.(types.Validatable); ok {
			if err := v.Err(); err != nil {
				return err
			}
		}

		// Fallback parse via string
		s := fmt.Sprintf("%v", value)
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			if f < minVal {
				return errors.ErrFieldBelowMinimum(int64(minVal))
			}
		} else {
			return errors.ErrInvalidNumericFormat
		}
	}

	return nil
}

func maxRule(value any, param string) error {
	maxVal, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return errors.ErrFieldInvalidParam(param)
	}

	val := reflect.ValueOf(value)

	switch val.Kind() {
	case reflect.String:
		if f, err := strconv.ParseFloat(val.String(), 64); err == nil {
			if f > maxVal {
				return errors.ErrFieldAboveMaximum(int64(maxVal))
			}
		} else {
			return errors.ErrInvalidNumericFormat
		}

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(val.Int()) > maxVal {
			return errors.ErrFieldAboveMaximum(int64(maxVal))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(val.Uint()) > maxVal {
			return errors.ErrFieldAboveMaximum(int64(maxVal))
		}

	case reflect.Float32, reflect.Float64:
		if val.Float() > maxVal {
			return errors.ErrFieldAboveMaximum(int64(maxVal))
		}

	default:
		// Cek jika value punya method Err() error
		if errVal, ok := value.(types.Validatable); ok {
			if err := errVal.Err(); err != nil {
				return err
			}
		}

		// Fallback parse dari string
		s := fmt.Sprintf("%v", value)
		if f, err := strconv.ParseFloat(s, 64); err == nil {
			if f > maxVal {
				return errors.ErrFieldAboveMaximum(int64(maxVal))
			}
		} else {
			return errors.ErrInvalidNumericFormat
		}
	}

	return nil
}

func minlenRule(value any, param string) error {
	min, _ := strconv.Atoi(param)
	if s, ok := value.(string); ok && len(s) < min {
		return errors.ErrFieldLengthBelowMinimum(min)
	}
	return nil
}

func maxlenRule(value any, param string) error {
	max, _ := strconv.Atoi(param)
	if s, ok := value.(string); ok && len(s) > max {
		return errors.ErrFieldLengthAboveMaximum(max)
	}
	return nil
}

func emailRule(value any, _ string) error {
	if s, ok := value.(string); ok && !emailRe.MatchString(s) {
		return errors.ErrFieldMustBeEmail
	}
	return nil
}

func digitRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !digitRe.MatchString(s) {
		return errors.ErrFieldMustBeDigit
	}
	return nil

}

func alphanunRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !alphanumRe.MatchString(s) {
		return errors.ErrFieldMustBeAlphanum
	}
	return nil
}

func alphabetRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !alphaRe.MatchString(s) {
		return errors.ErrFieldMustBeAlphabet
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
		return errors.ErrInvalidDateFormat
	}

	const format = "2006-01-02"

	if _, err := time.Parse(format, s); err != nil {
		return errors.ErrInvalidDateFormat
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
		return errors.ErrInvalidDatetimeFormat
	}

	layout := "2006-01-02 15:04:05"

	_, err := time.Parse(layout, s)
	if err != nil {
		return errors.ErrInvalidDatetimeFormat
	}
	return nil

}

func nameRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !nameRe.MatchString(s) {
		return errors.ErrFieldMustBeName
	}
	return nil
}

func textRule(value any, _ string) error {
	if s, ok := value.(string); !ok || !textRe.MatchString(s) {
		return errors.ErrFieldMustBeText
	}
	return nil
}

func oneOfRule(value any, param string) error {
	// Pecah parameter jadi daftar nilai dipisah "|"
	allowed := splitByPipe(param)

	valStr := fmt.Sprintf("%v", value)

	for _, v := range allowed {
		if valStr == v {
			return nil
		}
	}

	return errors.ErrFieldMustBeOneOf(allowed)
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
