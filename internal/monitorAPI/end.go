package monitorapi

type EndMonitor struct {
	Id         int    `json:"ID"`
	Output     string `json:"output"`
	ExitStatus int    `json:"exitStatus"`
}

type Success struct {
	Success bool `json:"Success"`
}
