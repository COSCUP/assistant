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
	HorizontalAlignmentTrailing HorizontalAlignment = "TRAINLING"
)

type Cell map[string]interface{}
type Row map[string]interface{}
type ColunmProperty map[string]interface{}
