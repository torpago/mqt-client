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

// VelocityWindow Defines the time period to which the `amount_limit` and `usage_limit` fields apply: * *MONTH* – one month; months begin on the first day of month at 00:00:00 ET.
type VelocityWindow string

// List of VelocityWindow
const (
	VELOCITYWINDOW_MONTH VelocityWindow = "MONTH"
)

// All allowed values of VelocityWindow enum
var AllowedVelocityWindowEnumValues = []VelocityWindow{
	"MONTH",
}

func (v *VelocityWindow) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VelocityWindow(value)
	for _, existing := range AllowedVelocityWindowEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VelocityWindow", value)
}

// NewVelocityWindowFromValue returns a pointer to a valid VelocityWindow
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVelocityWindowFromValue(v string) (*VelocityWindow, error) {
	ev := VelocityWindow(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VelocityWindow: valid values are %v", v, AllowedVelocityWindowEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VelocityWindow) IsValid() bool {
	for _, existing := range AllowedVelocityWindowEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VelocityWindow value
func (v VelocityWindow) Ptr() *VelocityWindow {
	return &v
}

type NullableVelocityWindow struct {
	value *VelocityWindow
	isSet bool
}

func (v NullableVelocityWindow) Get() *VelocityWindow {
	return v.value
}

func (v *NullableVelocityWindow) Set(val *VelocityWindow) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityWindow) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityWindow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityWindow(val *VelocityWindow) *NullableVelocityWindow {
	return &NullableVelocityWindow{value: val, isSet: true}
}

func (v NullableVelocityWindow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityWindow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

