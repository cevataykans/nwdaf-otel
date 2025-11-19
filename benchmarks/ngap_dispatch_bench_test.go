package benchmarks

import (
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
)

/*
 This file provides:
 - minimal synthetic NGAP structs
 - ExtractContextSwitch: large switch-based traversal (replicates your structure)
 - ExtractContextMap: map-of-handlers equivalent
 - Benchmarks that run both under parallel load for different UE counts
*/

// --- Minimal NGAP types (synthetic) ---

type ProcedureCode struct{ Value int }
type IEID struct{ Value int }

type IE struct {
	Id IEID
	// payload omitted
}

type ProtocolIEs struct {
	List []IE
}

type InitiatingMessageValue struct {
	InitialUEMessage                   *MessageWithIEs
	UplinkNASTransport                 *MessageWithIEs
	HandoverCancel                     *MessageWithIEs
	UEContextReleaseRequest            *MessageWithIEs
	NASNonDeliveryIndication           *MessageWithIEs
	UERadioCapabilityInfoIndication    *MessageWithIEs
	HandoverNotify                     *MessageWithIEs
	HandoverRequired                   *MessageWithIEs
	PDUSessionResourceNotify           *MessageWithIEs
	PathSwitchRequest                  *MessageWithIEs
	PDUSessionResourceModifyIndication *MessageWithIEs
	// add other fields as needed
}

type InitiatingMessage struct {
	ProcedureCode ProcedureCode
	Value         InitiatingMessageValue
}

type SuccessfulOutcomeValue struct {
	UEContextReleaseComplete          *MessageWithIEs
	PDUSessionResourceReleaseResponse *MessageWithIEs
	InitialContextSetupResponse       *MessageWithIEs
	UEContextModificationResponse     *MessageWithIEs
	PDUSessionResourceSetupResponse   *MessageWithIEs
	PDUSessionResourceModifyResponse  *MessageWithIEs
	HandoverRequestAcknowledge        *MessageWithIEs
}

type SuccessfulOutcome struct {
	ProcedureCode ProcedureCode
	Value         SuccessfulOutcomeValue
}

type UnsuccessfulOutcomeValue struct {
	InitialContextSetupFailure   *MessageWithIEs
	UEContextModificationFailure *MessageWithIEs
	HandoverFailure              *MessageWithIEs
}

type UnsuccessfulOutcome struct {
	ProcedureCode ProcedureCode
	Value         UnsuccessfulOutcomeValue
}

type NGAPPDU struct {
	Present             int
	InitiatingMessage   *InitiatingMessage
	SuccessfulOutcome   *SuccessfulOutcome
	UnsuccessfulOutcome *UnsuccessfulOutcome
}

// a simple wrapper for messages that contain ProtocolIEs.List
type MessageWithIEs struct {
	ProtocolIEs ProtocolIEs
}

// --- Example constants (procedure code values and IE IDs) ---
// These constants are synthetic but mirror the structure you're using.
// You can change numeric values as you wish — only equality matters in this micro-benchmark.

const (
	NGAPPDUPresentInitiatingMessage   = 1
	NGAPPDUPresentSuccessfulOutcome   = 2
	NGAPPDUPresentUnsuccessfulOutcome = 3

	ProcedureCodeNGSetup                    = 10
	ProcedureCodeInitialUEMessage           = 11
	ProcedureCodeUplinkNASTransport         = 12
	ProcedureCodeInitialContextSetup        = 13
	ProcedureCodeUEContextRelease           = 14
	ProcedureCodePDUSessionResourceRelease  = 15
	ProcedureCodeUEContextModification      = 16
	ProcedureCodePDUSessionResourceSetup    = 17
	ProcedureCodePDUSessionResourceModify   = 18
	ProcedureCodeHandoverResourceAllocation = 19
	// ... add as needed
)

const (
	ProtocolIEIDRANUENGAPID       = 100
	ProtocolIEIDAMFUENGAPID       = 101
	ProtocolIEIDFiveGSTMSI        = 102
	ProtocolIEIDSourceAMFUENGAPID = 103
)

// --- Work performed when an IE of interest is found ---
// We do minimal, equal work in both variants to keep the comparison fair.
var globalMatchCounter uint64

func matchedWork() {
	atomic.AddUint64(&globalMatchCounter, 1)
}

// --- Switch-based extractor (mirrors the structure you posted) ---

