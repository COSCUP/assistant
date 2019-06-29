package assistant

import (
	"encoding/json"
	"github.com/COSCUP/assistant/program-fetcher"
	"strings"
)

const (
	ContextKeySelectedSession = "selected_session"
)

type UserStorage map[string]interface{}

func NewUserStorageFromDialogflowRequest(input *DialogflowRequest) *UserStorage {
	s := UserStorage{}
	if input.OriginalDetectIntentRequest.Payload.User.UserStorage == "" {
		return &s
	}

	json.Unmarshal([]byte(input.OriginalDetectIntentRequest.Payload.User.UserStorage), &s)
	return &s
}

func (s *UserStorage) EncodeToString() string {

	data, _ := json.Marshal(s)
	return string(data)
}

func (s UserStorage) addFavorite(id string) {
	favList, ok := s["favorite_list"].([]interface{})
	if !ok {
		s["favorite_list"] = []string{id}
		return
	}

	for _, v := range favList {
		if v == id {
			return
		}
	}
	s["favorite_list"] = append(favList, id)

}

func (s UserStorage) removeFavorite(id string) {
	favList, ok := s["favorite_list"].([]interface{})
	if !ok {
		s["favorite_list"] = []string{}
		return
	}

	for i, v := range favList {
		if v == id {
			s["favorite_list"] = append(favList[:i], favList[i+1:]...)
			return
		}
	}

}
func (s UserStorage) getFavoriteList() []interface{} {
	favList, ok := s["favorite_list"].([]interface{})
	if !ok {
		return []interface{}{}
	}
	return favList
}

func (s UserStorage) isSessionIdInFavorite(id string) bool {
	favList, ok := s["favorite_list"].([]interface{})
	if !ok {
		return false
	}
	for _, v := range favList {
		if v == id {
			return true
		}
	}
	return false
}

// func (s *UserStorage) AddPreviousRequestSessionList([]) {

// }

type ConversationToken struct {
	innerStorage []interface{} // Diagramflow storge token data in array?
	mapStorage   map[string]interface{}
}

// usually array
func NewConversationTokenFromDialogflowRequest(input *DialogflowRequest) *ConversationToken {
	ret := ConversationToken{
		mapStorage: map[string]interface{}{},
	}

	json.Unmarshal([]byte(input.OriginalDetectIntentRequest.Payload.Conversation.ConversationToken), &ret.innerStorage)

	return &ret
}

func (s *ConversationToken) EncodeToString() string {
	list := []interface{}{}

	for _, value := range s.innerStorage {
		if strings.HasPrefix(value.(string), "map_storage") {
			continue
		}

		list = append(list, value)
	}

	mapStorageData, _ := json.Marshal(s.mapStorage)
	list = append(list, string(mapStorageData))
	data, _ := json.Marshal(list)
	return string(data)
}

func (s *ConversationToken) AddPreviousDisplaySessionList(sessionList []fetcher.Session) {
	idList := []string{}

	for _, session := range sessionList {
		idList = append(idList, session.ID)
	}

	s.mapStorage["pervious_display_session_list"] = idList

}
