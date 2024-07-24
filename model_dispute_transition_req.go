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
	"bytes"
	"fmt"
)

// checks if the DisputeTransitionReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DisputeTransitionReq{}

// DisputeTransitionReq Information about a transaction dispute update request.
type DisputeTransitionReq struct {
	// Updated amount of the dispute, based on the resolution.
	Amount float32 `json:"amount"`
	// Additional information on the dispute update (for example, a reason for the dispute update).
	Notes *string `json:"notes,omitempty"`
	Status DisputeStatus `json:"status"`
	// Unique identifier of the dispute update.
	Token *string `json:"token,omitempty"`
}

type _DisputeTransitionReq DisputeTransitionReq

// NewDisputeTransitionReq instantiates a new DisputeTransitionReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisputeTransitionReq(amount float32, status DisputeStatus) *DisputeTransitionReq {
	this := DisputeTransitionReq{}
	this.Amount = amount
	this.Status = status
	return &this
}

// NewDisputeTransitionReqWithDefaults instantiates a new DisputeTransitionReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisputeTransitionReqWithDefaults() *DisputeTransitionReq {
	this := DisputeTransitionReq{}
	return &this
}

// GetAmount returns the Amount field value
func (o *DisputeTransitionReq) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *DisputeTransitionReq) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *DisputeTransitionReq) SetAmount(v float32) {
	o.Amount = v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *DisputeTransitionReq) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeTransitionReq) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *DisputeTransitionReq) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *DisputeTransitionReq) SetNotes(v string) {
	o.Notes = &v
}

// GetStatus returns the Status field value
func (o *DisputeTransitionReq) GetStatus() DisputeStatus {
	if o == nil {
		var ret DisputeStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DisputeTransitionReq) GetStatusOk() (*DisputeStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DisputeTransitionReq) SetStatus(v DisputeStatus) {
	o.Status = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *DisputeTransitionReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisputeTransitionReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *DisputeTransitionReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *DisputeTransitionReq) SetToken(v string) {
	o.Token = &v
}

func (o DisputeTransitionReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DisputeTransitionReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["status"] = o.Status
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *DisputeTransitionReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"status",
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

	varDisputeTransitionReq := _DisputeTransitionReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDisputeTransitionReq)

	if err != nil {
		return err
	}

	*o = DisputeTransitionReq(varDisputeTransitionReq)

	return err
}

type NullableDisputeTransitionReq struct {
	value *DisputeTransitionReq
	isSet bool
}

func (v NullableDisputeTransitionReq) Get() *DisputeTransitionReq {
	return v.value
}

func (v *NullableDisputeTransitionReq) Set(val *DisputeTransitionReq) {
	v.value = val
	v.isSet = true
}

func (v NullableDisputeTransitionReq) IsSet() bool {
	return v.isSet
}

func (v *NullableDisputeTransitionReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisputeTransitionReq(val *DisputeTransitionReq) *NullableDisputeTransitionReq {
	return &NullableDisputeTransitionReq{value: val, isSet: true}
}

func (v NullableDisputeTransitionReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisputeTransitionReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


