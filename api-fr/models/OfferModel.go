package models

import "gorm.io/gorm"

type Offer struct {
	gorm.Model

	DispatcherID string `gorm:"index" json:"dispatcher_id"` // Chave estrangeira para Dispatcher
	//CarrierID    string `gorm:"column:carrier_id;not null" json:"carrier_id"` // Chave estrangeira para Carrier
	//Carrier              Carrier      `gorm:"foreignkey:CarrierID;references:ID"`           // Relacionamento com Carrier
	CostPrice            float64      `gorm:"type:decimal(10,2)" json:"cost_price"`     // Preço de custo
	DeliveryTime         DeliveryTime `gorm:"embedded" json:"delivery_time"`            // Estrutura de tempo de entrega embutida
	Expiration           string       `gorm:"type:varchar(255)" json:"expiration"`      // Data de expiração
	FinalPrice           float64      `gorm:"type:decimal(10,2)" json:"final_price"`    // Preço final
	HomeDelivery         bool         `gorm:"type:boolean" json:"home_delivery"`        // Entrega domiciliar
	Modal                string       `gorm:"type:varchar(50)" json:"modal"`            // Modalidade
	Offer                int          `gorm:"type:int" json:"offer"`                    // Posição da oferta
	OriginalDeliveryTime DeliveryTime `gorm:"embedded" json:"original_delivery_time"`   // Tempo de entrega original
	Service              string       `gorm:"type:varchar(255)" json:"service"`         // Nome do serviço
	SimulationType       int          `gorm:"type:int" json:"simulation_type"`          // Tipo de simulação
	TableReference       string       `gorm:"type:varchar(255)" json:"table_reference"` // Referência da tabela
	Weights              Weights      `gorm:"embedded" json:"weights"`                  // Estrutura de peso
	CompanyName          []Carrier    `gorm:"many2many:carrier_company_name" json:"company_name"`
}

func (c *Offer) TableName() string {
	return "offers"
}

func (c *Offer) Save() error {
	db, err := InitDB()
	if err != nil {
		return err
	}

	if c.ID == 0 {
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

func FindLastQuoteDesc(records *[]*Offer, limit int) error {

	db, err := InitDB()
	if err != nil {
		return err
	}

	return db.
		// Select([]string{"ult", "date"}).
		Order("createdAt desc").
		Limit(limit).
		Find(records).Error

}
