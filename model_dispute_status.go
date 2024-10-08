/*
Core API

Marqeta's Core API endpoints, conveniently annotated to enable code generation (including SDKs), test cases, and documentation. Currently in beta.

API version: 3.0.19
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// DisputeStatus Status of the dispute.  * `ACTIVE` - The dispute is active and awaiting resolution. * `REVERSED` - The dispute has been reversed and is no longer active. * `AH_WON` - The account holder won the dispute. * `AH_LOST` - The account holder lost the dispute.
type DisputeStatus string

// List of DisputeStatus
const (
	DISPUTESTATUS_ACTIVE DisputeStatus = "ACTIVE"
	DISPUTESTATUS_REVERSED DisputeStatus = "REVERSED"
	DISPUTESTATUS_AH_WON DisputeStatus = "AH_WON"
	DISPUTESTATUS_AH_LOST DisputeStatus = "AH_LOST"
	DISPUTESTATUS_REFUNDED DisputeStatus = "REFUNDED"
)

// All allowed values of DisputeStatus enum
var AllowedDisputeStatusEnumValues = []DisputeStatus{
	"ACTIVE",
	"REVERSED",
	"AH_WON",
	"AH_LOST",
	"REFUNDED",
}

func (v *DisputeStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DisputeStatus(value)
	for _, existing := range AllowedDisputeStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DisputeStatus", value)
}

// NewDisputeStatusFromValue returns a pointer to a valid DisputeStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDisputeStatusFromValue(v string) (*DisputeStatus, error) {
	ev := DisputeStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DisputeStatus: valid values are %v", v, AllowedDisputeStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DisputeStatus) IsValid() bool {
	for _, existing := range AllowedDisputeStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DisputeStatus value
func (v DisputeStatus) Ptr() *DisputeStatus {
	return &v
}

type NullableDisputeStatus struct {
	value *DisputeStatus
	isSet bool
}

func (v NullableDisputeStatus) Get() *DisputeStatus {
	return v.value
}

func (v *NullableDisputeStatus) Set(val *DisputeStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDisputeStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDisputeStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisputeStatus(val *DisputeStatus) *NullableDisputeStatus {
	return &NullableDisputeStatus{value: val, isSet: true}
}

func (v NullableDisputeStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisputeStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

