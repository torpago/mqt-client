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

// ProductSubType Subtype of the credit product type.  `CREDIT_CARD` - Card that enables the cardholder to make purchases on credit.
type ProductSubType string

// List of ProductSubType
const (
	PRODUCTSUBTYPE_CREDIT_CARD ProductSubType = "CREDIT_CARD"
)

// All allowed values of ProductSubType enum
var AllowedProductSubTypeEnumValues = []ProductSubType{
	"CREDIT_CARD",
}

func (v *ProductSubType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductSubType(value)
	for _, existing := range AllowedProductSubTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductSubType", value)
}

// NewProductSubTypeFromValue returns a pointer to a valid ProductSubType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductSubTypeFromValue(v string) (*ProductSubType, error) {
	ev := ProductSubType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductSubType: valid values are %v", v, AllowedProductSubTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductSubType) IsValid() bool {
	for _, existing := range AllowedProductSubTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductSubType value
func (v ProductSubType) Ptr() *ProductSubType {
	return &v
}

type NullableProductSubType struct {
	value *ProductSubType
	isSet bool
}

func (v NullableProductSubType) Get() *ProductSubType {
	return v.value
}

func (v *NullableProductSubType) Set(val *ProductSubType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductSubType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductSubType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductSubType(val *ProductSubType) *NullableProductSubType {
	return &NullableProductSubType{value: val, isSet: true}
}

func (v NullableProductSubType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductSubType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

