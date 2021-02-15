package queries

type EsQuery struct {
	Equals []FieldValue `json:"equals"`
}

type FieldValue struct {
	Field string      `field`
	Value interface{} `json:"value"`
}
