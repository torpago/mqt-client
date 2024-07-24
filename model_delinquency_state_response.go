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

// checks if the DelinquencyStateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DelinquencyStateResponse{}

// DelinquencyStateResponse Contains details of the delinquency state of an account.
type DelinquencyStateResponse struct {
	// Unique identifier of the credit account.
	AccountToken string `json:"account_token"`
	// One or more delinquency buckets for an account. Each delinquency bucket represents a billing cycle during which the account was delinquent.
	Buckets []DelinquencyBucketResponse `json:"buckets,omitempty"`
	// Amount that is due for the current billing cycle.
	CurrentDue decimal.Decimal `json:"current_due"`
	// Date and time when the account was last made current on the Marqeta platform, in UTC.  If the account was never delinquent, this field returns the date and time the account was created on the Marqeta platform, in UTC.  If `is_delinquent` is `true`, a null value is returned.
	DateAccountCurrent NullableTime `json:"date_account_current,omitempty"`
	// Date and time when the account last fell delinquent on the Marqeta platform, in UTC.  If `is_delinquent` is `false`, a null value is returned.
	DateAccountDelinquent NullableTime `json:"date_account_delinquent,omitempty"`
	// A value of `true` indicates that the account is currently delinquent.
	IsDelinquent bool `json:"is_delinquent"`
	// Total number of days that the account is past due.
	TotalDaysPastDue int32 `json:"total_days_past_due"`
	// Total amount that is due for the current billing cycle; the sum of `total_past_due_amount` and `current_due_amount`.
	TotalDue decimal.Decimal `json:"total_due"`
	// Total amount that is past due.
	TotalPastDue decimal.Decimal `json:"total_past_due"`
}

type _DelinquencyStateResponse DelinquencyStateResponse

// NewDelinquencyStateResponse instantiates a new DelinquencyStateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDelinquencyStateResponse(accountToken string, currentDue decimal.Decimal, isDelinquent bool, totalDaysPastDue int32, totalDue decimal.Decimal, totalPastDue decimal.Decimal) *DelinquencyStateResponse {
	this := DelinquencyStateResponse{}
	this.AccountToken = accountToken
	this.CurrentDue = currentDue
	this.IsDelinquent = isDelinquent
	this.TotalDaysPastDue = totalDaysPastDue
	this.TotalDue = totalDue
	this.TotalPastDue = totalPastDue
	return &this
}

// NewDelinquencyStateResponseWithDefaults instantiates a new DelinquencyStateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDelinquencyStateResponseWithDefaults() *DelinquencyStateResponse {
	this := DelinquencyStateResponse{}
	return &this
}

// GetAccountToken returns the AccountToken field value
func (o *DelinquencyStateResponse) GetAccountToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountToken
}

// GetAccountTokenOk returns a tuple with the AccountToken field value
// and a boolean to check if the value has been set.
func (o *DelinquencyStateResponse) GetAccountTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountToken, true
}

// SetAccountToken sets field value
func (o *DelinquencyStateResponse) SetAccountToken(v string) {
	o.AccountToken = v
}

// GetBuckets returns the Buckets field value if set, zero value otherwise.
func (o *DelinquencyStateResponse) GetBuckets() []DelinquencyBucketResponse {
	if o == nil || IsNil(o.Buckets) {
		var ret []DelinquencyBucketResponse
		return ret
	}
	return o.Buckets
}

// GetBucketsOk returns a tuple with the Buckets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DelinquencyStateResponse) GetBucketsOk() ([]DelinquencyBucketResponse, bool) {
	if o == nil || IsNil(o.Buckets) {
		return nil, false
	}
	return o.Buckets, true
}

// HasBuckets returns a boolean if a field has been set.
func (o *DelinquencyStateResponse) HasBuckets() bool {
	if o != nil && !IsNil(o.Buckets) {
		return true
	}

	return false
}

// SetBuckets gets a reference to the given []DelinquencyBucketResponse and assigns it to the Buckets field.
func (o *DelinquencyStateResponse) SetBuckets(v []DelinquencyBucketResponse) {
	o.Buckets = v
}

// GetCurrentDue returns the CurrentDue field value
func (o *DelinquencyStateResponse) GetCurrentDue() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.CurrentDue
}

// GetCurrentDueOk returns a tuple with the CurrentDue field value
// and a boolean to check if the value has been set.
func (o *DelinquencyStateResponse) GetCurrentDueOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentDue, true
}

// SetCurrentDue sets field value
func (o *DelinquencyStateResponse) SetCurrentDue(v decimal.Decimal) {
	o.CurrentDue = v
}

// GetDateAccountCurrent returns the DateAccountCurrent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DelinquencyStateResponse) GetDateAccountCurrent() time.Time {
	if o == nil || IsNil(o.DateAccountCurrent.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateAccountCurrent.Get()
}

// GetDateAccountCurrentOk returns a tuple with the DateAccountCurrent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DelinquencyStateResponse) GetDateAccountCurrentOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateAccountCurrent.Get(), o.DateAccountCurrent.IsSet()
}

// HasDateAccountCurrent returns a boolean if a field has been set.
func (o *DelinquencyStateResponse) HasDateAccountCurrent() bool {
	if o != nil && o.DateAccountCurrent.IsSet() {
		return true
	}

	return false
}

// SetDateAccountCurrent gets a reference to the given NullableTime and assigns it to the DateAccountCurrent field.
func (o *DelinquencyStateResponse) SetDateAccountCurrent(v time.Time) {
	o.DateAccountCurrent.Set(&v)
}
// SetDateAccountCurrentNil sets the value for DateAccountCurrent to be an explicit nil
func (o *DelinquencyStateResponse) SetDateAccountCurrentNil() {
	o.DateAccountCurrent.Set(nil)
}

