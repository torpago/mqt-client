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

// OriginalCreditAuthPlusCaptureEvent struct for OriginalCreditAuthPlusCaptureEvent
type OriginalCreditAuthPlusCaptureEvent struct {
	CardTransactionSimulation
	// Unique identifier of the card. Useful when a single account holder has multiple cards.
	CardToken string `json:"card_token"`
}

// NewOriginalCreditAuthPlusCaptureEvent instantiates a new OriginalCreditAuthPlusCaptureEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginalCreditAuthPlusCaptureEvent(cardToken string, amount float32) *OriginalCreditAuthPlusCaptureEvent {
	this := OriginalCreditAuthPlusCaptureEvent{}
	this.CardToken = cardToken
	this.Amount = &amount
	var isPreauthorization bool = false
	this.IsPreauthorization = &isPreauthorization
	var forcePost bool = false
	this.ForcePost = &forcePost
	return &this
}

// NewOriginalCreditAuthPlusCaptureEventWithDefaults instantiates a new OriginalCreditAuthPlusCaptureEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginalCreditAuthPlusCaptureEventWithDefaults() *OriginalCreditAuthPlusCaptureEvent {
	this := OriginalCreditAuthPlusCaptureEvent{}
	return &this
}

// GetCardToken returns the CardToken field value
func (o *OriginalCreditAuthPlusCaptureEvent) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *OriginalCreditAuthPlusCaptureEvent) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *OriginalCreditAuthPlusCaptureEvent) SetCardToken(v string) {
	o.CardToken = v
}

func (o OriginalCreditAuthPlusCaptureEvent) MarshalJSON() ([]byte, error) {
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
	return json.Marshal(toSerialize)
}

type NullableOriginalCreditAuthPlusCaptureEvent struct {
	value *OriginalCreditAuthPlusCaptureEvent
	isSet bool
}

func (v NullableOriginalCreditAuthPlusCaptureEvent) Get() *OriginalCreditAuthPlusCaptureEvent {
	return v.value
}

func (v *NullableOriginalCreditAuthPlusCaptureEvent) Set(val *OriginalCreditAuthPlusCaptureEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginalCreditAuthPlusCaptureEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginalCreditAuthPlusCaptureEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginalCreditAuthPlusCaptureEvent(val *OriginalCreditAuthPlusCaptureEvent) *NullableOriginalCreditAuthPlusCaptureEvent {
	return &NullableOriginalCreditAuthPlusCaptureEvent{value: val, isSet: true}
}

func (v NullableOriginalCreditAuthPlusCaptureEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginalCreditAuthPlusCaptureEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


