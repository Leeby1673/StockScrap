package cmd

import (
	"fmt"
	"stockscrap/scrap"

	"github.com/spf13/cobra"
)

func CatchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "catch",
		Short: "抓取股票資訊",
		Run: func(cmd *cobra.Command, args []string) {
			// 實現 爬取功能、存取功能、line notify 功能
			scrap.Scraper()
			fmt.Println("成功查看股票資訊")
		},
	}

	return cmd
}
