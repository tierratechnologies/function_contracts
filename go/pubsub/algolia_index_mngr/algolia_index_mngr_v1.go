package algolia_index_mngr

var Topic = "algolia_index_mngr_v1"

type PerformAction string
type SearchIndex string

const (
	WriteAction  PerformAction = "write"
	DeleteAction PerformAction = "delete"

	PersonIndex SearchIndex = "PERSON"
)

type PSAlgoliaIndexMngrWriteV1 struct {
	Topic     string        `json:"topic"`
	ObjectId  string        `json:"objectId"`
	IndexName SearchIndex   `json:"indexName"`
	Action    PerformAction `json:"action"`
	Data      interface{}   `json:"data"`
}

type PersonSearchDataV1 struct {
	FirstName string `json:"first_name`
}
