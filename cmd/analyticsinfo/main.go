package main

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"nwdaf-otel/clients/prometheus"
	"nwdaf-otel/repository"
	"nwdaf-otel/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO: add logger & metrics
func main() {
	// TODO: parse flags -> path for config

	shutdownChn := make(chan struct{})
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		close(shutdownChn)
	}()

	srv := server.NewAnalyticsInfoServer()
	srv.Setup()
	errChan := srv.Start(shutdownChn)
	log.Println("Server started")

	nrfClientTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	nrfClient := http.Client{
		Timeout:   time.Second * 5,
		Transport: nrfClientTransport,
	}
	res, err := nrfClient.Get("http://nrf:29510/nnrf-nfm/v1/nf-instances?limit=10")
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
	log.Println(res.StatusCode)
	log.Println("Successfully connected NRF client and ready to register.")

	log.Println("Creating DB")
	repo, err := repository.NewSQLiteRepo()
	if err != nil {
		log.Fatal(err)
	}

	err = repo.Setup()
	if err != nil {
		log.Fatal(err)
	}

	count, err := repo.Size()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Successfully initialized DB with %d rows", count)

	promClient, err := prometheus.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	go queryResources(promClient, repo, shutdownChn)

	err = <-errChan
	if err != nil {
		log.Printf("server listen err: %v\n", err)
	}
	log.Println("Application Finished!")
}

func queryResources(client *prometheus.Client, repo repository.Repository, shutdownChn chan struct{}) {
	curSeconds := time.Now().UTC()
	startDelay := time.Second * 30

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

	time.Sleep(startDelay)
	for {
		select {
		case <-shutdownChn:
			return
		default:
		}

		old := time.Now()
		statistics := make([]prometheus.MetricResults, 0)
		for _, service := range services {

			start, end := curSeconds.Add(-1*time.Second), curSeconds
			metrics, err := client.QueryMetrics(service, start, end, time.Second)
			if err != nil {
				log.Printf("Error querying metrics %v for service %v\n", err, service)
			}
			avgDuration, err := client.QueryTraces(service, start, end)
			if err != nil {
				log.Printf("Error querying traces: %v for service %v\n", err, service)
			}

			curMetrics := prometheus.MetricResults{
				Service:                     service,
				Timestamp:                   curSeconds.Unix(),
				CpuTotalSeconds:             metrics.CpuTotalSeconds,
				MemoryTotalBytes:            metrics.MemoryTotalBytes,
				NetworkReceiveBytesTotal:    metrics.NetworkReceiveBytesTotal,
				NetworkTransmitBytesTotal:   metrics.NetworkTransmitBytesTotal,
				NetworkReceivePacketsTotal:  metrics.NetworkReceivePacketsTotal,
				NetworkTransmitPacketsTotal: metrics.NetworkTransmitPacketsTotal,
				AvgTraceDuration:            avgDuration,
			}
			statistics = append(statistics, curMetrics)
		}

		err := repo.InsertBatch(statistics)
		if err != nil {
			log.Printf("Error inserting batch: %v\n", err)
		}

		curSeconds = curSeconds.Add(time.Second)
		cur := time.Now()
		diff := cur.Sub(old)
		if diff < time.Second {
			time.Sleep(time.Second - diff)
		}
	}
}
