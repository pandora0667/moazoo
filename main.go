package main

import (
	"runtime"
	. "./start"
)

func main()  {

	runtime.GOMAXPROCS(runtime.NumCPU())

	Logo()
	Spec()

}
