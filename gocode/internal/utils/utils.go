package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/timecode/CloudflarePagesPoC/gocode/internal/conf"
)

// CheckProdArg ...
func CheckProdArg(in []string) {
	// Add / Update actual worker
	args := os.Args[1:]
	if len(args) > 0 && args[0] == "prod" {
		conf.Prod = true
	}
}

// LoadAPIToken ...
func LoadAPIToken() {
	env := conf.EnvCfWorkersAPIToken
	conf.CfAPIToken = os.Getenv(env)
	if len(conf.CfAPIToken) < 1 {
		err := fmt.Errorf("error: ENV has invalid %s\n"+
			"export %s=???", env, env)
		log.Fatal(err)
	}
}

// LoadAccountID ...
func LoadAccountID() {
	env := conf.EnvCfAccountID
	conf.CfAccountID = os.Getenv(env)
	if len(conf.CfAccountID) < 1 {
		err := fmt.Errorf("error: ENV has invalid %s\n"+
			"export %s=???", env, env)
		log.Fatal(err)
	}
}

// LoadZoneID ...
func LoadZoneID() {
	env := conf.EnvCfZoneID
	conf.CfZoneID = os.Getenv(env)
	if len(conf.CfZoneID) < 1 {
		err := fmt.Errorf("error: ENV has invalid %s\n"+
			"export %s=???", env, env)
		log.Fatal(err)
	}
}

// InitRequest ...
func InitRequest(method string, url string, contentType string, body io.Reader) (request *http.Request, err error) {
	if request, err = http.NewRequest(method, url, body); err != nil {
		return
	}
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", conf.CfAPIToken))
	request.Header.Set("Content-Type", contentType)
	return
}
