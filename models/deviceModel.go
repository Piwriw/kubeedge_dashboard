package models

type PropertiesBody struct {
	Name string `json:"name"`
	Type struct {
		String struct {
			AccessMode   string `json:"accessMode"`
			DefaultValue string `json:"defaultValue"`
		} `json:"string"`
	} `json:"type"`
}

type DeviceModel struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		ResourceVersion string `json:"resourceVersion"`
		Name            string `json:"name"`
	} `json:"metadata"`
	Spec struct {
		Properties []PropertiesBody `json:"properties"`
	} `json:"spec"`
	Status interface{}
}

type DeviceTwins struct {
	PropertyName string `json:"propertyName"`
	Desired      struct {
		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`
		Value string `json:"value"`
	} `json:"desired"`
	Reported struct {
		Metadata struct {
			Type string `json:"type"`
		} `json:"metadata"`
		Value string `json:"value"`
	} `json:"reported"`
}

type Device struct {
	ApiVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Metadata   struct {
		Name            string `json:"name"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Spec struct {
		DeviceModelRef struct {
			Name string `json:"name"`
		} `json:"deviceModelRef"`
		NodeSelector struct {
			NodeSelectorTerms []MatchExpressions `json:"nodeSelectorTerms"`
		} `json:"nodeSelector"`
	} `json:"spec"`
	Status struct {
		Twins []DeviceTwins `json:"twins"`
	} `json:"status"`
}

type MatchExpressions struct {
	MatchExpressions []MatchExpressionsBody `json:"matchExpressions"`
}

type MatchExpressionsBody struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}
