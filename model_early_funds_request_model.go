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

// checks if the EarlyFundsRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EarlyFundsRequestModel{}

// EarlyFundsRequestModel struct for EarlyFundsRequestModel
type EarlyFundsRequestModel struct {
	// Unique identifier of the ACH transfer.
	BankTransferToken *string `json:"bank_transfer_token,omitempty"`
}

// NewEarlyFundsRequestModel instantiates a new EarlyFundsRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarlyFundsRequestModel() *EarlyFundsRequestModel {
	this := EarlyFundsRequestModel{}
	return &this
}

// NewEarlyFundsRequestModelWithDefaults instantiates a new EarlyFundsRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarlyFundsRequestModelWithDefaults() *EarlyFundsRequestModel {
	this := EarlyFundsRequestModel{}
	return &this
}

// GetBankTransferToken returns the BankTransferToken field value if set, zero value otherwise.
func (o *EarlyFundsRequestModel) GetBankTransferToken() string {
	if o == nil || IsNil(o.BankTransferToken) {
		var ret string
		return ret
	}
	return *o.BankTransferToken
}

// GetBankTransferTokenOk returns a tuple with the BankTransferToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarlyFundsRequestModel) GetBankTransferTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BankTransferToken) {
		return nil, false
	}
	return o.BankTransferToken, true
}

// HasBankTransferToken returns a boolean if a field has been set.
func (o *EarlyFundsRequestModel) HasBankTransferToken() bool {
	if o != nil && !IsNil(o.BankTransferToken) {
		return true
	}

	return false
}

// SetBankTransferToken gets a reference to the given string and assigns it to the BankTransferToken field.
func (o *EarlyFundsRequestModel) SetBankTransferToken(v string) {
	o.BankTransferToken = &v
}

func (o EarlyFundsRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EarlyFundsRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BankTransferToken) {
		toSerialize["bank_transfer_token"] = o.BankTransferToken
	}
	return toSerialize, nil
}

type NullableEarlyFundsRequestModel struct {
	value *EarlyFundsRequestModel
	isSet bool
}

func (v NullableEarlyFundsRequestModel) Get() *EarlyFundsRequestModel {
	return v.value
}

func (v *NullableEarlyFundsRequestModel) Set(val *EarlyFundsRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEarlyFundsRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEarlyFundsRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarlyFundsRequestModel(val *EarlyFundsRequestModel) *NullableEarlyFundsRequestModel {
	return &NullableEarlyFundsRequestModel{value: val, isSet: true}
}

func (v NullableEarlyFundsRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarlyFundsRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


