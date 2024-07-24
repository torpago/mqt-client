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

// checks if the ClearingModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClearingModel{}

// ClearingModel struct for ClearingModel
type ClearingModel struct {
	Amount float32 `json:"amount"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	ForcePost *bool `json:"force_post,omitempty"`
	IsRefund *bool `json:"is_refund,omitempty"`
	Mid *string `json:"mid,omitempty"`
	NetworkFees []NetworkFeeModel `json:"network_fees,omitempty"`
	OriginalTransactionToken string `json:"original_transaction_token"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

type _ClearingModel ClearingModel

// NewClearingModel instantiates a new ClearingModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClearingModel(amount float32, originalTransactionToken string) *ClearingModel {
	this := ClearingModel{}
	this.Amount = amount
	var forcePost bool = false
	this.ForcePost = &forcePost
	var isRefund bool = false
	this.IsRefund = &isRefund
	this.OriginalTransactionToken = originalTransactionToken
	return &this
}

// NewClearingModelWithDefaults instantiates a new ClearingModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClearingModelWithDefaults() *ClearingModel {
	this := ClearingModel{}
	var forcePost bool = false
	this.ForcePost = &forcePost
	var isRefund bool = false
	this.IsRefund = &isRefund
	return &this
}

// GetAmount returns the Amount field value
func (o *ClearingModel) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ClearingModel) SetAmount(v float32) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *ClearingModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *ClearingModel) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *ClearingModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetForcePost returns the ForcePost field value if set, zero value otherwise.
func (o *ClearingModel) GetForcePost() bool {
	if o == nil || IsNil(o.ForcePost) {
		var ret bool
		return ret
	}
	return *o.ForcePost
}

// GetForcePostOk returns a tuple with the ForcePost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetForcePostOk() (*bool, bool) {
	if o == nil || IsNil(o.ForcePost) {
		return nil, false
	}
	return o.ForcePost, true
}

// HasForcePost returns a boolean if a field has been set.
func (o *ClearingModel) HasForcePost() bool {
	if o != nil && !IsNil(o.ForcePost) {
		return true
	}

	return false
}

// SetForcePost gets a reference to the given bool and assigns it to the ForcePost field.
func (o *ClearingModel) SetForcePost(v bool) {
	o.ForcePost = &v
}

// GetIsRefund returns the IsRefund field value if set, zero value otherwise.
func (o *ClearingModel) GetIsRefund() bool {
	if o == nil || IsNil(o.IsRefund) {
		var ret bool
		return ret
	}
	return *o.IsRefund
}

// GetIsRefundOk returns a tuple with the IsRefund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetIsRefundOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRefund) {
		return nil, false
	}
	return o.IsRefund, true
}

// HasIsRefund returns a boolean if a field has been set.
func (o *ClearingModel) HasIsRefund() bool {
	if o != nil && !IsNil(o.IsRefund) {
		return true
	}

	return false
}

// SetIsRefund gets a reference to the given bool and assigns it to the IsRefund field.
func (o *ClearingModel) SetIsRefund(v bool) {
	o.IsRefund = &v
}

// GetMid returns the Mid field value if set, zero value otherwise.
func (o *ClearingModel) GetMid() string {
	if o == nil || IsNil(o.Mid) {
		var ret string
		return ret
	}
	return *o.Mid
}

// GetMidOk returns a tuple with the Mid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetMidOk() (*string, bool) {
	if o == nil || IsNil(o.Mid) {
		return nil, false
	}
	return o.Mid, true
}

// HasMid returns a boolean if a field has been set.
func (o *ClearingModel) HasMid() bool {
	if o != nil && !IsNil(o.Mid) {
		return true
	}

	return false
}

// SetMid gets a reference to the given string and assigns it to the Mid field.
func (o *ClearingModel) SetMid(v string) {
	o.Mid = &v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *ClearingModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || IsNil(o.NetworkFees) {
		var ret []NetworkFeeModel
		return ret
	}
	return o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetNetworkFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || IsNil(o.NetworkFees) {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *ClearingModel) HasNetworkFees() bool {
	if o != nil && !IsNil(o.NetworkFees) {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *ClearingModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = v
}

// GetOriginalTransactionToken returns the OriginalTransactionToken field value
func (o *ClearingModel) GetOriginalTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalTransactionToken
}

// GetOriginalTransactionTokenOk returns a tuple with the OriginalTransactionToken field value
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetOriginalTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalTransactionToken, true
}

// SetOriginalTransactionToken sets field value
func (o *ClearingModel) SetOriginalTransactionToken(v string) {
	o.OriginalTransactionToken = v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *ClearingModel) GetWebhook() Webhook {
	if o == nil || IsNil(o.Webhook) {
		var ret Webhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClearingModel) GetWebhookOk() (*Webhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *ClearingModel) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given Webhook and assigns it to the Webhook field.
func (o *ClearingModel) SetWebhook(v Webhook) {
	o.Webhook = &v
}

func (o ClearingModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClearingModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	if !IsNil(o.ForcePost) {
		toSerialize["force_post"] = o.ForcePost
	}
	if !IsNil(o.IsRefund) {
		toSerialize["is_refund"] = o.IsRefund
	}
	if !IsNil(o.Mid) {
		toSerialize["mid"] = o.Mid
	}
	if !IsNil(o.NetworkFees) {
		toSerialize["network_fees"] = o.NetworkFees
	}
	toSerialize["original_transaction_token"] = o.OriginalTransactionToken
	if !IsNil(o.Webhook) {
		toSerialize["webhook"] = o.Webhook
	}
	return toSerialize, nil
}

func (o *ClearingModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"original_transaction_token",
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

	varClearingModel := _ClearingModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClearingModel)

	if err != nil {
		return err
	}

	*o = ClearingModel(varClearingModel)

	return err
}

type NullableClearingModel struct {
	value *ClearingModel
	isSet bool
}

func (v NullableClearingModel) Get() *ClearingModel {
	return v.value
}

func (v *NullableClearingModel) Set(val *ClearingModel) {
	v.value = val
	v.isSet = true
}

func (v NullableClearingModel) IsSet() bool {
	return v.isSet
}

func (v *NullableClearingModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClearingModel(val *ClearingModel) *NullableClearingModel {
	return &NullableClearingModel{value: val, isSet: true}
}

func (v NullableClearingModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClearingModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


