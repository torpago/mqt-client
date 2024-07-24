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

// checks if the RedemptionsPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedemptionsPage{}

// RedemptionsPage Return paginated redemptions for a reward program.
type RedemptionsPage struct {
	// Contains one or more redemptions.
	Data []RedemptionsResponse `json:"data,omitempty"`
	// Number of resources returned.
	Count int32 `json:"count"`
	// Sort order index of the last resource in the returned array.
	EndIndex int64 `json:"end_index"`
	// A value of `true` indicates that more unreturned resources exist.
	IsMore bool `json:"is_more"`
	// Sort order index of the first resource in the returned array.
	StartIndex int64 `json:"start_index"`
}

type _RedemptionsPage RedemptionsPage

// NewRedemptionsPage instantiates a new RedemptionsPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedemptionsPage(count int32, endIndex int64, isMore bool, startIndex int64) *RedemptionsPage {
	this := RedemptionsPage{}
	this.Count = count
	this.EndIndex = endIndex
	this.IsMore = isMore
	this.StartIndex = startIndex
	return &this
}

// NewRedemptionsPageWithDefaults instantiates a new RedemptionsPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedemptionsPageWithDefaults() *RedemptionsPage {
	this := RedemptionsPage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RedemptionsPage) GetData() []RedemptionsResponse {
	if o == nil || IsNil(o.Data) {
		var ret []RedemptionsResponse
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedemptionsPage) GetDataOk() ([]RedemptionsResponse, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RedemptionsPage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RedemptionsResponse and assigns it to the Data field.
func (o *RedemptionsPage) SetData(v []RedemptionsResponse) {
	o.Data = v
}

// GetCount returns the Count field value
func (o *RedemptionsPage) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *RedemptionsPage) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *RedemptionsPage) SetCount(v int32) {
	o.Count = v
}

// GetEndIndex returns the EndIndex field value
func (o *RedemptionsPage) GetEndIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value
// and a boolean to check if the value has been set.
func (o *RedemptionsPage) GetEndIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndIndex, true
}

// SetEndIndex sets field value
func (o *RedemptionsPage) SetEndIndex(v int64) {
	o.EndIndex = v
}

// GetIsMore returns the IsMore field value
func (o *RedemptionsPage) GetIsMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMore
}

// GetIsMoreOk returns a tuple with the IsMore field value
// and a boolean to check if the value has been set.
func (o *RedemptionsPage) GetIsMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMore, true
}

// SetIsMore sets field value
func (o *RedemptionsPage) SetIsMore(v bool) {
	o.IsMore = v
}

// GetStartIndex returns the StartIndex field value
func (o *RedemptionsPage) GetStartIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *RedemptionsPage) GetStartIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *RedemptionsPage) SetStartIndex(v int64) {
	o.StartIndex = v
}

func (o RedemptionsPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedemptionsPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["count"] = o.Count
	toSerialize["end_index"] = o.EndIndex
	toSerialize["is_more"] = o.IsMore
	toSerialize["start_index"] = o.StartIndex
	return toSerialize, nil
}

func (o *RedemptionsPage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"end_index",
		"is_more",
		"start_index",
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

	varRedemptionsPage := _RedemptionsPage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRedemptionsPage)

	if err != nil {
		return err
	}

	*o = RedemptionsPage(varRedemptionsPage)

	return err
}

type NullableRedemptionsPage struct {
	value *RedemptionsPage
	isSet bool
}

func (v NullableRedemptionsPage) Get() *RedemptionsPage {
	return v.value
}

func (v *NullableRedemptionsPage) Set(val *RedemptionsPage) {
	v.value = val
	v.isSet = true
}

func (v NullableRedemptionsPage) IsSet() bool {
	return v.isSet
}

func (v *NullableRedemptionsPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedemptionsPage(val *RedemptionsPage) *NullableRedemptionsPage {
	return &NullableRedemptionsPage{value: val, isSet: true}
}

func (v NullableRedemptionsPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedemptionsPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


