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

type AlgoliaIndexMngrData struct {
	Topic     string        `json:"topic"`
	ObjectId  string        `json:"objectId"`
	IndexName SearchIndex   `json:"indexName"`
	Action    PerformAction `json:"action"`
	Data      interface{}   `json:"data"`
}

type PersonSearchData struct {
	FirstName    string `json:"first_name`
	FamilyName   string `json:"family_name"`
	EmailAddress string `json:"email_address"`
	uid          string `json:"uid"`
}
