package models

// Story json item
type Story struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

// StoryOption option:[]
type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
