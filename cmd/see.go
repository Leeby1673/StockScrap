package cmd

import (
	"fmt"
	"stockscrap/scrap"

	"github.com/spf13/cobra"
)

var priceLimit int

var seecmd = &cobra.Command{
	Use:   "see [stockSymbols...]",
	Short: "查看股票資訊",
	Run: func(cmd *cobra.Command, args []string) {

		// 若輸入 -p=N flag 顯示 N 價格以下的股票
		if priceLimit == 0 {
			// 實現查看股票資訊
			scrap.Reader(args)
		} else {
			scrap.PriceReader(priceLimit)
		}
		fmt.Println("成功查看股票資訊")
	},
}
