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

// CalculationType Type of calculation for the reward.  * `NET_BALANCE` - The reward is calculated based on the net balance for a billing cycle, which is the total amount spent during a billing cycle, minus any refunds or disputes.
type CalculationType string

// List of CalculationType
const (
	CALCULATIONTYPE_NET_BALANCE CalculationType = "NET_BALANCE"
)

// All allowed values of CalculationType enum
var AllowedCalculationTypeEnumValues = []CalculationType{
	"NET_BALANCE",
}

func (v *CalculationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CalculationType(value)
	for _, existing := range AllowedCalculationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CalculationType", value)
}

// NewCalculationTypeFromValue returns a pointer to a valid CalculationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCalculationTypeFromValue(v string) (*CalculationType, error) {
	ev := CalculationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CalculationType: valid values are %v", v, AllowedCalculationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CalculationType) IsValid() bool {
	for _, existing := range AllowedCalculationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CalculationType value
func (v CalculationType) Ptr() *CalculationType {
	return &v
}

type NullableCalculationType struct {
	value *CalculationType
	isSet bool
}

func (v NullableCalculationType) Get() *CalculationType {
	return v.value
}

func (v *NullableCalculationType) Set(val *CalculationType) {
	v.value = val
	v.isSet = true
}

func (v NullableCalculationType) IsSet() bool {
	return v.isSet
}

func (v *NullableCalculationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCalculationType(val *CalculationType) *NullableCalculationType {
	return &NullableCalculationType{value: val, isSet: true}
}

func (v NullableCalculationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCalculationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

