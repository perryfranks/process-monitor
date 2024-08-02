package monitorapi

type EndMonitor struct {
	Id          string `json:"ID"`
	End         bool   `json:"end"`
	ReturnValue string `json:"returnValue"`
}
