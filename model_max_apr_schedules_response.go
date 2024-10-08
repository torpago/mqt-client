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

// checks if the MaxAPRSchedulesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaxAPRSchedulesResponse{}

// MaxAPRSchedulesResponse struct for MaxAPRSchedulesResponse
type MaxAPRSchedulesResponse struct {
	// Date and time when the override APR ends, in UTC.
	EndDate time.Time `json:"end_date"`
	// Reason for the override APR.
	Reason string `json:"reason"`
	// Date and time when the override APR goes into effect, in UTC.
	StartDate time.Time `json:"start_date"`
	// The APR percentage value. This is the value of the fixed rate during the override period. The APR value must adhere to the constraints of the main schedule, such as maximum allowable values.
	Value decimal.Decimal `json:"value"`
}

type _MaxAPRSchedulesResponse MaxAPRSchedulesResponse

// NewMaxAPRSchedulesResponse instantiates a new MaxAPRSchedulesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaxAPRSchedulesResponse(endDate time.Time, reason string, startDate time.Time, value decimal.Decimal) *MaxAPRSchedulesResponse {
	this := MaxAPRSchedulesResponse{}
	this.EndDate = endDate
	this.Reason = reason
	this.StartDate = startDate
	this.Value = value
	return &this
}

// NewMaxAPRSchedulesResponseWithDefaults instantiates a new MaxAPRSchedulesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaxAPRSchedulesResponseWithDefaults() *MaxAPRSchedulesResponse {
	this := MaxAPRSchedulesResponse{}
	return &this
}

// GetEndDate returns the EndDate field value
func (o *MaxAPRSchedulesResponse) GetEndDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value
// and a boolean to check if the value has been set.
func (o *MaxAPRSchedulesResponse) GetEndDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndDate, true
}

// SetEndDate sets field value
func (o *MaxAPRSchedulesResponse) SetEndDate(v time.Time) {
	o.EndDate = v
}

// GetReason returns the Reason field value
func (o *MaxAPRSchedulesResponse) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *MaxAPRSchedulesResponse) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *MaxAPRSchedulesResponse) SetReason(v string) {
	o.Reason = v
}

// GetStartDate returns the StartDate field value
func (o *MaxAPRSchedulesResponse) GetStartDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value
// and a boolean to check if the value has been set.
func (o *MaxAPRSchedulesResponse) GetStartDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDate, true
}

// SetStartDate sets field value
func (o *MaxAPRSchedulesResponse) SetStartDate(v time.Time) {
	o.StartDate = v
}

// GetValue returns the Value field value
func (o *MaxAPRSchedulesResponse) GetValue() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MaxAPRSchedulesResponse) GetValueOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MaxAPRSchedulesResponse) SetValue(v decimal.Decimal) {
	o.Value = v
}

func (o MaxAPRSchedulesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaxAPRSchedulesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end_date"] = o.EndDate
	toSerialize["reason"] = o.Reason
	toSerialize["start_date"] = o.StartDate
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *MaxAPRSchedulesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"end_date",
		"reason",
		"start_date",
		"value",
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

	varMaxAPRSchedulesResponse := _MaxAPRSchedulesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMaxAPRSchedulesResponse)

	if err != nil {
		return err
	}

	*o = MaxAPRSchedulesResponse(varMaxAPRSchedulesResponse)

	return err
}

type NullableMaxAPRSchedulesResponse struct {
	value *MaxAPRSchedulesResponse
	isSet bool
}

func (v NullableMaxAPRSchedulesResponse) Get() *MaxAPRSchedulesResponse {
	return v.value
}

func (v *NullableMaxAPRSchedulesResponse) Set(val *MaxAPRSchedulesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMaxAPRSchedulesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMaxAPRSchedulesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaxAPRSchedulesResponse(val *MaxAPRSchedulesResponse) *NullableMaxAPRSchedulesResponse {
	return &NullableMaxAPRSchedulesResponse{value: val, isSet: true}
}

func (v NullableMaxAPRSchedulesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaxAPRSchedulesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


