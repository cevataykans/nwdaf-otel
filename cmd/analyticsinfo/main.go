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

	go queryTraces(promClient)
	//go queryMetrics(promClient)

	err = <-errChan
	if err != nil {
		log.Printf("server shutdown err: %v\n", err)
	}
	log.Println("Application Finished!")
}

func queryTraces(elasticClient *prometheus.Client) {
	curSeconds := time.Now().UTC().Unix()
	remainingSeconds := 60 - (curSeconds % 60)
	nextMin := curSeconds + remainingSeconds
	log.Printf("Sleeping for %v seconds\nCalculated query time for next min is: %v\n", remainingSeconds, time.Unix(nextMin, 0))
	time.Sleep(time.Duration(remainingSeconds) * time.Second)

	// services is a list of container names used for filtering queried metrics.
	services := []string{
		"bessd",
		"amf",
		"ausf",
		"nrf",
		"nssf",
		"pcf",
		"smf",
		"udm",
		"udr",
		"webui",
		"simapp",
	}

	// print for one hour metrics
	for i := 0; i < 60; i++ {

		if i == 0 {
			log.Println("Verifying Query Behavior - One Entry for WEBUI Must be Printed")
			for _, service := range services {
				log.Printf("Tracing service: %v\n", service)
				err := elasticClient.QueryTraces(
					service,
					time.Unix(1754041395, 0),
					time.Unix(1754041396, 0),
				)
				if err != nil {
					log.Printf("Error querying traces %v: %v\n", service, err)
				}
			}
			log.Println("!!!!!! Check behavior !!!!!!")
			continue
		}

		old := time.Now()
		log.Printf("Current Time: %v\n", old)
		for _, service := range services {
			log.Printf("Tracing service: %v\n", service)
			err := elasticClient.QueryTraces(
				service,
				time.Unix(nextMin-60, 1),
				time.Unix(nextMin, 0),
			)
			if err != nil {
				log.Printf("Error querying traces %v: %v\n", service, err)
			}
		}
		nextMin += 60
		cur := time.Now()
		log.Printf("Query Time: %v, sleep time: %v\n", cur.Sub(old), time.Minute-cur.Sub(old))
		time.Sleep(time.Minute - cur.Sub(old))
	}
	log.Println("Loop Complete!")
}

func queryMetrics(promClient *prometheus.Client) {

	curSeconds := time.Now().UTC().Unix()
	remainingSeconds := 60 - (curSeconds % 60)
	nextMin := curSeconds + remainingSeconds
	log.Printf("Sleeping for %v seconds\nCalculated query time for next min is: %v\n", remainingSeconds, time.Unix(nextMin, 0))
	time.Sleep(time.Duration(remainingSeconds) * time.Second)

	// services is a list of container names used for filtering queried metrics.
	services := []string{
		"bessd",
		"amf",
		"ausf",
		"nrf",
		"nssf",
		"pcf",
		"smf",
		"udm",
		"udr",
	}

	// print for one hour metrics
	for i := 0; i < 60; i++ {
		old := time.Now()
		log.Printf("Current Time: %v\n", old)
		for _, service := range services {
			err := promClient.QueryMetrics(
				service,
				time.Unix(nextMin-60, 1),
				time.Unix(nextMin, 0),
				time.Minute)
			if err != nil {
				log.Printf("Error querying metrics %v: %v\n", service, err)
			}
		}
		nextMin += 60
		cur := time.Now()
		log.Printf("Query Time: %v, sleep time: %v\n", cur.Sub(old), time.Minute-cur.Sub(old))
		time.Sleep(time.Minute - cur.Sub(old))
	}
	log.Println("Loop Complete!")
}
