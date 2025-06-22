package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

// TODO: add logger & metrics
func main() {
	// TODO: parse flags -> path for config

	//nwdafService := nwdaf.New()
	//nwdafService.SetupServices()
	//nwdafErrs := nwdafService.Start()
	//
	//for {
	//	// Main loop where errors are handled from clients & servers
	//	err := <-nwdafErrs
	//	if err != nil {
	//		// NWDAF unexpectedly shutdown
	//		log.Fatal(err)
	//	}
	//
	//	// TODO: also handle shutdown signal for graceful termination
	//}
	//// TODO: Setup NRF
	fmt.Println("Hello World")

	nrfClientTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	nrfClient := http.Client{
		Timeout:   time.Second * 5,
		Transport: nrfClientTransport,
	}
	res, err := nrfClient.Get("https://nrf:29510/nnrf-nfm/v1/nf-instances?limit=10")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	fmt.Println("Successfully connected NRF client and ready to register.")
}
