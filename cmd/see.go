package cmd

import (
	"fmt"
	"stockscrap/scrap"

	"github.com/spf13/cobra"
)

func SeeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "see",
		Short: "查看股票資訊",
		Run: func(cmd *cobra.Command, args []string) {
			// 實現 爬取功能、存取功能、line notify 功能
			scrap.Scraper()
			fmt.Println("成功查看股票資訊")
		},
	}

	return cmd
}
