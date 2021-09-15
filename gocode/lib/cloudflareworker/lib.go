package cloudflareworker

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/timecode/CloudflarePagesPoC/gocode/internal/conf"
	"github.com/timecode/CloudflarePagesPoC/gocode/internal/utils"
)

const (
	WorkerJSDir = "./gocode/internal/cloudflare-workers"

	WorkerBuildDataPlaceHolder = "POPULATE BUILD DATA HERE"
	WorkerScriptNameAPI        = "api-shadowcryptic-com"
)

var (
	WorkerAPIIn  = fmt.Sprintf("%s/%s", WorkerJSDir, "api-shadowcryptic-com-TEMPLATE.js")
	WorkerAPIOut = fmt.Sprintf("%s/%s", WorkerJSDir, "tmp/api-shadowcryptic-com.js")
)

type Worker struct {
	ScriptName string   `json:"id,omitempty"`
	Script     string   `json:"script,omitempty"`
	Size       string   `json:"size,omitempty"`
	ETag       string   `json:"etag,omitempty"`
	UsageModel string   `json:"usage_model,omitempty"`
	Handlers   []string `json:"handlers,omitempty"`
	Routes     []*Route `json:"routes,omitempty"`
	Created    string   `json:"created_on,omitempty"`
	Modified   string   `json:"modified_on,omitempty"`
}

type WorkersRequest struct {
	Success  bool      `json:"success,omitempty"`
	Errors   []string  `json:"errors,omitempty"`
	Messages []string  `json:"messages,omitempty"`
	Result   []*Worker `json:"result,omitempty"`
}

