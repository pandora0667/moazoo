package main

import (
	"runtime"
	. "./start"
	network "./network"
)

func main()  {

	runtime.GOMAXPROCS(runtime.NumCPU())

	Logo()
	Spec()

	network.PrivateIP()
	network.PublicIP()

}
