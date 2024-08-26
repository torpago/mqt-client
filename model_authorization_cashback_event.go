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

// AuthorizationCashbackEvent struct for AuthorizationCashbackEvent
type AuthorizationCashbackEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	CardToken string `json:"card_token"`
	CardAcceptor *TransactionCardAcceptor `json:"card_acceptor,omitempty"`
}

// NewAuthorizationCashbackEvent instantiates a new AuthorizationCashbackEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizationCashbackEvent(cardToken string, amount float32, cashBackAmount float32) *AuthorizationCashbackEvent {
	this := AuthorizationCashbackEvent{}
	this.CardToken = cardToken
	this.Amount = &amount
	this.CashBackAmount = &cashBackAmount
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// NewAuthorizationCashbackEventWithDefaults instantiates a new AuthorizationCashbackEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizationCashbackEventWithDefaults() *AuthorizationCashbackEvent {
	this := AuthorizationCashbackEvent{}
	return &this
}

// GetCardToken returns the CardToken field value
func (o *AuthorizationCashbackEvent) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *AuthorizationCashbackEvent) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *AuthorizationCashbackEvent) SetCardToken(v string) {
	o.CardToken = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *AuthorizationCashbackEvent) GetCardAcceptor() TransactionCardAcceptor {
	if o == nil || o.CardAcceptor == nil {
		var ret TransactionCardAcceptor
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizationCashbackEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool) {
	if o == nil || o.CardAcceptor == nil {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *AuthorizationCashbackEvent) HasCardAcceptor() bool {
	if o != nil && o.CardAcceptor != nil {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given TransactionCardAcceptor and assigns it to the CardAcceptor field.
func (o *AuthorizationCashbackEvent) SetCardAcceptor(v TransactionCardAcceptor) {
	o.CardAcceptor = &v
}

func (o AuthorizationCashbackEvent) MarshalJSON() ([]byte, error) {
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
		toSerialize["card_token"] = o.CardToken
	}
	if o.CardAcceptor != nil {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	return json.Marshal(toSerialize)
}

type NullableAuthorizationCashbackEvent struct {
	value *AuthorizationCashbackEvent
	isSet bool
}

func (v NullableAuthorizationCashbackEvent) Get() *AuthorizationCashbackEvent {
	return v.value
}

func (v *NullableAuthorizationCashbackEvent) Set(val *AuthorizationCashbackEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationCashbackEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationCashbackEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationCashbackEvent(val *AuthorizationCashbackEvent) *NullableAuthorizationCashbackEvent {
	return &NullableAuthorizationCashbackEvent{value: val, isSet: true}
}

func (v NullableAuthorizationCashbackEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationCashbackEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


