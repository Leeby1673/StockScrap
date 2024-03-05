package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"stockscrap/scrap"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var ongoing bool
var lineNotifyPercent int

var catchcmd = &cobra.Command{
	Use:   "catch [stockSymbols...]",
	Short: "抓取股票資訊 catch <args>, args 替換成股票代碼, 例: AVGO",
	Run: func(cmd *cobra.Command, args []string) {
		// 需除錯 catch -o 沒有參數的情況
		// 優先序列 給予參數 > 要哪個模式 以及觸發 line通知
		// 持續性監測
		if ongoing {
			scrap.OngoingScraper(args)

			ticker := time.NewTicker(15 * time.Second)
			defer ticker.Stop()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			// fmt.Println("監測開始")
			for {
				select {
				case <-ticker.C:
					scrap.OngoingScraper(args)
				case <-quit:
					fmt.Println("結束監測")
					return
				}
			}
		}

		// 一次性抓取
		if len(args) > 0 {
			scrap.Scraper(args)
			fmt.Println("成功抓取股票資訊")
		} else {
			fmt.Println("輸入想找的股票代碼, 連這都不會還想發財?")
		}

	},
}
