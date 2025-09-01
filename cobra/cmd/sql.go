package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var SqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "生成SQL模型",
	Long:  "自动生成SQL模型",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Sql gen success !!!")
	},
}
