package monitorapi

type StartMonitor struct {
	Name      string `json:"procName"`
	Workspace string `json:"workspaceName"`
	User      string `json:"user"`
	Pid       string `json:"PID"`
}

type StartReturn struct {
	Id      int  `json:"ID"`
	Success bool `json:"success"`
}
