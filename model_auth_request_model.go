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

// checks if the AuthRequestModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthRequestModel{}

// AuthRequestModel struct for AuthRequestModel
type AuthRequestModel struct {
	Amount decimal.Decimal `json:"amount"`
	CardAcceptor *CardAcceptorModel `json:"card_acceptor,omitempty"`
	CardOptions *CardOptions `json:"card_options,omitempty"`
	CardToken string `json:"card_token"`
	CashBackAmount *decimal.Decimal `json:"cash_back_amount,omitempty"`
	IsPreAuth *bool `json:"is_pre_auth,omitempty"`
	Mid string `json:"mid"`
	NetworkFees []NetworkFeeModel `json:"network_fees,omitempty"`
	NetworkMetadata *NetworkMetadata `json:"network_metadata,omitempty"`
	Pin *string `json:"pin,omitempty"`
	TransactionOptions *TransactionOptions `json:"transaction_options,omitempty"`
	Webhook *Webhook `json:"webhook,omitempty"`
}

type _AuthRequestModel AuthRequestModel

// NewAuthRequestModel instantiates a new AuthRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthRequestModel(amount decimal.Decimal, cardToken string, mid string) *AuthRequestModel {
	this := AuthRequestModel{}
	this.Amount = amount
	this.CardToken = cardToken
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	this.Mid = mid
	return &this
}

// NewAuthRequestModelWithDefaults instantiates a new AuthRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthRequestModelWithDefaults() *AuthRequestModel {
	this := AuthRequestModel{}
	var isPreAuth bool = false
	this.IsPreAuth = &isPreAuth
	return &this
}

// GetAmount returns the Amount field value
func (o *AuthRequestModel) GetAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *AuthRequestModel) SetAmount(v decimal.Decimal) {
	o.Amount = v
}

// GetCardAcceptor returns the CardAcceptor field value if set, zero value otherwise.
func (o *AuthRequestModel) GetCardAcceptor() CardAcceptorModel {
	if o == nil || IsNil(o.CardAcceptor) {
		var ret CardAcceptorModel
		return ret
	}
	return *o.CardAcceptor
}

// GetCardAcceptorOk returns a tuple with the CardAcceptor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCardAcceptorOk() (*CardAcceptorModel, bool) {
	if o == nil || IsNil(o.CardAcceptor) {
		return nil, false
	}
	return o.CardAcceptor, true
}

// HasCardAcceptor returns a boolean if a field has been set.
func (o *AuthRequestModel) HasCardAcceptor() bool {
	if o != nil && !IsNil(o.CardAcceptor) {
		return true
	}

	return false
}

// SetCardAcceptor gets a reference to the given CardAcceptorModel and assigns it to the CardAcceptor field.
func (o *AuthRequestModel) SetCardAcceptor(v CardAcceptorModel) {
	o.CardAcceptor = &v
}

// GetCardOptions returns the CardOptions field value if set, zero value otherwise.
func (o *AuthRequestModel) GetCardOptions() CardOptions {
	if o == nil || IsNil(o.CardOptions) {
		var ret CardOptions
		return ret
	}
	return *o.CardOptions
}

// GetCardOptionsOk returns a tuple with the CardOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCardOptionsOk() (*CardOptions, bool) {
	if o == nil || IsNil(o.CardOptions) {
		return nil, false
	}
	return o.CardOptions, true
}

// HasCardOptions returns a boolean if a field has been set.
func (o *AuthRequestModel) HasCardOptions() bool {
	if o != nil && !IsNil(o.CardOptions) {
		return true
	}

	return false
}

// SetCardOptions gets a reference to the given CardOptions and assigns it to the CardOptions field.
func (o *AuthRequestModel) SetCardOptions(v CardOptions) {
	o.CardOptions = &v
}

// GetCardToken returns the CardToken field value
func (o *AuthRequestModel) GetCardToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardToken
}

// GetCardTokenOk returns a tuple with the CardToken field value
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCardTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardToken, true
}