// UnsetDateAccountCurrent ensures that no value is present for DateAccountCurrent, not even an explicit nil
func (o *DelinquencyStateResponse) UnsetDateAccountCurrent() {
	o.DateAccountCurrent.Unset()
}

// GetDateAccountDelinquent returns the DateAccountDelinquent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DelinquencyStateResponse) GetDateAccountDelinquent() time.Time {
	if o == nil || IsNil(o.DateAccountDelinquent.Get()) {
		var ret time.Time
		return ret
	}
	return *o.DateAccountDelinquent.Get()
}

// GetDateAccountDelinquentOk returns a tuple with the DateAccountDelinquent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DelinquencyStateResponse) GetDateAccountDelinquentOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.DateAccountDelinquent.Get(), o.DateAccountDelinquent.IsSet()
}

// HasDateAccountDelinquent returns a boolean if a field has been set.
func (o *DelinquencyStateResponse) HasDateAccountDelinquent() bool {
	if o != nil && o.DateAccountDelinquent.IsSet() {
		return true
	}

	return false
}

// SetDateAccountDelinquent gets a reference to the given NullableTime and assigns it to the DateAccountDelinquent field.
func (o *DelinquencyStateResponse) SetDateAccountDelinquent(v time.Time) {
	o.DateAccountDelinquent.Set(&v)
}
// SetDateAccountDelinquentNil sets the value for DateAccountDelinquent to be an explicit nil
func (o *DelinquencyStateResponse) SetDateAccountDelinquentNil() {
	o.DateAccountDelinquent.Set(nil)
}

// UnsetDateAccountDelinquent ensures that no value is present for DateAccountDelinquent, not even an explicit nil
func (o *DelinquencyStateResponse) UnsetDateAccountDelinquent() {
	o.DateAccountDelinquent.Unset()
}

// GetIsDelinquent returns the IsDelinquent field value
func (o *DelinquencyStateResponse) GetIsDelinquent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDelinquent
}

// GetIsDelinquentOk returns a tuple with the IsDelinquent field value
// and a boolean to check if the value has been set.
func (o *DelinquencyStateResponse) GetIsDelinquentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsDelinquent, true
}

// SetIsDelinquent sets field value
func (o *DelinquencyStateResponse) SetIsDelinquent(v bool) {
	o.IsDelinquent = v
}

// GetTotalDaysPastDue returns the TotalDaysPastDue field value
func (o *DelinquencyStateResponse) GetTotalDaysPastDue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalDaysPastDue
}

// GetTotalDaysPastDueOk returns a tuple with the TotalDaysPastDue field value
// and a boolean to check if the value has been set.
func (o *DelinquencyStateResponse) GetTotalDaysPastDueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDaysPastDue, true
}

// SetTotalDaysPastDue sets field value
func (o *DelinquencyStateResponse) SetTotalDaysPastDue(v int32) {
	o.TotalDaysPastDue = v
}

// GetTotalDue returns the TotalDue field value
func (o *DelinquencyStateResponse) GetTotalDue() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.TotalDue
}

// GetTotalDueOk returns a tuple with the TotalDue field value
// and a boolean to check if the value has been set.
func (o *DelinquencyStateResponse) GetTotalDueOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalDue, true
}

// SetTotalDue sets field value
func (o *DelinquencyStateResponse) SetTotalDue(v decimal.Decimal) {
	o.TotalDue = v
}

// GetTotalPastDue returns the TotalPastDue field value
func (o *DelinquencyStateResponse) GetTotalPastDue() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.TotalPastDue
}

// GetTotalPastDueOk returns a tuple with the TotalPastDue field value
// and a boolean to check if the value has been set.
func (o *DelinquencyStateResponse) GetTotalPastDueOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalPastDue, true
}

// SetTotalPastDue sets field value
func (o *DelinquencyStateResponse) SetTotalPastDue(v decimal.Decimal) {
	o.TotalPastDue = v
}

func (o DelinquencyStateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DelinquencyStateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_token"] = o.AccountToken
	if !IsNil(o.Buckets) {
		toSerialize["buckets"] = o.Buckets
	}
	toSerialize["current_due"] = o.CurrentDue
	if o.DateAccountCurrent.IsSet() {
		toSerialize["date_account_current"] = o.DateAccountCurrent.Get()
	}
	if o.DateAccountDelinquent.IsSet() {
		toSerialize["date_account_delinquent"] = o.DateAccountDelinquent.Get()
	}
	toSerialize["is_delinquent"] = o.IsDelinquent
	toSerialize["total_days_past_due"] = o.TotalDaysPastDue
	toSerialize["total_due"] = o.TotalDue
	toSerialize["total_past_due"] = o.TotalPastDue
	return toSerialize, nil
}

func (o *DelinquencyStateResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_token",
		"current_due",
		"is_delinquent",
		"total_days_past_due",
		"total_due",
		"total_past_due",
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

	varDelinquencyStateResponse := _DelinquencyStateResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDelinquencyStateResponse)

	if err != nil {
		return err
	}

	*o = DelinquencyStateResponse(varDelinquencyStateResponse)

	return err
}

type NullableDelinquencyStateResponse struct {
	value *DelinquencyStateResponse
	isSet bool
}

func (v NullableDelinquencyStateResponse) Get() *DelinquencyStateResponse {
	return v.value
}

func (v *NullableDelinquencyStateResponse) Set(val *DelinquencyStateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDelinquencyStateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDelinquencyStateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDelinquencyStateResponse(val *DelinquencyStateResponse) *NullableDelinquencyStateResponse {
	return &NullableDelinquencyStateResponse{value: val, isSet: true}
}

func (v NullableDelinquencyStateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDelinquencyStateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


