package golmy

import (
	"fmt"

	"github.com/spf13/cobra"
)

var flag string

// 建立 golmy 根命令
func Golmy() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "golmy",
		Short: "股票出的去, 錢進的來, 投資發大財",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("golmy 啟動!")
		},
	}

	rootCmd.PersistentFlags().StringVarP(&flag, "flag", "f", "", "測試側")
	rootCmd.AddCommand(SeeCmd())
	rootCmd.AddCommand(DownCmd())

	return rootCmd
}

// func Execute() {
// 	if err := rootCmd.Execute(); err != nil {
// 		fmt.Println("注定當韭菜, 你的韭菜編號:", err)
// 		os.Exit(1)
// 	}
// }
