package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godev90/validator/errors"
)

type Date struct {
	t   time.Time
	s   string
	err error
}

const dateLayout = "2006-01-02"

// Set assigns a value to Date, accepts time.Time or string
func (d *Date) Set(val any) error {
	d.err = nil

	switch v := val.(type) {
	case time.Time:
		d.t = v
		d.s = v.Format(dateLayout)

	case string:
		d.s = v
		t, err := time.Parse(dateLayout, v)
		if err != nil {
			d.err = errors.ErrInvalidDateFormat
			d.t = time.Time{}
			return nil
		}
		d.t = t

	default:
		d.err = errors.ErrFieldUnsupportedDataType
		d.t = time.Time{}
		d.s = ""
	}

	return nil
}

func (d Date) String() string {
	return d.s
}

func (d Date) Time() time.Time {
	return d.t
}

func (d Date) IsZero() bool {
	return d.t.IsZero()
}

func (d Date) Valid() bool {
	return d.err == nil
}

func (d Date) Err() error {
	return d.err
}

func (d Date) Validate() error {
	return d.err
}

// UnmarshalJSON parses date from JSON string
func (d *Date) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		d.t = time.Time{}
		d.s = ""
		d.err = errors.ErrInvalidDateFormat
		return nil
	}

	_ = d.Set(str)
	return nil
}

// UnmarshalText parses date from text (e.g., query param)
func (d *Date) UnmarshalText(text []byte) error {
	return d.Set(string(text))
}

// MarshalJSON serializes the date to JSON
func (d Date) MarshalJSON() ([]byte, error) {
	if !d.Valid() {
		return json.Marshal(nil)
	}
	return json.Marshal(d.t.Format(dateLayout))
}

// Value for sql.Valuer
func (d Date) Value() (driver.Value, error) {
	if !d.Valid() {
		return nil, d.err
	}
	return d.t.Format(dateLayout), nil
}

// Scan implements sql.Scanner
func (d *Date) Scan(value any) error {
	d.err = nil

	switch v := value.(type) {
	case time.Time:
		d.Set(v)

	case []byte:
		_ = d.Set(string(v))

	case string:
		_ = d.Set(v)

	default:
		d.err = fmt.Errorf("unsupported type for Date: %T", value)
	}

	return nil
}

func NewDate(value time.Time) Date {
	return Date{
		t:   value,
		s:   value.Format(dateLayout),
		err: nil,
	}
}
