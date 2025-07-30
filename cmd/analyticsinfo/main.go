package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"nwdaf-otel/clients/prometheus"
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

	promClient, err := prometheus.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	go queryCPUMetrics(promClient)

	err = <-errChan
	if err != nil {
		log.Printf("server shutdown err: %v\n", err)
	}
	log.Println("Application Finished!")
}

func queryCPUMetrics(promClient *prometheus.Client) {

	curSeconds := time.Now().UTC().Unix()
	remainingSeconds := 60 - (curSeconds % 60)
	nextMin := curSeconds + remainingSeconds
	log.Printf("Sleeping for %v seconds\nCalculated query time for next min is: %v\n", remainingSeconds, time.Unix(nextMin, 0))
	time.Sleep(time.Duration(remainingSeconds) * time.Second)

	// print for one hour metrics
	for i := 0; i < 60; i++ {
		old := time.Now()
		log.Printf("Current Time: %v\n", old)
		err := promClient.QueryCPUTotalSeconds(
			time.Unix(nextMin-60, 1),
			time.Unix(nextMin, 0),
			time.Minute)
		if err != nil {
			log.Printf("Error querying prom for CPU Total Seconds: %v\nExiting loop.\n", err)
			break
		}
		nextMin += 60

		cur := time.Now()
		log.Printf("Sleeping amount: %v\n", cur.Sub(old))
		time.Sleep(60 - cur.Sub(old))
	}
	log.Println("Loop Complete!")
}
