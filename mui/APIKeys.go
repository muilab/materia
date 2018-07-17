package mui

// APIKeys are global API keys for several services
var APIKeys APIKeysData

// APIKeysData represents the data for the API keys.
type APIKeysData struct {
	Google struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"google"`

	Facebook struct {
		ID     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"facebook"`
}
