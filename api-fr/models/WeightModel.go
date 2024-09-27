package models

type Weights struct {
	Real  int64    `json:"real"`
	Used  int64    `json:"used"`
	Cubed *float64 `json:"cubed,omitempty"`
}

func (w *Weights) TableName() string {
	return "weights"
}
