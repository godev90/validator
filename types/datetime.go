package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godev90/validator/errors"
)

type Datetime time.Time

const datetimeLayout = "2006-01-02 15:04:05"

func (dt Datetime) String() string {
	t := time.Time(dt)
	return t.Format(datetimeLayout)
}

func (dt *Datetime) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return errors.ErrFieldUnsupportedDataType
	}

	t, err := time.Parse(datetimeLayout, str)
	if err != nil {
		return err
	}

	*dt = Datetime(t)
	return nil
}

func (dt *Datetime) UnmarshalText(text []byte) error {
	t, err := time.Parse(datetimeLayout, string(text))
	if err != nil {
		return err
	}
	*dt = Datetime(t)
	return nil
}

func (d Datetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (dt Datetime) Time() time.Time {
	return time.Time(dt)
}

func (dt Datetime) Value() (driver.Value, error) {
	t := time.Time(dt)
	return t.Format(datetimeLayout), nil
}

func (dt *Datetime) Scan(value any) error {
	switch v := value.(type) {
	case time.Time:
		*dt = Datetime(v)
	case []byte:
		t, err := time.Parse(datetimeLayout, string(v))
		if err != nil {
			return err
		}
		*dt = Datetime(t)
	case string:
		t, err := time.Parse(datetimeLayout, v)
		if err != nil {
			return err
		}
		*dt = Datetime(t)
	default:
		return fmt.Errorf("unsupported type for Datetime: %T", value)
	}
	return nil
}
