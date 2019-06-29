package assistant

// https://assistant.google.com/services/a/uid/0000007475a0f5b8?hl=zh_TW

import (
// "github.com/COSCUP/assistant/program-fetcher"
// log "github.com/Sirupsen/logrus"
// "math/rand"
)

type QuitIntentProcessor struct {
}

func (QuitIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Intent Quit"
}

func (p QuitIntentProcessor) displayMessage(input *DialogflowRequest) string {
	if IsInActivity(getUserTime(input)) {
		return "再見，希望您滿載而歸。"
	} else {
		return "再見，祝您有美好的一天。"
	}

}

func (p QuitIntentProcessor) speechMessage(input *DialogflowRequest) string {
	if IsInActivity(getUserTime(input)) {
		return `再見，希望您滿載而歸。`

	} else {
		return "再見，祝您有美好的一天。"
	}
}

func (p QuitIntentProcessor) getSuggsetionItemFromRoomName(roomName string) map[string]interface{} {
	return getSuggestionPayload(roomName + "的議程什麼時候開始")
}

func (p QuitIntentProcessor) shouldPromptStoreLink(input *DialogflowRequest) bool {
	return false
}

func (p QuitIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {

	items := []map[string]interface{}{
		getSimpleResponsePayload(p.speechMessage(input), p.displayMessage(input)),
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
	}

	if p.shouldPromptStoreLink(input) {
		items = append(items,
			getBasicCardResponsePayload("請為我們評分", "", "如果有任何意見或是需要改進的地方，歡迎利用商店連結給我們反饋。",
				"https://coscup.org/2019/_nuxt/img/c2f9236.png", "image", "商店連結", "https://assistant.google.com/services/a/uid/0000007475a0f5b8?hl=zh_TW", "CROPPED"))
	}

	return map[string]interface{}{
		"expectUserResponse": false,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": items,
			// "linkOutSuggestion": getLinkOutSuggestionPayload("tih", "https://www.tih.tw"),
		},
	}
}
