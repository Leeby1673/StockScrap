package linenotify

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// 判斷式 價格低於設定閥值, 才通知

func Linenotify(symbol string, pricechangepct float64) {
	// Line Motify Token
	accessToken := "WEiPdLsMq4TPfiPIMgDy6h5mty52ujTNISgNxJOD3Tg"

	// 要發送的訊息內容
	// 代入股票代號、顯示漲跌百分比
	message := fmt.Sprintf("%s 都跌 %.2f%% 了, 還不All in?", symbol, pricechangepct)
	message = url.QueryEscape(message)
	// 建立 POST 的請求內容
	requestBody := bytes.NewBufferString("message=" + message)

	// 建立 HTTP POST 請求
	request, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", requestBody)
	if err != nil {
		log.Fatal(err)
	}

	// 設置請求 Header(標頭), 包括 Token
	request.Header.Set("Authorization", "Bearer "+accessToken)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 建立 HTTP 客戶端
	client := &http.Client{}

	// 發送請求，獲得回覆
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("發送請求失敗: ", err)
	}

	defer response.Body.Close()

	// 檢查回應狀態碼
	if response.StatusCode != http.StatusOK {
		log.Fatal("Line Notify 訊息傳送失敗, 狀態碼: ", response.StatusCode)
	}

	// 顯示成功訊息
	println("成功了! 穩賺的吧?")
}
