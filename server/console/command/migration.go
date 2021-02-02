package command

import (
	"afkser/database"
	"afkser/initialize"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var migrationCmd = &cobra.Command{
	Use:   "migration",
	Short: "初始化数据库",
	Long:  "运行初始化数据库，生成表并生成迁移数据",
	Run: func(cmd *cobra.Command, args []string) {
		initialize.Init()
		database.Migration()
	},
}

func init() {
	rootCmd.AddCommand(migrationCmd)
}
