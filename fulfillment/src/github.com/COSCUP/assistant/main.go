package assistant

// import "cloud.google.com/go/dialogflow/apiv2"
import (
	"encoding/json"
	"net/http"
)

type DialogflowRequest struct {
	ResponseID  string `json:"responseId"`
	QueryResult struct {
		QueryText  string `json:"queryText"`
		Parameters struct {
			RoomName string `json:"RoomName"`
		} `json:"parameters"`
		AllRequiredParamsPresent string `json:"allRequiredParamsPresent"`
		FulfillmentText          string `json:"fulfillmentText"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
		} `json:"fulfillmentMessages"`
		OutputContexts []struct {
			Name          string `json:"name"`
			LifespanCount string `json:"lifespanCount,omitempty"`
			Parameters    struct {
				RoomName         string `json:"RoomName"`
				RoomNameOriginal string `json:"RoomName.original"`
			} `json:"parameters"`
		} `json:"outputContexts"`
		Intent struct {
			Name        string `json:"name"`
			DisplayName string `json:"displayName"`
		} `json:"intent"`
		IntentDetectionConfidence string `json:"intentDetectionConfidence"`
		LanguageCode              string `json:"languageCode"`
	} `json:"queryResult"`
	OriginalDetectIntentRequest struct {
		Source  string `json:"source"`
		Version string `json:"version"`
		Payload struct {
			User struct {
				UserID   string `json:"userId"`
				Locale   string `json:"locale"`
				LastSeen string `json:"lastSeen"`
			} `json:"user"`
			Conversation struct {
				ConversationID    string `json:"conversationId"`
				Type              string `json:"type"`
				ConversationToken string `json:"conversationToken"`
			} `json:"conversation"`
			Inputs []struct {
				Intent    string `json:"intent"`
				RawInputs []struct {
					InputType string `json:"inputType"`
					Query     string `json:"query"`
				} `json:"rawInputs"`
				Arguments []struct {
					Name      string `json:"name"`
					RawText   string `json:"rawText"`
					TextValue string `json:"textValue"`
				} `json:"arguments"`
			} `json:"inputs"`
			Surface struct {
				Capabilities []struct {
					Name string `json:"name"`
				} `json:"capabilities"`
			} `json:"surface"`
			IsInSandbox       bool `json:"isInSandbox"`
			AvailableSurfaces []struct {
				Capabilities []struct {
					Name string `json:"name"`
				} `json:"capabilities"`
			} `json:"availableSurfaces"`
		} `json:"payload"`
	} `json:"originalDetectIntentRequest"`
	Session string `json:"session"`
}

type DialogflowResponse struct {
	FulfillmentText string `json:"fulfillmentText"`
	// FulfillmentMessages []struct {
	// 	Card struct {
	// 		Title    string `json:"title"`
	// 		Subtitle string `json:"subtitle"`
	// 		ImageURI string `json:"imageUri"`
	// 		Buttons  []struct {
	// 			Text     string `json:"text"`
	// 			Postback string `json:"postback"`
	// 		} `json:"buttons"`
	// 	} `json:"card"`
	// } `json:"fulfillmentMessages"`
	Source  string                 `json:"source",emitempty`
	Payload map[string]interface{} `json:"payload"`
	// Google  `json:"google"`
	// Google struct {
	// 	ExpectUserResponse string `json:"expectUserResponse"`
	// 	RichResponse       struct {
	// 		Items []struct {
	// 			SimpleResponse struct {
	// 				TextToSpeech string `json:"textToSpeech"`
	// 			} `json:"simpleResponse"`
	// 		} `json:"items"`
	// 	} `json:"richResponse"`
	// } `json:"google",emitempty`
	// Facebook struct {
	// 	Text string `json:"text"`
	// } `json:"facebook"`
	// Slack struct {
	// 	Text string `json:"text"`
	// } `json:"slack"`
	// } `json:"payload",emitempty`
	// OutputContexts []struct {
	// 	Name          string `json:"name"`
	// 	LifespanCount string `json:"lifespanCount"`
	// 	Parameters    struct {
	// 		Param string `json:"param"`
	// 	} `json:"parameters"`
	// } `json:"outputContexts",emitempty`
	// FollowupEventInput *struct {
	// 	Name         string `json:"name"`
	// 	LanguageCode string `json:"languageCode"`
	// 	Parameters   struct {
	// 		Param string `json:"param"`
	// 	} `json:"parameters"`
	// } `json:"followupEventInput",emitempty`
}

func NewDialogflowResponseWithTestMessage(msg string) *DialogflowResponse {
	return &DialogflowResponse{}
}

func RequestHandler(w http.ResponseWriter, r *http.Request, data []byte) {
	request := DialogflowRequest{}
	json.Unmarshal(data, &request)

	for _, ip := range intentProcessorList {
		if ip.Name() == request.QueryResult.Intent.DisplayName {
			// match
			payload := ip.Payload(&request)
			response := DialogflowResponse{
				FulfillmentText: "test",
				Source:          request.OriginalDetectIntentRequest.Source,
				Payload: map[string]interface{}{
					"google": payload,
				},
			}
			writeDialogflowResponse(w, &response)

			return
		}
	}
	response := DialogflowResponse{
		FulfillmentText: "intent " + request.QueryResult.Intent.DisplayName + " not implement",
		Source:          request.OriginalDetectIntentRequest.Source,
		Payload:         map[string]interface{}{},
	}

	// r2, _ := json.Marshal(&map[string]interface{}{
	// 	"fulfillmentText": "texs",
	// })
	// w.Write(r2)
	writeDialogflowResponse(w, &response)

}

func writeDialogflowResponse(w http.ResponseWriter, dr *DialogflowResponse) {
	data, _ := json.Marshal(dr)
	w.Write(data)
}
