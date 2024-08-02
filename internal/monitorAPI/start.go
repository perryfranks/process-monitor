package monitorapi

type StartMonitor struct {
	Name      string `json:"procName"`
	Workspace string `json:"workspaceName"`
}

type StartReturn struct {
	Id      int  `json:"ID"`
	Success bool `json:"success"`
}
