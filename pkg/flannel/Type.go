package flannel

type Subnet struct {
	PublicIP    string `json:"PublicIP"`
	PublicIPv6  string `json:"PublicIPv6"`
	BackendType string `json:"BackendType"`
	BackendData struct {
		VNI     int    `json:"VNI"`
		VtepMAC string `json:"VtepMAC"`
	} `json:"BackendData"`
}