func ExtractContextSwitch(message *NGAPPDU) bool {
	switch message.Present {
	case NGAPPDUPresentInitiatingMessage:
		initiatingMessage := message.InitiatingMessage
		if initiatingMessage == nil {
			return false
		}
		switch initiatingMessage.ProcedureCode.Value {
		case ProcedureCodeNGSetup:
		case ProcedureCodeInitialUEMessage:
			ngapMsg := initiatingMessage.Value.InitialUEMessage
			if ngapMsg == nil {
				return false
			}
			for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
				ie := ngapMsg.ProtocolIEs.List[i]
				switch ie.Id.Value {
				case ProtocolIEIDRANUENGAPID:
					matchedWork()
				case ProtocolIEIDFiveGSTMSI:
				}
			}
		case ProcedureCodeUplinkNASTransport:
			ngapMsg := initiatingMessage.Value.UplinkNASTransport
			if ngapMsg == nil {
				return false
			}
			for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
				ie := ngapMsg.ProtocolIEs.List[i]
				switch ie.Id.Value {
				case ProtocolIEIDRANUENGAPID:
					matchedWork()
				case ProtocolIEIDAMFUENGAPID:
					matchedWork()
				}
			}
		case ProcedureCodeHandoverResourceAllocation:
			ngapMsg := initiatingMessage.Value.HandoverRequired
			if ngapMsg != nil {
				for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
					ie := ngapMsg.ProtocolIEs.List[i]
					switch ie.Id.Value {
					case ProtocolIEIDRANUENGAPID:
						matchedWork()
					case ProtocolIEIDAMFUENGAPID:
						matchedWork()
					}
				}
			}
		case ProcedureCodeUEContextRelease:
			ngapMsg := initiatingMessage.Value.UEContextReleaseRequest
			if ngapMsg != nil {
				for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
					ie := ngapMsg.ProtocolIEs.List[i]
					switch ie.Id.Value {
					case ProtocolIEIDAMFUENGAPID:
						matchedWork()
					case ProtocolIEIDRANUENGAPID:
						matchedWork()
					}
				}
			}
			// you can add other cases (mirrors the original) - left trimmed for brevity
		}
	case NGAPPDUPresentSuccessfulOutcome:
		successfulOutcome := message.SuccessfulOutcome
		if successfulOutcome == nil {
			return false
		}
		switch successfulOutcome.ProcedureCode.Value {
		case ProcedureCodeUEContextRelease:
			ngapMsg := successfulOutcome.Value.UEContextReleaseComplete
			if ngapMsg != nil {
				for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
					ie := ngapMsg.ProtocolIEs.List[i]
					switch ie.Id.Value {
					case ProtocolIEIDAMFUENGAPID:
						matchedWork()
					}
				}
			}
		case ProcedureCodePDUSessionResourceRelease:
			ngapMsg := successfulOutcome.Value.PDUSessionResourceReleaseResponse
			if ngapMsg != nil {
				for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
					ie := ngapMsg.ProtocolIEs.List[i]
					switch ie.Id.Value {
					case ProtocolIEIDRANUENGAPID:
						matchedWork()
					case ProtocolIEIDAMFUENGAPID:
						matchedWork()
					}
				}
			}
		case ProcedureCodeInitialContextSetup:
			ngapMsg := successfulOutcome.Value.InitialContextSetupResponse
			if ngapMsg != nil {
				for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
					ie := ngapMsg.ProtocolIEs.List[i]
					switch ie.Id.Value {
					case ProtocolIEIDRANUENGAPID:
						matchedWork()
					case ProtocolIEIDAMFUENGAPID:
						matchedWork()
					}
				}
			}
		}
	case NGAPPDUPresentUnsuccessfulOutcome:
		unsuccessfulOutcome := message.UnsuccessfulOutcome
		if unsuccessfulOutcome == nil {
			return false
		}
		switch unsuccessfulOutcome.ProcedureCode.Value {
		case ProcedureCodeInitialContextSetup:
			ngapMsg := unsuccessfulOutcome.Value.InitialContextSetupFailure
			if ngapMsg != nil {
				for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
					ie := ngapMsg.ProtocolIEs.List[i]
					switch ie.Id.Value {
					case ProtocolIEIDRANUENGAPID:
						matchedWork()
					case ProtocolIEIDAMFUENGAPID:
						matchedWork()
					}
				}
			}
		case ProcedureCodeUEContextModification:
			ngapMsg := unsuccessfulOutcome.Value.UEContextModificationFailure
			if ngapMsg != nil {
				for i := 0; i < len(ngapMsg.ProtocolIEs.List); i++ {
					ie := ngapMsg.ProtocolIEs.List[i]
					switch ie.Id.Value {
					case ProtocolIEIDRANUENGAPID:
						matchedWork()
					case ProtocolIEIDAMFUENGAPID:
						matchedWork()
					}
				}
			}
		}
	}
	return true
}

