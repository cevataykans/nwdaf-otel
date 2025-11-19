package server

import (
	"github.com/omec-project/aper"
	"github.com/omec-project/ngap/ngapType"
	"sync"
	"testing"
)

func TestDecoderConcurrency(t *testing.T) {
	// Example NGAP PDU: NGSetupRequest (encoded in ASN.1 PER)
	var ngSetupRequest = []byte{
		0x00, 0x10, 0x00, 0x05, 0x00, 0x00, 0x1b, 0x00,
		0x00, 0x00, 0x00, 0x01, 0x00, 0x20, 0x10, 0x01,
		0x00, 0x01, 0x02, 0x40, 0x00, 0x40, 0x40, 0x01,
		0x02, 0x00, 0x02, 0x40, 0x02, 0x01, 0x01, 0x40,
	}

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			pdu := &ngapType.NGAPPDU{}
			_ = aper.UnmarshalWithParams(ngSetupRequest, pdu, "valueExt,valueLB:0,valueUB:2")
		}()
	}
	wg.Wait()
}
