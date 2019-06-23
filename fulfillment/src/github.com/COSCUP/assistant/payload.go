package assistant

type ImageDisplayOptions string

const (
	ImageDisplayOptionsDefault ImageDisplayOptions = "DEFAULT"
	ImageDisplayOptionsWhite   ImageDisplayOptions = "WHITE"
	ImageDisplayOptionsCropped ImageDisplayOptions = "CROPPED"
)

type HorizontalAlignment string

const (
	HorizontalAlignmentCenter   HorizontalAlignment = "CENTER"
	HorizontalAlignmentLeading  HorizontalAlignment = "LEADING"
	HorizontalAlignmentTrailing HorizontalAlignment = "TRAILING"
)

type Cell map[string]interface{}
type Row map[string]interface{}
type ColunmProperty map[string]interface{}

func getListItemPayload(key string) map[string]interface{} {
	return map[string]interface{}{
		"title": key + "接下來的議程是 ",
		"optionInfo": map[string]interface{}{
			"key":      key + "接下來的議程是KEY",
			"synonyms": []string{"s1" + key},
		},
		"description": "d1",
		"image": map[string]interface{}{
			"url":               "https://coscup.org/2019/_nuxt/img/c2f9236.png",
			"accessibilityText": "text",
		},
	}
}

func getListSystemIntentPayload() map[string]interface{} {
	return map[string]interface{}{
		"intent": "actions.intent.OPTION",
		"data": map[string]interface{}{
			"@type": "type.googleapis.com/google.actions.v2.OptionValueSpec",
			"listSelect": map[string]interface{}{
				"title": "list title",
				"items": []map[string]interface{}{
					getListItemPayload("IB101"),
					getListItemPayload("IB102"),
					getListItemPayload("IB103"),
				},
			},
		},
	}
}

func getLinkOutSuggestionPayload(name, url string) map[string]interface{} {
	return map[string]interface{}{
		"destinationName": name,
		"url":             url,
	}
}

func getSuggestionPayload(title string) map[string]interface{} {
	return map[string]interface{}{
		"title": title,
	}
}

func getSimpleResponsePayload(textToSpeech string, displayText string) map[string]interface{} {
	return map[string]interface{}{
		"simpleResponse": map[string]interface{}{
			"textToSpeech": textToSpeech,
			"displayText":  displayText,
		},
	}
}

func getBasicCardResponsePayload(title, subtitle, formattedText,
	imageUrl, imageAccessibilityText,
	buttonTitle, buttonOpenUrl, imageDisplayOptions string) map[string]interface{} {
	return map[string]interface{}{
		"basicCard": map[string]interface{}{
			"title":         title,
			"subtitle":      subtitle,
			"formattedText": formattedText,
			"image": map[string]interface{}{
				"url":               imageUrl,
				"accessibilityText": imageAccessibilityText,
			},
			"buttons": []map[string]interface{}{
				{
					"title": buttonTitle,
					"image": map[string]interface{}{
						"url":               imageUrl,
						"accessibilityText": imageAccessibilityText,
					},
					"openUrlAction": map[string]interface{}{
						"url": buttonOpenUrl,
					},
				},
				// {
				// 	"title": buttonTitle,
				// 	"openUrlAction": map[string]interface{}{
				// 		"url": buttonOpenUrl,
				// 	},
				// },
			},
			"imageDisplayOptions": imageDisplayOptions,
		},
	}
}

// {

//            "basicCard": {
//              "title": "Title: this is a title",
//              "subtitle": "This is a subtitle",
//              "formattedText": "This is a basic card.  Text in a basic card can include \"quotes\" and\n        most other unicode characters including emoji 📱.  Basic cards also support\n        some markdown formatting like *emphasis* or _italics_, **strong** or\n        __bold__, and ***bold itallic*** or ___strong emphasis___ as well as other\n        things like line  \nbreaks",
//              "image": {
//                "url": "https://example.com/image.png",
//                "accessibilityText": "Image alternate text"
//              },
//              "buttons": [
//                {
//                  "title": "This is a button",
//                  "openUrlAction": {
//                    "url": "https://assistant.google.com/"
//                  }
//                }
//              ],
//              "imageDisplayOptions": "CROPPED"
//            }
//          }

func getCellPayload(text string) Cell {
	return map[string]interface{}{
		"text": text,
	}
}

func getRowPayload(cells []Cell, dividerAfter bool) Row {
	return map[string]interface{}{
		"cells":        cells,
		"dividerAfter": dividerAfter,
	}
}

func getColumnPropertyPayload(header string, horizontalAlignment HorizontalAlignment) ColunmProperty {
	return map[string]interface{}{
		"header":              header,
		"horizontalAlignment": horizontalAlignment,
	}
}

func getTableCardResponsePayload(title, subtitle string, rows []Row, columnProperties []ColunmProperty,
	imageUrl, imageAccessibilityText,
	buttonTitle, buttonOpenUrl, imageDisplayOptions string) map[string]interface{} {
	return map[string]interface{}{
		"tableCard": map[string]interface{}{
			"title":    title,
			"subtitle": subtitle,
			"image": map[string]interface{}{
				"url":               imageUrl,
				"accessibilityText": imageAccessibilityText,
			},
			"rows":             rows,
			"columnProperties": columnProperties,
			"buttons": []map[string]interface{}{
				{
					"title": buttonTitle,
					"image": map[string]interface{}{
						"url":               imageUrl,
						"accessibilityText": imageAccessibilityText,
					},
					"openUrlAction": map[string]interface{}{
						"url": buttonOpenUrl,
					},
				},
				// {
				// 	"title": buttonTitle,
				// 	"openUrlAction": map[string]interface{}{
				// 		"url": buttonOpenUrl,
				// 	},
				// },
			},
			"imageDisplayOptions": imageDisplayOptions,
		},
	}
}

// {
//             "tableCard": {
//               "title": "Table Title",
//               "subtitle": "Table Subtitle",
//               "image": {
//                 "url": "https://developers.google.com/actions/images/badges/XPM_BADGING_GoogleAssistant_VER.png",
//                 "accessibilityText": "Alt Text"
//               },
//               "rows": [
//                 {
//                   "cells": [
//                     {
//                       "text": "row 1 item 1"
//                     },
//                     {
//                       "text": "row 1 item 2"
//                     },
//                     {
//                       "text": "row 1 item 3"
//                     }
//                   ],
//                   "dividerAfter": false
//                 },
//                 {
//                   "cells": [
//                     {
//                       "text": "row 2 item 1"
//                     },
//                     {
//                       "text": "row 2 item 2"
//                     },
//                     {
//                       "text": "row 2 item 3"
//                     }
//                   ],
//                   "dividerAfter": true
//                 },
//                 {
//                   "cells": [
//                     {
//                       "text": "row 2 item 1"
//                     },
//                     {
//                       "text": "row 2 item 2"
//                     },
//                     {
//                       "text": "row 2 item 3"
//                     }
//                   ]
//                 }
//               ],
//               "columnProperties": [
//                 {
//                   "header": "header 1",
//                   "horizontalAlignment": "CENTER"
//                 },
//                 {
//                   "header": "header 2",
//                   "horizontalAlignment": "LEADING"
//                 },
//                 {
//                   "header": "header 3",
//                   "horizontalAlignment": "TRAILING"
//                 }
//               ],
//               "buttons": [
//                 {
//                   "title": "Button Text",
//                   "openUrlAction": {
//                     "url": "https://assistant.google.com"
//                   }
//                 }
//               ]
//             }
//           }
