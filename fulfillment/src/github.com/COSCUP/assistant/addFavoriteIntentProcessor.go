package assistant

import (
	"fmt"
	"github.com/COSCUP/assistant/program-fetcher"
	log "github.com/Sirupsen/logrus"
)

type AddFavoriteIntentProcessor struct {
}

func (AddFavoriteIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Intent Add Favorite"
}

func (p AddFavoriteIntentProcessor) displayMessage(sessionTitle string) string {
	return "已幫您加入「" + sessionTitle + "」，目前列表如下："
}

func (p AddFavoriteIntentProcessor) speechMessage(sessionTitle string) string {
	return "已幫您把議程加入訂閱列表，目前列表如下："
}

func (p AddFavoriteIntentProcessor) getSuggsetion(favList []interface{}) []map[string]interface{} {

	ret := []map[string]interface{}{}
	prog, _ := fetcher.GetPrograms()
	for _, it := range favList {
		sessionInfo := prog.GetSessionByID(it.(string))
		dt := "第一天"
		if IsDayTwo(sessionInfo.Start) {
			dt = "第二天"
		}
		text := dt + sessionInfo.End.Format("15點4分之後還有什麼議程")
		ret = append(ret, getSuggestionPayload(text))

	}
	ret = append(ret, getSuggestionPayload("你會做什麼"))
	ret = append(ret, getSuggestionPayload("好了謝謝"))

	return ret
}

func (p AddFavoriteIntentProcessor) getListSystemIntentPayload() map[string]interface{} {
	// list item must be 2 ~ 30

	return getListSystemIntentPayload(
		"COSCUP 2019 Day 1  11:15",
		// "議程導覽",
		[]ListItem{
			getListItemPayload("XXX", "KKK", "DDD", []string{"X1", "X2"}, getImagePayload("https://coscup.org/2019/_nuxt/img/c2f9236.png", "dd")),
			getListItemPayload("XXX1", "KKK2", "DDD3", []string{"X12", "X21"}, getImagePayload("https://coscup.org/2019/_nuxt/img/c2f9236.png", "dd")),
		},
	)
}

func (p AddFavoriteIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {

	userStorage := NewUserStorageFromDialogflowRequest(input)

	contextSelectedSession := input.Context("selected_session")
	log.Println("selected session:", contextSelectedSession, contextSelectedSession["id"])
	selectedID := contextSelectedSession["id"].(string)

	prog, _ := fetcher.GetPrograms()
	sessionInfo := prog.GetSessionByID(selectedID)
	title := sessionInfo.Zh.Title
	userStorage.addFavorite(selectedID)
	favList := userStorage.getFavoriteList()
	log.Println("favlist length:", len(favList))

	ret := map[string]interface{}{
		"expectUserResponse": true,
		"userStorage":        userStorage.EncodeToString(),

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(title), p.displayMessage(title)),
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
			"suggestions": p.getSuggsetion(favList),
			// "linkOutSuggestion": getLinkOutSuggestionPayload("tih", "https://www.tih.tw"),
		},
		"outputContexts": map[string]interface{}{
			"pervious_session_list": map[string]interface{}{
				"list": favList,
			},
		},
	}

	if len(favList) >= 2 {
		ll := []ListItem{}

		prog, _ := fetcher.GetPrograms()
		for i, id := range favList {

			sessionInfo := prog.GetSessionByID(id.(string))
			title := fmt.Sprintf("%d. ", i+1) + sessionInfo.Zh.Title
			desc := sessionInfo.Zh.Description
			dt := "D1"
			if IsDayTwo(sessionInfo.Start) {
				dt = "D2"
			}
			timeLine := dt + " " + sessionInfo.Start.Format("15:04") + "~" + sessionInfo.End.Format("15:04")
			subTitle := sessionInfo.Room + " " + timeLine
			sessionPhotoUrl := sessionInfo.SpeakerPhotoUrl()

			item := getListItemPayload(title, id.(string),
				subTitle+"\n"+desc, []string{title},
				getImagePayload(sessionPhotoUrl, "講者照片"))
			ll = append(ll, item)

		}

		ret["systemIntent"] = getListSystemIntentPayload("興趣列表", ll)

	} else if len(favList) == 1 {
		// card
		return p.PayloadWithOneFavorite(input, favList, userStorage)
	}
	return ret
}

func (p AddFavoriteIntentProcessor) PayloadWithOneFavorite(input *DialogflowRequest, favList []interface{}, userStorage *UserStorage) map[string]interface{} {
	sessId := favList[0].(string)
	prog, _ := fetcher.GetPrograms()
	sessionInfo := prog.GetSessionByID(sessId)
	title := sessionInfo.Zh.Title
	desc := sessionInfo.Zh.Description
	timeLine := sessionInfo.Start.Format("15:04") + "~" + sessionInfo.End.Format("15:04")
	subTitle := sessionInfo.Room + " " + timeLine
	sessionPhotoUrl := sessionInfo.SpeakerPhotoUrl()

	ret := map[string]interface{}{
		"expectUserResponse": true,
		"userStorage":        userStorage.EncodeToString(),

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.displayMessage(title), p.displayMessage(title)),

				getBasicCardResponsePayload(
					title,
					subTitle,
					desc,
					sessionPhotoUrl, "講者照片",
					"議程網頁", "https://coscup.org/2019/programs/"+sessionInfo.ID, "CROPPED"),
			},
			"suggestions": p.getSuggsetion(favList),
		},

		"outputContexts": map[string]interface{}{
			"selected_session": map[string]interface{}{
				"id": sessId,
			},
		},
	}
	return ret
}
