package cmd

import (
	"fmt"
	"stockscrap/scrap"

	"github.com/spf13/cobra"
)

func DownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down [stockSymbols...]",
		Short: "刪除股票資訊",
		Run: func(cmd *cobra.Command, args []string) {
			// 實現 刪除資料庫全部股票的功能
			if len(args) > 0 {
				scrap.Deleter(args)
				fmt.Println("成功刪除股票功能")
			} else {
				fmt.Println("輸入想刪除的股票代號, 去做功課!")
			}
		},
	}

	return cmd
}
