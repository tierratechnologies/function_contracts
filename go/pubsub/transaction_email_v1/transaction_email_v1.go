package transaction_email_v1

var Topic = "transaction_email_v1"

type TransactionEmailData struct {
	Topic              string                 `json:"topic,omitempty"`
	TemplateName       string                 `json:"template_name,omitempty"`
	To                 string                 `json:"to,omitempty"`
	From               string                 `json:"from"`
	ReceiverFirstName  string                 `json:"receiver_first_name,omitempty"`
	ReceiverFamilyName string                 `json:"receiver_family_name,omitempty"`
	Subject            string                 `json:"subject,omitempty"`
	TestMode           bool                   `json:"test_mode"`
	Variables          map[string]interface{} `json:"variables"`
	FSDocPath          string                 `json:"fs_doc_path"` // the path where the transaction email data  is located
}
