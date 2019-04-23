package pubsub_mailgun_v1

var Topic = "send_email_v1"

type GenerateLogbookLicenceMsgV1 struct {
	Topic string `json:"topic"`

	To      string `json:"to"`
	From    string `json:"from"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}
