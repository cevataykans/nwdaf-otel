package nrf

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"time"
)

type NRFClient struct {
	id uuid.UUID

	httpClient *http.Client
}

func NewNFClient() *NRFClient {
	nrfClientTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	nrfClient := http.Client{
		Timeout:   time.Second * 5,
		Transport: nrfClientTransport,
	}
	return &NRFClient{
		id:         uuid.New(),
		httpClient: &nrfClient,
	}
}

func (c *NRFClient) testConnection() error {
	res, err := c.httpClient.Get("http://nrf:29510/nnrf-nfm/v1/nf-instances?limit=10")
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_, _ = io.Copy(io.Discard, Body)
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return fmt.Errorf("NRF server returned status %d", res.StatusCode)
	}
	log.Println("Successfully connected NRF client and ready to register.")
	return nil
}

func (c *NRFClient) StartNFRegistration(stop chan struct{}) {
	go c.registerNWDAF(stop)
}

func (c *NRFClient) createRegistrationMsg() NFInstance {
	return NFInstance{
		NfInstanceId: c.id.String(),
		NfType:       "NWDAF",
		NfStatus:     "REGISTERED",
		HeartBeat:    60,
		NfServices: []NFService{
			{
				ServiceInstanceId: "nwdaf-1",
				ServiceName:       "nwdaf-analytics-info",
				Versions: []Version{
					{
						ApiFullVersion:  "v1",
						ApiVersionInUri: "1.0.0",
					},
				},
				Scheme:          "http",
				NfServiceStatus: "REGISTERED",
				IpEndPoints: []IPEndpoint{
					{
						Ipv4Address: "nwdaf-analytics-info",
						Port:        80,
					},
				},
			},
		},
	}
}

func (c *NRFClient) registerNWDAF(stop chan struct{}) {
	select {
	case <-stop:
		return
	default:
		err := c.testConnection()
		if err != nil {
			log.Printf("NRF is not available: %v", err)
			time.Sleep(5 * time.Second)
			go c.registerNWDAF(stop)
			return
		}

		registrationMsg := c.createRegistrationMsg()
		nrfRegistrationURL := fmt.Sprintf("http://nrf:29510/nnrf-nfm/v1/nf-instances/%s", registrationMsg.NfInstanceId)
		registrationBody, err := json.Marshal(registrationMsg)
		if err != nil {
			log.Printf("failed to marshal NFInstance: %v", err)
			return
		}

		req, err := http.NewRequest(http.MethodPut, nrfRegistrationURL, bytes.NewBuffer(registrationBody))
		if err != nil {
			log.Printf("cannot create registration request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.httpClient.Do(req)
		if err != nil {
			log.Printf("registration request failed: %w", err)
			return
		}
		defer resp.Body.Close()

		switch resp.StatusCode {
		case http.StatusOK:
			log.Println("NWDAF profile successfully updated")
		case http.StatusCreated:
			log.Println("NWDAF profile successfully created")
		default:
			log.Printf("registration request returned status %d, retring after 5 sec", resp.StatusCode)
			time.Sleep(5 * time.Second)
			go c.registerNWDAF(stop)
			return
		}

		// TODO: start pinging
	}
}

func (c *NRFClient) startHeartbeat(stop chan struct{}) {

}
