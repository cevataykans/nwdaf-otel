// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2025, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.4
 */

package analyticsinfo

import 	"time"


// DataNotification - Represents a Data Subscription Notification.
type DataNotification struct {

	// List of notifications of AMF events.
	AmfEventNotifs []AmfEventNotification `json:"amfEventNotifs,omitempty"`

	// List of notifications of SMF events.
	SmfEventNotifs []NsmfEventExposureNotification `json:"smfEventNotifs,omitempty"`

	// List of notifications of UDM events.
	UdmEventNotifs []MonitoringReport `json:"udmEventNotifs,omitempty"`

	// List of notifications of NEF events.
	NefEventNotifs []NefEventExposureNotif `json:"nefEventNotifs,omitempty"`

	// List of notifications of AF events.
	AfEventNotifs []AfEventExposureNotif `json:"afEventNotifs,omitempty"`

	// List of notifications of NRF events.
	NrfEventNotifs []NotificationData `json:"nrfEventNotifs,omitempty"`

	// List of notifications of NSACF events.
	NsacfEventNotifs []SacEventReport `json:"nsacfEventNotifs,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp,omitempty"`
}

// AssertDataNotificationRequired checks if the required fields are not zero-ed
func AssertDataNotificationRequired(obj DataNotification) error {
	for _, el := range obj.AmfEventNotifs {
		if err := AssertAmfEventNotificationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SmfEventNotifs {
		if err := AssertNsmfEventExposureNotificationRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UdmEventNotifs {
		if err := AssertMonitoringReportRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NefEventNotifs {
		if err := AssertNefEventExposureNotifRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AfEventNotifs {
		if err := AssertAfEventExposureNotifRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NrfEventNotifs {
		if err := AssertNotificationDataRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NsacfEventNotifs {
		if err := AssertSacEventReportRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertDataNotificationConstraints checks if the values respects the defined constraints
func AssertDataNotificationConstraints(obj DataNotification) error {
	for _, el := range obj.AmfEventNotifs {
		if err := AssertAmfEventNotificationConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SmfEventNotifs {
		if err := AssertNsmfEventExposureNotificationConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.UdmEventNotifs {
		if err := AssertMonitoringReportConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NefEventNotifs {
		if err := AssertNefEventExposureNotifConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.AfEventNotifs {
		if err := AssertAfEventExposureNotifConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NrfEventNotifs {
		if err := AssertNotificationDataConstraints(el); err != nil {
			return err
		}
	}
	for _, el := range obj.NsacfEventNotifs {
		if err := AssertSacEventReportConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
