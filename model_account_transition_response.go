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

// checks if the AccountTransitionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountTransitionResponse{}

// AccountTransitionResponse Information on the credit account transition.
type AccountTransitionResponse struct {
	// Unique identifier of the credit account for which to transition a status.
	AccountToken string `json:"account_token"`
	// Date and time when the transition record was created on Marqeta's credit platform, in UTC.
	CreatedTime time.Time `json:"created_time"`
	OriginalStatus AccountStatusEnum `json:"original_status"`
	Status AccountStatusEnum `json:"status"`
	// Unique identifier of the credit account transition.
	Token string `json:"token"`
}

type _AccountTransitionResponse AccountTransitionResponse

// NewAccountTransitionResponse instantiates a new AccountTransitionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountTransitionResponse(accountToken string, createdTime time.Time, originalStatus AccountStatusEnum, status AccountStatusEnum, token string) *AccountTransitionResponse {
	this := AccountTransitionResponse{}
	this.AccountToken = accountToken
	this.CreatedTime = createdTime
	this.OriginalStatus = originalStatus
	this.Status = status
	this.Token = token
	return &this
}

// NewAccountTransitionResponseWithDefaults instantiates a new AccountTransitionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountTransitionResponseWithDefaults() *AccountTransitionResponse {
	this := AccountTransitionResponse{}
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *AccountTransitionResponse) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *AccountTransitionResponse) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *AccountTransitionResponse) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *AccountTransitionResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *AccountTransitionResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *AccountTransitionResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetOriginalStatus returns the OriginalStatus field value
func (o *AccountTransitionResponse) GetOriginalStatus() AccountStatusEnum {
	if o == nil {
		var ret AccountStatusEnum
		return ret
	}

	return o.OriginalStatus
}

// GetOriginalStatusOk returns a tuple with the OriginalStatus field value
// and a boolean to check if the value has been set.
func (o *AccountTransitionResponse) GetOriginalStatusOk() (*AccountStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalStatus, true
}

// SetOriginalStatus sets field value
func (o *AccountTransitionResponse) SetOriginalStatus(v AccountStatusEnum) {
	o.OriginalStatus = v
}

// GetStatus returns the Status field value
func (o *AccountTransitionResponse) GetStatus() AccountStatusEnum {
	if o == nil {
		var ret AccountStatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AccountTransitionResponse) GetStatusOk() (*AccountStatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AccountTransitionResponse) SetStatus(v AccountStatusEnum) {
	o.Status = v
}

// GetToken returns the Token field value
func (o *AccountTransitionResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *AccountTransitionResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *AccountTransitionResponse) SetToken(v string) {
	o.Token = v
}

func (o AccountTransitionResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountTransitionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_token"] = o.AccountToken
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["original_status"] = o.OriginalStatus
	toSerialize["status"] = o.Status
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *AccountTransitionResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_token",
		"created_time",
		"original_status",
		"status",
		"token",
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

	varAccountTransitionResponse := _AccountTransitionResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountTransitionResponse)

	if err != nil {
		return err
	}

	*o = AccountTransitionResponse(varAccountTransitionResponse)

	return err
}

type NullableAccountTransitionResponse struct {
	value *AccountTransitionResponse
	isSet bool
}

func (v NullableAccountTransitionResponse) Get() *AccountTransitionResponse {
	return v.value
}

func (v *NullableAccountTransitionResponse) Set(val *AccountTransitionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountTransitionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountTransitionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountTransitionResponse(val *AccountTransitionResponse) *NullableAccountTransitionResponse {
	return &NullableAccountTransitionResponse{value: val, isSet: true}
}

func (v NullableAccountTransitionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountTransitionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


