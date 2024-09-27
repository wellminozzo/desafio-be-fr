package models

import (
	"strconv"

	"gorm.io/gorm"
)

type Carrier struct {
	gorm.Model

	//ID               string `gorm:"type:varchar(255)" json:"id"`                // Chave primária
	CompanyName      string `gorm:"type:varchar(255)" json:"company_name"`      // Nome da empresa
	Logo             string `gorm:"type:varchar(255)" json:"logo"`              // URL do logotipo
	Name             string `gorm:"type:varchar(255)" json:"name"`              // Nome da transportadora
	Reference        int    `gorm:"type:int" json:"reference"`                  // Referência única
	RegisteredNumber string `gorm:"type:varchar(14)" json:"registered_number"`  // CNPJ
	StateInscription string `gorm:"type:varchar(255)" json:"state_inscription"` // Inscrição estadual
	//Offers           []Offer `json:"offers"`
}

func (c *Carrier) TableName() string {
	return "carriers"
}

func (c *Carrier) Save() error {

	db, err := InitDB()
	if err != nil {
		return err
	}

	if c.CompanyName == strconv.Itoa(0) {
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
