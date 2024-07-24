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

// checks if the DigitalWalletTokenHash type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DigitalWalletTokenHash{}

// DigitalWalletTokenHash Contains identifiers of the digital wallet token resource and the card resource.
type DigitalWalletTokenHash struct {
	// Unique identifier of the card resource to use for the provisioning request.
	CardToken *string `json:"card_token,omitempty"`
	// Unique identifier of the digital wallet token resource.
	Token string `json:"token"`
}

type _DigitalWalletTokenHash DigitalWalletTokenHash

// NewDigitalWalletTokenHash instantiates a new DigitalWalletTokenHash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDigitalWalletTokenHash(token string) *DigitalWalletTokenHash {
	this := DigitalWalletTokenHash{}
	this.Token = token
	return &this
}

// NewDigitalWalletTokenHashWithDefaults instantiates a new DigitalWalletTokenHash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDigitalWalletTokenHashWithDefaults() *DigitalWalletTokenHash {
	this := DigitalWalletTokenHash{}
	return &this
}

// GetCardToken returns the CardToken field value if set, zero value otherwise.
func (o *DigitalWalletTokenHash) GetCardToken() string {
	if o == nil || IsNil(o.CardToken) {
		var ret string
		return ret
	}
	return *o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenHash) GetCardTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CardToken) {
		return nil, false
	}
	return o.CardToken, true
}

// HasCardToken returns a boolean if a field has been set.
func (o *DigitalWalletTokenHash) HasCardToken() bool {
	if o != nil && !IsNil(o.CardToken) {
		return true
	}

	return false
}

// SetCardToken gets a reference to the given string and assigns it to the CardToken field.
func (o *DigitalWalletTokenHash) SetCardToken(v string) {
	o.CardToken = &v
}

// GetToken returns the Token field value
func (o *DigitalWalletTokenHash) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DigitalWalletTokenHash) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DigitalWalletTokenHash) SetToken(v string) {
	o.Token = v
}

func (o DigitalWalletTokenHash) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DigitalWalletTokenHash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardToken) {
		toSerialize["card_token"] = o.CardToken
	}
	toSerialize["token"] = o.Token
	return toSerialize, nil
}

func (o *DigitalWalletTokenHash) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDigitalWalletTokenHash := _DigitalWalletTokenHash{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDigitalWalletTokenHash)

	if err != nil {
		return err
	}

	*o = DigitalWalletTokenHash(varDigitalWalletTokenHash)

	return err
}

type NullableDigitalWalletTokenHash struct {
	value *DigitalWalletTokenHash
	isSet bool
}

func (v NullableDigitalWalletTokenHash) Get() *DigitalWalletTokenHash {
	return v.value
}

func (v *NullableDigitalWalletTokenHash) Set(val *DigitalWalletTokenHash) {
	v.value = val
	v.isSet = true
}

func (v NullableDigitalWalletTokenHash) IsSet() bool {
	return v.isSet
}

func (v *NullableDigitalWalletTokenHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDigitalWalletTokenHash(val *DigitalWalletTokenHash) *NullableDigitalWalletTokenHash {
	return &NullableDigitalWalletTokenHash{value: val, isSet: true}
}

func (v NullableDigitalWalletTokenHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDigitalWalletTokenHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


