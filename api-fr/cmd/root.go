package cmd

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "Main api frete rapido",
	Long:  `This project has a api test for frete rapido`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'main -h' or 'go run main.go -h' to view all command options")
	},
}

func init() {

	rootCmd.AddCommand(MigrateCmd)
	rootCmd.AddCommand(MetricsCmd)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithFields(logrus.Fields{
			"error": fmt.Sprintf("%+v\n", err),
		}).Error("cmd error on run command")
		os.Exit(1)
	}
}
