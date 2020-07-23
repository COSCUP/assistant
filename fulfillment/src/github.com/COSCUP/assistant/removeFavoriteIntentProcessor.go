package assistant

import (
	"fmt"
	"github.com/COSCUP/assistant/program-fetcher"
	log "github.com/Sirupsen/logrus"
)

type RemoveFavoriteIntentProcessor struct {
}

func (RemoveFavoriteIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Intent Remove Favorite"
}

func (p RemoveFavoriteIntentProcessor) displayMessage(sessionTitle string) string {
	return "已幫您移除「" + sessionTitle + "」，目前列表如下："
}

func (p RemoveFavoriteIntentProcessor) speechMessage(sessionTitle string) string {
	return "已幫您把議程移出訂閱列表，目前列表如下："
}

func (p RemoveFavoriteIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}
}

func (p RemoveFavoriteIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {

	userStorage := NewUserStorageFromDialogflowRequest(input)

	selectedID := ""

	perviousDisplayedSessionListInfo := input.Context("pervious_session_list")
	log.Println("perviousDisplayedSessionList:", perviousDisplayedSessionListInfo)
	number := input.SelectedNumber()
	if number != 0 && perviousDisplayedSessionListInfo["list"] != nil {
		// remove from list
		list := perviousDisplayedSessionListInfo["list"].([]interface{})
		if number >= 1 && len(list) > number-1 {
			//
			selectedID = list[number-1].(string)
		}
	} else {
		// remove from single program

		contextSelectedSession := input.Context("selected_session")
		log.Println("selected session:", contextSelectedSession, contextSelectedSession["id"])
		selectedID = contextSelectedSession["id"].(string)
	}

	prog, _ := fetcher.GetPrograms()
	sessionInfo := prog.GetSessionByID(selectedID)
	title := sessionInfo.Zh.Title
	userStorage.removeFavorite(selectedID)
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
			"suggestions": p.getSuggsetion(),
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

		for i, id := range favList {

			prog, _ := fetcher.GetPrograms()
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

			item := getListItemPayload(title, id.(string), subTitle+"\n"+desc, []string{title}, getImagePayload(sessionPhotoUrl, "講者照片"))
			ll = append(ll, item)

		}

		ret["systemIntent"] = getListSystemIntentPayload("興趣列表", ll)

	} else if len(favList) == 1 {
		// card
		return p.PayloadWithOneFavorite(input, favList, userStorage, title)
	} else {
		// no favorite
		return p.PayloadWithNoFavoriteList(input, userStorage, title)
	}
	return ret
}

func (p RemoveFavoriteIntentProcessor) PayloadWithOneFavorite(input *DialogflowRequest, favList []interface{}, userStorage *UserStorage, removedSessionTitle string) map[string]interface{} {
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
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(removedSessionTitle), p.displayMessage(removedSessionTitle)),

				getBasicCardResponsePayload(
					title,
					subTitle,
					desc,
					sessionPhotoUrl, "講者照片",
					"議程網頁", "https://coscup.org/2020/zh-TW/agenda/"+sessionInfo.ID, "CROPPED"),
			},
			"suggestions": p.getSuggsetion(),
		},
	}
	return ret
}

func (p RemoveFavoriteIntentProcessor) displayMessageWithNoFavoriteList(title string) string {
	return "已幫您移除「" + title + "」，您目前沒有訂閱議程，可以透過查詢某個時段的議程進行加入，您可以問我「IB101接下來有哪些議程」。"
}

func (p RemoveFavoriteIntentProcessor) speechMessageWithNoFavoriteList(title string) string {
	return "已幫您將議程移出訂閱列表，您目前沒有任何訂閱議程，可以透過查詢某個時段的議程進行加入，您可以問我「IB101接下來有哪些議程」。"
}

func (p RemoveFavoriteIntentProcessor) getSuggsetionWithNoFavoriteList() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("IB101接下來有哪些議程"),
		// getSuggestionPayload("321"),
	}
}
func (p RemoveFavoriteIntentProcessor) PayloadWithNoFavoriteList(input *DialogflowRequest, userStorage *UserStorage, title string) map[string]interface{} {

	ret := map[string]interface{}{
		"expectUserResponse": true,
		"userStorage":        userStorage.EncodeToString(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessageWithNoFavoriteList(title), p.displayMessageWithNoFavoriteList(title)),
			},
			"suggestions": p.getSuggsetionWithNoFavoriteList(),
		},
	}
	return ret
}
