package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCatch(t *testing.T) {

	// 設置緩衝區
	actual := new(bytes.Buffer)

	// 設定 root 的標準輸出、標準錯誤輸出重新定向到緩衝區
	RootCmd.SetOut(actual)
	RootCmd.SetErr(actual)

	// 設置命令的參數
	RootCmd.SetArgs([]string{"catch", "AAPL", "-l=3"})

	// 執行命令
	RootCmd.Execute()

	// 我們期待的輸出
	expected := "成功抓取股票資訊"

	// 用 assert 套件來驗證結果是否一致，若沒有則顯示錯誤訊息
	assert.Equal(t, expected, actual.String(), "期待值 != 實際值")
}
