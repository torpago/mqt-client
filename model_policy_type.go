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

// PolicyType Type of policy.
type PolicyType string

// List of PolicyType
const (
	POLICYTYPE_APR PolicyType = "APR"
	POLICYTYPE_DOCUMENT PolicyType = "DOCUMENT"
	POLICYTYPE_FEE PolicyType = "FEE"
	POLICYTYPE_OFFER PolicyType = "OFFER"
	POLICYTYPE_PRODUCT PolicyType = "PRODUCT"
	POLICYTYPE_REWARD PolicyType = "REWARD"
	POLICYTYPE_ALL PolicyType = "ALL"
)

// All allowed values of PolicyType enum
var AllowedPolicyTypeEnumValues = []PolicyType{
	"APR",
	"DOCUMENT",
	"FEE",
	"OFFER",
	"PRODUCT",
	"REWARD",
	"ALL",
}

func (v *PolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PolicyType(value)
	for _, existing := range AllowedPolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PolicyType", value)
}

// NewPolicyTypeFromValue returns a pointer to a valid PolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPolicyTypeFromValue(v string) (*PolicyType, error) {
	ev := PolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PolicyType: valid values are %v", v, AllowedPolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PolicyType) IsValid() bool {
	for _, existing := range AllowedPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PolicyType value
func (v PolicyType) Ptr() *PolicyType {
	return &v
}

type NullablePolicyType struct {
	value *PolicyType
	isSet bool
}

func (v NullablePolicyType) Get() *PolicyType {
	return v.value
}

func (v *NullablePolicyType) Set(val *PolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyType(val *PolicyType) *NullablePolicyType {
	return &NullablePolicyType{value: val, isSet: true}
}

func (v NullablePolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

