package assistant

type HelpIntentProcessor struct {
}

func (HelpIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Intent Help"
}

func (p HelpIntentProcessor) displayMessage() string {
	return "現在是 Day 1 我可以告訴您下個議程什麼時候開始，或者是 Telegram 連結。"
}

func (p HelpIntentProcessor) speechMessage() string {
	return `<speak>現在是 <sub alias="哋萬">Day 1</sub> <break time="200ms"/>我可以告訴您下個議程什麼時候開始，或者是 貼了古拉姆 連結。 </speak>`
}

func (p HelpIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("好了謝謝"),
		getSuggestionPayload("IB101的議程什麼時候開始"),
		getSuggestionPayload("IB201的議程什麼時候開始"),
		// getSuggestionPayload("321"),
	}
}

func (p HelpIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {
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
