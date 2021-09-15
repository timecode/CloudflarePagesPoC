package main

import (
	"fmt"
	"log"

	"github.com/timecode/CloudflarePagesPoC/gocode/internal/conf"
	"github.com/timecode/CloudflarePagesPoC/gocode/internal/utils"
	lib "github.com/timecode/CloudflarePagesPoC/gocode/lib/cloudflareworker"
)

func main() {

	// fmt.Printf("%#v\n", lib.Hello())

	var (
		err         error
		workerStage string
		scriptname  string
		route       string
	)

	utils.LoadAPIToken()
	utils.LoadAccountID()
	utils.LoadZoneID()
	utils.LoadProdStatus()

	// Add / Update actual worker

	workerStage = "dev"
	if conf.Prod {
		workerStage = "prod"
	}

	// Add / Update API worker
	// CREATE
	route = ""
	scriptname = lib.WorkerScriptNameAPI
	if !conf.Prod {
		scriptname = fmt.Sprintf("%s-%s", scriptname, conf.ZoneSubDomainDev)
		route = fmt.Sprintf("%s-%s.", conf.ZoneSubDomainAPI, conf.ZoneSubDomainDev)
	} else {
		route = fmt.Sprintf("%s.", conf.ZoneSubDomainAPI)
	}
	route = fmt.Sprintf("%s%s/*", route, conf.ZoneName)
	fmt.Printf("... creating %s API cloudflare worker\n", workerStage)
	if err = lib.CreateAPIWorkerJS(); err != nil {
		log.Fatal(err)
	}
	// UPLOAD
	fmt.Printf("... uploading %s API cloudflare worker\n", workerStage)
	if err = lib.UploadWorker(lib.WorkerAPIOut, scriptname); err != nil {
		log.Fatal(err)
	}
	// HOOK UP
	fmt.Printf("... adding route for %s API cloudflare worker\n", workerStage)
	if err = lib.CreateWorkerRoute(route, scriptname); err != nil {
		log.Fatal(err)
	}

	// OPTIONALLY, LIST WORKERS
	var workers map[string]*lib.Worker
	if workers, err = lib.ListWorkers(); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	for k, v := range workers {
		fmt.Printf("%s  %s  %s\n", v.Created, v.Modified, k)
		for _, route := range v.Routes {
			fmt.Printf("\t%s\n", route.Pattern)
		}
		fmt.Println()
	}
}
