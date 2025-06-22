package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godev90/validator/errors"
)

type Datetime struct {
	t   time.Time
	s   string
	err error
}

const datetimeLayout = "2006-01-02 15:04:05"

// Set assigns a value to Datetime, accepts string or time.Time
func (d *Datetime) Set(val any) error {
	d.err = nil

	switch v := val.(type) {
	case time.Time:
		d.t = v
		d.s = v.Format(datetimeLayout)

	case string:
		d.s = v
		t, err := time.Parse(datetimeLayout, v)
		if err != nil {
			d.err = errors.ErrInvalidDatetimeFormat
			d.t = time.Time{}
			return nil
		}
		d.t = t

	default:
		d.t = time.Time{}
		d.s = ""
		d.err = errors.ErrFieldUnsupportedDataType
	}

	return nil
}

func (d Datetime) String() string {
	return d.s
}

func (d Datetime) Time() time.Time {
	return d.t
}

func (d Datetime) IsZero() bool {
	return d.t.IsZero()
}

func (d Datetime) Valid() bool {
	return d.err == nil
}

func (d Datetime) Err() error {
	return d.err
}

func (d Datetime) Validate() error {
	return d.err
}

// UnmarshalJSON parses from JSON string
func (d *Datetime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		d.err = errors.ErrInvalidDatetimeFormat
		d.t = time.Time{}
		d.s = ""
		return nil
	}
	_ = d.Set(str)
	return nil
}

// UnmarshalText parses from text (e.g. query param)
func (d *Datetime) UnmarshalText(text []byte) error {
	return d.Set(string(text))
}

// MarshalJSON serializes to JSON
func (d Datetime) MarshalJSON() ([]byte, error) {
	if !d.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(d.t.Format(datetimeLayout))
}

// Value for sql.Valuer
func (d Datetime) Value() (driver.Value, error) {
	if !d.Valid() {
		return nil, d.err
	}
	return d.t.Format(datetimeLayout), nil
}

// Scan implements sql.Scanner
func (d *Datetime) Scan(value any) error {
	d.err = nil

	switch v := value.(type) {
	case time.Time:
		d.Set(v)

	case []byte:
		_ = d.Set(string(v))

	case string:
		_ = d.Set(v)

	default:
		d.err = fmt.Errorf("unsupported type for Datetime: %T", value)
	}

	return nil
}

func NewDatetime(value time.Time) Datetime {
	return Datetime{
		t:   value,
		s:   value.Format(datetimeLayout),
		err: nil,
	}
}
