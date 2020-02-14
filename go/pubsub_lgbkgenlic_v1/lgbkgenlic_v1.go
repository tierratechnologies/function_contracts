// This package contains the information for
// the pubsub messages to generate a licence

package pubsub_lgbkgenlic_v1

var Topic = "generate_logbook_licence_v1"

type SubscriptionType string

const (
	TsOperatorSub SubscriptionType = "transportOp"
	DriverSub     SubscriptionType = "driver"
)

type GenerateLogbookLicenceMsgV1 struct {
	Topic            string           `json:"topic"`
	UserId           string           `json:"userId"`
	IssuerId         string           `json:"issuerId"`
	SubscriptionId   string           `json:"subscriptionId"`
	SubscriptionType SubscriptionType `json:"subscriptionType"`
	AccountName      string           `json:"accountName"`
	AddonAssigned    bool             `json:"addonAssigned"`
}
