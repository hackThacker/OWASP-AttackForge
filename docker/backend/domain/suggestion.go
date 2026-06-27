package domain

type Suggestion struct {
	Text     string `json:"text"`
	Category string `json:"category"`
	Icon     string `json:"icon"`
}

type SuggestionUsecase interface {
	GetSuggestions() ([]*Suggestion, error)
}