// SetCardToken sets field value
func (o *AuthRequestModel) SetCardToken(v string) {
	o.CardToken = v
}

// GetCashBackAmount returns the CashBackAmount field value if set, zero value otherwise.
func (o *AuthRequestModel) GetCashBackAmount() decimal.Decimal {
	if o == nil || IsNil(o.CashBackAmount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.CashBackAmount
}

// GetCashBackAmountOk returns a tuple with the CashBackAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetCashBackAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.CashBackAmount) {
		return nil, false
	}
	return o.CashBackAmount, true
}

// HasCashBackAmount returns a boolean if a field has been set.
func (o *AuthRequestModel) HasCashBackAmount() bool {
	if o != nil && !IsNil(o.CashBackAmount) {
		return true
	}

	return false
}

// SetCashBackAmount gets a reference to the given decimal.Decimal and assigns it to the CashBackAmount field.
func (o *AuthRequestModel) SetCashBackAmount(v decimal.Decimal) {
	o.CashBackAmount = &v
}

// GetIsPreAuth returns the IsPreAuth field value if set, zero value otherwise.
func (o *AuthRequestModel) GetIsPreAuth() bool {
	if o == nil || IsNil(o.IsPreAuth) {
		var ret bool
		return ret
	}
	return *o.IsPreAuth
}

// GetIsPreAuthOk returns a tuple with the IsPreAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetIsPreAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPreAuth) {
		return nil, false
	}
	return o.IsPreAuth, true
}

// HasIsPreAuth returns a boolean if a field has been set.
func (o *AuthRequestModel) HasIsPreAuth() bool {
	if o != nil && !IsNil(o.IsPreAuth) {
		return true
	}

	return false
}

// SetIsPreAuth gets a reference to the given bool and assigns it to the IsPreAuth field.
func (o *AuthRequestModel) SetIsPreAuth(v bool) {
	o.IsPreAuth = &v
}

// GetMid returns the Mid field value
func (o *AuthRequestModel) GetMid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mid
}

// GetMidOk returns a tuple with the Mid field value
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetMidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mid, true
}

// SetMid sets field value
func (o *AuthRequestModel) SetMid(v string) {
	o.Mid = v
}

// GetNetworkFees returns the NetworkFees field value if set, zero value otherwise.
func (o *AuthRequestModel) GetNetworkFees() []NetworkFeeModel {
	if o == nil || IsNil(o.NetworkFees) {
		var ret []NetworkFeeModel
		return ret
	}
	return o.NetworkFees
}

// GetNetworkFeesOk returns a tuple with the NetworkFees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetNetworkFeesOk() ([]NetworkFeeModel, bool) {
	if o == nil || IsNil(o.NetworkFees) {
		return nil, false
	}
	return o.NetworkFees, true
}

// HasNetworkFees returns a boolean if a field has been set.
func (o *AuthRequestModel) HasNetworkFees() bool {
	if o != nil && !IsNil(o.NetworkFees) {
		return true
	}

	return false
}

// SetNetworkFees gets a reference to the given []NetworkFeeModel and assigns it to the NetworkFees field.
func (o *AuthRequestModel) SetNetworkFees(v []NetworkFeeModel) {
	o.NetworkFees = v
}

// GetNetworkMetadata returns the NetworkMetadata field value if set, zero value otherwise.
func (o *AuthRequestModel) GetNetworkMetadata() NetworkMetadata {
	if o == nil || IsNil(o.NetworkMetadata) {
		var ret NetworkMetadata
		return ret
	}
	return *o.NetworkMetadata
}

// GetNetworkMetadataOk returns a tuple with the NetworkMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetNetworkMetadataOk() (*NetworkMetadata, bool) {
	if o == nil || IsNil(o.NetworkMetadata) {
		return nil, false
	}
	return o.NetworkMetadata, true
}