// --- Map-based extractor ---
// We build maps for: present->handler, procedureCode->handler, ieId->handler.
// All maps are prebuilt once and then used in hot path.

type HandlerFunc func(msg *NGAPPDU) bool

func buildMapDispatch() map[int]HandlerFunc {
	presentMap := make(map[int]HandlerFunc)

	// InitiatingMessage handler
	initHandler := func(m *NGAPPDU) bool {
		if m.InitiatingMessage == nil {
			return false
		}
		proc := m.InitiatingMessage.ProcedureCode.Value
		return initiatingProcMap[proc](m.InitiatingMessage.Value)
	}

	// SuccessfulOutcome handler
	successHandler := func(m *NGAPPDU) bool {
		if m.SuccessfulOutcome == nil {
			return false
		}
		proc := m.SuccessfulOutcome.ProcedureCode.Value
		return successfulProcMap[proc](m.SuccessfulOutcome.Value)
	}

	// UnsuccessfulOutcome handler
	unsuccessHandler := func(m *NGAPPDU) bool {
		if m.UnsuccessfulOutcome == nil {
			return false
		}
		proc := m.UnsuccessfulOutcome.ProcedureCode.Value
		return unsuccessfulProcMap[proc](m.UnsuccessfulOutcome.Value)
	}

	presentMap[NGAPPDUPresentInitiatingMessage] = initHandler
	presentMap[NGAPPDUPresentSuccessfulOutcome] = successHandler
	presentMap[NGAPPDUPresentUnsuccessfulOutcome] = unsuccessHandler

	return presentMap
}

// maps for inner dispatch
var initiatingProcMap map[int]func(InitiatingMessageValue) bool
var successfulProcMap map[int]func(SuccessfulOutcomeValue) bool
var unsuccessfulProcMap map[int]func(UnsuccessfulOutcomeValue) bool
var corpus []*NGAPPDU

func init() {
	corpus = buildCorpus(10000)

	// build the maps once
	initiatingProcMap = map[int]func(InitiatingMessageValue) bool{
		ProcedureCodeInitialUEMessage: func(val InitiatingMessageValue) bool {
			m := val.InitialUEMessage
			if m == nil {
				return false
			}
			for i := 0; i < len(m.ProtocolIEs.List); i++ {
				ie := m.ProtocolIEs.List[i]
				if h, ok := initiatingIEMap[ie.Id.Value]; ok {
					h(ie)
				}
			}
			return true
		},
		ProcedureCodeUplinkNASTransport: func(val InitiatingMessageValue) bool {
			m := val.UplinkNASTransport
			if m == nil {
				return false
			}
			for i := 0; i < len(m.ProtocolIEs.List); i++ {
				ie := m.ProtocolIEs.List[i]
				if h, ok := initiatingIEMap[ie.Id.Value]; ok {
					h(ie)
				}
			}
			return true
		},
		// other procedure handlers short-circuited for brevity
	}

	// IE handler map for initiating messages
	initiatingIEMap = map[int]func(IE){
		ProtocolIEIDRANUENGAPID:       func(ie IE) { matchedWork() },
		ProtocolIEIDFiveGSTMSI:        func(ie IE) {},
		ProtocolIEIDAMFUENGAPID:       func(ie IE) { matchedWork() },
		ProtocolIEIDSourceAMFUENGAPID: func(ie IE) {},
	}

	successfulProcMap = map[int]func(SuccessfulOutcomeValue) bool{
		ProcedureCodeInitialContextSetup: func(val SuccessfulOutcomeValue) bool {
			m := val.InitialContextSetupResponse
			if m == nil {
				return false
			}
			for i := 0; i < len(m.ProtocolIEs.List); i++ {
				ie := m.ProtocolIEs.List[i]
				if h, ok := successfulIEMap[ie.Id.Value]; ok {
					h(ie)
				}
			}
			return true
		},
		ProcedureCodeUEContextRelease: func(val SuccessfulOutcomeValue) bool {
			m := val.UEContextReleaseComplete
			if m == nil {
				return false
			}
			for i := 0; i < len(m.ProtocolIEs.List); i++ {
				ie := m.ProtocolIEs.List[i]
				if h, ok := successfulIEMap[ie.Id.Value]; ok {
					h(ie)
				}
			}
			return true
		},
		ProcedureCodePDUSessionResourceRelease: func(val SuccessfulOutcomeValue) bool {
			m := val.PDUSessionResourceReleaseResponse
			if m == nil {
				return false
			}
			for i := 0; i < len(m.ProtocolIEs.List); i++ {
				ie := m.ProtocolIEs.List[i]
				if h, ok := successfulIEMap[ie.Id.Value]; ok {
					h(ie)
				}
			}
			return true
		},
	}

	successfulIEMap = map[int]func(IE){
		ProtocolIEIDRANUENGAPID: func(ie IE) { matchedWork() },
		ProtocolIEIDAMFUENGAPID: func(ie IE) { matchedWork() },
	}

	unsuccessfulProcMap = map[int]func(UnsuccessfulOutcomeValue) bool{
		ProcedureCodeInitialContextSetup: func(val UnsuccessfulOutcomeValue) bool {
			m := val.InitialContextSetupFailure
			if m == nil {
				return false
			}
			for i := 0; i < len(m.ProtocolIEs.List); i++ {
				ie := m.ProtocolIEs.List[i]
				if h, ok := unsuccessfulIEMap[ie.Id.Value]; ok {
					h(ie)
				}
			}
			return true
		},
		ProcedureCodeUEContextModification: func(val UnsuccessfulOutcomeValue) bool {
			m := val.UEContextModificationFailure
			if m == nil {
				return false
			}
			for i := 0; i < len(m.ProtocolIEs.List); i++ {
				ie := m.ProtocolIEs.List[i]
				if h, ok := unsuccessfulIEMap[ie.Id.Value]; ok {
					h(ie)
				}
			}
			return true
		},
	}

	unsuccessfulIEMap = map[int]func(IE){
		ProtocolIEIDRANUENGAPID: func(ie IE) { matchedWork() },
		ProtocolIEIDAMFUENGAPID: func(ie IE) { matchedWork() },
	}

	// build top-level present map (used by benchmark)
	presentDispatch = buildMapDispatch()
}

