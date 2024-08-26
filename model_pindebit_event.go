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

// PindebitEvent struct for PindebitEvent
type PindebitEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	CardToken string `json:"card_token"`
	CardAcceptor *TransactionCardAcceptor `json:"card_acceptor,omitempty"`
}

// NewPindebitEvent instantiates a new PindebitEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPindebitEvent(cardToken string, amount float32) *PindebitEvent {
	this := PindebitEvent{}
	this.CardToken = cardToken
	this.Amount = &amount
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// NewPindebitEventWithDefaults instantiates a new PindebitEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPindebitEventWithDefaults() *PindebitEvent {
	this := PindebitEvent{}
	return &this
}

// GetCardToken returns the CardToken field value
func (o *PindebitEvent) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *PindebitEvent) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *PindebitEvent) SetCardToken(v string) {
	o.CardToken = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *PindebitEvent) GetCardAcceptor() TransactionCardAcceptor {
	if o == nil || o.CardAcceptor == nil {
		var ret TransactionCardAcceptor
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PindebitEvent) GetCardAcceptorOk() (*TransactionCardAcceptor, bool) {
	if o == nil || o.CardAcceptor == nil {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *PindebitEvent) HasCardAcceptor() bool {
	if o != nil && o.CardAcceptor != nil {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given TransactionCardAcceptor and assigns it to the CardAcceptor field.
func (o *PindebitEvent) SetCardAcceptor(v TransactionCardAcceptor) {
	o.CardAcceptor = &v
}

func (o PindebitEvent) MarshalJSON() ([]byte, error) {
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

type NullablePindebitEvent struct {
	value *PindebitEvent
	isSet bool
}

func (v NullablePindebitEvent) Get() *PindebitEvent {
	return v.value
}

func (v *NullablePindebitEvent) Set(val *PindebitEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePindebitEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePindebitEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePindebitEvent(val *PindebitEvent) *NullablePindebitEvent {
	return &NullablePindebitEvent{value: val, isSet: true}
}

func (v NullablePindebitEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePindebitEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


