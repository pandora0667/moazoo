package main

import (
	"log"
	"runtime"
)

func main()  {

	log.Println("Multi protocol IoT server platform MOAZOO Version 0.1 ALPHA")

	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Server CPU Core : ",  runtime.NumCPU())


}
