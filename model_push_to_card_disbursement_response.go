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
	"github.com/shopspring/decimal"
)

// checks if the PushToCardDisbursementResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushToCardDisbursementResponse{}

// PushToCardDisbursementResponse struct for PushToCardDisbursementResponse
type PushToCardDisbursementResponse struct {
	Amount *decimal.Decimal `json:"amount,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	CurrencyCode *string `json:"currency_code,omitempty"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	Memo *string `json:"memo,omitempty"`
	PaymentInstrumentToken *string `json:"payment_instrument_token,omitempty"`
	Status *string `json:"status,omitempty"`
	Tags *string `json:"tags,omitempty"`
	Token *string `json:"token,omitempty"`
}

type _PushToCardDisbursementResponse PushToCardDisbursementResponse

// NewPushToCardDisbursementResponse instantiates a new PushToCardDisbursementResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushToCardDisbursementResponse(createdTime time.Time, lastModifiedTime time.Time) *PushToCardDisbursementResponse {
	this := PushToCardDisbursementResponse{}
	this.CreatedTime = createdTime
	this.LastModifiedTime = lastModifiedTime
	return &this
}

// NewPushToCardDisbursementResponseWithDefaults instantiates a new PushToCardDisbursementResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushToCardDisbursementResponseWithDefaults() *PushToCardDisbursementResponse {
	this := PushToCardDisbursementResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *PushToCardDisbursementResponse) GetAmount() decimal.Decimal {
	if o == nil || IsNil(o.Amount) {
		var ret decimal.Decimal
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *PushToCardDisbursementResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given decimal.Decimal and assigns it to the Amount field.
func (o *PushToCardDisbursementResponse) SetAmount(v decimal.Decimal) {
	o.Amount = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *PushToCardDisbursementResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *PushToCardDisbursementResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *PushToCardDisbursementResponse) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *PushToCardDisbursementResponse) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *PushToCardDisbursementResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *PushToCardDisbursementResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *PushToCardDisbursementResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetMemo returns the Memo field value if set, zero value otherwise.
func (o *PushToCardDisbursementResponse) GetMemo() string {
	if o == nil || IsNil(o.Memo) {
		var ret string
		return ret
	}
	return *o.Memo
}

// GetMemoOk returns a tuple with the Memo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetMemoOk() (*string, bool) {
	if o == nil || IsNil(o.Memo) {
		return nil, false
	}
	return o.Memo, true
}

// HasMemo returns a boolean if a field has been set.
func (o *PushToCardDisbursementResponse) HasMemo() bool {
	if o != nil && !IsNil(o.Memo) {
		return true
	}

	return false
}

// SetMemo gets a reference to the given string and assigns it to the Memo field.
func (o *PushToCardDisbursementResponse) SetMemo(v string) {
	o.Memo = &v
}

// GetPaymentInstrumentToken returns the PaymentInstrumentToken field value if set, zero value otherwise.
func (o *PushToCardDisbursementResponse) GetPaymentInstrumentToken() string {
	if o == nil || IsNil(o.PaymentInstrumentToken) {
		var ret string
		return ret
	}
	return *o.PaymentInstrumentToken
}

// GetPaymentInstrumentTokenOk returns a tuple with the PaymentInstrumentToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetPaymentInstrumentTokenOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentInstrumentToken) {
		return nil, false
	}
	return o.PaymentInstrumentToken, true
}

// HasPaymentInstrumentToken returns a boolean if a field has been set.
func (o *PushToCardDisbursementResponse) HasPaymentInstrumentToken() bool {
	if o != nil && !IsNil(o.PaymentInstrumentToken) {
		return true
	}

	return false
}

// SetPaymentInstrumentToken gets a reference to the given string and assigns it to the PaymentInstrumentToken field.
func (o *PushToCardDisbursementResponse) SetPaymentInstrumentToken(v string) {
	o.PaymentInstrumentToken = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PushToCardDisbursementResponse) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PushToCardDisbursementResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PushToCardDisbursementResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PushToCardDisbursementResponse) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PushToCardDisbursementResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *PushToCardDisbursementResponse) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *PushToCardDisbursementResponse) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushToCardDisbursementResponse) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *PushToCardDisbursementResponse) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *PushToCardDisbursementResponse) SetToken(v string) {
	o.Token = &v
}

func (o PushToCardDisbursementResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushToCardDisbursementResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	toSerialize["created_time"] = o.CreatedTime
	if !IsNil(o.CurrencyCode) {
		toSerialize["currency_code"] = o.CurrencyCode
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	if !IsNil(o.Memo) {
		toSerialize["memo"] = o.Memo
	}
	if !IsNil(o.PaymentInstrumentToken) {
		toSerialize["payment_instrument_token"] = o.PaymentInstrumentToken
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *PushToCardDisbursementResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created_time",
		"last_modified_time",
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

	varPushToCardDisbursementResponse := _PushToCardDisbursementResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPushToCardDisbursementResponse)

	if err != nil {
		return err
	}

	*o = PushToCardDisbursementResponse(varPushToCardDisbursementResponse)

	return err
}

type NullablePushToCardDisbursementResponse struct {
	value *PushToCardDisbursementResponse
	isSet bool
}

func (v NullablePushToCardDisbursementResponse) Get() *PushToCardDisbursementResponse {
	return v.value
}

func (v *NullablePushToCardDisbursementResponse) Set(val *PushToCardDisbursementResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePushToCardDisbursementResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePushToCardDisbursementResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushToCardDisbursementResponse(val *PushToCardDisbursementResponse) *NullablePushToCardDisbursementResponse {
	return &NullablePushToCardDisbursementResponse{value: val, isSet: true}
}

func (v NullablePushToCardDisbursementResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushToCardDisbursementResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


