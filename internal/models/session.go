package models

type Session struct {
	ID       string   `json:"id"`
	StartTs  int64    `json:"startTs"`
	EndTs    int64    `json:"endTs"`
	Duration int64    `json:"duration"`
	Events   []Event  `json:"events"`
	Repos    []string `json:"repos"`
}
