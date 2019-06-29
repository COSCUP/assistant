package assistant

const (
	IntentFallback      = "Default Fallback Intent"
	IntentWelcomeIntent = "Default WelcomeIntent"
)

type IntentProcessor interface {
	Name() string
	Payload(request *DialogflowRequest) map[string]interface{}
}

var intentProcessorList = []IntentProcessor{
	WelcomeIntentProcessor{},
	HelpIntentProcessor{},
	RegisterIntentProcessor{},
	AskProgramListByRoomIntentProcessor{},
	AskProgramListByTimeIntentProcessor{},
	AskProgramByProgramIntentProcessor{},

	AddFavoriteIntentProcessor{},
	QueryFavoriteListIntentProcessor{},
	RemoveFavoriteIntentProcessor{},

	DefaultFallbackIntent{},
	LocationByLocationNameIntentProcessor{},
	QuitIntentProcessor{},
}
