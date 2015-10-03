package main

import (
	"runtime"

	"gitlab.com/notionapp/api/config"
	"gitlab.com/notionapp/api/log"
	"gitlab.com/notionapp/api/routes"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.InitEnvs()
	log.Init()
	routes.Init()
	select {}
}
