package models

type Notification struct {
	Teacher    Teacher   `json:"teacher"`
	Message    string    `json:"notification"`
	Recipients []Student `json:"recipients"`
}
