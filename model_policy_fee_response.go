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

// checks if the PolicyFeeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyFeeResponse{}

// PolicyFeeResponse Contains information on a fee policy.
type PolicyFeeResponse struct {
	Account *PolicyFeeAccount `json:"account,omitempty"`
	// Date and time when the fee policy was created on Marqeta's credit platform, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// Description of the fee policy.
	Description *string `json:"description,omitempty"`
	// Name of the fee policy.
	Name *string `json:"name,omitempty"`
	// Unique identifier of the fee policy.
	Token *string `json:"token,omitempty" validate:"regexp=(?!^ +$)^.+$"`
	// Date and time when the fee policy was last updated on Marqeta's credit platform, in UTC.
	UpdatedTime *time.Time `json:"updated_time,omitempty"`
}

// NewPolicyFeeResponse instantiates a new PolicyFeeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyFeeResponse() *PolicyFeeResponse {
	this := PolicyFeeResponse{}
	return &this
}

// NewPolicyFeeResponseWithDefaults instantiates a new PolicyFeeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyFeeResponseWithDefaults() *PolicyFeeResponse {
	this := PolicyFeeResponse{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *PolicyFeeResponse) GetAccount() PolicyFeeAccount {
	if o == nil || IsNil(o.Account) {
		var ret PolicyFeeAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeResponse) GetAccountOk() (*PolicyFeeAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *PolicyFeeResponse) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given PolicyFeeAccount and assigns it to the Account field.
func (o *PolicyFeeResponse) SetAccount(v PolicyFeeAccount) {
	o.Account = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *PolicyFeeResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *PolicyFeeResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *PolicyFeeResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PolicyFeeResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PolicyFeeResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PolicyFeeResponse) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PolicyFeeResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PolicyFeeResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PolicyFeeResponse) SetName(v string) {
	o.Name = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PolicyFeeResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PolicyFeeResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PolicyFeeResponse) SetToken(v string) {
	o.Token = &v
}

// GetUpdatedTime returns the UpdatedTime field value if set, zero value otherwise.
func (o *PolicyFeeResponse) GetUpdatedTime() time.Time {
	if o == nil || IsNil(o.UpdatedTime) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedTime
}

// GetUpdatedTimeOk returns a tuple with the UpdatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyFeeResponse) GetUpdatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedTime) {
		return nil, false
	}
	return o.UpdatedTime, true
}

// HasUpdatedTime returns a boolean if a field has been set.
func (o *PolicyFeeResponse) HasUpdatedTime() bool {
	if o != nil && !IsNil(o.UpdatedTime) {
		return true
	}

	return false
}

// SetUpdatedTime gets a reference to the given time.Time and assigns it to the UpdatedTime field.
func (o *PolicyFeeResponse) SetUpdatedTime(v time.Time) {
	o.UpdatedTime = &v
}

func (o PolicyFeeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyFeeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.UpdatedTime) {
		toSerialize["updated_time"] = o.UpdatedTime
	}
	return toSerialize, nil
}

type NullablePolicyFeeResponse struct {
	value *PolicyFeeResponse
	isSet bool
}

func (v NullablePolicyFeeResponse) Get() *PolicyFeeResponse {
	return v.value
}

func (v *NullablePolicyFeeResponse) Set(val *PolicyFeeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyFeeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyFeeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyFeeResponse(val *PolicyFeeResponse) *NullablePolicyFeeResponse {
	return &NullablePolicyFeeResponse{value: val, isSet: true}
}

func (v NullablePolicyFeeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyFeeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


