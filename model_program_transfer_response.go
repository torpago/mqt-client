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

// checks if the ProgramTransferResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramTransferResponse{}

// ProgramTransferResponse Contains information about a program transfer, which moves funds from an account holder's GPA to a program funding source.
type ProgramTransferResponse struct {
	// Amount of program transfer.
	Amount float32 `json:"amount"`
	// Unique identifier of the business account holder. Returned if `user_token` is not specified.
	BusinessToken *string `json:"business_token,omitempty"`
	// Date and time when the program transfer object was created, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// Three-digit ISO 4217 currency code.
	CurrencyCode string `json:"currency_code"`
	// Contains attributes that define characteristics of one or more fees.
	Fees []FeeDetail `json:"fees,omitempty"`
	JitFunding *JitFundingApi `json:"jit_funding,omitempty"`
	// Additional description of the program transfer.
	Memo *string `json:"memo,omitempty"`
	// Comma-delimited list of tags describing the program transfer.
	Tags *string `json:"tags,omitempty"`
	// Unique identifier of the program transfer.
	Token *string `json:"token,omitempty"`
	// Unique identifier of the transaction.
	TransactionToken string `json:"transaction_token"`
	// Unique identifier of the program transfer type.
	TypeToken string `json:"type_token"`
	// Unique identifier of the user account holder. Returned if `business_token` is not specified.
	UserToken *string `json:"user_token,omitempty"`
}

type _ProgramTransferResponse ProgramTransferResponse

// NewProgramTransferResponse instantiates a new ProgramTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramTransferResponse(amount float32, currencyCode string, transactionToken string, typeToken string) *ProgramTransferResponse {
	this := ProgramTransferResponse{}
	this.Amount = amount
	this.CurrencyCode = currencyCode
	this.TransactionToken = transactionToken
	this.TypeToken = typeToken
	return &this
}

// NewProgramTransferResponseWithDefaults instantiates a new ProgramTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramTransferResponseWithDefaults() *ProgramTransferResponse {
	this := ProgramTransferResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *ProgramTransferResponse) GetAmount() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetAmountOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *ProgramTransferResponse) SetAmount(v float32) {
	o.Amount = v
}

// GetBusinessToken returns the BusinessToken field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetBusinessToken() string {
	if o == nil || IsNil(o.BusinessToken) {
		var ret string
		return ret
	}
	return *o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetBusinessTokenOk() (*string, bool) {
	if o == nil || IsNil(o.BusinessToken) {
		return nil, false
	}
	return o.BusinessToken, true
}

// HasBusinessToken returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasBusinessToken() bool {
	if o != nil && !IsNil(o.BusinessToken) {
		return true
	}

	return false
}

// SetBusinessToken gets a reference to the given string and assigns it to the BusinessToken field.
func (o *ProgramTransferResponse) SetBusinessToken(v string) {
	o.BusinessToken = &v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *ProgramTransferResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *ProgramTransferResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *ProgramTransferResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetFees() []FeeDetail {
	if o == nil || IsNil(o.Fees) {
		var ret []FeeDetail
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetFeesOk() ([]FeeDetail, bool) {
	if o == nil || IsNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasFees() bool {
	if o != nil && !IsNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []FeeDetail and assigns it to the Fees field.
func (o *ProgramTransferResponse) SetFees(v []FeeDetail) {
	o.Fees = v
}

// GetJitFunding returns the JitFunding field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetJitFunding() JitFundingApi {
	if o == nil || IsNil(o.JitFunding) {
		var ret JitFundingApi
		return ret
	}
	return *o.JitFunding
}

// GetJitFundingOk returns a tuple with the JitFunding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetJitFundingOk() (*JitFundingApi, bool) {
	if o == nil || IsNil(o.JitFunding) {
		return nil, false
	}
	return o.JitFunding, true
}

// HasJitFunding returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasJitFunding() bool {
	if o != nil && !IsNil(o.JitFunding) {
		return true
	}

	return false
}

// SetJitFunding gets a reference to the given JitFundingApi and assigns it to the JitFunding field.
func (o *ProgramTransferResponse) SetJitFunding(v JitFundingApi) {
	o.JitFunding = &v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *ProgramTransferResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *ProgramTransferResponse) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *ProgramTransferResponse) SetToken(v string) {
	o.Token = &v
}

// GetTransactionToken returns the TransactionToken field value
func (o *ProgramTransferResponse) GetTransactionToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionToken
}

// GetTransactionTokenOk returns a tuple with the TransactionToken field value
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetTransactionTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionToken, true
}

// SetTransactionToken sets field value
func (o *ProgramTransferResponse) SetTransactionToken(v string) {
	o.TransactionToken = v
}

// GetTypeToken returns the TypeToken field value
func (o *ProgramTransferResponse) GetTypeToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TypeToken
}

// GetTypeTokenOk returns a tuple with the TypeToken field value
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetTypeTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeToken, true
}

// SetTypeToken sets field value
func (o *ProgramTransferResponse) SetTypeToken(v string) {
	o.TypeToken = v
}

// GetUserToken returns the UserToken field value if set, zero value otherwise.
func (o *ProgramTransferResponse) GetUserToken() string {
	if o == nil || IsNil(o.UserToken) {
		var ret string
		return ret
	}
	return *o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramTransferResponse) GetUserTokenOk() (*string, bool) {
	if o == nil || IsNil(o.UserToken) {
		return nil, false
	}
	return o.UserToken, true
}

// HasUserToken returns a boolean if a field has been set.
func (o *ProgramTransferResponse) HasUserToken() bool {
	if o != nil && !IsNil(o.UserToken) {
		return true
	}

	return false
}

// SetUserToken gets a reference to the given string and assigns it to the UserToken field.
func (o *ProgramTransferResponse) SetUserToken(v string) {
	o.UserToken = &v
}

func (o ProgramTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramTransferResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	if !IsNil(o.BusinessToken) {
		toSerialize["business_token"] = o.BusinessToken
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !IsNil(o.JitFunding) {
		toSerialize["jit_funding"] = o.JitFunding
	}
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	toSerialize["transaction_token"] = o.TransactionToken
	toSerialize["type_token"] = o.TypeToken
	if !IsNil(o.UserToken) {
		toSerialize["user_token"] = o.UserToken
	}
	return toSerialize, nil
}

func (o *ProgramTransferResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"currency_code",
		"transaction_token",
		"type_token",
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

	varProgramTransferResponse := _ProgramTransferResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProgramTransferResponse)

	if err != nil {
		return err
	}

	*o = ProgramTransferResponse(varProgramTransferResponse)

	return err
}

type NullableProgramTransferResponse struct {
	value *ProgramTransferResponse
	isSet bool
}

func (v NullableProgramTransferResponse) Get() *ProgramTransferResponse {
	return v.value
}

func (v *NullableProgramTransferResponse) Set(val *ProgramTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramTransferResponse(val *ProgramTransferResponse) *NullableProgramTransferResponse {
	return &NullableProgramTransferResponse{value: val, isSet: true}
}

func (v NullableProgramTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


