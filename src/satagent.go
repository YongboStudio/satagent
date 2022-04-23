package main

import (
	"flag"
	"fmt"
	"github.com/jakecoffman/cron"
	"github.com/yongbostudio/satagent/satagent/src/funcs"
	"github.com/yongbostudio/satagent/satagent/src/common"
	"github.com/yongbostudio/satagent/satagent/src/http"
	"os"
	"runtime"
	//"sync"
)

// Init config
var Version = "0.8.0"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	version := flag.Bool("v", false, "show version")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}
	common.ParseConfig(Version)
	go funcs.ClearArchive()
	c := cron.New()
	c.AddFunc("*/60 * * * * *", func() {
		go funcs.Ping()
		go funcs.Mapping()
		if common.Cfg.Mode["Type"] == "cloud" {
			go funcs.StartCloudMonitor()
		}
	}, "ping")
	c.AddFunc("0 0 * * * *", func() {
		go funcs.ClearArchive()
	}, "mtc")
	c.Start()
	http.StartHttp()
}
