package stackpointio

// Solution is a application or process running with or on a kubernetes cluster
type Solution struct {
	ID       int    `json:"pk"`
	Solution string `json:"solution"`
	Keyset   int    `json:"keyset,omitempty"` // only for turbonomic and sysdig
	MaxNodes int    `json:"max_nodes"`        // only for autoscaler
}
