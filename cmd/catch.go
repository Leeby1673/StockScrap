package cmd

import (
	"fmt"
	"stockscrap/scrap"

	"github.com/spf13/cobra"
)

var catchcmd = &cobra.Command{
	Use:   "catch [stockSymbols...]",
	Short: "抓取股票資訊 catch <args>, args 替換成股票代碼, 例: AVGO",
	Run: func(cmd *cobra.Command, args []string) {
		// 實現 爬取功能、存取功能、line notify 功能
		// 可輸入複數參數, 獲取複數的股票
		if len(args) > 0 {
			scrap.Scraper(args)
			fmt.Println("成功抓取股票資訊")
		} else {
			fmt.Println("輸入想找的股票代碼, 連這都不會還想發財?")
		}

	},
}
