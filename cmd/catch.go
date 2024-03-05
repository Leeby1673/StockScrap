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
		if len(args) <= 0 {
			// 沒有給參數
			fmt.Println("輸入想找的股票代碼, 連這都不會還想發財?")
		} else if len(args) > 0 && !ongoing {
			// 有參數、一次性抓取
			scrap.Scraper(args, lineNotifyPercent)
			fmt.Println("成功抓取股票資訊")
		} else if len(args) > 0 && ongoing {
			// 有參數、持續性監測
			scrap.OngoingScraper(args, lineNotifyPercent)

			ticker := time.NewTicker(15 * time.Second)
			defer ticker.Stop()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

			// fmt.Println("監測開始")
			for {
				select {
				case <-ticker.C:
					scrap.OngoingScraper(args, lineNotifyPercent)
				case <-quit:
					fmt.Println("結束監測")
					return
				}
			}
		}

	},
}
