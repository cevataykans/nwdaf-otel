package nrf

type Config struct {
	NRFUri string `json:"nrfUri"`
	// in seconds
	HeartbeatInterval int `json:"heartbeatInterval"`
	HeartbeatTimeout  int `json:"heartbeatTimeout"`
}
