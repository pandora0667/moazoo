package main

import (
	network "./network"
	. "./start"
	"runtime"
)

func main()  {

	runtime.GOMAXPROCS(runtime.NumCPU())

	Logo()
	Spec()

	network.PrivateIP()
	network.PublicIP()


}
