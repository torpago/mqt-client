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

// checks if the MerchantScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantScope{}

// MerchantScope Defines the group of merchants to which the velocity control applies.  Populate no more than one field of the `merchant_scope` object. If no fields are populated, the velocity control applies to all merchants.
type MerchantScope struct {
	// Merchant Category Code (MCC). Identifies the type of products or services provided by the merchant.  Enter a value to control spending on a particular type of product or service.
	Mcc *string `json:"mcc,omitempty"`
	// Token identifying a group of MCCs. Enter a value to control spending on a group of product or service types.  Send a `GET` request to `/mccgroups` to retrieve MCC group tokens.
	MccGroup *string `json:"mcc_group,omitempty"`
	// Merchant identifier (MID). Identifies a specific merchant.  Enter a value to control spending with a particular merchant.
	Mid *string `json:"mid,omitempty"`
}

// NewMerchantScope instantiates a new MerchantScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantScope() *MerchantScope {
	this := MerchantScope{}
	return &this
}

// NewMerchantScopeWithDefaults instantiates a new MerchantScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantScopeWithDefaults() *MerchantScope {
	this := MerchantScope{}
	return &this
}

// GetMcc returns the Mcc field value if set, zero value otherwise.
func (o *MerchantScope) GetMcc() string {
	if o == nil || IsNil(o.Mcc) {
		var ret string
		return ret
	}
	return *o.Mcc
}

// GetMccOk returns a tuple with the Mcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantScope) GetMccOk() (*string, bool) {
	if o == nil || IsNil(o.Mcc) {
		return nil, false
	}
	return o.Mcc, true
}

// HasMcc returns a boolean if a field has been set.
func (o *MerchantScope) HasMcc() bool {
	if o != nil && !IsNil(o.Mcc) {
		return true
	}

	return false
}

// SetMcc gets a reference to the given string and assigns it to the Mcc field.
func (o *MerchantScope) SetMcc(v string) {
	o.Mcc = &v
}

// GetMccGroup returns the MccGroup field value if set, zero value otherwise.
func (o *MerchantScope) GetMccGroup() string {
	if o == nil || IsNil(o.MccGroup) {
		var ret string
		return ret
	}
	return *o.MccGroup
}

// GetMccGroupOk returns a tuple with the MccGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantScope) GetMccGroupOk() (*string, bool) {
	if o == nil || IsNil(o.MccGroup) {
		return nil, false
	}
	return o.MccGroup, true
}

// HasMccGroup returns a boolean if a field has been set.
func (o *MerchantScope) HasMccGroup() bool {
	if o != nil && !IsNil(o.MccGroup) {
		return true
	}

	return false
}

// SetMccGroup gets a reference to the given string and assigns it to the MccGroup field.
func (o *MerchantScope) SetMccGroup(v string) {
	o.MccGroup = &v
}

// GetMid returns the Mid field value if set, zero value otherwise.
func (o *MerchantScope) GetMid() string {
	if o == nil || IsNil(o.Mid) {
		var ret string
		return ret
	}
	return *o.Mid
}

// GetMidOk returns a tuple with the Mid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantScope) GetMidOk() (*string, bool) {
	if o == nil || IsNil(o.Mid) {
		return nil, false
	}
	return o.Mid, true
}

// HasMid returns a boolean if a field has been set.
func (o *MerchantScope) HasMid() bool {
	if o != nil && !IsNil(o.Mid) {
		return true
	}

	return false
}

// SetMid gets a reference to the given string and assigns it to the Mid field.
func (o *MerchantScope) SetMid(v string) {
	o.Mid = &v
}

func (o MerchantScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mcc) {
		toSerialize["mcc"] = o.Mcc
	}
	if !IsNil(o.MccGroup) {
		toSerialize["mcc_group"] = o.MccGroup
	}
	if !IsNil(o.Mid) {
		toSerialize["mid"] = o.Mid
	}
	return toSerialize, nil
}

type NullableMerchantScope struct {
	value *MerchantScope
	isSet bool
}

func (v NullableMerchantScope) Get() *MerchantScope {
	return v.value
}

func (v *NullableMerchantScope) Set(val *MerchantScope) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantScope) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantScope(val *MerchantScope) *NullableMerchantScope {
	return &NullableMerchantScope{value: val, isSet: true}
}

func (v NullableMerchantScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


