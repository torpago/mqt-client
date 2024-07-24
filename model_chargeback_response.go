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
	"time"
	"bytes"
	"fmt"
)

// checks if the ChargebackResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargebackResponse{}

// ChargebackResponse Contains the chargeback object associated with this transaction if a chargeback has been initiated.
type ChargebackResponse struct {
	// Amount of the chargeback.
	Amount float32 `json:"amount"`
	// Channel the chargeback came through.
	Channel string `json:"channel"`
	// Date and time when the chargeback was created. Not returned for transactions when the associated chargeback is in the `INITIATED` state.
	CreatedTime time.Time `json:"created_time"`
	// Whether to credit the user for the chargeback amount.
	CreditUser bool `json:"credit_user"`
	// Date and time when the chargeback was last modified. Not returned for transactions when the associated chargeback is in the `INITIATED` state.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Additional comments about the chargeback.
	Memo *string `json:"memo,omitempty"`
	// Network handling the chargeback.
	Network string `json:"network"`
	// Network-assigned identifier of the chargeback.
	NetworkCaseId *string `json:"network_case_id,omitempty"`
	// Identifies the standardized reason for the chargeback.
	ReasonCode *string `json:"reason_code,omitempty"`
	// State of the case.
	State string `json:"state"`
	// Unique identifier of the chargeback.
	Token string `json:"token"`
	// Unique identifier of the transaction being charged back.
	TransactionToken string `json:"transaction_token"`
}

type _ChargebackResponse ChargebackResponse

// NewChargebackResponse instantiates a new ChargebackResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargebackResponse(amount float32, channel string, createdTime time.Time, creditUser bool, lastModifiedTime time.Time, network string, state string, token string, transactionToken string) *ChargebackResponse {
	this := ChargebackResponse{}
	this.Amount = amount
	this.Channel = channel
	this.CreatedTime = createdTime
	this.CreditUser = creditUser
	this.LastModifiedTime = lastModifiedTime
	this.Network = network
	this.State = state
	this.Token = token
	this.TransactionToken = transactionToken
	return &this
}

// NewChargebackResponseWithDefaults instantiates a new ChargebackResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargebackResponseWithDefaults() *ChargebackResponse {
	this := ChargebackResponse{}
	var creditUser bool = false
	this.CreditUser = creditUser
	return &this
}

// GetAmount returns the Amount field value
func (o *ChargebackResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ChargebackResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetChannel returns the Channel field value
func (o *ChargebackResponse) GetChannel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Channel
}

// GetChannelOk returns a tuple with the Channel field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetChannelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channel, true
}

// SetChannel sets field value
func (o *ChargebackResponse) SetChannel(v string) {
	o.Channel = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *ChargebackResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *ChargebackResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCreditUser returns the CreditUser field value
func (o *ChargebackResponse) GetCreditUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CreditUser
}

// GetCreditUserOk returns a tuple with the CreditUser field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetCreditUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditUser, true
}

// SetCreditUser sets field value
func (o *ChargebackResponse) SetCreditUser(v bool) {
	o.CreditUser = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *ChargebackResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *ChargebackResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *ChargebackResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *ChargebackResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *ChargebackResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetNetwork returns the Network field value
func (o *ChargebackResponse) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetNetworkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *ChargebackResponse) SetNetwork(v string) {
	o.Network = v
}

// GetNetworkCaseId returns the NetworkCaseId field value if set, zero value otherwise.
func (o *ChargebackResponse) GetNetworkCaseId() string {
	if o == nil || IsNil(o.NetworkCaseId) {
		var ret string
		return ret
	}
	return *o.NetworkCaseId
}

// GetNetworkCaseIdOk returns a tuple with the NetworkCaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetNetworkCaseIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkCaseId) {
		return nil, false
	}
	return o.NetworkCaseId, true
}

// HasNetworkCaseId returns a boolean if a field has been set.
func (o *ChargebackResponse) HasNetworkCaseId() bool {
	if o != nil && !IsNil(o.NetworkCaseId) {
		return true
	}

	return false
}

// SetNetworkCaseId gets a reference to the given string and assigns it to the NetworkCaseId field.
func (o *ChargebackResponse) SetNetworkCaseId(v string) {
	o.NetworkCaseId = &v
}

// GetReasonCode returns the ReasonCode field value if set, zero value otherwise.
func (o *ChargebackResponse) GetReasonCode() string {
	if o == nil || IsNil(o.ReasonCode) {
		var ret string
		return ret
	}
	return *o.ReasonCode
}

// GetReasonCodeOk returns a tuple with the ReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetReasonCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReasonCode) {
		return nil, false
	}
	return o.ReasonCode, true
}

// HasReasonCode returns a boolean if a field has been set.
func (o *ChargebackResponse) HasReasonCode() bool {
	if o != nil && !IsNil(o.ReasonCode) {
		return true
	}

	return false
}

// SetReasonCode gets a reference to the given string and assigns it to the ReasonCode field.
func (o *ChargebackResponse) SetReasonCode(v string) {
	o.ReasonCode = &v
}

// GetState returns the State field value
func (o *ChargebackResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *ChargebackResponse) SetState(v string) {
	o.State = v
}

// GetToken returns the Token field value
func (o *ChargebackResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *ChargebackResponse) SetToken(v string) {
	o.Token = v
}

// GetTransactionToken returns the TransactionToken field value
func (o *ChargebackResponse) GetTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionToken
}

// GetTransactionTokenOk returns a tuple with the TransactionToken field value
// and a boolean to check if the value has been set.
func (o *ChargebackResponse) GetTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionToken, true
}

// SetTransactionToken sets field value
func (o *ChargebackResponse) SetTransactionToken(v string) {
	o.TransactionToken = v
}

func (o ChargebackResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargebackResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["channel"] = o.Channel
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["credit_user"] = o.CreditUser
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	toSerialize["network"] = o.Network
	if !IsNil(o.NetworkCaseId) {
		toSerialize["network_case_id"] = o.NetworkCaseId
	}
	if !IsNil(o.ReasonCode) {
		toSerialize["reason_code"] = o.ReasonCode
	}
	toSerialize["state"] = o.State
	toSerialize["token"] = o.Token
	toSerialize["transaction_token"] = o.TransactionToken
	return toSerialize, nil
}

func (o *ChargebackResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"channel",
		"created_time",
		"credit_user",
		"last_modified_time",
		"network",
		"state",
		"token",
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

	varChargebackResponse := _ChargebackResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChargebackResponse)

	if err != nil {
		return err
	}

	*o = ChargebackResponse(varChargebackResponse)

	return err
}

type NullableChargebackResponse struct {
	value *ChargebackResponse
	isSet bool
}

func (v NullableChargebackResponse) Get() *ChargebackResponse {
	return v.value
}

func (v *NullableChargebackResponse) Set(val *ChargebackResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChargebackResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChargebackResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargebackResponse(val *ChargebackResponse) *NullableChargebackResponse {
	return &NullableChargebackResponse{value: val, isSet: true}
}

func (v NullableChargebackResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargebackResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


