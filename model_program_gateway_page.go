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

// checks if the ProgramGatewayPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramGatewayPage{}

// ProgramGatewayPage Returns paginated Program Gateways.
type ProgramGatewayPage struct {
	// Number of resources returned.
	Count int32 `json:"count"`
	// Contains one or more Program Gateway objects.
	Data []ProgramGatewayResponse `json:"data"`
	// Sort order index of the last resource in the returned array.
	EndIndex int32 `json:"end_index"`
	// A value of `true` indicates that more unreturned resources exist.
	IsMore bool `json:"is_more"`
	// Sort order index of the first resource in the returned array.
	StartIndex int32 `json:"start_index"`
}

type _ProgramGatewayPage ProgramGatewayPage

// NewProgramGatewayPage instantiates a new ProgramGatewayPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramGatewayPage(count int32, data []ProgramGatewayResponse, endIndex int32, isMore bool, startIndex int32) *ProgramGatewayPage {
	this := ProgramGatewayPage{}
	this.Count = count
	this.Data = data
	this.EndIndex = endIndex
	this.IsMore = isMore
	this.StartIndex = startIndex
	return &this
}

// NewProgramGatewayPageWithDefaults instantiates a new ProgramGatewayPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramGatewayPageWithDefaults() *ProgramGatewayPage {
	this := ProgramGatewayPage{}
	return &this
}

// GetCount returns the Count field value
func (o *ProgramGatewayPage) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ProgramGatewayPage) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ProgramGatewayPage) SetCount(v int32) {
	o.Count = v
}

// GetData returns the Data field value
func (o *ProgramGatewayPage) GetData() []ProgramGatewayResponse {
	if o == nil {
		var ret []ProgramGatewayResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProgramGatewayPage) GetDataOk() ([]ProgramGatewayResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ProgramGatewayPage) SetData(v []ProgramGatewayResponse) {
	o.Data = v
}

// GetEndIndex returns the EndIndex field value
func (o *ProgramGatewayPage) GetEndIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value
// and a boolean to check if the value has been set.
func (o *ProgramGatewayPage) GetEndIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndIndex, true
}

// SetEndIndex sets field value
func (o *ProgramGatewayPage) SetEndIndex(v int32) {
	o.EndIndex = v
}

// GetIsMore returns the IsMore field value
func (o *ProgramGatewayPage) GetIsMore() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMore
}

// GetIsMoreOk returns a tuple with the IsMore field value
// and a boolean to check if the value has been set.
func (o *ProgramGatewayPage) GetIsMoreOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMore, true
}

// SetIsMore sets field value
func (o *ProgramGatewayPage) SetIsMore(v bool) {
	o.IsMore = v
}

// GetStartIndex returns the StartIndex field value
func (o *ProgramGatewayPage) GetStartIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value
// and a boolean to check if the value has been set.
func (o *ProgramGatewayPage) GetStartIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartIndex, true
}

// SetStartIndex sets field value
func (o *ProgramGatewayPage) SetStartIndex(v int32) {
	o.StartIndex = v
}

func (o ProgramGatewayPage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramGatewayPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["count"] = o.Count
	toSerialize["data"] = o.Data
	toSerialize["end_index"] = o.EndIndex
	toSerialize["is_more"] = o.IsMore
	toSerialize["start_index"] = o.StartIndex
	return toSerialize, nil
}

func (o *ProgramGatewayPage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"count",
		"data",
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

	varProgramGatewayPage := _ProgramGatewayPage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varProgramGatewayPage)

	if err != nil {
		return err
	}

	*o = ProgramGatewayPage(varProgramGatewayPage)

	return err
}

type NullableProgramGatewayPage struct {
	value *ProgramGatewayPage
	isSet bool
}

func (v NullableProgramGatewayPage) Get() *ProgramGatewayPage {
	return v.value
}

func (v *NullableProgramGatewayPage) Set(val *ProgramGatewayPage) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramGatewayPage) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramGatewayPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramGatewayPage(val *ProgramGatewayPage) *NullableProgramGatewayPage {
	return &NullableProgramGatewayPage{value: val, isSet: true}
}

func (v NullableProgramGatewayPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramGatewayPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


