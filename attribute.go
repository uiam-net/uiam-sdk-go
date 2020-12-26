package uiamsdk

import (
	"database/sql/driver"
	"encoding/json"
)

type Attribute map[string]interface{}

func (attr Attribute) Value() (driver.Value, error) {
	if attr == nil {
		return nil, nil
	}

	bt, err := json.Marshal(attr)
	if err != nil {
		return nil, err
	}

	return string(bt), nil
}

func (attr *Attribute) Scan(v interface{}) error {
	str := string(v.([]byte))
	return json.Unmarshal([]byte(str), &attr)
}
