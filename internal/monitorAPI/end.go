package monitorapi

type EndMonitor struct {
	Id          int    `json:"ID"`
	ReturnValue string `json:"returnValue"`
}

type Success struct {
	Success bool `json:"Success"`
}
