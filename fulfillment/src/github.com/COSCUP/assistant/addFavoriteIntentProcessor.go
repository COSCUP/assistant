package assistant

import (
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

func (p AddFavoriteIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}
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

	return map[string]interface{}{
		"expectUserResponse": true,
		"userStorage":        userStorage.EncodeToString(),

		"systemIntent": p.getListSystemIntentPayload(),
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
	}
}
