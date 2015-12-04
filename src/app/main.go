package main

import (
	"math/rand"
	"notion/config"
	"notion/db"
	"notion/log"
	"notion/routes"
	"notion/ws"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	config.InitEnvs()
	log.Init()
	db.Init()
	ws.InitCommitter()
	routes.Init()
	select {}
}
