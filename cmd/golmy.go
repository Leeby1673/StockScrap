package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	catchcmd.Flags().BoolVarP(&ongoing, "ongoing", "o", false, "啟動持續監測，不會把資料存進資料庫")
	catchcmd.Flags().Float64VarP(&lineNotifyPercent, "ineNotifyPercent", "l", 0, "設定多少漲跌幅，觸發 Line Notify")
	seecmd.Flags().IntVarP(&priceLimit, "printLimit", "p", 0, "查看股票價格設定 N 以下")
	RootCmd.AddCommand(catchcmd)
	RootCmd.AddCommand(seecmd)
	RootCmd.AddCommand(downcmd)
}

// 建立 golmy 根命令

var RootCmd = &cobra.Command{
	Use:   "golmy",
	Short: "股票出的去, 錢進的來, 投資發大財",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("golmy 啟動!")
	},
}
