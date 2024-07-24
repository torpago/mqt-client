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
	"bytes"
	"fmt"
)

// checks if the ConfigFeeScheduleEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigFeeScheduleEntry{}

// ConfigFeeScheduleEntry Contains information on a fee schedule.
type ConfigFeeScheduleEntry struct {
	// Date and time when the fee goes into effect, in UTC.
	EffectiveDate *time.Time `json:"effective_date,omitempty"`
	Method FeeMethod `json:"method"`
	// Amount of the fee.
	Value float32 `json:"value"`
}

type _ConfigFeeScheduleEntry ConfigFeeScheduleEntry

// NewConfigFeeScheduleEntry instantiates a new ConfigFeeScheduleEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigFeeScheduleEntry(method FeeMethod, value float32) *ConfigFeeScheduleEntry {
	this := ConfigFeeScheduleEntry{}
	this.Method = method
	this.Value = value
	return &this
}

// NewConfigFeeScheduleEntryWithDefaults instantiates a new ConfigFeeScheduleEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigFeeScheduleEntryWithDefaults() *ConfigFeeScheduleEntry {
	this := ConfigFeeScheduleEntry{}
	return &this
}

// GetEffectiveDate returns the EffectiveDate field value if set, zero value otherwise.
func (o *ConfigFeeScheduleEntry) GetEffectiveDate() time.Time {
	if o == nil || IsNil(o.EffectiveDate) {
		var ret time.Time
		return ret
	}
	return *o.EffectiveDate
}

// GetEffectiveDateOk returns a tuple with the EffectiveDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigFeeScheduleEntry) GetEffectiveDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EffectiveDate) {
		return nil, false
	}
	return o.EffectiveDate, true
}

// HasEffectiveDate returns a boolean if a field has been set.
func (o *ConfigFeeScheduleEntry) HasEffectiveDate() bool {
	if o != nil && !IsNil(o.EffectiveDate) {
		return true
	}

	return false
}

// SetEffectiveDate gets a reference to the given time.Time and assigns it to the EffectiveDate field.
func (o *ConfigFeeScheduleEntry) SetEffectiveDate(v time.Time) {
	o.EffectiveDate = &v
}

// GetMethod returns the Method field value
func (o *ConfigFeeScheduleEntry) GetMethod() FeeMethod {
	if o == nil {
		var ret FeeMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *ConfigFeeScheduleEntry) GetMethodOk() (*FeeMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *ConfigFeeScheduleEntry) SetMethod(v FeeMethod) {
	o.Method = v
}

// GetValue returns the Value field value
func (o *ConfigFeeScheduleEntry) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *ConfigFeeScheduleEntry) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ConfigFeeScheduleEntry) SetValue(v float32) {
	o.Value = v
}

func (o ConfigFeeScheduleEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigFeeScheduleEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EffectiveDate) {
		toSerialize["effective_date"] = o.EffectiveDate
	}
	toSerialize["method"] = o.Method
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *ConfigFeeScheduleEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"method",
		"value",
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

	varConfigFeeScheduleEntry := _ConfigFeeScheduleEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varConfigFeeScheduleEntry)

	if err != nil {
		return err
	}

	*o = ConfigFeeScheduleEntry(varConfigFeeScheduleEntry)

	return err
}

type NullableConfigFeeScheduleEntry struct {
	value *ConfigFeeScheduleEntry
	isSet bool
}

func (v NullableConfigFeeScheduleEntry) Get() *ConfigFeeScheduleEntry {
	return v.value
}

func (v *NullableConfigFeeScheduleEntry) Set(val *ConfigFeeScheduleEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigFeeScheduleEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigFeeScheduleEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigFeeScheduleEntry(val *ConfigFeeScheduleEntry) *NullableConfigFeeScheduleEntry {
	return &NullableConfigFeeScheduleEntry{value: val, isSet: true}
}

func (v NullableConfigFeeScheduleEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigFeeScheduleEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


