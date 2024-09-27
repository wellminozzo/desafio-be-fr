package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wellminozzo/desafio-be-fr/models"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := models.InitDB() // Função para inicializar o banco de dados
		if err != nil {
			log.Fatalf("migrate error: %v", err) // Melhora a captura do erro
		}

		if db == nil {
			log.Fatal("Database connection failed, db is nil")
		}

		// Executa as migrações
		err = db.AutoMigrate(&models.Dispatcher{}, &models.Offer{}, &models.Carrier{}, &models.DeliveryTime{}, &models.Weights{})
		if err != nil {
			log.Fatalf("migrate error: %v", err)
		}

		log.Println("Migration completed successfully")
	},
}
