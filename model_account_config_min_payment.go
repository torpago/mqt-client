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
	"time"
)

// checks if the AccountConfigMinPayment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountConfigMinPayment{}

// AccountConfigMinPayment Contains configurations for a minimum payment override on a credit account, which overrides the minimum payment configurations on the associated credit product.
type AccountConfigMinPayment struct {
	// Whether the minimum payment override is currently active.
	Active *bool `json:"active,omitempty"`
	// Flat amount of the minimum payment override.
	MinPaymentFlatAmount *float32 `json:"min_payment_flat_amount,omitempty"`
	// Percentage of the total statement balance used to calculate the minimum payment override amount.
	MinPaymentPercentage *float32 `json:"min_payment_percentage,omitempty"`
	// Date and time when the minimum payment override ends, in UTC.
	OverrideEndTime *time.Time `json:"override_end_time,omitempty"`
	// Date and time when the minimum payment override starts, in UTC.
	OverrideStartTime *time.Time `json:"override_start_time,omitempty"`
}

// NewAccountConfigMinPayment instantiates a new AccountConfigMinPayment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountConfigMinPayment() *AccountConfigMinPayment {
	this := AccountConfigMinPayment{}
	return &this
}

// NewAccountConfigMinPaymentWithDefaults instantiates a new AccountConfigMinPayment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountConfigMinPaymentWithDefaults() *AccountConfigMinPayment {
	this := AccountConfigMinPayment{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *AccountConfigMinPayment) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigMinPayment) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *AccountConfigMinPayment) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *AccountConfigMinPayment) SetActive(v bool) {
	o.Active = &v
}

// GetMinPaymentFlatAmount returns the MinPaymentFlatAmount field value if set, zero value otherwise.
func (o *AccountConfigMinPayment) GetMinPaymentFlatAmount() float32 {
	if o == nil || IsNil(o.MinPaymentFlatAmount) {
		var ret float32
		return ret
	}
	return *o.MinPaymentFlatAmount
}

// GetMinPaymentFlatAmountOk returns a tuple with the MinPaymentFlatAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigMinPayment) GetMinPaymentFlatAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.MinPaymentFlatAmount) {
		return nil, false
	}
	return o.MinPaymentFlatAmount, true
}

// HasMinPaymentFlatAmount returns a boolean if a field has been set.
func (o *AccountConfigMinPayment) HasMinPaymentFlatAmount() bool {
	if o != nil && !IsNil(o.MinPaymentFlatAmount) {
		return true
	}

	return false
}

// SetMinPaymentFlatAmount gets a reference to the given float32 and assigns it to the MinPaymentFlatAmount field.
func (o *AccountConfigMinPayment) SetMinPaymentFlatAmount(v float32) {
	o.MinPaymentFlatAmount = &v
}

// GetMinPaymentPercentage returns the MinPaymentPercentage field value if set, zero value otherwise.
func (o *AccountConfigMinPayment) GetMinPaymentPercentage() float32 {
	if o == nil || IsNil(o.MinPaymentPercentage) {
		var ret float32
		return ret
	}
	return *o.MinPaymentPercentage
}

// GetMinPaymentPercentageOk returns a tuple with the MinPaymentPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigMinPayment) GetMinPaymentPercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.MinPaymentPercentage) {
		return nil, false
	}
	return o.MinPaymentPercentage, true
}

// HasMinPaymentPercentage returns a boolean if a field has been set.
func (o *AccountConfigMinPayment) HasMinPaymentPercentage() bool {
	if o != nil && !IsNil(o.MinPaymentPercentage) {
		return true
	}

	return false
}

// SetMinPaymentPercentage gets a reference to the given float32 and assigns it to the MinPaymentPercentage field.
func (o *AccountConfigMinPayment) SetMinPaymentPercentage(v float32) {
	o.MinPaymentPercentage = &v
}

// GetOverrideEndTime returns the OverrideEndTime field value if set, zero value otherwise.
func (o *AccountConfigMinPayment) GetOverrideEndTime() time.Time {
	if o == nil || IsNil(o.OverrideEndTime) {
		var ret time.Time
		return ret
	}
	return *o.OverrideEndTime
}

// GetOverrideEndTimeOk returns a tuple with the OverrideEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigMinPayment) GetOverrideEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OverrideEndTime) {
		return nil, false
	}
	return o.OverrideEndTime, true
}

// HasOverrideEndTime returns a boolean if a field has been set.
func (o *AccountConfigMinPayment) HasOverrideEndTime() bool {
	if o != nil && !IsNil(o.OverrideEndTime) {
		return true
	}

	return false
}

// SetOverrideEndTime gets a reference to the given time.Time and assigns it to the OverrideEndTime field.
func (o *AccountConfigMinPayment) SetOverrideEndTime(v time.Time) {
	o.OverrideEndTime = &v
}

// GetOverrideStartTime returns the OverrideStartTime field value if set, zero value otherwise.
func (o *AccountConfigMinPayment) GetOverrideStartTime() time.Time {
	if o == nil || IsNil(o.OverrideStartTime) {
		var ret time.Time
		return ret
	}
	return *o.OverrideStartTime
}

// GetOverrideStartTimeOk returns a tuple with the OverrideStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountConfigMinPayment) GetOverrideStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OverrideStartTime) {
		return nil, false
	}
	return o.OverrideStartTime, true
}

// HasOverrideStartTime returns a boolean if a field has been set.
func (o *AccountConfigMinPayment) HasOverrideStartTime() bool {
	if o != nil && !IsNil(o.OverrideStartTime) {
		return true
	}

	return false
}

// SetOverrideStartTime gets a reference to the given time.Time and assigns it to the OverrideStartTime field.
func (o *AccountConfigMinPayment) SetOverrideStartTime(v time.Time) {
	o.OverrideStartTime = &v
}

func (o AccountConfigMinPayment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountConfigMinPayment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.MinPaymentFlatAmount) {
		toSerialize["min_payment_flat_amount"] = o.MinPaymentFlatAmount
	}
	if !IsNil(o.MinPaymentPercentage) {
		toSerialize["min_payment_percentage"] = o.MinPaymentPercentage
	}
	if !IsNil(o.OverrideEndTime) {
		toSerialize["override_end_time"] = o.OverrideEndTime
	}
	if !IsNil(o.OverrideStartTime) {
		toSerialize["override_start_time"] = o.OverrideStartTime
	}
	return toSerialize, nil
}

type NullableAccountConfigMinPayment struct {
	value *AccountConfigMinPayment
	isSet bool
}

func (v NullableAccountConfigMinPayment) Get() *AccountConfigMinPayment {
	return v.value
}

func (v *NullableAccountConfigMinPayment) Set(val *AccountConfigMinPayment) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountConfigMinPayment) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountConfigMinPayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountConfigMinPayment(val *AccountConfigMinPayment) *NullableAccountConfigMinPayment {
	return &NullableAccountConfigMinPayment{value: val, isSet: true}
}

func (v NullableAccountConfigMinPayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountConfigMinPayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


