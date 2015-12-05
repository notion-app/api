package main

import (
	"math/rand"
	"notion/cache"
	"notion/config"
	"notion/db"
	"notion/log"
	"notion/routes"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UTC().UnixNano())
	config.InitEnvs()
	log.Init()
	db.Init()
	cache.InitDumps()
	routes.Init()
	select {}
}
