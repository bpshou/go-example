package cmd

import "github.com/spf13/cobra"

func Execute() {
	c := &cobra.Command{}

	c.AddCommand(GormCmd)
	c.AddCommand(SqlCmd)
	// 执行
	c.Execute()
}
