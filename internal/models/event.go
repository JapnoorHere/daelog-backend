package models

import "go.mongodb.org/mongo-driver/v2/bson"

type EventType string

const (
	EventTypeGitCommit   EventType = "GIT_COMMIT"
	EventTypeGitPush     EventType = "GIT_PUSH"
	EventTypeGitPull     EventType = "GIT_PULL"
	EventTypeFileSave    EventType = "FILE_SAVE"
	EventTypeWindowFocus EventType = "WINDOW_FOCUS"
	EventTypeTerminalCmd EventType = "TERMINAL_CMD"
	EventTypeBrowserTab  EventType = "BROWSER_TAB"
)

type Event struct {
	ID        bson.ObjectID `bson:"_id,omitempty"  json:"id"`
	SessionID string        `bson:"sessionId"      json:"sessionId"`
	EventType EventType     `bson:"eventType"      json:"eventType"`
	Payload   any           `bson:"payload"        json:"payload"`
	Repo      string        `bson:"repo"           json:"repo"`
	FilePath  string        `bson:"filePath"       json:"filePath"`
	Ts        int64         `bson:"ts"             json:"ts"`
	Synced    bool          `bson:"synced"         json:"synced"`
	CreatedAt int64         `bson:"createdAt"      json:"createdAt"`
}

type CreateEventRequest struct {
	SessionID string    `json:"sessionId" binding:"required"`
	EventType EventType `json:"eventType" binding:"required"`
	Payload   any       `json:"payload"   binding:"required"`
	Repo      string    `json:"repo"`
	FilePath  string    `json:"filePath"`
	Ts        int64     `json:"ts"        binding:"required"`
}
