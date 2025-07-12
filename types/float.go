package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/godev90/validator/errors"
)

type Float struct {
	s   string
	f   float64
	err error
}

func (f Float) String() string {
	return f.s
}

func (f *Float) Set(val any) error {
	f.err = nil // reset error
	switch v := val.(type) {
	case float64:
		f.f = v
		f.s = strconv.FormatFloat(v, 'f', -1, 64)
	case float32:
		return f.Set(float64(v))
	case int:
		return f.Set(float64(v))
	case int64:
		return f.Set(float64(v))
	case json.Number:
		f.s = v.String()
		parsed, err := v.Float64()
		f.f = parsed

		if err != nil {
			f.err = errors.ErrInvalidFloatNumber
		}
	case string:
		if strings.TrimSpace(v) == "" {
			f.s = ""
			f.f = 0
			f.err = nil // treat empty as NULL, not error
			return nil
		}
		f.s = v
		parsed, err := strconv.ParseFloat(v, 64)
		f.f = parsed
		if err != nil {
			f.err = errors.ErrInvalidFloatNumber
		}

	default:
		f.s = ""
		f.f = 0
		f.err = errors.ErrInvalidFloatNumber
	}
	return nil
}

func (f *Float) UnmarshalJSON(data []byte) error {
	var raw json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		f.err = err
		return nil
	}

	var floatVal float64
	if err := json.Unmarshal(raw, &floatVal); err == nil {
		_ = f.Set(floatVal)
		return nil
	}

	var strVal string
	if err := json.Unmarshal(raw, &strVal); err == nil {
		_ = f.Set(strVal)
		return nil
	}

	f.err = errors.ErrInvalidFloatNumber
	return nil
}

func (f *Float) UnmarshalText(text []byte) error {
	_ = f.Set(string(text))
	return nil
}

func (f Float) MarshalJSON() ([]byte, error) {

	if !f.Valid() {
		return json.Marshal(nil)
	}

	return json.Marshal(f.f)
}

func (f Float) Float64() float64 {
	return f.f
}

func (f Float) Value() (driver.Value, error) {
	if f.err != nil {
		return nil, f.err
	}
	return f.f, nil
}

func (f *Float) Scan(value any) error {

	_ = f.Set(value)
	return nil
}

func (f Float) Err() error {
	return f.err
}

func (f Float) Valid() bool {
	return f.err == nil
}

func NewFloat(value float64) Float {
	return Float{
		s:   fmt.Sprintf("%.2f", value),
		f:   value,
		err: nil,
	}
}
