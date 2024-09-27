package models

import (
	"database/sql"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:project@tcp(database:3306)/project?charset=utf8mb4&parseTime=True&loc=Local"
	// jdbc:mariadb://localhost:9091/project
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//gormDB.Model(&Dispatcher{}).Where("active = ?", true).Update("id", "1")
	//gormDB.Model(&Carrier{}).Where("active = ?", true).Update("id", "1")

	err = gormDB.Migrator().DropTable(&Dispatcher{}, &Carrier{}, &Offer{}, &Weights{}, &DeliveryTime{})
	if err != nil {
		log.Fatal("Erro ao remover tabelas:", err)
	}

	err = gormDB.AutoMigrate(&Dispatcher{}, &Carrier{}, &Offer{}, &Weights{}, &DeliveryTime{})
	if err != nil {
		log.Fatal("Erro ao criar tabelas:", err)
	}

	return gormDB, nil
}
