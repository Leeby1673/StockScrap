package cmd

import (
	"fmt"
	"stockscrap/scrap"

	"github.com/spf13/cobra"
)

func DownCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "down",
		Short: "成功刪除股票資訊",
		Run: func(cmd *cobra.Command, args []string) {
			// 實現 刪除資料庫全部股票的功能
			scrap.Deleter()
			fmt.Println("成功刪除股票功能")
		},
	}

	return cmd
}
