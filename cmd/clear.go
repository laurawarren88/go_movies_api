package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear the log file",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := loadConfig()
		if err != nil {
			fmt.Println("Failed to load config:", err)
			return
		}

		if err := os.Truncate(config.LogFilePath, 0); err != nil {
			fmt.Println("Failed to clear log:", err)
		} else {
			fmt.Println(Red + "Log file cleared." + Reset)
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
