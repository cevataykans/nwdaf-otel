package main

import "fmt"

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
}
