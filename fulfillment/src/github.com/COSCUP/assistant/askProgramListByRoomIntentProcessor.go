package assistant

import (
	"github.com/COSCUP/assistant/program-fetcher"
	log "github.com/Sirupsen/logrus"
	"sort"
)

type AskProgramListByRoomIntentProcessor struct {
}

func (AskProgramListByRoomIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/e90810fe-1e68-4536-9b74-a246d9615cc8"

	return "Intent Ask Program List by Room"
}

func (p AskProgramListByRoomIntentProcessor) displayMessage(input *DialogflowRequest) string {
	t := getUserTime(input.UserId())
	timeString := t.Format("現在是1月02日15點04分")
	roomName := input.RoomName()

	return timeString + "，" + roomName.String() + "的下個議程「加密/解密 雜湊看 PHP 版本的演進」在13:00開始。"
}

func (p AskProgramListByRoomIntentProcessor) speechMessage(input *DialogflowRequest) string {
	t := getUserTime(input.UserId())
	timeString := t.Format("現在是1月02日15點04分")
	roomName := input.RoomName()

	return timeString + "，" + roomName.String() + "的下個議程「加密/解密 雜湊看 PHP 版本的演進」在13:00開始。"
}

func (p AskProgramListByRoomIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}
}

func (p AskProgramListByRoomIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {
	t := getUserTime(input.UserId())
	roomName := input.RoomName()
	log.Println("user time: ", t)
	coscupPrograms, _ := fetcher.GetPrograms()

	log.Println("sessions length: ", len(coscupPrograms.Sessions))

	filited := []fetcher.Session{}
	for _, session := range coscupPrograms.Sessions {
		log.Println("comparing:", session.Room, roomName)
		if session.Room != roomName.String() {
			continue
		}

		filited = append(filited, session)
	}

	log.Println("filited sessions length: ", len(filited))
	sort.Sort(fetcher.ByStartTime(filited))

	rs := []Row{}

	for _, session := range filited {

		title := session.Zh.Title
		timeLine := session.Start.Format("15:04") + "~" + session.End.Format("15:04")

		rs = append(rs,
			getRowPayload([]Cell{
				getCellPayload(title),
				getCellPayload(timeLine),
			}, true),
		)
	}

	title := "Room " + roomName

	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(input), p.displayMessage(input)),
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
