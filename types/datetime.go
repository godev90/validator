package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"

	"github.com/godev90/validator/errors"
)

type Datetime struct {
	t   time.Time
	s   string
	err error
}

func (d *Datetime) Set(val any) error {
	d.err = nil

	switch v := val.(type) {
	case time.Time:
		d.t = v
		d.s = v.Format(datetimeLayout)

	case string:
		layouts := []string{
			datetimeLayout,
			datetimeISOZLayout,
			time.RFC3339,
			dateLayout,
		}

		for _, layout := range layouts {
			var t time.Time
			var err error

			if strings.HasSuffix(layout, "Z") || layout == time.RFC3339 {
				t, err = time.Parse(layout, v)
			} else {
				t, err = time.ParseInLocation(layout, v, time.Local)
			}

			if err == nil {
				d.t = t
				d.s = t.Format(datetimeLayout)
				return nil
			}

		}

		d.err = errors.ErrInvalidDatetimeFormat
		d.t = time.Time{}
		d.s = ""

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
		d.err = d.Set(string(v))

	case sql.RawBytes:
		d.err = d.Set(string([]byte(v)))

	case string:
		d.err = d.Set(v)

	default:
		d.err = errors.ErrUnsuppotedDatetime
	}

	return d.err
}

func NewDatetime(value time.Time) Datetime {
	return Datetime{
		t:   value,
		s:   value.Format(datetimeLayout),
		err: nil,
	}
}
