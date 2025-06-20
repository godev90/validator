package types

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"

	"github.com/godev90/validator/errors"
)

type Integer string

func (i Integer) String() string {
	return string(i)
}

func (i *Integer) Set(val int64) {
	*i = Integer(strconv.FormatInt(val, 10))
}

func (i *Integer) UnmarshalJSON(data []byte) error {
	var intVal int64
	if err := json.Unmarshal(data, &intVal); err == nil {
		*i = Integer(strconv.FormatInt(intVal, 10))
		return nil
	}

	var strVal string
	if err := json.Unmarshal(data, &strVal); err == nil {
		*i = Integer(strVal)
		return nil
	}

	return errors.ErrFieldUnsupportedDataType
}

func (i *Integer) UnmarshalText(text []byte) error {
	s := string(text)
	if _, err := strconv.ParseInt(s, 10, 64); err != nil {
		return errors.ErrFieldUnsupportedDataType
	}
	*i = Integer(s)
	return nil
}

func (i Integer) MarshalJSON() ([]byte, error) {
	val, err := strconv.ParseInt(string(i), 10, 64)
	if err != nil {
		return nil, errors.ErrFieldUnsupportedDataType
	}
	return json.Marshal(val)
}

func (i Integer) Integer() int64 {
	val, _ := strconv.ParseInt(string(i), 10, 64)
	return val
}

func (i Integer) Value() (driver.Value, error) {
	return strconv.ParseInt(string(i), 10, 64)
}

func (i *Integer) Scan(value any) error {
	switch v := value.(type) {
	case int64:
		*i = Integer(strconv.FormatInt(v, 10))
	case []byte:
		*i = Integer(string(v))
	case string:
		*i = Integer(v)
	default:
		return errors.ErrFieldUnsupportedDataType
	}
	return nil
}
