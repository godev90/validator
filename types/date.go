package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"

	"github.com/godev90/validator/errors"
)

type Date time.Time

const dateLayout = "2006-01-02"

func (d Date) String() string {
	t := time.Time(d)
	return t.Format(dateLayout)
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return errors.ErrFieldUnsupportedDataType
	}

	t, err := time.Parse(dateLayout, str)
	if err != nil {
		return err
	}

	*d = Date(t)
	return nil
}

func (d *Date) UnmarshalText(text []byte) error {
	t, err := time.Parse(dateLayout, string(text))
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d Date) Time() time.Time {
	return time.Time(d)
}

func (d Date) Value() (driver.Value, error) {
	t := time.Time(d)
	return t.Format(dateLayout), nil
}

func (d *Date) Scan(value any) error {
	switch v := value.(type) {
	case time.Time:
		*d = Date(v)
	case []byte:
		t, err := time.Parse(dateLayout, string(v))
		if err != nil {
			return err
		}
		*d = Date(t)
	case string:
		t, err := time.Parse(dateLayout, v)
		if err != nil {
			return err
		}
		*d = Date(t)
	default:
		return fmt.Errorf("unsupported type for Date: %T", value)
	}
	return nil
}
