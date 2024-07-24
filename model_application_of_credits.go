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
	"bytes"
	"fmt"
)

// checks if the ApplicationOfCredits type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationOfCredits{}

// ApplicationOfCredits Contains information on the cycle type and billing cycle day when credits are applied in the daily balance calculation.
type ApplicationOfCredits struct {
	CycleType CycleType `json:"cycle_type"`
	// Day of the billing cycle when credits are applied.
	Day int32 `json:"day"`
}

type _ApplicationOfCredits ApplicationOfCredits

// NewApplicationOfCredits instantiates a new ApplicationOfCredits object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOfCredits(cycleType CycleType, day int32) *ApplicationOfCredits {
	this := ApplicationOfCredits{}
	this.CycleType = cycleType
	this.Day = day
	return &this
}

// NewApplicationOfCreditsWithDefaults instantiates a new ApplicationOfCredits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOfCreditsWithDefaults() *ApplicationOfCredits {
	this := ApplicationOfCredits{}
	return &this
}

// GetCycleType returns the CycleType field value
func (o *ApplicationOfCredits) GetCycleType() CycleType {
	if o == nil {
		var ret CycleType
		return ret
	}

	return o.CycleType
}

// GetCycleTypeOk returns a tuple with the CycleType field value
// and a boolean to check if the value has been set.
func (o *ApplicationOfCredits) GetCycleTypeOk() (*CycleType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CycleType, true
}

// SetCycleType sets field value
func (o *ApplicationOfCredits) SetCycleType(v CycleType) {
	o.CycleType = v
}

// GetDay returns the Day field value
func (o *ApplicationOfCredits) GetDay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Day
}

// GetDayOk returns a tuple with the Day field value
// and a boolean to check if the value has been set.
func (o *ApplicationOfCredits) GetDayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Day, true
}

// SetDay sets field value
func (o *ApplicationOfCredits) SetDay(v int32) {
	o.Day = v
}

func (o ApplicationOfCredits) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationOfCredits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cycle_type"] = o.CycleType
	toSerialize["day"] = o.Day
	return toSerialize, nil
}

func (o *ApplicationOfCredits) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cycle_type",
		"day",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApplicationOfCredits := _ApplicationOfCredits{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationOfCredits)

	if err != nil {
		return err
	}

	*o = ApplicationOfCredits(varApplicationOfCredits)

	return err
}

type NullableApplicationOfCredits struct {
	value *ApplicationOfCredits
	isSet bool
}

func (v NullableApplicationOfCredits) Get() *ApplicationOfCredits {
	return v.value
}

func (v *NullableApplicationOfCredits) Set(val *ApplicationOfCredits) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOfCredits) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOfCredits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOfCredits(val *ApplicationOfCredits) *NullableApplicationOfCredits {
	return &NullableApplicationOfCredits{value: val, isSet: true}
}

func (v NullableApplicationOfCredits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOfCredits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


