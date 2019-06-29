package assistant

import (
	"github.com/COSCUP/assistant/program-fetcher"

	log "github.com/Sirupsen/logrus"
	"sort"
	"time"
)

type AskProgramListByTimeIntentProcessor struct {
}

func (AskProgramListByTimeIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/daae629f-90f2-48d8-9827-52ea12d20bf7"
	return "Intent Ask Program List by Time"
}

func (p AskProgramListByTimeIntentProcessor) displayMessage(t *time.Time) string {
	timeStr := t.Format("15點04分")

	return timeStr + "之後的議程資訊如下："
}

func (p AskProgramListByTimeIntentProcessor) speechMessage(t *time.Time) string {
	timeStr := t.Format("15點04分")
	return timeStr + "之後的議程資訊如下："
}

func (p AskProgramListByTimeIntentProcessor) getSuggsetion() []map[string]interface{} {
	return []map[string]interface{}{
		getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}
}

func (p AskProgramListByTimeIntentProcessor) getListSystemIntentPayload(listTitle string, sessions []fetcher.Session) map[string]interface{} {
	// list item must be 2 ~ 30
	retList := []ListItem{}
	for _, sessionInfo := range sessions {
		title := sessionInfo.Zh.Title
		desc := sessionInfo.Zh.Description
		timeLine := getSessionTimeLineWithDay(&sessionInfo)
		subTitle := sessionInfo.Room + " " + timeLine
		sessionPhotoUrl := sessionInfo.SpeakerPhotoUrl()

		item := getListItemPayload(title,
			sessionInfo.ID,
			subTitle+"\n"+desc, []string{title}, getImagePayload(sessionPhotoUrl, "講者照片"))

		retList = append(retList, item)
	}

	return getListSystemIntentPayload(
		listTitle,
		// "議程導覽",
		retList,
	)
}

func (p AskProgramListByTimeIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {

	t := input.Time()
	dayType := input.DayType()

	// if t == nil {
	// 	tt := time.Now()
	// 	t = &tt
	// }

	if dayType == "第一天" && t == nil {
		ymdString := getDay1StartTime().Format("2006-01-02")
		timeString := " 00:00:00+0800"

		tt, _ := time.Parse("2006-01-02 15:04:05Z0700", ymdString+timeString)
		t = &tt
	} else if dayType == "第一天" && t != nil {
		// 第一天下午三點之後
		ymdString := getDay1StartTime().Format("2006-01-02")
		timeString := t.Format(" 15:04:05Z0700")

		tt, _ := time.Parse("2006-01-02 15:04:05Z0700", ymdString+timeString)
		t = &tt
	} else if dayType == "第二天" && t == nil {
		// 第二天有什麼議程
		ymdString := getDay2StartTime().Format("2006-01-02")
		timeString := " 00:00:00+0800"

		tt, _ := time.Parse("2006-01-02 15:04:05Z0700", ymdString+timeString)
		t = &tt
	} else if dayType == "第二天" && t != nil {
		// 第二天下午三點之後
		ymdString := getDay2StartTime().Format("2006-01-02")
		log.Println("ymdString:", ymdString, getDay2StartTime())
		timeString := t.Format(" 15:04:05Z0700")

		tt, _ := time.Parse("2006-01-02 15:04:05Z0700", ymdString+timeString)
		t = &tt
	}

	coscupPrograms, _ := fetcher.GetPrograms()

	log.Println("filter start time:", t)
	log.Println("sessions length: ", len(coscupPrograms.Sessions))
	filtered := []fetcher.Session{}
	for _, s := range coscupPrograms.Sessions {
		if s.Start.Before(*t) {
			continue
		}
		filtered = append(filtered, s)
	}

	sort.Sort(fetcher.ByStartTime(filtered))

	dayTypeString := "Day 1"
	if IsDayTwo(*t) {
		dayTypeString = "Day 2"
	}

	listTitle := "COSCUP 2019 " + dayTypeString + " " + t.Format("15:04")

	return map[string]interface{}{
		"expectUserResponse": true,
		"systemIntent":       p.getListSystemIntentPayload(listTitle, filtered[:15]),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(t), p.displayMessage(t)),
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
