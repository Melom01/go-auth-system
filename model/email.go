package model

type VerificationEmail struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Subject  string `json:"subject"`
	Body     string `json:"body"`
}
