package model

type Email struct {
	ReceiverEmail string `json:"receiverEmail"`
	Username      string `json:"username"`
	Subject       string `json:"subject"`
	Body          string `json:"body"`
}
