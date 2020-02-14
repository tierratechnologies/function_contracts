package algolia_index_mngr

var Topic = "algolia_index_mngr_v1"

type PerformAction string

const (
	WriteIndex  PerformAction = "write"
	DeleteIndex PerformAction = "delete"
)

type PSAlgoliaIndexMngrWriteV1 struct {
	Topic     string        `json:"topic"`
	ObjectId  string        `json:"objectId"`
	IndexName string        `json:"indexName"`
	Action    PerformAction `json:"action"`
	Data      interface{}   `json:"data"`
}