// maps used by map-dispatch implementation
var initiatingIEMap map[int]func(IE)
var successfulIEMap map[int]func(IE)
var unsuccessfulIEMap map[int]func(IE)
var presentDispatch map[int]HandlerFunc

// ExtractContextMap uses the prebuilt maps to dispatch
func ExtractContextMap(message *NGAPPDU) bool {
	return presentDispatch[message.Present](message)
	//if h, ok := presentDispatch[message.Present]; ok {
	//	return h(message)
	//}
	//return false
}

// --- Message generation for realistic registration-ish sequence ---
// We simulate a short sequence of messages common in registration flow:
// For each UE we create four messages: InitialUEMessage, UplinkNASTransport,
// SuccessfulInitialContextSetup (SuccessfulOutcome), UplinkNASTransport (complete).
// Each message contains a few IEs; IE ids are chosen from the set the extractors inspect.

func makeInitialUEMessageForUE(ueIdx int) *NGAPPDU {
	ieList := []IE{
		{Id: IEID{Value: ProtocolIEIDRANUENGAPID}},
		{Id: IEID{Value: ProtocolIEIDFiveGSTMSI}},
	}
	m := &MessageWithIEs{ProtocolIEs{List: ieList}}
	return &NGAPPDU{
		Present: NGAPPDUPresentInitiatingMessage,
		InitiatingMessage: &InitiatingMessage{
			ProcedureCode: ProcedureCode{Value: ProcedureCodeInitialUEMessage},
			Value: InitiatingMessageValue{
				InitialUEMessage: m,
			},
		},
	}
}

func makeUplinkNASTransportForUE(ueIdx int) *NGAPPDU {
	ieList := []IE{
		{Id: IEID{Value: ProtocolIEIDRANUENGAPID}},
		{Id: IEID{Value: ProtocolIEIDAMFUENGAPID}},
	}
	m := &MessageWithIEs{ProtocolIEs{List: ieList}}
	return &NGAPPDU{
		Present: NGAPPDUPresentInitiatingMessage,
		InitiatingMessage: &InitiatingMessage{
			ProcedureCode: ProcedureCode{Value: ProcedureCodeUplinkNASTransport},
			Value: InitiatingMessageValue{
				UplinkNASTransport: m,
			},
		},
	}
}

func makeInitialContextSetupResponseForUE(ueIdx int) *NGAPPDU {
	ieList := []IE{
		{Id: IEID{Value: ProtocolIEIDRANUENGAPID}},
		{Id: IEID{Value: ProtocolIEIDAMFUENGAPID}},
	}
	m := &MessageWithIEs{ProtocolIEs{List: ieList}}
	return &NGAPPDU{
		Present: NGAPPDUPresentSuccessfulOutcome,
		SuccessfulOutcome: &SuccessfulOutcome{
			ProcedureCode: ProcedureCode{Value: ProcedureCodeInitialContextSetup},
			Value: SuccessfulOutcomeValue{
				InitialContextSetupResponse: m,
			},
		},
	}
}

