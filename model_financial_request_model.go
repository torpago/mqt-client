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

// checks if the FinancialRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FinancialRequestModel{}

// FinancialRequestModel struct for FinancialRequestModel
type FinancialRequestModel struct {
	Amount float32 `json:"amount"`
	CardAcceptor CardAcceptorModel `json:"card_acceptor"`
	CardToken string `json:"card_token"`
	CashBackAmount *float32 `json:"cash_back_amount,omitempty"`
	IsPreAuth *bool `json:"is_pre_auth,omitempty"`
	Mid string `json:"mid"`
	Pin *string `json:"pin,omitempty"`
	TransactionOptions *TransactionOptions `json:"transaction_options,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

type _FinancialRequestModel FinancialRequestModel

// NewFinancialRequestModel instantiates a new FinancialRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFinancialRequestModel(amount float32, cardAcceptor CardAcceptorModel, cardToken string, mid string) *FinancialRequestModel {
	this := FinancialRequestModel{}
	this.Amount = amount
	this.CardAcceptor = cardAcceptor
	this.CardToken = cardToken
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	this.Mid = mid
	return &this
}

// NewFinancialRequestModelWithDefaults instantiates a new FinancialRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFinancialRequestModelWithDefaults() *FinancialRequestModel {
	this := FinancialRequestModel{}
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	return &this
}

// GetAmount returns the Amount field value
func (o *FinancialRequestModel) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FinancialRequestModel) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value
func (o *FinancialRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil {
		var ret CardAcceptorModel
		return ret
	}

	return o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardAcceptor, true
}

// SetCardAcceptor sets field value
func (o *FinancialRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = v
}

// GetCardToken returns the CardToken field value
func (o *FinancialRequestModel) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *FinancialRequestModel) SetCardToken(v string) {
	o.CardToken = v
}

// GetCashBackAmount returns the CashBackAmount field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetCashBackAmount() float32 {
	if o == nil || IsNil(o.CashBackAmount) {
		var ret float32
		return ret
	}
	return *o.CashBackAmount
}

// GetCashBackAmountOk returns a tuple with the CashBackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetCashBackAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.CashBackAmount) {
		return nil, false
	}
	return o.CashBackAmount, true
}

// HasCashBackAmount returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasCashBackAmount() bool {
	if o != nil && !IsNil(o.CashBackAmount) {
		return true
	}

	return false
}

// SetCashBackAmount gets a reference to the given float32 and assigns it to the CashBackAmount field.
func (o *FinancialRequestModel) SetCashBackAmount(v float32) {
	o.CashBackAmount = &v
}

// GetIsPreAuth returns the IsPreAuth field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetIsPreAuth() bool {
	if o == nil || IsNil(o.IsPreAuth) {
		var ret bool
		return ret
	}
	return *o.IsPreAuth
}

// GetIsPreAuthOk returns a tuple with the IsPreAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetIsPreAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPreAuth) {
		return nil, false
	}
	return o.IsPreAuth, true
}

// HasIsPreAuth returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasIsPreAuth() bool {
	if o != nil && !IsNil(o.IsPreAuth) {
		return true
	}

	return false
}

// SetIsPreAuth gets a reference to the given bool and assigns it to the IsPreAuth field.
func (o *FinancialRequestModel) SetIsPreAuth(v bool) {
	o.IsPreAuth = &v
}

// GetMid returns the Mid field value
func (o *FinancialRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *FinancialRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *FinancialRequestModel) SetPin(v string) {
	o.Pin = &v
}

// GetTransactionOptions returns the TransactionOptions field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetTransactionOptions() TransactionOptions {
	if o == nil || IsNil(o.TransactionOptions) {
		var ret TransactionOptions
		return ret
	}
	return *o.TransactionOptions
}

// GetTransactionOptionsOk returns a tuple with the TransactionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetTransactionOptionsOk() (*TransactionOptions, bool) {
	if o == nil || IsNil(o.TransactionOptions) {
		return nil, false
	}
	return o.TransactionOptions, true
}

// HasTransactionOptions returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasTransactionOptions() bool {
	if o != nil && !IsNil(o.TransactionOptions) {
		return true
	}

	return false
}

// SetTransactionOptions gets a reference to the given TransactionOptions and assigns it to the TransactionOptions field.
func (o *FinancialRequestModel) SetTransactionOptions(v TransactionOptions) {
	o.TransactionOptions = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *FinancialRequestModel) GetWebhook() Webhook {
	if o == nil || IsNil(o.Webhook) {
		var ret Webhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FinancialRequestModel) GetWebhookOk() (*Webhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *FinancialRequestModel) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given Webhook and assigns it to the Webhook field.
func (o *FinancialRequestModel) SetWebhook(v Webhook) {
	o.Webhook = &v
}

func (o FinancialRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FinancialRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["card_acceptor"] = o.CardAcceptor
	toSerialize["card_token"] = o.CardToken
	if !IsNil(o.CashBackAmount) {
		toSerialize["cash_back_amount"] = o.CashBackAmount
	}
	if !IsNil(o.IsPreAuth) {
		toSerialize["is_pre_auth"] = o.IsPreAuth
	}
	toSerialize["mid"] = o.Mid
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.TransactionOptions) {
		toSerialize["transaction_options"] = o.TransactionOptions
	}
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

func (o *FinancialRequestModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"card_acceptor",
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

	varFinancialRequestModel := _FinancialRequestModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFinancialRequestModel)

	if err != nil {
		return err
	}

	*o = FinancialRequestModel(varFinancialRequestModel)

	return err
}

type NullableFinancialRequestModel struct {
	value *FinancialRequestModel
	isSet bool
}

func (v NullableFinancialRequestModel) Get() *FinancialRequestModel {
	return v.value
}

func (v *NullableFinancialRequestModel) Set(val *FinancialRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFinancialRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFinancialRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFinancialRequestModel(val *FinancialRequestModel) *NullableFinancialRequestModel {
	return &NullableFinancialRequestModel{value: val, isSet: true}
}

func (v NullableFinancialRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFinancialRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


