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

// checks if the CardLifeCycle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CardLifeCycle{}

// CardLifeCycle Defines characteristics of the lifecycle of cards of this card product type.
type CardLifeCycle struct {
	// A value of `true` indicates that cards of this card product type are active once they are issued.
	ActivateUponIssue *bool `json:"activate_upon_issue,omitempty"`
	// Sequence of three digits that defines various services, differentiates card usage in international or domestic interchange, designates personal identification number (PIN) and authorization requirements, and identifies card restrictions. The following values are commonly used:  *First digit*  * *1* — International interchange OK * *2* — International interchange, use IC (chip) where feasible * *5* — National interchange only except under bilateral agreement * *6* — National interchange only except under bilateral agreement, use IC (chip) where feasible * *7* — No interchange except under bilateral agreement (closed loop) * *9* — Test  *Second digit*  * *0* — Normal * *2* — Contact issuer via online means * *4* — Contact issuer via online means except under bilateral agreement  *Third digit*  * *0* — No restrictions, PIN required * *1* — No restrictions * *2* — Goods and services only (no cash) * *3* — ATM only, PIN required * *4* — Cash only * *5* — Goods and services only (no cash), PIN required * *6* — No restrictions, use PIN where feasible * *7* — Goods and services only (no cash), use PIN where feasible
	CardServiceCode *int32 `json:"card_service_code,omitempty"`
	ExpirationOffset *ExpirationOffsetWithMinimum `json:"expiration_offset,omitempty"`
	// Normally, the `expiration_offset` is measured from the date of issue. Set this field to `true` to measure `expiration_offset` from the date of activation instead.
	UpdateExpirationUponActivation *bool `json:"update_expiration_upon_activation,omitempty"`
}

// NewCardLifeCycle instantiates a new CardLifeCycle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCardLifeCycle() *CardLifeCycle {
	this := CardLifeCycle{}
	var activateUponIssue bool = false
	this.ActivateUponIssue = &activateUponIssue
	var cardServiceCode int32 = 101
	this.CardServiceCode = &cardServiceCode
	var updateExpirationUponActivation bool = false
	this.UpdateExpirationUponActivation = &updateExpirationUponActivation
	return &this
}

// NewCardLifeCycleWithDefaults instantiates a new CardLifeCycle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCardLifeCycleWithDefaults() *CardLifeCycle {
	this := CardLifeCycle{}
	var activateUponIssue bool = false
	this.ActivateUponIssue = &activateUponIssue
	var cardServiceCode int32 = 101
	this.CardServiceCode = &cardServiceCode
	var updateExpirationUponActivation bool = false
	this.UpdateExpirationUponActivation = &updateExpirationUponActivation
	return &this
}

// GetActivateUponIssue returns the ActivateUponIssue field value if set, zero value otherwise.
func (o *CardLifeCycle) GetActivateUponIssue() bool {
	if o == nil || IsNil(o.ActivateUponIssue) {
		var ret bool
		return ret
	}
	return *o.ActivateUponIssue
}

// GetActivateUponIssueOk returns a tuple with the ActivateUponIssue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardLifeCycle) GetActivateUponIssueOk() (*bool, bool) {
	if o == nil || IsNil(o.ActivateUponIssue) {
		return nil, false
	}
	return o.ActivateUponIssue, true
}

// HasActivateUponIssue returns a boolean if a field has been set.
func (o *CardLifeCycle) HasActivateUponIssue() bool {
	if o != nil && !IsNil(o.ActivateUponIssue) {
		return true
	}

	return false
}

// SetActivateUponIssue gets a reference to the given bool and assigns it to the ActivateUponIssue field.
func (o *CardLifeCycle) SetActivateUponIssue(v bool) {
	o.ActivateUponIssue = &v
}

// GetCardServiceCode returns the CardServiceCode field value if set, zero value otherwise.
func (o *CardLifeCycle) GetCardServiceCode() int32 {
	if o == nil || IsNil(o.CardServiceCode) {
		var ret int32
		return ret
	}
	return *o.CardServiceCode
}

// GetCardServiceCodeOk returns a tuple with the CardServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardLifeCycle) GetCardServiceCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.CardServiceCode) {
		return nil, false
	}
	return o.CardServiceCode, true
}

