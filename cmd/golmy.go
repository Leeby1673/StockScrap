package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// 建立 golmy 根命令
func Golmy() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "golmy",
		Short: "股票出的去, 錢進的來, 投資發大財",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("golmy 啟動!")
		},
	}

	seecmd.Flags().IntVarP(&priceLimit, "printLimit", "p", 0, "查看股票價格設定 N 以下")
	rootCmd.AddCommand(catchcmd)
	rootCmd.AddCommand(seecmd)
	rootCmd.AddCommand(downcmd)

	return rootCmd
}
