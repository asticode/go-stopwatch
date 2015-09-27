package stopwatch

type Configuration struct {
	Id      string  `json:"id"`
	Enabler Enabler `json:"enabler"`
}

type Enabler struct {
	Headers     map[string]string `json:"headers"`
	IpAddresses []string          `json:"ip_addresses"`
}
