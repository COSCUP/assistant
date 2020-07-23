package assistant

import (
	"fmt"
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

func (p AskProgramListByRoomIntentProcessor) displayMessage(input *DialogflowRequest, session *fetcher.Session) string {
	t := getUserTime(input)
	timeString := t.Format("現在是1月02日15點04分")
	roomName := input.RoomName()
	if session == nil {
		return timeString + "，" + roomName.String() + "找不到下一個議程。"
	}

	timeAbbr := session.Start.Format("15:04")

	return timeString + "，" + roomName.String() + "的下個議程「" + session.Zh.Title + "」將在" + timeAbbr + "開始。"
}

func (p AskProgramListByRoomIntentProcessor) speechMessage(input *DialogflowRequest, session *fetcher.Session) string {
	t := getUserTime(input)
	timeString := t.Format("現在是1月02日15點04分")
	roomName := input.RoomName()
	if session == nil {
		return timeString + "，" + roomName.String() + "找不到下一個議程。"
	}

	timeAbbr := session.Start.Format("15:04")

	return timeString + "，" + roomName.String() + "的下個議程「" + session.Zh.Title + "」將在" + timeAbbr + "開始。"
}

func (p AskProgramListByRoomIntentProcessor) getSuggsetion(input *DialogflowRequest, sessionLength int) []map[string]interface{} {
	ret := []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}

	if sessionLength > 3 {
		ret = append(ret, getSuggestionPayload("告訴我第二場那場的資訊"))
	}

	if input.RoomName() != "IB201" {
		ret = append(ret, getSuggestionPayload("IB101的議程資訊"))
	} else {
		ret = append(ret, getSuggestionPayload("IB201的議程資訊"))
	}

	return ret
}

func (p AskProgramListByRoomIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {
	userStorage := NewUserStorageFromDialogflowRequest(input)
	// userConversationToken := NewConversationTokenFromDialogflowRequest(input)
	log.Println("user storage: ", userStorage)

	t := getUserTime(input)
	roomName := input.RoomName()
	log.Println("user time: ", t)
	coscupPrograms, _ := fetcher.GetPrograms()

	log.Println("sessions length: ", len(coscupPrograms.Sessions))

	filtered := []fetcher.Session{}
	for _, session := range coscupPrograms.Sessions {
		// log.Println("comparing:", session.Room, roomName)
		if session.Room != roomName.String() {
			continue
		}

		filtered = append(filtered, session)
	}

	log.Println("filtered sessions length: ", len(filtered))
	sort.Sort(fetcher.ByStartTime(filtered))

	rs := []Row{}

	for i, session := range filtered {

		title := fmt.Sprintf("%d. %s", i+1, session.Zh.Title)
		timeLine := session.Start.Format("15:04") + "~" + session.End.Format("15:04")

		rs = append(rs,
			getRowPayload([]Cell{
				getCellPayload(title),
				getCellPayload(timeLine),
			}, true),
		)
	}

	title := "Room " + roomName
	var firstSession *fetcher.Session
	if len(filtered) > 0 {
		firstSession = &filtered[0]
	}

	sessionIdList := []string{}

	for _, session := range filtered {
		sessionIdList = append(sessionIdList, session.ID)
	}

	// userConversationToken.AddPreviousDisplaySessionList(filtered)
	return map[string]interface{}{
		"expectUserResponse": true,
		"userStorage":        userStorage.EncodeToString(),
		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(input, firstSession), p.displayMessage(input, firstSession)),
				// getBasicCardResponsePayload("title", "subtitle", "formattedText",
				// 	"https://coscup.org/2019/_nuxt/img/c2f9236.png", "image", "按鈕", "https://www.tih.tw", "CROPPED"),

				// getSimpleResponsePayload("123", "321"),
				getTableCardResponsePayload(string(title), "簡易議程表",
					rs,
					[]ColunmProperty{
						getColumnPropertyPayload("名稱", HorizontalAlignmentLeading),
						getColumnPropertyPayload("開始時間", HorizontalAlignmentTrailing),
					},
					"https://coscup.org/2019/_nuxt/img/c2f9236.png", "COSCUP LOGO", "議程網頁", "https://coscup.org/2020/zh-TW/agenda", "CROPPED",
				),
			},
			"suggestions": p.getSuggsetion(input, len(filtered)),
			// "linkOutSuggestion": getLinkOutSuggestionPayload("tih", "https://www.tih.tw"),
		},

		"outputContexts": map[string]interface{}{
			"pervious_session_list": map[string]interface{}{
				"list": sessionIdList,
			},
		},
	}
}
