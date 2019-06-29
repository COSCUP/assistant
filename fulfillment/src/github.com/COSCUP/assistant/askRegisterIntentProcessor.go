package assistant

type RegisterIntentProcessor struct {
}

func (RegisterIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Intent Ask Register"
}

func (p RegisterIntentProcessor) displayMessage() string {
	return "以下是註冊資訊："
}

func (p RegisterIntentProcessor) speechMessage() string {
	return "報名期間為2019年7月11日到8月10日，地點在臺灣科技大學，免費報名。"
}

func (p RegisterIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}
}

func (p RegisterIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {
	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(), p.displayMessage()),
				getBasicCardResponsePayload(
					"COSCUP 2019 註冊資訊", "",
					// "地點： 台灣科技大學 國際大樓\n"+
					// "註冊時間： 2019 年 7 月 11  日 ~ 8 月 10 日\n"+
					// "票價： 免費",
					"報名期間為2019年7月11日到8月10日\n，地點在臺灣科技大學，免費報名。",
					"https://t.kfs.io/upload_images/98464/Screenshot_2019-05-25_COSCUP_2019_large.png",
					"報名介紹圖片", "報名", "https://coscup2019.kktix.cc/events/coscup2019regist", "CROPPED"),

				// getSimpleResponsePayload("123", "321"),
				// getTableCardResponsePayload("title", "subtitle",
				// 	[]Row{
				// 		getRowPayload([]Cell{getCellPayload("1"), getCellPayload("2"), getCellPayload("3")}, true),
				// 		getRowPayload([]Cell{getCellPayload("12"), getCellPayload("23"), getCellPayload("34")}, true),
				// 	},
				// 	[]ColunmProperty{
				// 		getColumnPropertyPayload("C1", HorizontalAlignmentCenter),
				// 		getColumnPropertyPayload("C2", HorizontalAlignmentCenter),
				// 		getColumnPropertyPayload("C3", HorizontalAlignmentCenter),
				// 	},
				// 	"https://coscup.org/2019/_nuxt/img/c2f9236.png", "image", "按鈕", "https://www.tih.tw", "CROPPED",
				// ),
			},
			"suggestions": p.getSuggsetion(),
			// "linkOutSuggestion": getLinkOutSuggestionPayload("tih", "https://www.tih.tw"),
		},
	}
}
