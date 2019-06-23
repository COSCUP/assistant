package assistant

type AskProgramByProgramIntentProcessor struct {
}

func (AskProgramByProgramIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/30d89d56-6dac-4649-a940-71c99eb69324"
	return "Intent Ask Program by Program"
}

func (p AskProgramByProgramIntentProcessor) displayMessage() string {
	return "「Software Packaging for Cross OS  Distribution : Build Debian Package in  Container for Example」的議程資訊如下："
}

func (p AskProgramByProgramIntentProcessor) speechMessage() string {
	return "議程資訊如下"
}

func (p AskProgramByProgramIntentProcessor) getSuggsetion() []map[string]interface{} {
	ret := []map[string]interface{}{
		getSuggestionPayload("🌟我有興趣"),
		getSuggestionPayload("IB503在哪"),
		getSuggestionPayload("IB503下一場議程什麼時候開始"),
		getSuggestionPayload("好了謝謝"),
		// getSuggestionPayload("你會做什麼"),
		// getSuggestionPayload("321"),
	}

	return ret
}

func (p AskProgramByProgramIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {

	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.speechMessage(), p.displayMessage()),
				getBasicCardResponsePayload(
					"FLOSS! not only Linux and hackers!!",
					"IB503 11:20 ~ 11:50",
					"因為工作需要，時常要打包 Debian package。"+
						"但開發人員平時熟悉、使用的作業系統，不一定是 Debian 系列的 Linux distribution。"+
						"單純為了打包 Debian package，而另外安裝一個 Debian 系列的作業系統，似乎是不合成本的作法。 "+
						"有賴於 Container 技術的興起與流行，它提供了一個乾淨、簡潔又輕便的方案。"+
						"而本次將分享開發人員在 Arch Linux 上用 docker 工具，"+
						"從 Docker Hub 取得相對輕巧的 Debian image 為基底，並將它執行成為一個 Container。且在裡頭準備相對應的「工具」、「權限」與「環境」的經驗，最後實際打包出一個 Debian package 當作範例。",
					"https://coscup.org/2019/_nuxt/img/c2f9236.png", "講者照片",
					"議程網頁", "https://coscup.org/2019/programs/63d5742d-0c03-4849-bfb2-24cbb68f66a1", "CROPPED"),

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
