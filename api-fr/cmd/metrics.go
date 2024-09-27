package cmd

import (
	"github.com/spf13/cobra"
)

var MetricsCmd = &cobra.Command{
	Use:   "get-metrics",
	Short: "Exibe as métricas de cotações",
	Long: `Este comando retorna várias métricas relacionadas às cotações, 
incluindo a contagem de cotações por empresa, a cotação mais barata e a mais cara.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Inicialize a conexão com o banco de dados
		// db, err := models.InitDB()
		// if err != nil {
		// 	fmt.Println("Erro ao conectar ao banco de dados:", err)
		// 	return
		// }

		// // Chama a função GetMetrics
		// metrics, err := GetMetrics(db)
		// if err != nil {
		// 	fmt.Println("Erro ao obter as métricas:", err)
		// 	return
		// // }

		// // Exibe as métricas
		// fmt.Printf("Métricas: %+v\n", metrics)
	},
}
