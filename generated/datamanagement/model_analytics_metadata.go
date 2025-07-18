// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// AnalyticsMetadata - Possible values are: - NUM_OF_SAMPLES: Number of data samples used for the generation of the output analytics. - DATA_WINDOW: Data time window of the data samples. - DATA_STAT_PROPS: Dataset statistical properties of the data used to generate the analytics. - STRATEGY: Output strategy used for the reporting of the analytics. - ACCURACY: Level of accuracy reached for the analytics. 
type AnalyticsMetadata struct {
}

// AssertAnalyticsMetadataRequired checks if the required fields are not zero-ed
func AssertAnalyticsMetadataRequired(obj AnalyticsMetadata) error {
	return nil
}

// AssertAnalyticsMetadataConstraints checks if the values respects the defined constraints
func AssertAnalyticsMetadataConstraints(obj AnalyticsMetadata) error {
	return nil
}
