package model

type Date []uint8

func (d Date) String() string {
	return string(d)
}

type NullDate struct {
	Date  Date
	Valid bool
}

func (d NullDate) Val() interface{} {
	if d.Date.String() == "" {
		return nil
	}

	return d.Date.String()
}

func (d *NullDate) Scan(val interface{}) error {
	if val == nil {
		d.Valid = false
		return nil
	}

	d.Valid = true
	d.Date = val.([]uint8)

	return nil
}
