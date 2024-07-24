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
)

// checks if the AccountConfigPaymentHolds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountConfigPaymentHolds{}

// AccountConfigPaymentHolds Contains configurations for a payment hold.
type AccountConfigPaymentHolds struct {
	// Number of days to hold an ACH payment.
	AchHoldDays *int32 `json:"ach_hold_days,omitempty"`
	// Number of days to hold a check payment.
	CheckHoldDays *int32 `json:"check_hold_days,omitempty"`
}

// NewAccountConfigPaymentHolds instantiates a new AccountConfigPaymentHolds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountConfigPaymentHolds() *AccountConfigPaymentHolds {
	this := AccountConfigPaymentHolds{}
	return &this
}

// NewAccountConfigPaymentHoldsWithDefaults instantiates a new AccountConfigPaymentHolds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountConfigPaymentHoldsWithDefaults() *AccountConfigPaymentHolds {
	this := AccountConfigPaymentHolds{}
	return &this
}

// GetAchHoldDays returns the AchHoldDays field value if set, zero value otherwise.
func (o *AccountConfigPaymentHolds) GetAchHoldDays() int32 {
	if o == nil || IsNil(o.AchHoldDays) {
		var ret int32
		return ret
	}
	return *o.AchHoldDays
}

// GetAchHoldDaysOk returns a tuple with the AchHoldDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigPaymentHolds) GetAchHoldDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.AchHoldDays) {
		return nil, false
	}
	return o.AchHoldDays, true
}

// HasAchHoldDays returns a boolean if a field has been set.
func (o *AccountConfigPaymentHolds) HasAchHoldDays() bool {
	if o != nil && !IsNil(o.AchHoldDays) {
		return true
	}

	return false
}

// SetAchHoldDays gets a reference to the given int32 and assigns it to the AchHoldDays field.
func (o *AccountConfigPaymentHolds) SetAchHoldDays(v int32) {
	o.AchHoldDays = &v
}

// GetCheckHoldDays returns the CheckHoldDays field value if set, zero value otherwise.
func (o *AccountConfigPaymentHolds) GetCheckHoldDays() int32 {
	if o == nil || IsNil(o.CheckHoldDays) {
		var ret int32
		return ret
	}
	return *o.CheckHoldDays
}

// GetCheckHoldDaysOk returns a tuple with the CheckHoldDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigPaymentHolds) GetCheckHoldDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckHoldDays) {
		return nil, false
	}
	return o.CheckHoldDays, true
}

// HasCheckHoldDays returns a boolean if a field has been set.
func (o *AccountConfigPaymentHolds) HasCheckHoldDays() bool {
	if o != nil && !IsNil(o.CheckHoldDays) {
		return true
	}

	return false
}

// SetCheckHoldDays gets a reference to the given int32 and assigns it to the CheckHoldDays field.
func (o *AccountConfigPaymentHolds) SetCheckHoldDays(v int32) {
	o.CheckHoldDays = &v
}

func (o AccountConfigPaymentHolds) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountConfigPaymentHolds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AchHoldDays) {
		toSerialize["ach_hold_days"] = o.AchHoldDays
	}
	if !IsNil(o.CheckHoldDays) {
		toSerialize["check_hold_days"] = o.CheckHoldDays
	}
	return toSerialize, nil
}

type NullableAccountConfigPaymentHolds struct {
	value *AccountConfigPaymentHolds
	isSet bool
}

func (v NullableAccountConfigPaymentHolds) Get() *AccountConfigPaymentHolds {
	return v.value
}

func (v *NullableAccountConfigPaymentHolds) Set(val *AccountConfigPaymentHolds) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountConfigPaymentHolds) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountConfigPaymentHolds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountConfigPaymentHolds(val *AccountConfigPaymentHolds) *NullableAccountConfigPaymentHolds {
	return &NullableAccountConfigPaymentHolds{value: val, isSet: true}
}

func (v NullableAccountConfigPaymentHolds) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountConfigPaymentHolds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


