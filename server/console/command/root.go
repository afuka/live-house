package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "console",
	Short: "这是一款终端命令行工具",
	Long:  "欢迎使用afkcommand终端工具",
}

// Execute 执行
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
