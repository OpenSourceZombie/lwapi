package lwapi

//VirtualServer struct type
type VirtualServer struct {
	CloudServerID interface{} `json:"cloudServerId"`
	Contract      struct {
		BillingCycle      int64   `json:"billingCycle"`
		BillingFrequency  string  `json:"billingFrequency"`
		Currency          string  `json:"currency"`
		EndsAt            string  `json:"endsAt"`
		ID                string  `json:"id"`
		PricePerFrequency float64 `json:"pricePerFrequency"`
		StartsAt          string  `json:"startsAt"`
	} `json:"contract"`
	CustomerID    string `json:"customerId"`
	DataCenter    string `json:"dataCenter"`
	FirewallState string `json:"firewallState"`
	Hardware      struct {
		CPU struct {
			Cores int64 `json:"cores"`
		} `json:"cpu"`
		Memory struct {
			Amount int64  `json:"amount"`
			Unit   string `json:"unit"`
		} `json:"memory"`
		Storage struct {
			Amount int64  `json:"amount"`
			Unit   string `json:"unit"`
		} `json:"storage"`
	} `json:"hardware"`
	ID  string `json:"id"`
	Ips []struct {
		IP      string `json:"ip"`
		Type    string `json:"type"`
		Version int    `json:"version"`
	} `json:"ips"`
	Iso struct {
		DisplayName string `json:"displayName"`
		ID          string `json:"id"`
		Name        string `json:"name"`
	} `json:"iso"`
	Reference       string `json:"reference"`
	ServiceOffering string `json:"serviceOffering"`
	SLA             string `json:"sla"`
	State           string `json:"state"`
	Template        string `json:"template"`
}

//Metadata struct type
type Metadata struct {
	Limit      int64 `json:"limit"`
	Offset     int64 `json:"offset"`
	TotalCount int64 `json:"totalCount"`
}

//VirtualServerList struct type
type VirtualServerList struct {
	VirtualServers []VirtualServer
	Metadata
}

//Metrics struct type
type Metrics struct {
	Metadata struct {
		Aggregation string `json:"aggregation"`
		From        string `json:"from"`
		Granularity string `json:"granularity"`
		To          string `json:"to"`
	} `json:"_metadata"`
	TrafficMetrics struct {
		DownPublic struct {
			Unit   string `json:"unit"`
			Values []struct {
				Timestamp string `json:"timestamp"`
				Value     int64  `json:"value"`
			} `json:"values"`
		} `json:"DOWN_PUBLIC"`
		UpPublic struct {
			Unit   string `json:"unit"`
			Values []struct {
				Timestamp string `json:"timestamp"`
				Value     int64  `json:"value"`
			} `json:"values"`
		} `json:"UP_PUBLIC"`
	} `json:"metrics"`
}

//AsyncResponse struct type
type AsyncResponse struct {
	CreatedAt string `json:"createdAt"`
	ID        string `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
}

//Credentials struct type
type Credentials struct {
	Type     string `json:"type"`
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

//CredentialsList struct type
type CredentialsList struct {
	Credentials []Credentials `json:"credentials"`
	Metadata
}
type ErrorResponse struct {
	ErrorCode    int64  `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Reference    string `json:"reference"`
	UserMessage  string `json:"userMessage"`
}
