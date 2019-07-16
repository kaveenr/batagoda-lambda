package telegram

// Used online tool :3 https://mholt.github.io/json-to-go/
type WebHookRequest struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		Date int `json:"date"`
		Chat struct {
			LastName  string `json:"last_name"`
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"chat"`
		MessageID int `json:"message_id"`
		From      struct {
			LastName  string `json:"last_name"`
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			Username  string `json:"username"`
		} `json:"from"`
		Text string `json:"text"`
	} `json:"message"`
}

type WebHookResponse struct {
	Method 			string 	`json:"method"`
	ChatID			int	`json:"chat_id"`
	Text			string 	`json:"text"`
}