package algolia_index_mngr_v1

var Topic = "algolia_index_mngr_v1"

type PerformAction string
type SearchIndex string

const (
	WriteAction  PerformAction = "write"
	DeleteAction PerformAction = "delete"
)

const (
	PersonIndex SearchIndex = "PERSON"
)

type SearchData struct {
	Topic     string        `json:"topic,omitempty"`
	IndexName SearchIndex   `json:"indexName,omitempty"`
	Action    PerformAction `json:"action,omitempty"`
}

type PersonData struct {
	ObjectId     string `json:"objectId,omitempty"`
	FirstName    string `json:"firstName,omitempty"`
	FamilyName   string `json:"familyName,omitempty"`
	EmailAddress string `json:"emailAddress,omitempty"`
	UID          string `json:"uid,omitempty"`
}

type PersonSearchData struct {
	SearchData
	PersonData
}
