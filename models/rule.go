package models

type RouterPointBody struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name interface{} `json:"name"`
	} `json:"metadata"`
	Spec struct {
		RuleEndpointType interface{} `json:"ruleEndpointType"`
		Properties       interface{} `json:"properties"`
	} `json:"spec"`
}

type RuleBody struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name interface{} `json:"name"`
	} `json:"metadata"`
	Spec struct {
		Source         interface{} `json:"source"`
		SourceResource interface{} `json:"sourceResource"`
		Target         interface{} `json:"target"`
		TargetResource interface{} `json:"targetResource"`
	} `json:"spec"`
}