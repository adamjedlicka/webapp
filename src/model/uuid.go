package model

import "database/sql/driver"
import "strings"
import "fmt"

const SizeOfUUID = 36

type UUID []byte

func (uuid *UUID) Scan(val interface{}) error {
	switch val.(type) {
	case string:
		v := strings.TrimSpace(val.(string))
		if len(v) != SizeOfUUID {
			return fmt.Errorf("wrong size of UUID")
		}
		*uuid = UUID(v)
	case []uint8:
		v := strings.TrimSpace(string(val.([]uint8)))
		if len(v) != SizeOfUUID {
			return fmt.Errorf("wrong size of UUID")
		}
		*uuid = UUID(v)
	default:
		return fmt.Errorf("wrong value for UUID")
	}

	return nil
}

func (uuid UUID) Value() (driver.Value, error) {
	return uuid.String(), nil
}

func (uuid UUID) String() string {
	return string(uuid)
}

type NullUUID struct {
	UUID  UUID
	Valid bool
}

func (uuid *NullUUID) Scan(val interface{}) error {
	switch val.(type) {
	case nil:
		uuid.Valid = false
	default:
		err := uuid.UUID.Scan(val)
		if err != nil {
			return err
		}
		uuid.Valid = true
	}

	return nil
}

func (uuid NullUUID) Value() (driver.Value, error) {
	if !uuid.Valid {
		return nil, nil
	}

	return uuid.UUID.String(), nil
}

func (uuid NullUUID) String() string {
	if !uuid.Valid {
		return ""
	}

	return uuid.UUID.String()
}
