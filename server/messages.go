package nwdaf

import "net/url"

type Event string

const (
	EventSliceLoadLevel Event = "SLICE_LOAD_LEVEL"
	// TODO: add remaining events
)

type Features struct {
	EneNA          bool
	Aggregation    bool
	AnaCtxTransfer bool
}

type Subscription struct {
}

// Sent when an event is triggered
type Notification struct {
}

type NnwdafEventsSubscriptionRequest struct {
	notificationURI    url.URL
	eventSubscriptions []EventSubscription
	evtReq             []EventReportingInformation
	prevSub            Subscription // depends on AnaCtxTransfer
	consNfInfo         int          // depends on AnaCtxTransfer
	notifCorrId        int          // Depend on EneNA

}

type EventSubscription struct {
	// MUST
	event              string
	notificationMethod string
	repetitionPeriod   int // depends on nofitication method PERIODIC

	// May include
	maxObjectNbr int
	maxSupiNbr   int
	startTs      int
	endTs        int
	accuracy     int

	// Depends on EneNA Feature
	timeAnaNeeded     int
	offsetPeriod      int
	accPerSubset      int
	histAnaTimePeriod int

	// Depends on Aggregation Feature
	anaMeta    int
	anaMetaInd int
}

type EventReportingInformation struct {
	notifMethod  string // overrides EventSubscription notificationMethod
	maxReportNbr int
	monDur       int
	repPeriod    int
	immRep       int
	sampRatio    int

	// If EneNA feature supported
	partitionCriteria int
	grpRepTime        int
	notifFlag         int
}
