package dtos

type Quiz struct {
	Question   string     `json:"question"`
	Options    []Option   `json:"options"`
	Explanation string     `json:"explanation"`
}

type Option struct {
	Text     string `json:"text"`
	IsAnswer bool   `json:"is_answer"`
}