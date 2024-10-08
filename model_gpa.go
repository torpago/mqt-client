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
	"github.com/shopspring/decimal"
)

// checks if the Gpa type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Gpa{}

// Gpa Defines the type of order.
type Gpa struct {
	// Available balance on the card after the reload has completed.  This value must be greater than or equal to the value of `trigger_amount`. Note that this is not the same as the amount added to the card, which will vary from reload to reload.
	ReloadAmount decimal.Decimal `json:"reload_amount"`
	// Threshold that determines when the reload happens.  The reload is triggered when the card balance falls below this amount.
	TriggerAmount decimal.Decimal `json:"trigger_amount"`
}

type _Gpa Gpa

// NewGpa instantiates a new Gpa object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGpa(reloadAmount decimal.Decimal, triggerAmount decimal.Decimal) *Gpa {
	this := Gpa{}
	this.ReloadAmount = reloadAmount
	this.TriggerAmount = triggerAmount
	return &this
}

// NewGpaWithDefaults instantiates a new Gpa object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGpaWithDefaults() *Gpa {
	this := Gpa{}
	return &this
}

// GetReloadAmount returns the ReloadAmount field value
func (o *Gpa) GetReloadAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.ReloadAmount
}

// GetReloadAmountOk returns a tuple with the ReloadAmount field value
// and a boolean to check if the value has been set.
func (o *Gpa) GetReloadAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReloadAmount, true
}

// SetReloadAmount sets field value
func (o *Gpa) SetReloadAmount(v decimal.Decimal) {
	o.ReloadAmount = v
}

// GetTriggerAmount returns the TriggerAmount field value
func (o *Gpa) GetTriggerAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.TriggerAmount
}

// GetTriggerAmountOk returns a tuple with the TriggerAmount field value
// and a boolean to check if the value has been set.
func (o *Gpa) GetTriggerAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerAmount, true
}

// SetTriggerAmount sets field value
func (o *Gpa) SetTriggerAmount(v decimal.Decimal) {
	o.TriggerAmount = v
}

func (o Gpa) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Gpa) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["reload_amount"] = o.ReloadAmount
	toSerialize["trigger_amount"] = o.TriggerAmount
	return toSerialize, nil
}

func (o *Gpa) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reload_amount",
		"trigger_amount",
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

	varGpa := _Gpa{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGpa)

	if err != nil {
		return err
	}

	*o = Gpa(varGpa)

	return err
}

type NullableGpa struct {
	value *Gpa
	isSet bool
}

func (v NullableGpa) Get() *Gpa {
	return v.value
}

func (v *NullableGpa) Set(val *Gpa) {
	v.value = val
	v.isSet = true
}

func (v NullableGpa) IsSet() bool {
	return v.isSet
}

func (v *NullableGpa) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGpa(val *Gpa) *NullableGpa {
	return &NullableGpa{value: val, isSet: true}
}

func (v NullableGpa) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGpa) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


