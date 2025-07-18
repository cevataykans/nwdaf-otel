// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2024, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.3
 */

package datamanagement




// CollectiveBehaviourFilter - Contains the collective behaviour filter information to be collected from UE.
type CollectiveBehaviourFilter struct {

	Type CollectiveBehaviourFilterType `json:"type"`

	// Value of the parameter type as in the type attribute.
	Value string `json:"value"`

	// Indicates whether request list of UE IDs that fulfill a collective behaviour within the area of interest. This attribute shall set to \"true\" if request the list of UE IDs, otherwise, set to \"false\". May only be present and sets to \"true\" if \"AfEvent\" sets to \"COLLECTIVE_BEHAVIOUR\". 
	ListOfUeInd bool `json:"listOfUeInd,omitempty"`
}

// AssertCollectiveBehaviourFilterRequired checks if the required fields are not zero-ed
func AssertCollectiveBehaviourFilterRequired(obj CollectiveBehaviourFilter) error {
	elements := map[string]interface{}{
		"type": obj.Type,
		"value": obj.Value,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCollectiveBehaviourFilterTypeRequired(obj.Type); err != nil {
		return err
	}
	return nil
}

// AssertCollectiveBehaviourFilterConstraints checks if the values respects the defined constraints
func AssertCollectiveBehaviourFilterConstraints(obj CollectiveBehaviourFilter) error {
	if err := AssertCollectiveBehaviourFilterTypeConstraints(obj.Type); err != nil {
		return err
	}
	return nil
}
