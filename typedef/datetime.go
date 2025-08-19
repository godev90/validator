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

var localTime *time.Location

func SetTimezone(timezone string) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		loc, _ = time.LoadLocation("Local")
	}
	localTime = loc
}

type Datetime struct {
	t   time.Time
	s   string
	err error
}

func (d *Datetime) Set(val any) error {
	d.err = nil

	switch v := val.(type) {
	case time.Time:
		d.t = v.In(localTime)
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
				t, err = time.ParseInLocation(layout, v, localTime)
			}

			if err == nil {
				d.t = t.In(localTime)
				d.s = t.Format(datetimeLayout)
				return nil
			}

		}

		d.err = faults.ErrInvalidDatetimeFormat
		d.t = time.Time{}
		d.s = ""

	default:
		d.t = time.Time{}
		d.s = ""
		d.err = faults.ErrInvalidDatetimeFormat
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
		d.err = faults.ErrInvalidDatetimeFormat
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
		d.err = faults.ErrInvalidDatetimeFormat
	}

	return d.err
}

func (d Datetime) ToProto() *timestamppb.Timestamp {
	if !d.Valid() {
		return nil
	}
	return timestamppb.New(d.t)
}

func NewDatetime(value time.Time) Datetime {
	value = value.In(localTime)
	return Datetime{
		t:   value,
		s:   value.Format(datetimeLayout),
		err: nil,
	}
}

func DatetimeNow() Datetime {
	var now time.Time = time.Now().In(localTime)
	return NewDatetime(now)
}

func DatetimeFromProto(ts *timestamppb.Timestamp) *Datetime {
	if ts == nil {
		return nil
	}
	date := NewDatetime(ts.AsTime().In(localTime))
	return &date
}
