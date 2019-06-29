package fetcher

import (
	"time"
)

type SessionLocalization struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Session struct {
	ID        string              `json:"id"`
	Type      string              `json:"type"`
	Room      string              `json:"room"`
	Broadcast string              `json:"broadcast"`
	Start     time.Time           `json:"start"`
	End       time.Time           `json:"end"`
	Qa        string              `json:"qa"`
	Slide     string              `json:"slide"`
	Live      string              `json:"live"`
	Record    string              `json:"record"`
	Zh        SessionLocalization `json:"zh"`
	En        SessionLocalization `json:"en"`
	Speakers  []string            `json:"speakers"`
	Tags      []string            `json:"tags"`
}

type SpeakerLocalization struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

type Speaker struct {
	ID     string              `json:"id"`
	Avatar string              `json:"avatar"`
	Zh     SpeakerLocalization `json:"zh"`
	En     SpeakerLocalization `json:"en"`
}

type SessionTypeLocalization struct {
	Name string `json:"name"`
}

type SessionType struct {
	ID string                  `json:"id"`
	Zh SessionTypeLocalization `json:"zh"`
	En SessionTypeLocalization `json:"en"`
}

type RoomLocalization struct {
	Name string `json:"name"`
}

type Room struct {
	ID string           `json:"id"`
	Zh RoomLocalization `json:"zh"`
	En RoomLocalization `json:"en"`
}

type TagLocalization struct {
	Name string `json:"name"`
}

type Tag struct {
	ID string          `json:"id"`
	Zh TagLocalization `json:"zh"`
	En TagLocalization `json:"en"`
}

type ProgramsResponedPayload struct {
	Sessions     []Session     `json:"sessions"`
	Speakers     []Speaker     `json:"speakers"`
	SessionTypes []SessionType `json:"session_types"`
	Rooms        []Room        `json:"rooms"`
	Tags         []Tag         `json:"tags"`
}

func (p *ProgramsResponedPayload) GetSessionByID(id string) *Session {
	for _, s := range p.Sessions {
		if s.ID == id {
			t := s
			return &t
		}
	}
	return nil
}

func (p *ProgramsResponedPayload) GetSpeakerByID(id string) *Speaker {
	for _, s := range p.Speakers {
		if s.ID == id {
			t := s
			return &t
		}
	}
	return nil
}

func (p *Session) SpeakerPhotoUrl() string {
	if len(p.Speakers) == 0 {

		return "https://coscup.org/2019/_nuxt/img/c2f9236.png"
	}
	speakerId := p.Speakers[0]

	prog, _ := GetPrograms()
	speaker := prog.GetSpeakerByID(speakerId)
	if speaker != nil {
		return speaker.Avatar
	}

	return "https://coscup.org/2019/_nuxt/img/c2f9236.png"
}
