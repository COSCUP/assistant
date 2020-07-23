package assistant

import (
	"github.com/COSCUP/assistant/program-fetcher"
	log "github.com/Sirupsen/logrus"
)

type DefaultFallbackIntent struct {
}

func (DefaultFallbackIntent) Name() string {
	// return "projects/coscup/agent/intents/30d89d56-6dac-4649-a940-71c99eb69324"
	return "Default Fallback Intent"
}

func (p DefaultFallbackIntent) displayMessage(sessionTitle string) string {
	return "「" + sessionTitle + "」的議程資訊如下："
}

func (p DefaultFallbackIntent) speechMessage(sessionTitle string) string {
	return "議程資訊如下"
}

func (p DefaultFallbackIntent) getSuggsetion() []map[string]interface{} {
	ret := []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		getSuggestionPayload("離開"),
		getSuggestionPayload("IB503下一場議程什麼時候開始"),
		// getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}

	return ret
}

func (p DefaultFallbackIntent) GetFallbackMessage(input *DialogflowRequest) string {
	return ""
}

func (p DefaultFallbackIntent) Payload(input *DialogflowRequest) map[string]interface{} {
	sessionId := input.GetSessionIdFromOptionResult()
	log.Println("session id:", sessionId)

	if sessionId != "" {
		return p.PayloadWithQueryProgram(input)
	}
	type response struct {
		suggestion []map[string]interface{}
		text       string
		ssml       string
	}

	responseList := []response{
		response{
			suggestion: p.getSuggsetion(),
			text:       "這邊是開源人年會，您可以換種方式說說看嗎？或者是幫忙在 GitHub 上開個 ISSUE 如何？",
			ssml:       `<speak>這邊是開源人年會，<break/>您可以換種方式說說看嗎？或者是幫忙在 基ハ 上開個 伊シュー 如 何？</speak>`,
		},
		response{
			suggestion: p.getSuggsetion(),
			text:       "這邊是開源人年會，我不是很確定您想要我做什麼，或者是您可以問我「你會做什麼」",
			ssml:       `<speak>這邊是開源人年會，<break/>我不是很確定您想要我做什麼，或者是您可以問我「你會做什麼」</speak>`,
		},
		response{
			suggestion: p.getSuggsetion(),
			text:       "我是會幫忙管理議程資訊的開源人年會，您可以說「離開」來離開對話。",
			ssml:       `<speak>我是會幫忙管理議程資訊的開源人年會，<break/>您可以說「離開」來離開對話。</speak>`,
		},
	}

	fallbackCount := input.Context("fallback_count")
	if fallbackCount == nil {
		fallbackCount = map[string]interface{}{}
	}

	counter, ok := fallbackCount["counter"].(float64)
	if !ok {
		counter = 0
	}

	resp := responseList[int(counter)%len(responseList)]
	fallbackCount["counter"] = counter + 1

	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(resp.ssml, resp.text),

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
			"suggestions":       resp.suggestion,
			"linkOutSuggestion": getLinkOutSuggestionPayload("GitHub 連結", "https://github.com/tihtw/coscup_action/issues"),
		},
		"outputContexts": map[string]interface{}{
			"fallback_count": fallbackCount,
		},
	}
}

func (p DefaultFallbackIntent) PayloadWithQueryProgram(input *DialogflowRequest) map[string]interface{} {
	sessionId := input.GetSessionIdFromOptionResult()
	log.Println("PayloadWithQueryProgram session id:", sessionId)
	prog, _ := fetcher.GetPrograms()
	sessionInfo := prog.GetSessionByID(sessionId)
	title := sessionInfo.Zh.Title
	desc := sessionInfo.Zh.Description
	timeLine := sessionInfo.Start.Format("15:04") + "~" + sessionInfo.End.Format("15:04")
	subTitle := sessionInfo.Room + " " + timeLine

	sessionPhotoUrl := sessionInfo.SpeakerPhotoUrl()

	inFavoriteList := NewUserStorageFromDialogflowRequest(input).isSessionIdInFavorite(sessionId)

	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(title), p.displayMessage(title)),
				getBasicCardResponsePayload(
					title,
					subTitle,
					desc,
					sessionPhotoUrl, "講者照片",
					"議程網頁", "https://coscup.org/2020/zh-TW/agenda/"+sessionInfo.ID, "CROPPED"),

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
			"suggestions": getSuggestionWithSession(inFavoriteList, sessionInfo),
		},
		"outputContexts": map[string]interface{}{
			"selected_session": map[string]interface{}{
				"id": sessionId,
			},

			"ask_program": map[string]interface{}{},
		},
	}
}
