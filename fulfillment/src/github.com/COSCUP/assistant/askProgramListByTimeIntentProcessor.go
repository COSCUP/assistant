package assistant

import (
	"github.com/COSCUP/assistant/program-fetcher"

	log "github.com/Sirupsen/logrus"
)

type AskProgramListByTimeIntentProcessor struct {
}

func (AskProgramListByTimeIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/daae629f-90f2-48d8-9827-52ea12d20bf7"
	return "Intent Ask Program List by Time"
}

func (p AskProgramListByTimeIntentProcessor) displayMessage() string {
	return "現在是上午11點15分，接下來的議程資訊如下："
}

func (p AskProgramListByTimeIntentProcessor) speechMessage() string {
	return "現在是上午11點15分，接下來的議程資訊如下："
}

func (p AskProgramListByTimeIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}
}

func (p AskProgramListByTimeIntentProcessor) getListSystemIntentPayload() map[string]interface{} {
	// list item must be 2 ~ 30
	coscupPrograms, _ := fetcher.GetPrograms()

	log.Println("sessions length: ", len(coscupPrograms.Sessions))

	filitedSession :=

		coscupPrograms.Sessions[:5]

	retList := []ListItem{}

	for _, session := range filitedSession {
		retList = append(retList,
			getListItemPayload(
				session.Zh.Title,
				session.ID,
				"XXX\n"+session.Zh.Description,
				[]string{}, getImagePayload("https://coscup.org/2019/_nuxt/img/c2f9236.png", "dd")))
	}

	return getListSystemIntentPayload(
		"COSCUP 2019 Day 1  11:15",
		// "議程導覽",
		retList,
	)
}

func (p AskProgramListByTimeIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {

	return map[string]interface{}{
		"expectUserResponse": true,
		"systemIntent":       p.getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(), p.displayMessage()),
				// getCarouselBrowsePayload([]CarouselBrowseItem{
				// 	getCarouselBrowseItemPayload("title", "desc", "footer", "https://www.tih.tw", getImagePayload("https://coscup.org/2019/_nuxt/img/c2f9236.png", "dd")),
				// 	getCarouselBrowseItemPayload("title", "desc", "footer", "https://www.tih.tw", getImagePayload("https://coscup.org/2019/_nuxt/img/c2f9236.png", "dd")),
				// 	getCarouselBrowseItemPayload("title", "desc", "footer", "https://www.tih.tw", getImagePayload("https://coscup.org/2019/_nuxt/img/c2f9236.png", "dd")),
				// }),
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
