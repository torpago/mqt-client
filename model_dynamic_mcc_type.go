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

// DynamicMccType Type of dynamic MCC.
type DynamicMccType string

// List of DynamicMccType
const (
	DYNAMICMCCTYPE_HIGHEST_SPEND DynamicMccType = "HIGHEST_SPEND"
)

// All allowed values of DynamicMccType enum
var AllowedDynamicMccTypeEnumValues = []DynamicMccType{
	"HIGHEST_SPEND",
}

func (v *DynamicMccType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DynamicMccType(value)
	for _, existing := range AllowedDynamicMccTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DynamicMccType", value)
}

// NewDynamicMccTypeFromValue returns a pointer to a valid DynamicMccType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDynamicMccTypeFromValue(v string) (*DynamicMccType, error) {
	ev := DynamicMccType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DynamicMccType: valid values are %v", v, AllowedDynamicMccTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DynamicMccType) IsValid() bool {
	for _, existing := range AllowedDynamicMccTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DynamicMccType value
func (v DynamicMccType) Ptr() *DynamicMccType {
	return &v
}

type NullableDynamicMccType struct {
	value *DynamicMccType
	isSet bool
}

func (v NullableDynamicMccType) Get() *DynamicMccType {
	return v.value
}

func (v *NullableDynamicMccType) Set(val *DynamicMccType) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicMccType) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicMccType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicMccType(val *DynamicMccType) *NullableDynamicMccType {
	return &NullableDynamicMccType{value: val, isSet: true}
}

func (v NullableDynamicMccType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicMccType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

