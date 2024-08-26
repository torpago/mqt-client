/*
Marqeta Core API

Simplified management of your payment programs

API version: 3.0.0
Contact: support@marqeta.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AuthorizationClearingAtmWithdrawalEvent struct for AuthorizationClearingAtmWithdrawalEvent
type AuthorizationClearingAtmWithdrawalEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	PrecedingRelatedTransactionToken string `json:"preceding_related_transaction_token"`
}

// NewAuthorizationClearingAtmWithdrawalEvent instantiates a new AuthorizationClearingAtmWithdrawalEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationClearingAtmWithdrawalEvent(precedingRelatedTransactionToken string, amount float32) *AuthorizationClearingAtmWithdrawalEvent {
	this := AuthorizationClearingAtmWithdrawalEvent{}
	this.Amount = &amount
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	this.PrecedingRelatedTransactionToken = precedingRelatedTransactionToken
	return &this
}

// NewAuthorizationClearingAtmWithdrawalEventWithDefaults instantiates a new AuthorizationClearingAtmWithdrawalEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationClearingAtmWithdrawalEventWithDefaults() *AuthorizationClearingAtmWithdrawalEvent {
	this := AuthorizationClearingAtmWithdrawalEvent{}
	return &this
}

// GetPrecedingRelatedTransactionToken returns the PrecedingRelatedTransactionToken field value
func (o *AuthorizationClearingAtmWithdrawalEvent) GetPrecedingRelatedTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrecedingRelatedTransactionToken
}

// GetPrecedingRelatedTransactionTokenOk returns a tuple with the PrecedingRelatedTransactionToken field value
// and a boolean to check if the value has been set.
func (o *AuthorizationClearingAtmWithdrawalEvent) GetPrecedingRelatedTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrecedingRelatedTransactionToken, true
}

// SetPrecedingRelatedTransactionToken sets field value
func (o *AuthorizationClearingAtmWithdrawalEvent) SetPrecedingRelatedTransactionToken(v string) {
	o.PrecedingRelatedTransactionToken = v
}

func (o AuthorizationClearingAtmWithdrawalEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedCardTransactionSimulation, errCardTransactionSimulation := json.Marshal(o.CardTransactionSimulation)
	if errCardTransactionSimulation != nil {
		return []byte{}, errCardTransactionSimulation
	}
	errCardTransactionSimulation = json.Unmarshal([]byte(serializedCardTransactionSimulation), &toSerialize)
	if errCardTransactionSimulation != nil {
		return []byte{}, errCardTransactionSimulation
	}
	if true {
		toSerialize["preceding_related_transaction_token"] = o.PrecedingRelatedTransactionToken
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationClearingAtmWithdrawalEvent struct {
	value *AuthorizationClearingAtmWithdrawalEvent
	isSet bool
}

func (v NullableAuthorizationClearingAtmWithdrawalEvent) Get() *AuthorizationClearingAtmWithdrawalEvent {
	return v.value
}

func (v *NullableAuthorizationClearingAtmWithdrawalEvent) Set(val *AuthorizationClearingAtmWithdrawalEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationClearingAtmWithdrawalEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationClearingAtmWithdrawalEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationClearingAtmWithdrawalEvent(val *AuthorizationClearingAtmWithdrawalEvent) *NullableAuthorizationClearingAtmWithdrawalEvent {
	return &NullableAuthorizationClearingAtmWithdrawalEvent{value: val, isSet: true}
}

func (v NullableAuthorizationClearingAtmWithdrawalEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationClearingAtmWithdrawalEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


