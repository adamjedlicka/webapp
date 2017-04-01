package model

import (
	"database/sql/driver"
	"fmt"
)

type Date string

func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

func (d *Date) Scan(val interface{}) error {
	switch val.(type) {
	case string:
		*d = Date(val.(string))
	case []uint8:
		*d = Date(val.([]uint8))
	default:
		return fmt.Errorf("invalid value as not-null date")
	}

	return nil
}

func (d Date) String() string {
	return string(d)
}

type NullDate struct {
	Date  Date
	Valid bool
}

func (d NullDate) Value() (driver.Value, error) {
	if d.Valid {
		return d.Date.String(), nil
	}
	return nil, nil
}

func (d *NullDate) Scan(val interface{}) error {
	switch val.(type) {
	case nil:
		d.Valid = false
	case string:
		if val.(string) == "" {
			d.Valid = false
		} else {
			d.Valid = true
			d.Date = Date(val.(string))
		}
	case []uint8:
		d.Valid = true
		d.Date = Date(val.([]uint8))
	default:
		return fmt.Errorf("invalid value as null date")
	}

	return nil
}

func (d NullDate) String() string {
	if !d.Valid {
		return ""
	}
	return d.Date.String()
}
