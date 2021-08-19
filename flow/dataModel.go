package flow

type ChangeListFlowInfo struct {
	// PrecheckinInfo PrecheckinInfo
	Producers []AutobuildInfo `json:"producers"`
	Consumers []AutobuildInfo `json:"consumers"`
}

type PrecheckinInfo struct {
	Stages []string `json:"stages"`
	Current string `json:"current"`
	Completed bool `json:"completed"`
}

type AutobuildInfo struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
}
