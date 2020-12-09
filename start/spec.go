package start

import (
	"log"
	"runtime"
)

func Spec()  {


	log.Println("OS Detected: ", runtime.GOOS)
	log.Println("Server CPU Core : ",  runtime.NumCPU())

}