// HasNetworkMetadata returns a boolean if a field has been set.
func (o *AuthRequestModel) HasNetworkMetadata() bool {
	if o != nil && !IsNil(o.NetworkMetadata) {
		return true
	}

	return false
}

// SetNetworkMetadata gets a reference to the given NetworkMetadata and assigns it to the NetworkMetadata field.
func (o *AuthRequestModel) SetNetworkMetadata(v NetworkMetadata) {
	o.NetworkMetadata = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *AuthRequestModel) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *AuthRequestModel) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *AuthRequestModel) SetPin(v string) {
	o.Pin = &v
}

// GetTransactionOptions returns the TransactionOptions field value if set, zero value otherwise.
func (o *AuthRequestModel) GetTransactionOptions() TransactionOptions {
	if o == nil || IsNil(o.TransactionOptions) {
		var ret TransactionOptions
		return ret
	}
	return *o.TransactionOptions
}

// GetTransactionOptionsOk returns a tuple with the TransactionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetTransactionOptionsOk() (*TransactionOptions, bool) {
	if o == nil || IsNil(o.TransactionOptions) {
		return nil, false
	}
	return o.TransactionOptions, true
}

// HasTransactionOptions returns a boolean if a field has been set.
func (o *AuthRequestModel) HasTransactionOptions() bool {
	if o != nil && !IsNil(o.TransactionOptions) {
		return true
	}

	return false
}

// SetTransactionOptions gets a reference to the given TransactionOptions and assigns it to the TransactionOptions field.
func (o *AuthRequestModel) SetTransactionOptions(v TransactionOptions) {
	o.TransactionOptions = &v
}

// GetWebhook returns the Webhook field value if set, zero value otherwise.
func (o *AuthRequestModel) GetWebhook() Webhook {
	if o == nil || IsNil(o.Webhook) {
		var ret Webhook
		return ret
	}
	return *o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthRequestModel) GetWebhookOk() (*Webhook, bool) {
	if o == nil || IsNil(o.Webhook) {
		return nil, false
	}
	return o.Webhook, true
}

// HasWebhook returns a boolean if a field has been set.
func (o *AuthRequestModel) HasWebhook() bool {
	if o != nil && !IsNil(o.Webhook) {
		return true
	}

	return false
}

// SetWebhook gets a reference to the given Webhook and assigns it to the Webhook field.
func (o *AuthRequestModel) SetWebhook(v Webhook) {
	o.Webhook = &v
}

func (o AuthRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthRequestModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.CardAcceptor) {
		toSerialize["card_acceptor"] = o.CardAcceptor
	}
	if !IsNil(o.CardOptions) {
		toSerialize["card_options"] = o.CardOptions
	}
	toSerialize["card_token"] = o.CardToken
	if !IsNil(o.CashBackAmount) {
		toSerialize["cash_back_amount"] = o.CashBackAmount
	}
	if !IsNil(o.IsPreAuth) {
		toSerialize["is_pre_auth"] = o.IsPreAuth
	}
	toSerialize["mid"] = o.Mid
	if !IsNil(o.NetworkFees) {
		toSerialize["network_fees"] = o.NetworkFees
	}
	if !IsNil(o.NetworkMetadata) {
		toSerialize["network_metadata"] = o.NetworkMetadata
	}
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

func (o *AuthRequestModel) UnmarshalJSON(data []byte) (err error) {
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

	varAuthRequestModel := _AuthRequestModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthRequestModel)

	if err != nil {
		return err
	}

	*o = AuthRequestModel(varAuthRequestModel)

	return err
}

type NullableAuthRequestModel struct {
	value *AuthRequestModel
	isSet bool
}

func (v NullableAuthRequestModel) Get() *AuthRequestModel {
	return v.value
}

func (v *NullableAuthRequestModel) Set(val *AuthRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthRequestModel(val *AuthRequestModel) *NullableAuthRequestModel {
	return &NullableAuthRequestModel{value: val, isSet: true}
}

func (v NullableAuthRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


