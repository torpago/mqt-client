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

// checks if the PushToCardRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushToCardRequest{}

// PushToCardRequest struct for PushToCardRequest
type PushToCardRequest struct {
	Address1 string `json:"address_1"`
	Address2 *string `json:"address_2,omitempty"`
	City string `json:"city"`
	Country string `json:"country"`
	Cvv string `json:"cvv"`
	ExpDate string `json:"exp_date"`
	NameOnCard string `json:"name_on_card"`
	Pan string `json:"pan"`
	PostalCode string `json:"postal_code"`
	State string `json:"state"`
	Token *string `json:"token,omitempty"`
	UserToken string `json:"user_token"`
}

type _PushToCardRequest PushToCardRequest

// NewPushToCardRequest instantiates a new PushToCardRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushToCardRequest(address1 string, city string, country string, cvv string, expDate string, nameOnCard string, pan string, postalCode string, state string, userToken string) *PushToCardRequest {
	this := PushToCardRequest{}
	this.Address1 = address1
	this.City = city
	this.Country = country
	this.Cvv = cvv
	this.ExpDate = expDate
	this.NameOnCard = nameOnCard
	this.Pan = pan
	this.PostalCode = postalCode
	this.State = state
	this.UserToken = userToken
	return &this
}

// NewPushToCardRequestWithDefaults instantiates a new PushToCardRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushToCardRequestWithDefaults() *PushToCardRequest {
	this := PushToCardRequest{}
	return &this
}

// GetAddress1 returns the Address1 field value
func (o *PushToCardRequest) GetAddress1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetAddress1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address1, true
}

// SetAddress1 sets field value
func (o *PushToCardRequest) SetAddress1(v string) {
	o.Address1 = v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *PushToCardRequest) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *PushToCardRequest) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *PushToCardRequest) SetAddress2(v string) {
	o.Address2 = &v
}

// GetCity returns the City field value
func (o *PushToCardRequest) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *PushToCardRequest) SetCity(v string) {
	o.City = v
}

// GetCountry returns the Country field value
func (o *PushToCardRequest) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *PushToCardRequest) SetCountry(v string) {
	o.Country = v
}

// GetCvv returns the Cvv field value
func (o *PushToCardRequest) GetCvv() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cvv
}

// GetCvvOk returns a tuple with the Cvv field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetCvvOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cvv, true
}

// SetCvv sets field value
func (o *PushToCardRequest) SetCvv(v string) {
	o.Cvv = v
}

// GetExpDate returns the ExpDate field value
func (o *PushToCardRequest) GetExpDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExpDate
}

// GetExpDateOk returns a tuple with the ExpDate field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetExpDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpDate, true
}

// SetExpDate sets field value
func (o *PushToCardRequest) SetExpDate(v string) {
	o.ExpDate = v
}

// GetNameOnCard returns the NameOnCard field value
func (o *PushToCardRequest) GetNameOnCard() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameOnCard
}

// GetNameOnCardOk returns a tuple with the NameOnCard field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetNameOnCardOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NameOnCard, true
}

// SetNameOnCard sets field value
func (o *PushToCardRequest) SetNameOnCard(v string) {
	o.NameOnCard = v
}

// GetPan returns the Pan field value
func (o *PushToCardRequest) GetPan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pan
}

// GetPanOk returns a tuple with the Pan field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetPanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pan, true
}

// SetPan sets field value
func (o *PushToCardRequest) SetPan(v string) {
	o.Pan = v
}

// GetPostalCode returns the PostalCode field value
func (o *PushToCardRequest) GetPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PostalCode, true
}

// SetPostalCode sets field value
func (o *PushToCardRequest) SetPostalCode(v string) {
	o.PostalCode = v
}

// GetState returns the State field value
func (o *PushToCardRequest) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PushToCardRequest) SetState(v string) {
	o.State = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PushToCardRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PushToCardRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PushToCardRequest) SetToken(v string) {
	o.Token = &v
}

// GetUserToken returns the UserToken field value
func (o *PushToCardRequest) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *PushToCardRequest) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *PushToCardRequest) SetUserToken(v string) {
	o.UserToken = v
}

func (o PushToCardRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushToCardRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address_1"] = o.Address1
	if !IsNil(o.Address2) {
		toSerialize["address_2"] = o.Address2
	}
	toSerialize["city"] = o.City
	toSerialize["country"] = o.Country
	toSerialize["cvv"] = o.Cvv
	toSerialize["exp_date"] = o.ExpDate
	toSerialize["name_on_card"] = o.NameOnCard
	toSerialize["pan"] = o.Pan
	toSerialize["postal_code"] = o.PostalCode
	toSerialize["state"] = o.State
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["user_token"] = o.UserToken
	return toSerialize, nil
}

func (o *PushToCardRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"address_1",
		"city",
		"country",
		"cvv",
		"exp_date",
		"name_on_card",
		"pan",
		"postal_code",
		"state",
		"user_token",
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

	varPushToCardRequest := _PushToCardRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPushToCardRequest)

	if err != nil {
		return err
	}

	*o = PushToCardRequest(varPushToCardRequest)

	return err
}

type NullablePushToCardRequest struct {
	value *PushToCardRequest
	isSet bool
}

func (v NullablePushToCardRequest) Get() *PushToCardRequest {
	return v.value
}

func (v *NullablePushToCardRequest) Set(val *PushToCardRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePushToCardRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePushToCardRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushToCardRequest(val *PushToCardRequest) *NullablePushToCardRequest {
	return &NullablePushToCardRequest{value: val, isSet: true}
}

func (v NullablePushToCardRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushToCardRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


