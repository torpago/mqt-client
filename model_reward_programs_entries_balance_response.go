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

// checks if the RewardProgramsEntriesBalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewardProgramsEntriesBalanceResponse{}

// RewardProgramsEntriesBalanceResponse struct for RewardProgramsEntriesBalanceResponse
type RewardProgramsEntriesBalanceResponse struct {
	// Date and time the reward entries balance was created on the Marqeta platform, in UTC.
	CreatedDate *time.Time `json:"created_date,omitempty"`
	// The ending date (or date-time) of a date range from which to return accrued rewards, in UTC. Reward entries created on or before this date count toward the total reward balance.
	EndDate time.Time `json:"end_date"`
	// Unique identifier of the reward program for which to retrieve the reward entries balance.
	RewardProgramToken string `json:"reward_program_token"`
	// The starting date (or date-time) of a date range from which to return accrued rewards, in UTC. Reward entries created on or after this date count toward the total reward balance.
	StartDate time.Time `json:"start_date"`
	// The total balance of rewards accrued within a date range.
	TotalRewardBalance decimal.Decimal `json:"total_reward_balance"`
}

type _RewardProgramsEntriesBalanceResponse RewardProgramsEntriesBalanceResponse

// NewRewardProgramsEntriesBalanceResponse instantiates a new RewardProgramsEntriesBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewardProgramsEntriesBalanceResponse(endDate time.Time, rewardProgramToken string, startDate time.Time, totalRewardBalance decimal.Decimal) *RewardProgramsEntriesBalanceResponse {
	this := RewardProgramsEntriesBalanceResponse{}
	this.EndDate = endDate
	this.RewardProgramToken = rewardProgramToken
	this.StartDate = startDate
	this.TotalRewardBalance = totalRewardBalance
	return &this
}

// NewRewardProgramsEntriesBalanceResponseWithDefaults instantiates a new RewardProgramsEntriesBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewardProgramsEntriesBalanceResponseWithDefaults() *RewardProgramsEntriesBalanceResponse {
	this := RewardProgramsEntriesBalanceResponse{}
	return &this
}

// GetCreatedDate returns the CreatedDate field value if set, zero value otherwise.
func (o *RewardProgramsEntriesBalanceResponse) GetCreatedDate() time.Time {
	if o == nil || IsNil(o.CreatedDate) {
		var ret time.Time
		return ret
	}
	return *o.CreatedDate
}

// GetCreatedDateOk returns a tuple with the CreatedDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewardProgramsEntriesBalanceResponse) GetCreatedDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedDate) {
		return nil, false
	}
	return o.CreatedDate, true
}

// HasCreatedDate returns a boolean if a field has been set.
func (o *RewardProgramsEntriesBalanceResponse) HasCreatedDate() bool {
	if o != nil && !IsNil(o.CreatedDate) {
		return true
	}

	return false
}

// SetCreatedDate gets a reference to the given time.Time and assigns it to the CreatedDate field.
func (o *RewardProgramsEntriesBalanceResponse) SetCreatedDate(v time.Time) {
	o.CreatedDate = &v
}

// GetEndDate returns the EndDate field value
func (o *RewardProgramsEntriesBalanceResponse) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsEntriesBalanceResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *RewardProgramsEntriesBalanceResponse) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetRewardProgramToken returns the RewardProgramToken field value
func (o *RewardProgramsEntriesBalanceResponse) GetRewardProgramToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RewardProgramToken
}

// GetRewardProgramTokenOk returns a tuple with the RewardProgramToken field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsEntriesBalanceResponse) GetRewardProgramTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RewardProgramToken, true
}

// SetRewardProgramToken sets field value
func (o *RewardProgramsEntriesBalanceResponse) SetRewardProgramToken(v string) {
	o.RewardProgramToken = v
}

// GetStartDate returns the StartDate field value
func (o *RewardProgramsEntriesBalanceResponse) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsEntriesBalanceResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *RewardProgramsEntriesBalanceResponse) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetTotalRewardBalance returns the TotalRewardBalance field value
func (o *RewardProgramsEntriesBalanceResponse) GetTotalRewardBalance() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.TotalRewardBalance
}

// GetTotalRewardBalanceOk returns a tuple with the TotalRewardBalance field value
// and a boolean to check if the value has been set.
func (o *RewardProgramsEntriesBalanceResponse) GetTotalRewardBalanceOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalRewardBalance, true
}

// SetTotalRewardBalance sets field value
func (o *RewardProgramsEntriesBalanceResponse) SetTotalRewardBalance(v decimal.Decimal) {
	o.TotalRewardBalance = v
}

func (o RewardProgramsEntriesBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewardProgramsEntriesBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedDate) {
		toSerialize["created_date"] = o.CreatedDate
	}
	toSerialize["end_date"] = o.EndDate
	toSerialize["reward_program_token"] = o.RewardProgramToken
	toSerialize["start_date"] = o.StartDate
	toSerialize["total_reward_balance"] = o.TotalRewardBalance
	return toSerialize, nil
}

func (o *RewardProgramsEntriesBalanceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end_date",
		"reward_program_token",
		"start_date",
		"total_reward_balance",
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

	varRewardProgramsEntriesBalanceResponse := _RewardProgramsEntriesBalanceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRewardProgramsEntriesBalanceResponse)

	if err != nil {
		return err
	}

	*o = RewardProgramsEntriesBalanceResponse(varRewardProgramsEntriesBalanceResponse)

	return err
}

type NullableRewardProgramsEntriesBalanceResponse struct {
	value *RewardProgramsEntriesBalanceResponse
	isSet bool
}

func (v NullableRewardProgramsEntriesBalanceResponse) Get() *RewardProgramsEntriesBalanceResponse {
	return v.value
}

func (v *NullableRewardProgramsEntriesBalanceResponse) Set(val *RewardProgramsEntriesBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRewardProgramsEntriesBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRewardProgramsEntriesBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewardProgramsEntriesBalanceResponse(val *RewardProgramsEntriesBalanceResponse) *NullableRewardProgramsEntriesBalanceResponse {
	return &NullableRewardProgramsEntriesBalanceResponse{value: val, isSet: true}
}

func (v NullableRewardProgramsEntriesBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewardProgramsEntriesBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


