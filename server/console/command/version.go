package command

import (
	"github.com/gookit/color"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本信息",
	Long:  "当前版本描述说明",
	Run: func(cmd *cobra.Command, args []string) {
		color.Green.Println("v0.0.3")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
