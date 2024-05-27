package model

type SlotUri struct {
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Image          string      `json:"image"`
	ExternalUrl    string      `json:"external_url"`
	BannerImageUrl string      `json:"banner_image_url"`
	YoutubeUrl     string      `json:"youtube_url"`
	Attributes     []Attribute `json:"attributes"`
}

type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}
