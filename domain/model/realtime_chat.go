package model

type RealtimeChatEvent struct {
	Id string `json:"messageId"`
	Text string `json:"text"`
	User UserResponse `json:"user"`
}

