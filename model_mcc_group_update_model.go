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
)

// checks if the MccGroupUpdateModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MccGroupUpdateModel{}

// MccGroupUpdateModel struct for MccGroupUpdateModel
type MccGroupUpdateModel struct {
	// Indicates whether the MCC group is active or inactive.
	Active *bool `json:"active,omitempty"`
	Config *Config `json:"config,omitempty"`
	// The set of merchant category codes that you want to include in this group. For each element, valid characters are 0-9, and the length must be 4 digits. You can also specify a range like \"9876-9880\". An MCC can belong to more than one group.  Updating the merchant category codes for the group completely replaces the group's existing codes. For example, if the current MCC group is `[\"1234\"]` and you want to add the 2345 code (while retaining the existing code), you must specify `[\"1234\", \"2345\"]` in this field.
	Mccs []string `json:"mccs,omitempty"`
	// The name of the MCC group.
	Name *string `json:"name,omitempty"`
}

// NewMccGroupUpdateModel instantiates a new MccGroupUpdateModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMccGroupUpdateModel() *MccGroupUpdateModel {
	this := MccGroupUpdateModel{}
	var active bool = false
	this.Active = &active
	return &this
}

// NewMccGroupUpdateModelWithDefaults instantiates a new MccGroupUpdateModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMccGroupUpdateModelWithDefaults() *MccGroupUpdateModel {
	this := MccGroupUpdateModel{}
	var active bool = false
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *MccGroupUpdateModel) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccGroupUpdateModel) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *MccGroupUpdateModel) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *MccGroupUpdateModel) SetActive(v bool) {
	o.Active = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *MccGroupUpdateModel) GetConfig() Config {
	if o == nil || IsNil(o.Config) {
		var ret Config
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccGroupUpdateModel) GetConfigOk() (*Config, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *MccGroupUpdateModel) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given Config and assigns it to the Config field.
func (o *MccGroupUpdateModel) SetConfig(v Config) {
	o.Config = &v
}

// GetMccs returns the Mccs field value if set, zero value otherwise.
func (o *MccGroupUpdateModel) GetMccs() []string {
	if o == nil || IsNil(o.Mccs) {
		var ret []string
		return ret
	}
	return o.Mccs
}

// GetMccsOk returns a tuple with the Mccs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccGroupUpdateModel) GetMccsOk() ([]string, bool) {
	if o == nil || IsNil(o.Mccs) {
		return nil, false
	}
	return o.Mccs, true
}

// HasMccs returns a boolean if a field has been set.
func (o *MccGroupUpdateModel) HasMccs() bool {
	if o != nil && !IsNil(o.Mccs) {
		return true
	}

	return false
}

// SetMccs gets a reference to the given []string and assigns it to the Mccs field.
func (o *MccGroupUpdateModel) SetMccs(v []string) {
	o.Mccs = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MccGroupUpdateModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MccGroupUpdateModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MccGroupUpdateModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MccGroupUpdateModel) SetName(v string) {
	o.Name = &v
}

func (o MccGroupUpdateModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MccGroupUpdateModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Mccs) {
		toSerialize["mccs"] = o.Mccs
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableMccGroupUpdateModel struct {
	value *MccGroupUpdateModel
	isSet bool
}

func (v NullableMccGroupUpdateModel) Get() *MccGroupUpdateModel {
	return v.value
}

func (v *NullableMccGroupUpdateModel) Set(val *MccGroupUpdateModel) {
	v.value = val
	v.isSet = true
}

func (v NullableMccGroupUpdateModel) IsSet() bool {
	return v.isSet
}

func (v *NullableMccGroupUpdateModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMccGroupUpdateModel(val *MccGroupUpdateModel) *NullableMccGroupUpdateModel {
	return &NullableMccGroupUpdateModel{value: val, isSet: true}
}

func (v NullableMccGroupUpdateModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMccGroupUpdateModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