type WorkerUploadRequest struct {
	Success  bool     `json:"success,omitempty"`
	Errors   []string `json:"errors,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type Route struct {
	ID         string `json:"id,omitempty"`
	Pattern    string `json:"pattern,omitempty"`
	ScriptName string `json:"script,omitempty"`
	FailOpen   bool   `json:"request_limit_fail_open,omitempty"`
}

type RoutesRequest struct {
	Result   []*Route `json:"result,omitempty"`
	Success  bool     `json:"success,omitempty"`
	Errors   []string `json:"errors,omitempty"`
	Messages []string `json:"messages,omitempty"`
}

type WorkerRouteRequest struct {
	Success  bool     `json:"success,omitempty"`
	Errors   []string `json:"errors,omitempty"`
	Messages []string `json:"messages,omitempty"`
	Result   *Route   `json:"result,omitempty"`
}

func Hello() string {
	return "cloudflareworker hello"
}

// UploadWorker ...
func UploadWorker(filename string, name string) (err error) {

	var (
		b         []byte
		request   *http.Request
		client    = &http.Client{}
		response  *http.Response
		bodyBytes []byte
		res       WorkerUploadRequest
	)

	if b, err = utils.LoadFile(filename); err != nil {
		return
	}

	url := fmt.Sprintf("%s/accounts/%s/workers/scripts", conf.CfAPIRoot, conf.CfAccountID)
	if request, err = utils.InitRequest("PUT", fmt.Sprintf("%s/%s", url, name),
		"application/javascript", bytes.NewReader(b),
	); err != nil {
		return
	}

	if response, err = client.Do(request); response.StatusCode != 200 {
		err = fmt.Errorf("error uploadWorker %s", response.Status)
		return
	}
	defer response.Body.Close()

	if response.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(response.Body)
	}

	if err = json.Unmarshal(bodyBytes, &res); err != nil {
		err = fmt.Errorf("error uploadWorker json.Unmarshal: %s\n%#v", string(bodyBytes), err)
		return
	}

	if !res.Success {
		err = fmt.Errorf("error uploadWorker: %s\n%s", res.Errors, res.Messages)
		return
	}

	return
}

// ListWorkers ...
func ListWorkers() (out map[string]*Worker, err error) {

	var (
		request   *http.Request
		client    = &http.Client{}
		response  *http.Response
		bodyBytes []byte
		res       WorkersRequest
	)

	url := fmt.Sprintf("%s/accounts/%s/workers/scripts", conf.CfAPIRoot, conf.CfAccountID)
	if request, err = utils.InitRequest("GET", url,
		"application/json", nil); err != nil {
		return
	}

	if response, err = client.Do(request); response.StatusCode != 200 {
		err = fmt.Errorf("error listWorkers %s", response.Status)
		return
	}
	defer response.Body.Close()

	if response.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(response.Body)
	}

	if err = json.Unmarshal(bodyBytes, &res); err != nil {
		err = fmt.Errorf("error listWorkers json.Unmarshal: %s\n%#v", string(bodyBytes), err)
		return
	}

	if !res.Success {
		err = fmt.Errorf("error listWorkers: %s\n%s", res.Errors, res.Messages)
		return
	}

	out = make(map[string]*Worker)
	for _, r := range res.Result {
		out[r.ScriptName] = r
		// if !r.FailOpen {
		// 	fmt.Printf("WARNING: FailOpen is '%v' for route %s\n", r.FailOpen, r.Pattern)
		// }
	}

	return
}

// CreateWorkerRoute ...
func CreateWorkerRoute(pattern string, script string) (err error) {

	var (
		routes    map[string]*Route
		request   *http.Request
		client    = &http.Client{}
		response  *http.Response
		bodyBytes []byte
		res       WorkerRouteRequest
	)

	type postData struct {
		Pattern  string `json:"pattern"`
		Script   string `json:"script"`
		FailOpen bool   `json:"request_limit_fail_open"`
	}

	// get routes so we know if we need to create or update (below)
	if routes, err = listWorkerRoutes(); err != nil {
		return
	}

	url := fmt.Sprintf("%s/zones/%s/workers/routes", conf.CfAPIRoot, conf.CfZoneID)
	method := "POST"
	if route, exists := routes[pattern]; exists {
		url = fmt.Sprintf("%s/%s", url, route.ID)
		method = "PUT"
	}

	jsonValue, _ := json.Marshal(postData{
		Pattern:  pattern,
		Script:   script,
		FailOpen: true,
	})

	if request, err = utils.InitRequest(method, url,
		"application/json", bytes.NewReader(jsonValue)); err != nil {
		return
	}

	if response, err = client.Do(request); response.StatusCode != 200 {
		err = fmt.Errorf("error createWorkerRoute %s", response.Status)
		return
	}
	defer response.Body.Close()

	if response.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(response.Body)
	}

	if err = json.Unmarshal(bodyBytes, &res); err != nil {
		err = fmt.Errorf("error createWorkerRoute json.Unmarshal: %s\n%#v", string(bodyBytes), err)
		return
	}

	if !res.Success {
		err = fmt.Errorf("error createWorkerRoute: %s\n%s", res.Errors, res.Messages)
		return
	}

	return
}

// listWorkerRoutes ...
func listWorkerRoutes() (out map[string]*Route, err error) {

	var (
		request   *http.Request
		client    = &http.Client{}
		response  *http.Response
		bodyBytes []byte
		res       RoutesRequest
	)

	url := fmt.Sprintf("%s/zones/%s/workers/routes", conf.CfAPIRoot, conf.CfZoneID)
	if request, err = utils.InitRequest("GET", url,
		"application/json", nil); err != nil {
		return
	}

	if response, err = client.Do(request); response.StatusCode != 200 {
		err = fmt.Errorf("error listWorkerRoutes %s", response.Status)
		return
	}
	defer response.Body.Close()

	if response.Body != nil {
		bodyBytes, _ = ioutil.ReadAll(response.Body)
	}

	if err = json.Unmarshal(bodyBytes, &res); err != nil {
		err = fmt.Errorf("error listWorkerRoutes json.Unmarshal: %s\n%#v", string(bodyBytes), err)
		return
	}

	if !res.Success {
		err = fmt.Errorf("error listWorkerRoutes: %s\n%s", res.Errors, res.Messages)
		return
	}

	out = make(map[string]*Route)
	for _, r := range res.Result {
		out[r.Pattern] = r
		if !r.FailOpen {
			fmt.Printf("WARNING: FailOpen is '%v' for route %s\n", r.FailOpen, r.Pattern)
		}
	}

	return
}

// CreateTestTimeWorkerJS ...
func CreateAPIWorkerJS() (err error) {

	// buildTimeStr := time.Now().UTC().Format("2006-01-02T15:04:05Z07:00")

	// load template
	var b []byte
	if b, err = utils.LoadFile(WorkerAPIIn); err != nil {
		return
	}

	// replace text placeholders
	s := string(b)
	t := time.Now().UTC()
	version := fmt.Sprintf("%s", t.Format(time.RFC3339))
	s = strings.Replace(s, WorkerBuildDataPlaceHolder, version, -1)
	b = []byte(s)

	// save for upload
	err = utils.SaveFile(WorkerAPIOut, b)
	return
}
