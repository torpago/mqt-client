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

// checks if the FeeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeeResponse{}

// FeeResponse Contains details about the fee.
type FeeResponse struct {
	// Specifies whether the fee is active.
	Active bool `json:"active"`
	// Amount of the fee.
	Amount decimal.Decimal `json:"amount"`
	// Specifies if the fee is a standalone fee or a real-time fee.
	Category *string `json:"category,omitempty"`
	// Date and time when the `fees` object was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Three-digit ISO 4217 currency code.
	CurrencyCode string `json:"currency_code"`
	FeeAttributes *FeeAttributes `json:"fee_attributes,omitempty"`
	// Date and time when the `fees` object was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Name of the fee.
	Name string `json:"name"`
	// Descriptive metadata about the fee.
	Tags *string `json:"tags,omitempty"`
	// Unique identifier of the `fees` object.
	Token string `json:"token"`
	// Specifies if the fee is a flat fee or a percentage of the transaction amount.
	Type *string `json:"type,omitempty"`
}

type _FeeResponse FeeResponse

// NewFeeResponse instantiates a new FeeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeResponse(active bool, amount decimal.Decimal, createdTime time.Time, currencyCode string, lastModifiedTime time.Time, name string, token string) *FeeResponse {
	this := FeeResponse{}
	this.Active = active
	this.Amount = amount
	this.CreatedTime = createdTime
	this.CurrencyCode = currencyCode
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	this.Token = token
	return &this
}

// NewFeeResponseWithDefaults instantiates a new FeeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeResponseWithDefaults() *FeeResponse {
	this := FeeResponse{}
	return &this
}

// GetActive returns the Active field value
func (o *FeeResponse) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *FeeResponse) SetActive(v bool) {
	o.Active = v
}

// GetAmount returns the Amount field value
func (o *FeeResponse) GetAmount() decimal.Decimal {
	if o == nil {
		var ret decimal.Decimal
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetAmountOk() (*decimal.Decimal, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *FeeResponse) SetAmount(v decimal.Decimal) {
	o.Amount = v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *FeeResponse) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *FeeResponse) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *FeeResponse) SetCategory(v string) {
	o.Category = &v
}

// GetCreatedTime returns the CreatedTime field value
func (o *FeeResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *FeeResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCurrencyCode returns the CurrencyCode field value
func (o *FeeResponse) GetCurrencyCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetCurrencyCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrencyCode, true
}

// SetCurrencyCode sets field value
func (o *FeeResponse) SetCurrencyCode(v string) {
	o.CurrencyCode = v
}

// GetFeeAttributes returns the FeeAttributes field value if set, zero value otherwise.
func (o *FeeResponse) GetFeeAttributes() FeeAttributes {
	if o == nil || IsNil(o.FeeAttributes) {
		var ret FeeAttributes
		return ret
	}
	return *o.FeeAttributes
}

// GetFeeAttributesOk returns a tuple with the FeeAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetFeeAttributesOk() (*FeeAttributes, bool) {
	if o == nil || IsNil(o.FeeAttributes) {
		return nil, false
	}
	return o.FeeAttributes, true
}

// HasFeeAttributes returns a boolean if a field has been set.
func (o *FeeResponse) HasFeeAttributes() bool {
	if o != nil && !IsNil(o.FeeAttributes) {
		return true
	}

	return false
}

// SetFeeAttributes gets a reference to the given FeeAttributes and assigns it to the FeeAttributes field.
func (o *FeeResponse) SetFeeAttributes(v FeeAttributes) {
	o.FeeAttributes = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *FeeResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *FeeResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetName returns the Name field value
func (o *FeeResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FeeResponse) SetName(v string) {
	o.Name = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *FeeResponse) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *FeeResponse) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *FeeResponse) SetTags(v string) {
	o.Tags = &v
}

// GetToken returns the Token field value
func (o *FeeResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *FeeResponse) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeeResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeeResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FeeResponse) SetType(v string) {
	o.Type = &v
}

func (o FeeResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["amount"] = o.Amount
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["currency_code"] = o.CurrencyCode
	if !IsNil(o.FeeAttributes) {
		toSerialize["fee_attributes"] = o.FeeAttributes
	}
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["name"] = o.Name
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["token"] = o.Token
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

func (o *FeeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"amount",
		"created_time",
		"currency_code",
		"last_modified_time",
		"name",
		"token",
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

	varFeeResponse := _FeeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFeeResponse)

	if err != nil {
		return err
	}

	*o = FeeResponse(varFeeResponse)

	return err
}

type NullableFeeResponse struct {
	value *FeeResponse
	isSet bool
}

func (v NullableFeeResponse) Get() *FeeResponse {
	return v.value
}

func (v *NullableFeeResponse) Set(val *FeeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeResponse(val *FeeResponse) *NullableFeeResponse {
	return &NullableFeeResponse{value: val, isSet: true}
}

func (v NullableFeeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


