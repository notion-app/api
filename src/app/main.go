package main

import (
	"runtime"
	"notion/config"
	"notion/db"
	"notion/log"
	"notion/routes"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.InitEnvs()
	log.Init()
	db.Init()
	routes.Init()
	select {}
}