// HasCardServiceCode returns a boolean if a field has been set.
func (o *CardLifeCycle) HasCardServiceCode() bool {
	if o != nil && !IsNil(o.CardServiceCode) {
		return true
	}

	return false
}

// SetCardServiceCode gets a reference to the given int32 and assigns it to the CardServiceCode field.
func (o *CardLifeCycle) SetCardServiceCode(v int32) {
	o.CardServiceCode = &v
}

// GetExpirationOffset returns the ExpirationOffset field value if set, zero value otherwise.
func (o *CardLifeCycle) GetExpirationOffset() ExpirationOffsetWithMinimum {
	if o == nil || IsNil(o.ExpirationOffset) {
		var ret ExpirationOffsetWithMinimum
		return ret
	}
	return *o.ExpirationOffset
}

// GetExpirationOffsetOk returns a tuple with the ExpirationOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardLifeCycle) GetExpirationOffsetOk() (*ExpirationOffsetWithMinimum, bool) {
	if o == nil || IsNil(o.ExpirationOffset) {
		return nil, false
	}
	return o.ExpirationOffset, true
}

// HasExpirationOffset returns a boolean if a field has been set.
func (o *CardLifeCycle) HasExpirationOffset() bool {
	if o != nil && !IsNil(o.ExpirationOffset) {
		return true
	}

	return false
}

// SetExpirationOffset gets a reference to the given ExpirationOffsetWithMinimum and assigns it to the ExpirationOffset field.
func (o *CardLifeCycle) SetExpirationOffset(v ExpirationOffsetWithMinimum) {
	o.ExpirationOffset = &v
}

// GetUpdateExpirationUponActivation returns the UpdateExpirationUponActivation field value if set, zero value otherwise.
func (o *CardLifeCycle) GetUpdateExpirationUponActivation() bool {
	if o == nil || IsNil(o.UpdateExpirationUponActivation) {
		var ret bool
		return ret
	}
	return *o.UpdateExpirationUponActivation
}

// GetUpdateExpirationUponActivationOk returns a tuple with the UpdateExpirationUponActivation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CardLifeCycle) GetUpdateExpirationUponActivationOk() (*bool, bool) {
	if o == nil || IsNil(o.UpdateExpirationUponActivation) {
		return nil, false
	}
	return o.UpdateExpirationUponActivation, true
}

// HasUpdateExpirationUponActivation returns a boolean if a field has been set.
func (o *CardLifeCycle) HasUpdateExpirationUponActivation() bool {
	if o != nil && !IsNil(o.UpdateExpirationUponActivation) {
		return true
	}

	return false
}

// SetUpdateExpirationUponActivation gets a reference to the given bool and assigns it to the UpdateExpirationUponActivation field.
func (o *CardLifeCycle) SetUpdateExpirationUponActivation(v bool) {
	o.UpdateExpirationUponActivation = &v
}

func (o CardLifeCycle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CardLifeCycle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivateUponIssue) {
		toSerialize["activate_upon_issue"] = o.ActivateUponIssue
	}
	if !IsNil(o.CardServiceCode) {
		toSerialize["card_service_code"] = o.CardServiceCode
	}
	if !IsNil(o.ExpirationOffset) {
		toSerialize["expiration_offset"] = o.ExpirationOffset
	}
	if !IsNil(o.UpdateExpirationUponActivation) {
		toSerialize["update_expiration_upon_activation"] = o.UpdateExpirationUponActivation
	}
	return toSerialize, nil
}

type NullableCardLifeCycle struct {
	value *CardLifeCycle
	isSet bool
}

func (v NullableCardLifeCycle) Get() *CardLifeCycle {
	return v.value
}

func (v *NullableCardLifeCycle) Set(val *CardLifeCycle) {
	v.value = val
	v.isSet = true
}

func (v NullableCardLifeCycle) IsSet() bool {
	return v.isSet
}

func (v *NullableCardLifeCycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCardLifeCycle(val *CardLifeCycle) *NullableCardLifeCycle {
	return &NullableCardLifeCycle{value: val, isSet: true}
}

func (v NullableCardLifeCycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCardLifeCycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


