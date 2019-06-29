package assistant

type LocationByLocationNameIntentProcessor struct {
}

func (LocationByLocationNameIntentProcessor) Name() string {
	// return "projects/coscup/agent/intents/7d85fb8f-3e3c-4776-a604-f9ca0f6627e6"
	return "Inetnt Ask Location by Location Name"
}

func (p LocationByLocationNameIntentProcessor) displayMessage(input *DialogflowRequest) string {

	if string(input.RoomName()) == "" {
		return "我知道您想問位置，但是不確定你要問哪邊的位置"
	}
	return string(input.RoomName()) + "的位置如下："

}

func (p LocationByLocationNameIntentProcessor) getSuggsetionItemFromRoomName(roomName string) map[string]interface{} {
	return getSuggestionPayload(roomName + "的議程什麼時候開始")
}

func (p LocationByLocationNameIntentProcessor) getSuggsetion(input *DialogflowRequest) []map[string]interface{} {

	ret := []map[string]interface{}{
		getSuggestionPayload("好了謝謝"),
		// getSuggestionPayload("321"),

	}

	if string(input.RoomName()) == "" {
		ret = append(ret, getSuggestionPayload("IB101在哪"))
		ret = append(ret, getSuggestionPayload("IB201在哪"))
		ret = append(ret, getSuggestionPayload("IB301在哪"))
		ret = append(ret, getSuggestionPayload("IB302在哪"))
		ret = append(ret, getSuggestionPayload("IB304在哪"))
		ret = append(ret, getSuggestionPayload("IB305在哪"))
		ret = append(ret, getSuggestionPayload("IB306在哪"))
		ret = append(ret, getSuggestionPayload("IB401在哪"))
		ret = append(ret, getSuggestionPayload("IB408在哪"))
		ret = append(ret, getSuggestionPayload("IB501在哪"))
		ret = append(ret, getSuggestionPayload("IB502在哪"))
		ret = append(ret, getSuggestionPayload("IB503在哪"))
		ret = append(ret, getSuggestionPayload("E2102在哪"))
	} else {

		ret = append(ret, getSuggestionPayload(string(input.RoomName())+"的下個議程什麼時候開始"))
	}

	return ret
}

func (p LocationByLocationNameIntentProcessor) getMapURl(room RoomNameType) string {
	switch room {
	case RoomNameTypeIB101:
		return "https://api2019.coscup.org/assets/1f.jpg"
	case RoomNameTypeIB201:
		return "https://api2019.coscup.org/assets/2f.jpg"
	case RoomNameTypeIB301:
		return "https://api2019.coscup.org/assets/3f.jpg"
	case RoomNameTypeIB302:
		return "https://api2019.coscup.org/assets/3f.jpg"
	case RoomNameTypeIB304:
		return "https://api2019.coscup.org/assets/3f.jpg"
	case RoomNameTypeIB305:
		return "https://api2019.coscup.org/assets/3f.jpg"
	case RoomNameTypeIB306:
		return "https://api2019.coscup.org/assets/3f.jpg"
	case RoomNameTypeIB401:
		return "https://api2019.coscup.org/assets/4f.jpg"
	case RoomNameTypeIB408:
		return "https://api2019.coscup.org/assets/4f.jpg"
	case RoomNameTypeIB501:
		return "https://api2019.coscup.org/assets/5f.jpg"
	case RoomNameTypeIB502:
		return "https://api2019.coscup.org/assets/5f.jpg"
	case RoomNameTypeIB503:
		return "https://api2019.coscup.org/assets/5f.jpg"
	case RoomNameTypeIE2102:
		return "https://api2019.coscup.org/assets/1f.jpg"
	}
	return "https://api2019.coscup.org/assets/1f.jpg"
}

func (p LocationByLocationNameIntentProcessor) getMapDesc(room RoomNameType) string {
	// 會需要用到這個敘述的人，可能是不容易看清楚地圖的人
	switch room {
	case RoomNameTypeIB101:
		return "IB101的位置在國際會議廳一樓"
	case RoomNameTypeIB201:
		return "IB201的位置在國際會議廳二樓，電梯出口左轉"
	case RoomNameTypeIB301:
		return "IB301的位置在國際會議廳三樓，電梯出口左轉"
	case RoomNameTypeIB302:
		return "IB302的位置在國際會議廳三樓，電梯出口左轉"
	case RoomNameTypeIB304:
		return "IB304的位置在國際會議廳三樓，電梯出口靠右直走第一間"
	case RoomNameTypeIB305:
		return "IB305的位置在國際會議廳三樓，電梯出口靠右直走第二間"
	case RoomNameTypeIB306:
		return "IB306的位置在國際會議廳三樓，電梯出口靠右直走第三間"
	case RoomNameTypeIB401:
		return "IB401的位置在國際會議廳四樓，電梯出口左轉"
	case RoomNameTypeIB408:
		return "IB408的位置在國際會議廳四樓"
	case RoomNameTypeIB501:
		return "IB501的位置在國際會議廳五樓，電梯出口左轉第三間"
	case RoomNameTypeIB502:
		return "IB502的位置在國際會議廳五樓，電梯出口左轉第二間"
	case RoomNameTypeIB503:
		return "IB503的位置在國際會議廳五樓，電梯出口左轉第一間"
	case RoomNameTypeIE2102:
		return "E2102的位置在工學院二館，國際會議廳對面"
	}
	return "我知道您想問位置，但是不確定你要問哪邊的位置"
}
func (p LocationByLocationNameIntentProcessor) Payload(input *DialogflowRequest) map[string]interface{} {

	room := RoomNameType(input.RoomName())
	url := p.getMapURl(room)
	return map[string]interface{}{
		"expectUserResponse": true,

		// "systemIntent": getListSystemIntentPayload(),
		"richResponse": map[string]interface{}{
			"items": []map[string]interface{}{
				getSimpleResponsePayload(p.getMapDesc(room), p.displayMessage(input)),
				getBasicCardResponsePayload("", "", "",
					url, "image", "MAP", url, string(ImageDisplayOptionsDefault)),
				getSimpleResponsePayload("還有什麼可以喂您服務的嗎？", "還有什麼可以為您服務的嗎？"),

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
			"suggestions": p.getSuggsetion(input),
			// "linkOutSuggestion": getLinkOutSuggestionPayload("tih", "https://www.tih.tw"),
		},
	}
}
