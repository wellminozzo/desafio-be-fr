package models

type DeliveryTime struct {
	Days          *int64  `json:"days,omitempty"`
	EstimatedDate *string `json:"estimated_date,omitempty"`
	Hours         *int64  `json:"hours,omitempty"`
	Minutes       *int64  `json:"minutes,omitempty"`
}

func (d *DeliveryTime) TableName() string {
	return "delivery_times"
}

func (d *DeliveryTime) Save() error {
	return nil
}
