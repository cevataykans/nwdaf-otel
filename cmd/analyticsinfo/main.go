package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"nwdaf-otel/server"
	"time"
)

// TODO: add logger & metrics
func main() {
	// TODO: parse flags -> path for config

	log.Println("Hello World")

	srv := server.NewAnalyticsInfoServer()
	srv.Setup()
	errChan := srv.Start()
	log.Println("Server started")

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
	defer func(Body io.ReadCloser) {
		_, _ = io.Copy(io.Discard, Body)
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)
	fmt.Println(res.StatusCode)
	fmt.Println("Successfully connected NRF client and ready to register.")

	err = <-errChan
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Application Finished!")
}
