package nrf

// NFService represents a service a Network Function Instance can offer.
type NFService struct {
	ServiceInstanceId string `json:"serviceInstanceId"`
	ServiceName       string `json:"serviceName"`
	Versions          []struct {
		ApiVersionInUri string `json:"apiVersionInUri"`
		ApiFullVersion  string `json:"apiFullVersion"`
	} `json:"versions"`
	Scheme          string `json:"scheme"`
	NfServiceStatus string `json:"nfServiceStatus"`
	IpEndPoints     []struct {
		Ipv4Address string `json:"ipv4Address"`
		Port        int    `json:"port"`
	} `json:"ipEndPoints"`
}

// NFInstance represents any Network Function (NF) registered in the 5G Core.
type NFInstance struct {
	// UUID
	NfInstanceId  string      `json:"nfInstanceId"`
	NfType        string      `json:"nfType"`
	NfStatus      string      `json:"nfStatus"`
	HeartBeat     int         `json:"heartBeatTimer"`
	NfServices    []NFService `json:"nfServices"`
	Ipv4Addresses []string    `json:"ipv4Addresses"`
}
