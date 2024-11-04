package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View the log file for previous searched data - just input number of lines",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		n, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid number:", err)
			return
		}

		config, err := loadConfig()
		if err != nil {
			fmt.Println("Failed to load config:", err)
			return
		}

		viewLastNEntries(config.LogFilePath, n)
	},
}

func init() {
	rootCmd.AddCommand(viewCmd)
}

func viewLastNEntries(filePath string, n int) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open log file at %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if n > len(lines) {
		n = len(lines)
	}
	fmt.Printf("Last %d entries from log:\n", n)
	fmt.Println(strings.Repeat("=", 30))
	for _, line := range lines[len(lines)-n:] {
		fmt.Println(line)
	}
}
