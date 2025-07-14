package typedef

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/godev90/validator/faults"
)

type Integer struct {
	s   string
	i   int64
	err error
}

func (i Integer) String() string {
	return i.s
}

func (i *Integer) Set(val any) error {
	i.err = nil // reset error
	switch v := val.(type) {
	case int:
		i.i = int64(v)
		i.s = strconv.FormatInt(i.i, 10)
	case int64:
		i.i = v
		i.s = strconv.FormatInt(v, 10)
	case json.Number:
		i.s = v.String()
		parsed, err := v.Int64()
		i.i = parsed

		if err != nil {
			i.err = faults.ErrInvalidIntegerNumber
		}
	case string:
		if strings.TrimSpace(v) == "" {
			i.s = ""
			i.i = 0
			i.err = nil // treat empty as NULL, not error
			return nil
		}
		i.s = v
		parsed, err := strconv.ParseInt(v, 10, 64)
		i.i = parsed
		if err != nil {
			i.err = faults.ErrInvalidIntegerNumber
		}

	default:
		i.s = ""
		i.i = 0
		i.err = faults.ErrInvalidIntegerNumber
	}
	return nil
}

func (i *Integer) UnmarshalJSON(data []byte) error {
	var raw json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		i.err = err
		return nil
	}

	var intVal int64
	if err := json.Unmarshal(raw, &intVal); err == nil {
		_ = i.Set(intVal)
		return nil
	}

	var strVal string
	if err := json.Unmarshal(raw, &strVal); err == nil {
		_ = i.Set(strVal)
		return nil
	}

	i.err = faults.ErrInvalidIntegerNumber
	return nil
}

func (i *Integer) UnmarshalText(text []byte) error {
	_ = i.Set(string(text))
	return nil
}

func (i Integer) MarshalJSON() ([]byte, error) {
	if !i.Valid() {
		return json.Marshal(nil)
	}

	return json.Marshal(i.i)
}

func (i Integer) Int64() int64 {
	return i.i
}

func (i Integer) Value() (driver.Value, error) {
	if i.err != nil {
		return nil, i.err
	}
	return i.i, nil
}

func (i *Integer) Scan(value any) error {
	_ = i.Set(value)
	return nil
}

func (i Integer) Err() error {
	return i.err
}

func (i Integer) Valid() bool {
	return i.err == nil
}

func NewInteger(value int64) Integer {
	return Integer{
		s:   fmt.Sprintf("%d", value),
		i:   value,
		err: nil,
	}
}
