package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var GormCmd = &cobra.Command{
	Use:   "gorm",
	Short: "Gorm 生成",
	Long:  `Gorm 生成数据`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Gorm gen success !!!")
	},
}
