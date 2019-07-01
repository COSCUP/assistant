package assistant

type WelcomeIntentProcessor struct {
}

func (WelcomeIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Default Welcome Intent"
}

func (p WelcomeIntentProcessor) displayMessage() string {
	return "您好，我是開源人年會小幫手，我可以告訴您關於開源人年會 COSCUP 的資訊，您可以從「你會做什麼」開始："
}

func (p WelcomeIntentProcessor) speechMessage() string {
	return "您好，我是開源人年會小幫手，我可以告訴您關於開源人年會 摳史髂頗 的資訊，您可以從「你會做什麼」開始："
}

func (p WelcomeIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		getSuggestionPayload("第一天有哪些議程"),
		getSuggestionPayload("第二天有哪些議程"),
		getSuggestionPayload("我的議程列表"),

		// getSuggestionPayload("321"),
	}
}

func (p WelcomeIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {
	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(), p.displayMessage()),
				// getBasicCardResponsePayload("title", "subtitle", "formattedText",
				// 	"https://coscup.org/2019/_nuxt/img/c2f9236.png", "image", "按鈕", "https://www.tih.tw", "CROPPED"),

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
