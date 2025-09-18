package cli

import (
	"gin/pkg/config"
	"gin/pkg/db"
	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed",
	Long:  "Seed",
}

var seedStartCmd = &cobra.Command{
	Use:   "start",
	Short: "start seeding",
	Long:  "This command start seeding",
	Run: func(
		cmd *cobra.Command,
		args []string,
	) {
		db.Seed(config.InitDB())
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
	seedCmd.AddCommand(seedStartCmd)
}
