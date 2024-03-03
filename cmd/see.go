package cmd

import (
	"fmt"
	"stockscrap/scrap"

	"github.com/spf13/cobra"
)

func SeeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "see [stockSymbols...]",
		Short: "查看股票資訊",
		Run: func(cmd *cobra.Command, args []string) {
			// 實現查看股票資訊
			scrap.Reader(args)
			fmt.Println("成功查看股票資訊")
		},
	}

	return cmd
}
