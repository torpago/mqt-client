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

// checks if the CreateRedemptionsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRedemptionsRequest{}

// CreateRedemptionsRequest struct for CreateRedemptionsRequest
type CreateRedemptionsRequest struct {
	// Amount of the reward redemption.
	Amount decimal.Decimal `json:"amount"`
	Destination *DestinationType `json:"destination,omitempty"`
	// A note explaining why the reward is being redeemed.
	Note *string `json:"note,omitempty"`
	// Unique identifier of the external account receiving the reward redemption. This token is equivalent to the <</core-api/payment-sources, payment source>> token.
	ReceivingAccountToken *string `json:"receiving_account_token,omitempty"`
	// Unique identifier of the reward redemption.
	Token *string `json:"token,omitempty"`
	Type RedemptionType `json:"type"`
}

type _CreateRedemptionsRequest CreateRedemptionsRequest

// NewCreateRedemptionsRequest instantiates a new CreateRedemptionsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRedemptionsRequest(amount decimal.Decimal, type_ RedemptionType) *CreateRedemptionsRequest {
	this := CreateRedemptionsRequest{}
	this.Amount = amount
	this.Type = type_
	return &this
}

// NewCreateRedemptionsRequestWithDefaults instantiates a new CreateRedemptionsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRedemptionsRequestWithDefaults() *CreateRedemptionsRequest {
	this := CreateRedemptionsRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CreateRedemptionsRequest) GetAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CreateRedemptionsRequest) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CreateRedemptionsRequest) SetAmount(v decimal.Decimal) {
	o.Amount = v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *CreateRedemptionsRequest) GetDestination() DestinationType {
	if o == nil || IsNil(o.Destination) {
		var ret DestinationType
		return ret
	}
	return *o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRedemptionsRequest) GetDestinationOk() (*DestinationType, bool) {
	if o == nil || IsNil(o.Destination) {
		return nil, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *CreateRedemptionsRequest) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given DestinationType and assigns it to the Destination field.
func (o *CreateRedemptionsRequest) SetDestination(v DestinationType) {
	o.Destination = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *CreateRedemptionsRequest) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRedemptionsRequest) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *CreateRedemptionsRequest) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *CreateRedemptionsRequest) SetNote(v string) {
	o.Note = &v
}

// GetReceivingAccountToken returns the ReceivingAccountToken field value if set, zero value otherwise.
func (o *CreateRedemptionsRequest) GetReceivingAccountToken() string {
	if o == nil || IsNil(o.ReceivingAccountToken) {
		var ret string
		return ret
	}
	return *o.ReceivingAccountToken
}

// GetReceivingAccountTokenOk returns a tuple with the ReceivingAccountToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRedemptionsRequest) GetReceivingAccountTokenOk() (*string, bool) {
	if o == nil || IsNil(o.ReceivingAccountToken) {
		return nil, false
	}
	return o.ReceivingAccountToken, true
}

// HasReceivingAccountToken returns a boolean if a field has been set.
func (o *CreateRedemptionsRequest) HasReceivingAccountToken() bool {
	if o != nil && !IsNil(o.ReceivingAccountToken) {
		return true
	}

	return false
}

// SetReceivingAccountToken gets a reference to the given string and assigns it to the ReceivingAccountToken field.
func (o *CreateRedemptionsRequest) SetReceivingAccountToken(v string) {
	o.ReceivingAccountToken = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CreateRedemptionsRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRedemptionsRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CreateRedemptionsRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CreateRedemptionsRequest) SetToken(v string) {
	o.Token = &v
}

// GetType returns the Type field value
func (o *CreateRedemptionsRequest) GetType() RedemptionType {
	if o == nil {
		var ret RedemptionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateRedemptionsRequest) GetTypeOk() (*RedemptionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateRedemptionsRequest) SetType(v RedemptionType) {
	o.Type = v
}

func (o CreateRedemptionsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRedemptionsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	if !IsNil(o.ReceivingAccountToken) {
		toSerialize["receiving_account_token"] = o.ReceivingAccountToken
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *CreateRedemptionsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"type",
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

	varCreateRedemptionsRequest := _CreateRedemptionsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateRedemptionsRequest)

	if err != nil {
		return err
	}

	*o = CreateRedemptionsRequest(varCreateRedemptionsRequest)

	return err
}

type NullableCreateRedemptionsRequest struct {
	value *CreateRedemptionsRequest
	isSet bool
}

func (v NullableCreateRedemptionsRequest) Get() *CreateRedemptionsRequest {
	return v.value
}

func (v *NullableCreateRedemptionsRequest) Set(val *CreateRedemptionsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRedemptionsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRedemptionsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRedemptionsRequest(val *CreateRedemptionsRequest) *NullableCreateRedemptionsRequest {
	return &NullableCreateRedemptionsRequest{value: val, isSet: true}
}

func (v NullableCreateRedemptionsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRedemptionsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


