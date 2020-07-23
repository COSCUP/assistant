package assistant

import (
	"fmt"
	"github.com/COSCUP/assistant/program-fetcher"
	log "github.com/Sirupsen/logrus"
)

type QueryFavoriteListIntentProcessor struct {
}

func (QueryFavoriteListIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Intent Query Favorite"
}

func (p QueryFavoriteListIntentProcessor) displayMessage(l int) string {
	return "您目前的訂閱列表如下，一共有" + fmt.Sprintf("%d", l) + "項："
}

func (p QueryFavoriteListIntentProcessor) speechMessage() string {
	return "您目前的訂閱列表如下："
}

func (p QueryFavoriteListIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		getSuggestionPayload("移除第一項議程"),
		getSuggestionPayload("11點55分之後有哪些議程"),

		// getSuggestionPayload("321"),
	}
}

func (p QueryFavoriteListIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {
	// favoriteList :=
	userStorage := NewUserStorageFromDialogflowRequest(input)
	favList := userStorage.getFavoriteList()

	ret := map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(), p.displayMessage(len(favList))),
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

	log.Println("favlist length:", len(favList))

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
		return p.PayloadWithOneFavorite(input, favList)
	} else {
		// emptyState
		return p.PayloadWithNoFavoriteList(input)

	}

	return ret
}

func (p QueryFavoriteListIntentProcessor) displayMessageWithOneFavorite() string {
	return "您有一項訂閱的議程："
}

func (p QueryFavoriteListIntentProcessor) speechMessageWithOneFavoite() string {
	return "您有一項訂閱的議程："
}

func (p QueryFavoriteListIntentProcessor) PayloadWithOneFavorite(input *DialogflowRequest, favList []interface{}) map[string]interface{} {
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

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessageWithOneFavoite(), p.displayMessageWithOneFavorite()),

				getBasicCardResponsePayload(
					title,
					subTitle,
					desc,
					sessionPhotoUrl, "講者照片",
					"議程網頁", "https://coscup.org/2020/zh-TW/agenda/"+sessionInfo.ID, "CROPPED"),
			},
			"suggestions": p.getSuggsetion(),
		},
		"outputContexts": map[string]interface{}{
			"selected_session": map[string]interface{}{
				"id": sessId,
			},
		},
	}
	return ret
}

func (p QueryFavoriteListIntentProcessor) displayMessageWithNoFavoriteList() string {
	return "您目前沒有訂閱任何議程，可以透過查詢某個時段的議程進行加入，您可以問我「IB101接下來有哪些議程」。"
}

func (p QueryFavoriteListIntentProcessor) speechMessageWithNoFavoriteList() string {
	return "您目前沒有訂閱任何議程，可以透過查詢某個時段的議程進行加入，您可以問我「IB101接下來有哪些議程」。"
}

func (p QueryFavoriteListIntentProcessor) getSuggsetionWithNoFavoriteList() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("IB101接下來有哪些議程"),
		// getSuggestionPayload("321"),
	}
}
func (p QueryFavoriteListIntentProcessor) PayloadWithNoFavoriteList(input *DialogflowRequest) map[string]interface{} {

	ret := map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessageWithNoFavoriteList(), p.displayMessageWithNoFavoriteList()),
			},
			"suggestions": p.getSuggsetionWithNoFavoriteList(),
		},
	}
	return ret
}
