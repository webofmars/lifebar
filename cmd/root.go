package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var totalHearts int
var emptyPercentage float64
var outputFile string

var rootCmd = &cobra.Command{
	Use:   "lifebar",
	Short: "Generates a life bar",
	Run: func(cmd *cobra.Command, args []string) {
		err := generateLifeBar(totalHearts, emptyPercentage, outputFile)
		if err != nil {
			fmt.Println("Error:", err)
		}
	},
}

func init() {
	rootCmd.Flags().IntVarP(&totalHearts, "total", "t", 10, "Total number of hearts")
	rootCmd.Flags().Float64VarP(&emptyPercentage, "empty", "e", 0.3, "Percentage of empty hearts")
	rootCmd.Flags().StringVarP(&outputFile, "output", "o", "lifebar.png", "Output file path")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	// You can add any additional configuration initialization logic here
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
