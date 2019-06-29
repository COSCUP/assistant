package assistant

// import "cloud.google.com/go/dialogflow/apiv2"
import (
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/Sirupsen/logrus"
)

type RoomNameType string

func (t RoomNameType) String() string {
	return string(t)
}

const (
	RoomNameTypeIB101  RoomNameType = "IB101"
	RoomNameTypeIB201  RoomNameType = "IB201"
	RoomNameTypeIB202  RoomNameType = "IB202"
	RoomNameTypeIB301  RoomNameType = "IB301"
	RoomNameTypeIB302  RoomNameType = "IB302"
	RoomNameTypeIB304  RoomNameType = "IB304"
	RoomNameTypeIB305  RoomNameType = "IB305"
	RoomNameTypeIB306  RoomNameType = "IB306"
	RoomNameTypeIB401  RoomNameType = "IB401"
	RoomNameTypeIB408  RoomNameType = "IB408"
	RoomNameTypeIB501  RoomNameType = "IB501"
	RoomNameTypeIB502  RoomNameType = "IB502"
	RoomNameTypeIB503  RoomNameType = "IB503"
	RoomNameTypeIE2102 RoomNameType = "IE2102"
)

type DialogflowContext struct {
	Name          string                 `json:"name"`
	LifespanCount int                    `json:"lifespanCount"`
	Parameters    map[string]interface{} `json:"parameters"`
}

type DialogflowRequest struct {
	ResponseID  string `json:"responseId"`
	QueryResult struct {
		QueryText                string                 `json:"queryText"`
		Parameters               map[string]interface{} `json:"parameters"`
		AllRequiredParamsPresent string                 `json:"allRequiredParamsPresent"`
		FulfillmentText          string                 `json:"fulfillmentText"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text"`
			} `json:"text"`
		} `json:"fulfillmentMessages"`
		OutputContexts []DialogflowContext `json:"outputContexts"`
		Intent         struct {
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
				UserID      string `json:"userId"`
				Locale      string `json:"locale"`
				LastSeen    string `json:"lastSeen"`
				UserStorage string `json:"userStorage"`
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
	OutputContexts []DialogflowContext `json:"outputContexts",emitempty`
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
		log.Println(" ", ip.Name(), " ? ", request.QueryResult.Intent.DisplayName)
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

			log.Println("outputContexts", payload["outputContexts"])

			contexts, ok := payload["outputContexts"].(map[string]interface{})
			if ok {
				dContexts := []DialogflowContext{}
				for key, context := range contexts {
					dContexts = append(dContexts,
						DialogflowContext{
							Name:          request.Session + "/contexts/" + key,
							LifespanCount: 5,
							Parameters:    context.(map[string]interface{}),
						})
				}

				response.OutputContexts = dContexts
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

func (r DialogflowRequest) RoomName() RoomNameType {
	return RoomNameType(r.QueryResult.Parameters["RoomName"].(string))
}

func (r DialogflowRequest) UserId() string {
	return r.OriginalDetectIntentRequest.Payload.Conversation.ConversationID
}

func (r DialogflowRequest) Context(key string) map[string]interface{} {
	for _, s := range r.QueryResult.OutputContexts {
		if strings.HasSuffix(s.Name, "/"+key) {
			return s.Parameters
		}
	}
	return nil

}

func (r DialogflowRequest) SelectedNumber() int {
	switch v := r.QueryResult.Parameters["number"].(type) {
	case string:
		if v == "" {
			return 0
		}
		// todo strconv
		return 0
	case float64:
		return int(v)
	default:
		return 0

	}

}

// func (r DialogflowRequest) AddContext(key string, lifespanCount int) {
// 	r.QueryResult.OutputContexts = append(r.QueryResult.OutputContexts,
// 		DialogflowContext{
// 			Name:          "projects/coscup/agent/sessions/ABwppHH_zjv8cpuXq5jWgZ3jX3Hw4D96diVEz4vD0cdKiiDu84s6fjNt0y1XHxCpGYk51sMw/contexts/test",
// 			LifespanCount: 5,
// 			Parameters: map[string]interface{}{
// 				"key": "value",
// 			},
// 		},
// 	)
// }

func writeDialogflowResponse(w http.ResponseWriter, dr *DialogflowResponse) {
	data, _ := json.Marshal(dr)
	w.Write(data)
}
