package models

import "strconv"

type Dispatcher struct {
	ID                         string  `json:"id"`
	RequestID                  string  `json:"request_id"`
	RegisteredNumberShipper    string  `json:"registered_number_shipper"`
	RegisteredNumberDispatcher string  `json:"registered_number_dispatcher"`
	ZipcodeOrigin              int64   `json:"zipcode_origin"`
	Offers                     []Offer `json:"offers"`
}

func (c *Dispatcher) TableName() string {
	return "dispatchers"
}

func (c *Dispatcher) Save() error {
	db, err := InitDB()
	if err != nil {
		return err
	}

	if c.ID == strconv.Itoa(0) {
		// create ....
		err = db.Create(&c).Error
		if err != nil {
			return err
		}
	} else {
		// update ...
		err = db.Save(&c).Error
		if err != nil {
			return err
		}
	}

	return nil
}
