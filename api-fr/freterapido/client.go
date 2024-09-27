package freterapido

type APIRequest struct {
	Shipper struct {
		RegisteredNumber string `json:"registered_number"`
		Token            string `json:"token"`
		PlatformCode     string `json:"platform_code"`
	} `json:"shipper"`
	Recipient struct {
		Type             int    `json:"type"`
		RegisteredNumber string `json:"registered_number"`
		StateInscription string `json:"state_inscription"`
		Country          string `json:"country"`
		Zipcode          int    `json:"zipcode"`
	} `json:"recipient"`
	Dispatchers []struct {
		RegisteredNumber string  `json:"registered_number"`
		Zipcode          int     `json:"zipcode"`
		TotalPrice       float64 `json:"total_price"`
		Volumes          []struct {
			Amount        int     `json:"amount"`
			AmountVolumes int     `json:"amount_volumes"`
			Category      string  `json:"category"`
			SKU           string  `json:"sku"`
			Tag           string  `json:"tag"`
			Description   string  `json:"description"`
			Height        float64 `json:"height"`
			Width         float64 `json:"width"`
			Length        float64 `json:"length"`
			UnitaryPrice  float64 `json:"unitary_price"`
			UnitaryWeight float64 `json:"unitary_weight"`
			Consolidate   bool    `json:"consolidate"`
			Overlaid      bool    `json:"overlaid"`
			Rotate        bool    `json:"rotate"`
		} `json:"volumes"`
	} `json:"dispatchers"`
	Channel        string `json:"channel"`
	Filter         int    `json:"filter"`
	Limit          int    `json:"limit"`
	Identification string `json:"identification"`
	Reverse        bool   `json:"reverse"`
	SimulationType []int  `json:"simulation_type"`
	Returns        struct {
		Composition  bool `json:"composition"`
		Volumes      bool `json:"volumes"`
		AppliedRules bool `json:"applied_rules"`
	} `json:"returns"`
}
