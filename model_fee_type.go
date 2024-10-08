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

// FeeType Type of fee.  *NOTE:* Only `RETURNED_PAYMENT_FEE`, `LATE_PAYMENT_FEE`, `ANNUAL_FEE`, and `MONTHLY_FEE` are currently supported. Do not pass other fees types.
type FeeType string

// List of FeeType
const (
	FEETYPE_FOREIGN_TRANSACTION_FEE FeeType = "FOREIGN_TRANSACTION_FEE"
	FEETYPE_OVER_LIMIT_FEE FeeType = "OVER_LIMIT_FEE"
	FEETYPE_LATE_PAYMENT_FEE FeeType = "LATE_PAYMENT_FEE"
	FEETYPE_RETURNED_PAYMENT_FEE FeeType = "RETURNED_PAYMENT_FEE"
	FEETYPE_CARD_REPLACEMENT_FEE FeeType = "CARD_REPLACEMENT_FEE"
	FEETYPE_MINIMUM_INTEREST_FEE FeeType = "MINIMUM_INTEREST_FEE"
	FEETYPE_MINIMUM_INTEREST_FEE_REVERSAL FeeType = "MINIMUM_INTEREST_FEE_REVERSAL"
	FEETYPE_ANNUAL_FEE FeeType = "ANNUAL_FEE"
	FEETYPE_MONTHLY_FEE FeeType = "MONTHLY_FEE"
)

// All allowed values of FeeType enum
var AllowedFeeTypeEnumValues = []FeeType{
	"FOREIGN_TRANSACTION_FEE",
	"OVER_LIMIT_FEE",
	"LATE_PAYMENT_FEE",
	"RETURNED_PAYMENT_FEE",
	"CARD_REPLACEMENT_FEE",
	"MINIMUM_INTEREST_FEE",
	"MINIMUM_INTEREST_FEE_REVERSAL",
	"ANNUAL_FEE",
	"MONTHLY_FEE",
}

func (v *FeeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FeeType(value)
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FeeType", value)
}

// NewFeeTypeFromValue returns a pointer to a valid FeeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFeeTypeFromValue(v string) (*FeeType, error) {
	ev := FeeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FeeType: valid values are %v", v, AllowedFeeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FeeType) IsValid() bool {
	for _, existing := range AllowedFeeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FeeType value
func (v FeeType) Ptr() *FeeType {
	return &v
}

type NullableFeeType struct {
	value *FeeType
	isSet bool
}

func (v NullableFeeType) Get() *FeeType {
	return v.value
}

func (v *NullableFeeType) Set(val *FeeType) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeType) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeType(val *FeeType) *NullableFeeType {
	return &NullableFeeType{value: val, isSet: true}
}

func (v NullableFeeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

