package main

import (
	"fmt"
	"os"
	"stockscrap/cmd"
)

func main() {
	rootCmd := cmd.Golmy()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("注定當韭菜, 你的韭菜編號:", err)
		os.Exit(1)
	}
}
