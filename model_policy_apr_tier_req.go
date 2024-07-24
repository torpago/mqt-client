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

// checks if the PolicyAprTierReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAprTierReq{}

// PolicyAprTierReq Request details for the APR for a risk tier.
type PolicyAprTierReq struct {
	// Number of percentage points added to the prime rate, used to calculate a variable APR value.
	MarginRate decimal.Decimal `json:"margin_rate"`
}

type _PolicyAprTierReq PolicyAprTierReq

// NewPolicyAprTierReq instantiates a new PolicyAprTierReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAprTierReq(marginRate decimal.Decimal) *PolicyAprTierReq {
	this := PolicyAprTierReq{}
	this.MarginRate = marginRate
	return &this
}

// NewPolicyAprTierReqWithDefaults instantiates a new PolicyAprTierReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAprTierReqWithDefaults() *PolicyAprTierReq {
	this := PolicyAprTierReq{}
	return &this
}

// GetMarginRate returns the MarginRate field value
func (o *PolicyAprTierReq) GetMarginRate() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.MarginRate
}

// GetMarginRateOk returns a tuple with the MarginRate field value
// and a boolean to check if the value has been set.
func (o *PolicyAprTierReq) GetMarginRateOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarginRate, true
}

// SetMarginRate sets field value
func (o *PolicyAprTierReq) SetMarginRate(v decimal.Decimal) {
	o.MarginRate = v
}

func (o PolicyAprTierReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAprTierReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["margin_rate"] = o.MarginRate
	return toSerialize, nil
}

func (o *PolicyAprTierReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"margin_rate",
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

	varPolicyAprTierReq := _PolicyAprTierReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPolicyAprTierReq)

	if err != nil {
		return err
	}

	*o = PolicyAprTierReq(varPolicyAprTierReq)

	return err
}

type NullablePolicyAprTierReq struct {
	value *PolicyAprTierReq
	isSet bool
}

func (v NullablePolicyAprTierReq) Get() *PolicyAprTierReq {
	return v.value
}

func (v *NullablePolicyAprTierReq) Set(val *PolicyAprTierReq) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAprTierReq) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAprTierReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAprTierReq(val *PolicyAprTierReq) *NullablePolicyAprTierReq {
	return &NullablePolicyAprTierReq{value: val, isSet: true}
}

func (v NullablePolicyAprTierReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAprTierReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


