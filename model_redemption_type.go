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

// RedemptionType Type of redemption.  * `EXTERNAL` - You issue the redemption on your external platform; Marqeta adjusts the reward program balance on the system of record. * `STATEMENT_CREDIT` - Marqeta issues the redemption as a statement credit on the associated account. + *NOTE*: This creates a new journal entry on the account and cannot be undone. * `ACH` - The reward redemption is issued as an ACH transfer to the receiving account.
type RedemptionType string

// List of RedemptionType
const (
	REDEMPTIONTYPE_EXTERNAL RedemptionType = "EXTERNAL"
	REDEMPTIONTYPE_STATEMENT_CREDIT RedemptionType = "STATEMENT_CREDIT"
	REDEMPTIONTYPE_ACH RedemptionType = "ACH"
)

// All allowed values of RedemptionType enum
var AllowedRedemptionTypeEnumValues = []RedemptionType{
	"EXTERNAL",
	"STATEMENT_CREDIT",
	"ACH",
}

func (v *RedemptionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RedemptionType(value)
	for _, existing := range AllowedRedemptionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RedemptionType", value)
}

// NewRedemptionTypeFromValue returns a pointer to a valid RedemptionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRedemptionTypeFromValue(v string) (*RedemptionType, error) {
	ev := RedemptionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RedemptionType: valid values are %v", v, AllowedRedemptionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RedemptionType) IsValid() bool {
	for _, existing := range AllowedRedemptionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RedemptionType value
func (v RedemptionType) Ptr() *RedemptionType {
	return &v
}

type NullableRedemptionType struct {
	value *RedemptionType
	isSet bool
}

func (v NullableRedemptionType) Get() *RedemptionType {
	return v.value
}

func (v *NullableRedemptionType) Set(val *RedemptionType) {
	v.value = val
	v.isSet = true
}

func (v NullableRedemptionType) IsSet() bool {
	return v.isSet
}

func (v *NullableRedemptionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedemptionType(val *RedemptionType) *NullableRedemptionType {
	return &NullableRedemptionType{value: val, isSet: true}
}

func (v NullableRedemptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedemptionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

