package transation_email_v1

var Topic = "email_deliverability_v1"

type TransactionEmailData struct {
	Topic        string                 `json:"topic,omitempty"`
	TemplateName string                 `json:"templateName,omitempty"`
	To           string                 `json:"to,omitempty"`
	From         string                 `json:"from"`
	Subject      string                 `json:"subject,omitempty"`
	TestMode     bool                   `json:"testMode"`
	Variables    map[string]interface{} `json:"variables"`
}
