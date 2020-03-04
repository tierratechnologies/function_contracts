package pubsub_transaction_email_v1

var Topic = "pubsub_transaction_email_v1"

type SendEmailMsgV1 struct {
	Topic                string `json:"topic"`
	To                   string `json:"to"`
	From                 string `json:"from"`
	Subject              string `json:"subject"`
	Text                 string `json:"message"`
	TransactionEmailPath string `json:"transaction_email_path"` // the location of the orgina
}
