package exporter

type ResourceMetric struct {
	Name     string          `json:"name"`
	Endpoint string          `json:"endpoint"`
	Query    []ResourceQuery `json:"query,omitempty"`
	Enable   bool            `json:"enable"`
}

type ResourceQuery struct {
	Kind          ResourceKind   `json:"kind,omitempty"`
	Phase         ResourcePhase  `json:"phase,omitempty"`
	LabelSelector []SelectorSpec `json:"labelSelector,omitempty"`
}

type ResourceKind string

const (
	Edge    ResourceKind = "Edge"
	Service ResourceKind = "Service"
	Human   ResourceKind = "Human"
)

type ResourcePhase string

const (
	NotReady    ResourcePhase = "NotReady"
	Pending     ResourcePhase = "Pending"
	Synchronous ResourcePhase = "Synchronous"
	Failed      ResourcePhase = "Failed"
)

type SelectorSpec struct {
	Key      string   `json:"key"`
	Value    string   `json:"value"`
	Operator Operator `json:"op"`
}

type Operator string

const (
	Eq   =   "Eq"
	Gt   =   "Gt"
	Lt   =   "Lt"
)

func NewResourceMetric() (*ResourceMetric, error) {
	return nil, nil
}

func (metric *ResourceMetric) Start() {

}
