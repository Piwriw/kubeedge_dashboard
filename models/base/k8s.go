package base

type CRDList struct {
	ApiVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Metadata   interface{}   `json:"metadata"`
	Items      []interface{} `json:"items"`
}

type CRD struct {
	ApiVersion string        `json:"apiVersion"`
	Kind       string        `json:"kind"`
	Metadata   interface{}   `json:"metadata"`
	Items      []interface{} `json:"items"`
	Spec       interface{}   `json:"spec"`
	Status     interface{}   `json:"status"`
}
