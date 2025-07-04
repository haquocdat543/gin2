package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "app",
	Long:  "app",
	Run: func(
		cmd *cobra.Command,
		args []string,
	) {
		fmt.Printf("App\n")
	},
}

func ExucuteCLI() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