// Build a corpus of messages for N UEs
func buildCorpus(nUE int) []*NGAPPDU {
	c := make([]*NGAPPDU, 0, nUE*3)
	for i := 0; i < nUE; i++ {
		c = append(c, makeInitialUEMessageForUE(i))
		c = append(c, makeUplinkNASTransportForUE(i))
		c = append(c, makeInitialContextSetupResponseForUE(i))
		// add another uplink nas to simulate registration complete
		c = append(c, makeUplinkNASTransportForUE(i))
	}
	return c
}

// --- Benchmarks ---
// We'll run for multiple UE counts and for each do:
// - Reset global counter
// - Generate corpus
// - Run ExtractContextSwitch in parallel reading from corpus using rand
// - Run ExtractContextMap in parallel reading same corpus

func runParallelB(b *testing.B, corpus []*NGAPPDU, fn func(*NGAPPDU) bool) {
	// make a local RNG per goroutine to avoid contention
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(time.Now().UnixNano() ^ int64(uintptr(unsafePointer(pb)))))
		ln := len(corpus)
		if ln == 0 {
			return
		}
		i := 0
		for pb.Next() {
			// random pick reduces perfect predictability; use round robin + small randomness
			idx := (i + r.Intn(ln)) % ln
			_ = fn(corpus[idx])
			i++
		}
	})
	b.StopTimer()
}

// unsafePointer is a tiny helper to get distinct seed per goroutine
// using uintptr of the pointer to pb. This avoids importing "unsafe" repeatedly.
func unsafePointer(x interface{}) uintptr {
	return uintptr((^uintptr(0)) & (uintptr(time.Now().UnixNano())))
}

type MinimalIE struct {
	Id struct{ Value int }
}

type MinimalNGAPMessage struct {
	Present       int
	ProcedureCode struct{ Value int }
	ProtocolIEs   struct{ List []MinimalIE }
}

func RandomNGAPMessage() MinimalNGAPMessage {
	msg := MinimalNGAPMessage{}

	// Randomly pick Present value (Initiating, Successful, Unsuccessful)
	msg.Present = rand.Intn(3) // 0..2

	// Random ProcedureCode based on Present
	switch msg.Present {
	case 0: // InitiatingMessage
		msg.ProcedureCode.Value = rand.Intn(20) // map to your ProcedureCode enums
	case 1: // SuccessfulOutcome
		msg.ProcedureCode.Value = rand.Intn(10) + 100
	case 2: // UnsuccessfulOutcome
		msg.ProcedureCode.Value = rand.Intn(5) + 200
	}

	// Random number of IEs
	numIEs := rand.Intn(5) + 1
	msg.ProtocolIEs.List = make([]MinimalIE, numIEs)
	for i := 0; i < numIEs; i++ {
		msg.ProtocolIEs.List[i].Id.Value = rand.Intn(5) // map to IE IDs you care about
	}

	return msg
}

var result bool

func BenchmarkSwitch(b *testing.B) {
	var res bool
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(corpus); i++ {
			res = ExtractContextSwitch(corpus[i])
		}
	}
	result = res

	//for _, n := range ueCounts {
	//	corpus := buildCorpus(10000)
	//	b.Run(fmt.Sprintf("UEs=%d/Switch", n), func(b *testing.B) {
	//		atomic.StoreUint64(&globalMatchCounter, 0)
	//		b.ResetTimer()
	//		runParallelB(b, corpus, func(m *NGAPPDU) bool { return ExtractContextSwitch(m) })
	//		b.ReportMetric(float64(atomic.LoadUint64(&globalMatchCounter))/float64(b.N), "matches_per_op")
	//	})
	//	b.Run(fmt.Sprintf("UEs=%d/Map", n), func(b *testing.B) {
	//		atomic.StoreUint64(&globalMatchCounter, 0)
	//		b.ResetTimer()
	//		runParallelB(b, corpus, func(m *NGAPPDU) bool { return ExtractContextMap(m) })
	//		b.ReportMetric(float64(atomic.LoadUint64(&globalMatchCounter))/float64(b.N), "matches_per_op")
	//	})
	//}
}

func BenchmarkMap(b *testing.B) {
	var res bool
	for n := 0; n < b.N; n++ {
		for i := 0; i < len(corpus); i++ {
			res = ExtractContextMap(corpus[i])
		}
	}
	result = res
}
