package nks

//NetworkComponent describes a network configuration
type NetworkComponent struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Cidr          string `json:"cidr"`
	ComponentType string `json:"component_type"`
	ProviderID    string `json:"provider_id"`
	VpcID         string `json:"vpcId"`
	Zone          string `json:"zone"`
}
