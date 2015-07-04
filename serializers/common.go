package serializers

type ListMeta struct {
	ObjectType string `json:"object"`
	Count      uint16 `json:"total_count"`
	More       bool   `json:"has_more"`
	URL        string `json:"url"`
}
