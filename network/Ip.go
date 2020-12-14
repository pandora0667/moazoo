package neteork

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
)

func PrivateIP()  {

	address, err := net.InterfaceAddrs()
	if err != nil {
		//_, _ = os.Stderr.WriteString("Check Network Connection: " + err.Error() + "\n")
		log.Println("Check Network Connection: " + err.Error())
		os.Exit(1)
	}
	for _, a := range address {
		if local, ok := a.(*net.IPNet); ok && !local.IP.IsLoopback() {
			if local.IP.To4() != nil {
				log.Printf("Server Local IP Address : " + local.IP.String())
				//_, _ = os.Stdout.WriteString(local.IP.String())
			}
		}
	}

}

func PublicIP()  {

	url := "https://api.ipify.org?format=text"
	// https://www.ipify.org
	// http://myexternalip.com
	// http://api.ident.me
	// http://whatismyipaddress.com/api

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Check Internet Connection " + err.Error())
	}
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Printf("Server Public IP Address : :%s", ip)

	defer resp.Body.Close()

}

