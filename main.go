package main

import (
	network "./network"
	server "./server"
	. "./start"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	Logo()
	Spec()

	network.PrivateIP()
	network.PublicIP()
	server.Start()

}
