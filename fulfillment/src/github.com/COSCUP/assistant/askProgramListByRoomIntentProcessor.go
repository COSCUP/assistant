package assistant

import (
// log "github.com/Sirupsen/logrus"
)

type AskProgramListByRoomIntentProcessor struct {
}

func (AskProgramListByRoomIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/e90810fe-1e68-4536-9b74-a246d9615cc8"

	return "Intent Ask Program List by Room"
}

func (p AskProgramListByRoomIntentProcessor) displayMessage() string {
	return "現在是上午11點，IB101的下個議程「加密/解密 雜湊看 PHP 版本的演進」在13:00開始。"
}

func (p AskProgramListByRoomIntentProcessor) speechMessage() string {
	return "現在是上午11點，IB101的下個議程「加密/解密 雜湊看 PHP 版本的演進」在13:00開始。"
}

func (p AskProgramListByRoomIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}
}

func (p AskProgramListByRoomIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {
	rs := []Row{
		getRowPayload([]Cell{
			getCellPayload("加密/解密 雜湊看 PHP 版本的演進"),
			getCellPayload("13:00 ~ 13:50"),
		}, true),
		getRowPayload([]Cell{
			getCellPayload("从 GRANK 到 GITRANK ， ..."),
			getCellPayload("15:00 ~ 15:50"),
		}, true),
	}

	roomName := input.RoomName()
	title := "Room " + roomName

	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(), p.displayMessage()),
				// getBasicCardResponsePayload("title", "subtitle", "formattedText",
				// 	"https://coscup.org/2019/_nuxt/img/c2f9236.png", "image", "按鈕", "https://www.tih.tw", "CROPPED"),

				// getSimpleResponsePayload("123", "321"),
				getTableCardResponsePayload(string(title), "簡易議程表",
					rs,
					[]ColunmProperty{
						getColumnPropertyPayload("名稱", HorizontalAlignmentLeading),
						getColumnPropertyPayload("時間", HorizontalAlignmentTrailing),
					},
					"https://coscup.org/2019/_nuxt/img/c2f9236.png", "COSCUP LOGO", "議程網頁", "https://coscup.org/2019/programs/", "CROPPED",
				),
			},
			"suggestions": p.getSuggsetion(),
			// "linkOutSuggestion": getLinkOutSuggestionPayload("tih", "https://www.tih.tw"),
		},
	}
}
