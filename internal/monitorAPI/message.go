package monitorapi

// Base message that will be deMuxed if endpoints are not enough. But I think they are
type Message struct {
	MessageType string      `json:"messageType"`
	Payload     interface{} `json:"payload"`
}

type CheckHealthMsg struct {
	Connection bool `json:"connection"`
}
