package algolia_index_mngr_v1

var Topic = "email_deliverability_v1"

type EmailDeliverabilityData struct {
	Topic        string `json:"topic,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	DocumentPath string `json:"documentPath,omitempty"`
}
