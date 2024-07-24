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

// checks if the FraudFeedbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FraudFeedbackRequest{}

// FraudFeedbackRequest of the fraud object
type FraudFeedbackRequest struct {
	// This is the party making a call.
	Actor string `json:"actor"`
	Amount string `json:"amount"`
	IsFraud bool `json:"is_fraud"`
	// This is the value of the status of the fraud.
	Status string `json:"status"`
	TransactionToken string `json:"transaction_token"`
}

type _FraudFeedbackRequest FraudFeedbackRequest

// NewFraudFeedbackRequest instantiates a new FraudFeedbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFraudFeedbackRequest(actor string, amount string, isFraud bool, status string, transactionToken string) *FraudFeedbackRequest {
	this := FraudFeedbackRequest{}
	this.Actor = actor
	this.Amount = amount
	this.IsFraud = isFraud
	this.Status = status
	this.TransactionToken = transactionToken
	return &this
}

// NewFraudFeedbackRequestWithDefaults instantiates a new FraudFeedbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFraudFeedbackRequestWithDefaults() *FraudFeedbackRequest {
	this := FraudFeedbackRequest{}
	return &this
}

// GetActor returns the Actor field value
func (o *FraudFeedbackRequest) GetActor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *FraudFeedbackRequest) GetActorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *FraudFeedbackRequest) SetActor(v string) {
	o.Actor = v
}

// GetAmount returns the Amount field value
func (o *FraudFeedbackRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FraudFeedbackRequest) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FraudFeedbackRequest) SetAmount(v string) {
	o.Amount = v
}

// GetIsFraud returns the IsFraud field value
func (o *FraudFeedbackRequest) GetIsFraud() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFraud
}

// GetIsFraudOk returns a tuple with the IsFraud field value
// and a boolean to check if the value has been set.
func (o *FraudFeedbackRequest) GetIsFraudOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFraud, true
}

// SetIsFraud sets field value
func (o *FraudFeedbackRequest) SetIsFraud(v bool) {
	o.IsFraud = v
}

// GetStatus returns the Status field value
func (o *FraudFeedbackRequest) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *FraudFeedbackRequest) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *FraudFeedbackRequest) SetStatus(v string) {
	o.Status = v
}

// GetTransactionToken returns the TransactionToken field value
func (o *FraudFeedbackRequest) GetTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionToken
}

// GetTransactionTokenOk returns a tuple with the TransactionToken field value
// and a boolean to check if the value has been set.
func (o *FraudFeedbackRequest) GetTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionToken, true
}

// SetTransactionToken sets field value
func (o *FraudFeedbackRequest) SetTransactionToken(v string) {
	o.TransactionToken = v
}

func (o FraudFeedbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FraudFeedbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["actor"] = o.Actor
	toSerialize["amount"] = o.Amount
	toSerialize["is_fraud"] = o.IsFraud
	toSerialize["status"] = o.Status
	toSerialize["transaction_token"] = o.TransactionToken
	return toSerialize, nil
}

func (o *FraudFeedbackRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"actor",
		"amount",
		"is_fraud",
		"status",
		"transaction_token",
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

	varFraudFeedbackRequest := _FraudFeedbackRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFraudFeedbackRequest)

	if err != nil {
		return err
	}

	*o = FraudFeedbackRequest(varFraudFeedbackRequest)

	return err
}

type NullableFraudFeedbackRequest struct {
	value *FraudFeedbackRequest
	isSet bool
}

func (v NullableFraudFeedbackRequest) Get() *FraudFeedbackRequest {
	return v.value
}

func (v *NullableFraudFeedbackRequest) Set(val *FraudFeedbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFraudFeedbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFraudFeedbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFraudFeedbackRequest(val *FraudFeedbackRequest) *NullableFraudFeedbackRequest {
	return &NullableFraudFeedbackRequest{value: val, isSet: true}
}

func (v NullableFraudFeedbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFraudFeedbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


