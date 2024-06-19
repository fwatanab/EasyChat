package models

type Message struct {
	ID           int64     `json:"id,omitempty"`
	Name         string    `json:"name"`
	InputMessage string    `json:"inputMessage"`
	Timestamp    string    `json:"timestamp,omitempty"`
	Type         string    `json:"type,omitempty"`
	Messages     []Message `json:"messages,omitempty"`
}
