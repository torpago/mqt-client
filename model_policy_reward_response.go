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

// checks if the PolicyRewardResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyRewardResponse{}

// PolicyRewardResponse Contains information on a reward policy.
type PolicyRewardResponse struct {
	// Date and time when the reward policy was created on Marqeta's credit platform, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// Description of the reward policy.
	Description *string `json:"description,omitempty"`
	// Name of the reward policy.
	Name *string `json:"name,omitempty"`
	// One or more reward rules
	Rules []PolicyRewardRule `json:"rules,omitempty"`
	// Unique identifier of the reward policy.
	Token *string `json:"token,omitempty" validate:"regexp=(?!^ +$)^.+$"`
	// Date and time when the reward policy was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}

// NewPolicyRewardResponse instantiates a new PolicyRewardResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRewardResponse() *PolicyRewardResponse {
	this := PolicyRewardResponse{}
	return &this
}

// NewPolicyRewardResponseWithDefaults instantiates a new PolicyRewardResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRewardResponseWithDefaults() *PolicyRewardResponse {
	this := PolicyRewardResponse{}
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *PolicyRewardResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRewardResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *PolicyRewardResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *PolicyRewardResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyRewardResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRewardResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyRewardResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyRewardResponse) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyRewardResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRewardResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyRewardResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyRewardResponse) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *PolicyRewardResponse) GetRules() []PolicyRewardRule {
	if o == nil || IsNil(o.Rules) {
		var ret []PolicyRewardRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRewardResponse) GetRulesOk() ([]PolicyRewardRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *PolicyRewardResponse) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []PolicyRewardRule and assigns it to the Rules field.
func (o *PolicyRewardResponse) SetRules(v []PolicyRewardRule) {
	o.Rules = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PolicyRewardResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRewardResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PolicyRewardResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PolicyRewardResponse) SetToken(v string) {
	o.Token = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *PolicyRewardResponse) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRewardResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *PolicyRewardResponse) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *PolicyRewardResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

func (o PolicyRewardResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRewardResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}
	return toSerialize, nil
}

type NullablePolicyRewardResponse struct {
	value *PolicyRewardResponse
	isSet bool
}

func (v NullablePolicyRewardResponse) Get() *PolicyRewardResponse {
	return v.value
}

func (v *NullablePolicyRewardResponse) Set(val *PolicyRewardResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRewardResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRewardResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRewardResponse(val *PolicyRewardResponse) *NullablePolicyRewardResponse {
	return &NullablePolicyRewardResponse{value: val, isSet: true}
}

func (v NullablePolicyRewardResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRewardResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


