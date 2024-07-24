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

// PaymentAllocationOrderEnum Ordered list of balance types to which payments are allocated, from first to last.
type PaymentAllocationOrderEnum string

// List of PaymentAllocationOrderEnum
const (
	PAYMENTALLOCATIONORDERENUM_INTEREST PaymentAllocationOrderEnum = "INTEREST"
	PAYMENTALLOCATIONORDERENUM_FEES PaymentAllocationOrderEnum = "FEES"
	PAYMENTALLOCATIONORDERENUM_PRINCIPAL PaymentAllocationOrderEnum = "PRINCIPAL"
)

// All allowed values of PaymentAllocationOrderEnum enum
var AllowedPaymentAllocationOrderEnumEnumValues = []PaymentAllocationOrderEnum{
	"INTEREST",
	"FEES",
	"PRINCIPAL",
}

func (v *PaymentAllocationOrderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentAllocationOrderEnum(value)
	for _, existing := range AllowedPaymentAllocationOrderEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentAllocationOrderEnum", value)
}

// NewPaymentAllocationOrderEnumFromValue returns a pointer to a valid PaymentAllocationOrderEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentAllocationOrderEnumFromValue(v string) (*PaymentAllocationOrderEnum, error) {
	ev := PaymentAllocationOrderEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentAllocationOrderEnum: valid values are %v", v, AllowedPaymentAllocationOrderEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentAllocationOrderEnum) IsValid() bool {
	for _, existing := range AllowedPaymentAllocationOrderEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentAllocationOrderEnum value
func (v PaymentAllocationOrderEnum) Ptr() *PaymentAllocationOrderEnum {
	return &v
}

type NullablePaymentAllocationOrderEnum struct {
	value *PaymentAllocationOrderEnum
	isSet bool
}

func (v NullablePaymentAllocationOrderEnum) Get() *PaymentAllocationOrderEnum {
	return v.value
}

func (v *NullablePaymentAllocationOrderEnum) Set(val *PaymentAllocationOrderEnum) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentAllocationOrderEnum) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentAllocationOrderEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentAllocationOrderEnum(val *PaymentAllocationOrderEnum) *NullablePaymentAllocationOrderEnum {
	return &NullablePaymentAllocationOrderEnum{value: val, isSet: true}
}

func (v NullablePaymentAllocationOrderEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentAllocationOrderEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

