package conf

const (
	EnvHugoEnvironment   = "HUGO_ENVIRONMENT"
	EnvCfAccountID       = "CF_ACCOUNT_ID"
	EnvCfZoneID          = "CF_ZONE_ID"
	EnvCfWorkersAPIToken = "CF_WORKERS_API_TOKEN"

	ZoneName         = "shadowcryptic.com"
	ZoneSubDomainDev = "dev"
	ZoneSubDomainAPI = "api-poc"

	// See docs at
	// https://developers.cloudflare.com/docs
	// https://api.cloudflare.com/#worker-script-properties
	// https://developers.cloudflare.com/workers/
	// https://developers.cloudflare.com/workers/#playground
	// Note Free Tier Worker limits
	// Request			100,000 requests/day
	//			   		   1000 requests/min
	// Worker memory	    128 MB
	// CPU runtime		     10 ms
	CfAPIRoot = "https://api.cloudflare.com/client/v4"
)

var (
	CfAPIToken      string
	CfAccountID     string
	CfZoneID        string
	HugoEnvironment string
	Prod            = false
)

func Config() string {
	return "conf Config"
}
