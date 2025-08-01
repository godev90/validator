package typedef

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"

	"github.com/godev90/validator/faults"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Date struct {
	t   time.Time
	s   string
	err error
}

const (
	dateLayout         = "2006-01-02"
	datetimeLayout     = "2006-01-02 15:04:05"
	datetimeISOZLayout = "2006-01-02T15:04:05Z"
)

func (d *Date) Set(val any) error {
	d.err = nil

	switch v := val.(type) {
	case time.Time:
		d.t = v
		d.s = v.Format(dateLayout)
		return nil

	case string:

		layouts := []string{
			dateLayout,
			datetimeISOZLayout,
			time.RFC3339,
			datetimeLayout,
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
				d.s = t.Format(dateLayout)
				return nil
			}
		}

		d.err = faults.ErrInvalidDateFormat
		d.t = time.Time{}
		d.s = ""

	default:
		d.err = faults.ErrInvalidDateFormat
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
		d.err = faults.ErrInvalidDateFormat
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
		d.err = d.Set(string(v))

	case sql.RawBytes:
		d.err = d.Set(string([]byte(v)))

	case string:
		d.err = d.Set(v)

	default:
		d.err = faults.ErrInvalidDateFormat
	}

	return d.err
}

func (d Date) ToProto() *timestamppb.Timestamp {
	if !d.Valid() {
		return nil
	}
	return timestamppb.New(d.t)
}

func NewDate(value time.Time) Date {
	return Date{
		t:   value,
		s:   value.Format(dateLayout),
		err: nil,
	}
}

func DateToday() Date {
	var now time.Time = time.Now()
	return NewDate(now)
}

func DateFromProto(ts *timestamppb.Timestamp) *Date {
	if ts == nil {
		return nil
	}
	date := NewDate(ts.AsTime())
	return &date
}
