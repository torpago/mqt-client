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

// checks if the WithdrawalRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WithdrawalRequestModel{}

// WithdrawalRequestModel struct for WithdrawalRequestModel
type WithdrawalRequestModel struct {
	AccountType *string `json:"account_type,omitempty"`
	Amount float32 `json:"amount"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	CardToken string `json:"card_token"`
	Mid string `json:"mid"`
	Pin *string `json:"pin,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

type _WithdrawalRequestModel WithdrawalRequestModel

// NewWithdrawalRequestModel instantiates a new WithdrawalRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithdrawalRequestModel(amount float32, cardToken string, mid string) *WithdrawalRequestModel {
	this := WithdrawalRequestModel{}
	this.Amount = amount
	this.CardToken = cardToken
	this.Mid = mid
	return &this
}

// NewWithdrawalRequestModelWithDefaults instantiates a new WithdrawalRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithdrawalRequestModelWithDefaults() *WithdrawalRequestModel {
	this := WithdrawalRequestModel{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *WithdrawalRequestModel) GetAccountType() string {
	if o == nil || IsNil(o.AccountType) {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetAccountTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AccountType) {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *WithdrawalRequestModel) HasAccountType() bool {
	if o != nil && !IsNil(o.AccountType) {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *WithdrawalRequestModel) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAmount returns the Amount field value
func (o *WithdrawalRequestModel) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WithdrawalRequestModel) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *WithdrawalRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *WithdrawalRequestModel) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *WithdrawalRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetCardToken returns the CardToken field value
func (o *WithdrawalRequestModel) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *WithdrawalRequestModel) SetCardToken(v string) {
	o.CardToken = v
}

// GetMid returns the Mid field value
func (o *WithdrawalRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *WithdrawalRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *WithdrawalRequestModel) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *WithdrawalRequestModel) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *WithdrawalRequestModel) SetPin(v string) {
	o.Pin = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *WithdrawalRequestModel) GetWebhook() Webhook {
	if o == nil || IsNil(o.Webhook) {
		var ret Webhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithdrawalRequestModel) GetWebhookOk() (*Webhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *WithdrawalRequestModel) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given Webhook and assigns it to the Webhook field.
func (o *WithdrawalRequestModel) SetWebhook(v Webhook) {
	o.Webhook = &v
}

func (o WithdrawalRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithdrawalRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountType) {
		toSerialize["account_type"] = o.AccountType
	}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	toSerialize["card_token"] = o.CardToken
	toSerialize["mid"] = o.Mid
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

func (o *WithdrawalRequestModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"card_token",
		"mid",
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

	varWithdrawalRequestModel := _WithdrawalRequestModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWithdrawalRequestModel)

	if err != nil {
		return err
	}

	*o = WithdrawalRequestModel(varWithdrawalRequestModel)

	return err
}

type NullableWithdrawalRequestModel struct {
	value *WithdrawalRequestModel
	isSet bool
}

func (v NullableWithdrawalRequestModel) Get() *WithdrawalRequestModel {
	return v.value
}

func (v *NullableWithdrawalRequestModel) Set(val *WithdrawalRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWithdrawalRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWithdrawalRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithdrawalRequestModel(val *WithdrawalRequestModel) *NullableWithdrawalRequestModel {
	return &NullableWithdrawalRequestModel{value: val, isSet: true}
}

func (v NullableWithdrawalRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithdrawalRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


