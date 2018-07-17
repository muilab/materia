package mui

import (
	"io/ioutil"
	"os"
	"path"

	jsoniter "github.com/json-iterator/go"
	"github.com/muilab/materia/mui/utils"
)

// Root is the full path to the root directory of the repository.
var Root = path.Join(os.Getenv("GOPATH"), "src/github.com/muilab/materia")

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

// init loads the API keys.
func init() {
	// Path for API keys
	apiKeysPath := path.Join(Root, "security/api-keys.json")

	// If the API keys file is not available, create a symlink to the default API keys
	if _, err := os.Stat(apiKeysPath); os.IsNotExist(err) {
		defaultAPIKeysPath := path.Join(Root, "security/default/api-keys.json")
		err := os.Link(defaultAPIKeysPath, apiKeysPath)
		utils.PanicOnError(err)
	}

	// Load API keys
	data, err := ioutil.ReadFile(apiKeysPath)
	utils.PanicOnError(err)

	// Parse JSON
	err = jsoniter.Unmarshal(data, &APIKeys)
	utils.PanicOnError(err)
}
