package domain

type RealtimeChatEvent struct {
	Id string `json:"messageId"`
	Text string `json:"text"`
	User UserModel `json:"user"`
}

