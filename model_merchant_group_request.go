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

// checks if the MerchantGroupRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MerchantGroupRequest{}

// MerchantGroupRequest struct for MerchantGroupRequest
type MerchantGroupRequest struct {
	// Indicates if the merchant group is active or not.
	Active *bool `json:"active,omitempty"`
	// Comma-separated list of alphanumeric merchant identifiers. You can include merchant identifiers in multiple merchant groups.
	Mids []string `json:"mids,omitempty"`
	// Name of the merchant group.
	Name string `json:"name"`
	// Unique identifier of the group.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
}

type _MerchantGroupRequest MerchantGroupRequest

// NewMerchantGroupRequest instantiates a new MerchantGroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMerchantGroupRequest(name string) *MerchantGroupRequest {
	this := MerchantGroupRequest{}
	var active bool = false
	this.Active = &active
	this.Name = name
	return &this
}

// NewMerchantGroupRequestWithDefaults instantiates a new MerchantGroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMerchantGroupRequestWithDefaults() *MerchantGroupRequest {
	this := MerchantGroupRequest{}
	var active bool = false
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *MerchantGroupRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGroupRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *MerchantGroupRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *MerchantGroupRequest) SetActive(v bool) {
	o.Active = &v
}

// GetMids returns the Mids field value if set, zero value otherwise.
func (o *MerchantGroupRequest) GetMids() []string {
	if o == nil || IsNil(o.Mids) {
		var ret []string
		return ret
	}
	return o.Mids
}

// GetMidsOk returns a tuple with the Mids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGroupRequest) GetMidsOk() ([]string, bool) {
	if o == nil || IsNil(o.Mids) {
		return nil, false
	}
	return o.Mids, true
}

// HasMids returns a boolean if a field has been set.
func (o *MerchantGroupRequest) HasMids() bool {
	if o != nil && !IsNil(o.Mids) {
		return true
	}

	return false
}

// SetMids gets a reference to the given []string and assigns it to the Mids field.
func (o *MerchantGroupRequest) SetMids(v []string) {
	o.Mids = v
}

// GetName returns the Name field value
func (o *MerchantGroupRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MerchantGroupRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MerchantGroupRequest) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *MerchantGroupRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MerchantGroupRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *MerchantGroupRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *MerchantGroupRequest) SetToken(v string) {
	o.Token = &v
}

func (o MerchantGroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MerchantGroupRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Mids) {
		toSerialize["mids"] = o.Mids
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *MerchantGroupRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varMerchantGroupRequest := _MerchantGroupRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMerchantGroupRequest)

	if err != nil {
		return err
	}

	*o = MerchantGroupRequest(varMerchantGroupRequest)

	return err
}

type NullableMerchantGroupRequest struct {
	value *MerchantGroupRequest
	isSet bool
}

func (v NullableMerchantGroupRequest) Get() *MerchantGroupRequest {
	return v.value
}

func (v *NullableMerchantGroupRequest) Set(val *MerchantGroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMerchantGroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMerchantGroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMerchantGroupRequest(val *MerchantGroupRequest) *NullableMerchantGroupRequest {
	return &NullableMerchantGroupRequest{value: val, isSet: true}
}

func (v NullableMerchantGroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMerchantGroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


