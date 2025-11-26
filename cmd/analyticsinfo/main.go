package main

import (
	"log"
	"nwdaf-otel/clients/nrf"
	"nwdaf-otel/clients/prometheus"
	"nwdaf-otel/repository"
	"nwdaf-otel/server"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// TODO: add logger library & export own metrics?
func main() {
	// TODO: have some configuration options, parse flags -> path for config

	shutdownChn := make(chan struct{})
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan
		close(shutdownChn)
	}()

	promClient, err := prometheus.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewAnalyticsInfoServer(promClient)
	srv.Setup()
	errChan := srv.Start(shutdownChn)
	log.Println("Analytics Server started")

	nrfClient := nrf.NewNFClient()
	nrfClient.StartNFRegistration(shutdownChn)

	//log.Println("Creating DB")
	//repo, err := repository.NewSQLiteRepo()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//err = repo.Setup()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//count, err := repo.Size()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("Successfully initialized DB with %d rows", count)
	go queryUDM(promClient, shutdownChn)

	//go queryResources(promClient, repo, shutdownChn)

	err = <-errChan
	if err != nil {
		log.Printf("server listen err: %v\n", err)
	}
	log.Println("Application Finished!")
}

func queryUDM(promClient prometheus.Client, shutdownChn chan struct{}) {

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

	log.Printf("nwdaf analytics will start recording from: %v seconds (unix) after delay: %v", curSeconds.Unix())
	time.Sleep(startDelay)
	for {
		select {
		case <-shutdownChn:
			log.Printf("nwdaf analytics stopped at: %v seconds (unix)", curSeconds.Unix())
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
