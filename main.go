package main

import (
	"runtime"

	"gitlab.com/notionapp/api/config"
	"gitlab.com/notionapp/api/log"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config.InitEnvs()
	log.Init()
}
