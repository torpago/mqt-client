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

// checks if the AccountCreditBalanceRefundReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCreditBalanceRefundReq{}

// AccountCreditBalanceRefundReq Contains details on a credit balance refund.
type AccountCreditBalanceRefundReq struct {
	// Amount of the credit balance refund.  The maximum refund amount is the amount that brings the account balance to $0. For example, $4000 is the maximum refund amount for a -$4000 account balance.
	Amount float32 `json:"amount"`
	CurrencyCode CurrencyCode `json:"currency_code"`
	// Description for the credit balance refund.
	Description string `json:"description"`
	Method RefundMethod `json:"method"`
	// Unique identifier of the credit balance refund.
	Token *string `json:"token,omitempty"`
}

type _AccountCreditBalanceRefundReq AccountCreditBalanceRefundReq

// NewAccountCreditBalanceRefundReq instantiates a new AccountCreditBalanceRefundReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreditBalanceRefundReq(amount float32, currencyCode CurrencyCode, description string, method RefundMethod) *AccountCreditBalanceRefundReq {
	this := AccountCreditBalanceRefundReq{}
	this.Amount = amount
	this.CurrencyCode = currencyCode
	this.Description = description
	this.Method = method
	return &this
}

// NewAccountCreditBalanceRefundReqWithDefaults instantiates a new AccountCreditBalanceRefundReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreditBalanceRefundReqWithDefaults() *AccountCreditBalanceRefundReq {
	this := AccountCreditBalanceRefundReq{}
	var currencyCode CurrencyCode = CURRENCYCODE_USD
	this.CurrencyCode = currencyCode
	return &this
}

// GetAmount returns the Amount field value
func (o *AccountCreditBalanceRefundReq) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AccountCreditBalanceRefundReq) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AccountCreditBalanceRefundReq) SetAmount(v float32) {
	o.Amount = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *AccountCreditBalanceRefundReq) GetCurrencyCode() CurrencyCode {
	if o == nil {
		var ret CurrencyCode
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *AccountCreditBalanceRefundReq) GetCurrencyCodeOk() (*CurrencyCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *AccountCreditBalanceRefundReq) SetCurrencyCode(v CurrencyCode) {
	o.CurrencyCode = v
}

// GetDescription returns the Description field value
func (o *AccountCreditBalanceRefundReq) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AccountCreditBalanceRefundReq) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AccountCreditBalanceRefundReq) SetDescription(v string) {
	o.Description = v
}

// GetMethod returns the Method field value
func (o *AccountCreditBalanceRefundReq) GetMethod() RefundMethod {
	if o == nil {
		var ret RefundMethod
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *AccountCreditBalanceRefundReq) GetMethodOk() (*RefundMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *AccountCreditBalanceRefundReq) SetMethod(v RefundMethod) {
	o.Method = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AccountCreditBalanceRefundReq) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreditBalanceRefundReq) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AccountCreditBalanceRefundReq) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AccountCreditBalanceRefundReq) SetToken(v string) {
	o.Token = &v
}

func (o AccountCreditBalanceRefundReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCreditBalanceRefundReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["currency_code"] = o.CurrencyCode
	toSerialize["description"] = o.Description
	toSerialize["method"] = o.Method
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *AccountCreditBalanceRefundReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency_code",
		"description",
		"method",
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

	varAccountCreditBalanceRefundReq := _AccountCreditBalanceRefundReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountCreditBalanceRefundReq)

	if err != nil {
		return err
	}

	*o = AccountCreditBalanceRefundReq(varAccountCreditBalanceRefundReq)

	return err
}

type NullableAccountCreditBalanceRefundReq struct {
	value *AccountCreditBalanceRefundReq
	isSet bool
}

func (v NullableAccountCreditBalanceRefundReq) Get() *AccountCreditBalanceRefundReq {
	return v.value
}

func (v *NullableAccountCreditBalanceRefundReq) Set(val *AccountCreditBalanceRefundReq) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreditBalanceRefundReq) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreditBalanceRefundReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreditBalanceRefundReq(val *AccountCreditBalanceRefundReq) *NullableAccountCreditBalanceRefundReq {
	return &NullableAccountCreditBalanceRefundReq{value: val, isSet: true}
}

func (v NullableAccountCreditBalanceRefundReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreditBalanceRefundReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


