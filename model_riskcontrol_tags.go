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

// checks if the RiskcontrolTags type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskcontrolTags{}

// RiskcontrolTags The RiskControl tags that were triggered by the transaction.
type RiskcontrolTags struct {
	// Name of the rule, as defined in the Real-Time Decisioning dashboard, that was triggered.
	RuleName *string `json:"rule_name,omitempty"`
	// Tag name defined in the rule definition in the Real-Time Decisioning dashboard.
	Tag *string `json:"tag,omitempty"`
	// Value associated with the tag.
	Values []string `json:"values,omitempty"`
}

// NewRiskcontrolTags instantiates a new RiskcontrolTags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskcontrolTags() *RiskcontrolTags {
	this := RiskcontrolTags{}
	return &this
}

// NewRiskcontrolTagsWithDefaults instantiates a new RiskcontrolTags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskcontrolTagsWithDefaults() *RiskcontrolTags {
	this := RiskcontrolTags{}
	return &this
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *RiskcontrolTags) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskcontrolTags) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *RiskcontrolTags) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *RiskcontrolTags) SetRuleName(v string) {
	o.RuleName = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *RiskcontrolTags) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskcontrolTags) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *RiskcontrolTags) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *RiskcontrolTags) SetTag(v string) {
	o.Tag = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *RiskcontrolTags) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskcontrolTags) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *RiskcontrolTags) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *RiskcontrolTags) SetValues(v []string) {
	o.Values = v
}

func (o RiskcontrolTags) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskcontrolTags) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RuleName) {
		toSerialize["rule_name"] = o.RuleName
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableRiskcontrolTags struct {
	value *RiskcontrolTags
	isSet bool
}

func (v NullableRiskcontrolTags) Get() *RiskcontrolTags {
	return v.value
}

func (v *NullableRiskcontrolTags) Set(val *RiskcontrolTags) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskcontrolTags) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskcontrolTags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskcontrolTags(val *RiskcontrolTags) *NullableRiskcontrolTags {
	return &NullableRiskcontrolTags{value: val, isSet: true}
}

func (v NullableRiskcontrolTags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskcontrolTags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


