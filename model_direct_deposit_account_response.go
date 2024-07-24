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

// checks if the DirectDepositAccountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DirectDepositAccountResponse{}

// DirectDepositAccountResponse struct for DirectDepositAccountResponse
type DirectDepositAccountResponse struct {
	AccountNumber string `json:"account_number"`
	AllowImmediateCredit bool `json:"allow_immediate_credit"`
	BusinessToken string `json:"business_token"`
	// yyyy-MM-ddTHH:mm:ssZ
	CreatedTime time.Time `json:"created_time"`
	// yyyy-MM-ddTHH:mm:ssZ
	LastModifiedTime time.Time `json:"last_modified_time"`
	RoutingNumber string `json:"routing_number"`
	State string `json:"state"`
	Token string `json:"token"`
	Type *string `json:"type,omitempty"`
	UserToken string `json:"user_token"`
}

type _DirectDepositAccountResponse DirectDepositAccountResponse

// NewDirectDepositAccountResponse instantiates a new DirectDepositAccountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDirectDepositAccountResponse(accountNumber string, allowImmediateCredit bool, businessToken string, createdTime time.Time, lastModifiedTime time.Time, routingNumber string, state string, token string, userToken string) *DirectDepositAccountResponse {
	this := DirectDepositAccountResponse{}
	this.AccountNumber = accountNumber
	this.AllowImmediateCredit = allowImmediateCredit
	this.BusinessToken = businessToken
	this.CreatedTime = createdTime
	this.LastModifiedTime = lastModifiedTime
	this.RoutingNumber = routingNumber
	this.State = state
	this.Token = token
	this.UserToken = userToken
	return &this
}

// NewDirectDepositAccountResponseWithDefaults instantiates a new DirectDepositAccountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDirectDepositAccountResponseWithDefaults() *DirectDepositAccountResponse {
	this := DirectDepositAccountResponse{}
	var allowImmediateCredit bool = false
	this.AllowImmediateCredit = allowImmediateCredit
	return &this
}

// GetAccountNumber returns the AccountNumber field value
func (o *DirectDepositAccountResponse) GetAccountNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetAccountNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountNumber, true
}

// SetAccountNumber sets field value
func (o *DirectDepositAccountResponse) SetAccountNumber(v string) {
	o.AccountNumber = v
}

// GetAllowImmediateCredit returns the AllowImmediateCredit field value
func (o *DirectDepositAccountResponse) GetAllowImmediateCredit() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowImmediateCredit
}

// GetAllowImmediateCreditOk returns a tuple with the AllowImmediateCredit field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetAllowImmediateCreditOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowImmediateCredit, true
}

// SetAllowImmediateCredit sets field value
func (o *DirectDepositAccountResponse) SetAllowImmediateCredit(v bool) {
	o.AllowImmediateCredit = v
}

// GetBusinessToken returns the BusinessToken field value
func (o *DirectDepositAccountResponse) GetBusinessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessToken
}

// GetBusinessTokenOk returns a tuple with the BusinessToken field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetBusinessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessToken, true
}

// SetBusinessToken sets field value
func (o *DirectDepositAccountResponse) SetBusinessToken(v string) {
	o.BusinessToken = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *DirectDepositAccountResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *DirectDepositAccountResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *DirectDepositAccountResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *DirectDepositAccountResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetRoutingNumber returns the RoutingNumber field value
func (o *DirectDepositAccountResponse) GetRoutingNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingNumber
}

// GetRoutingNumberOk returns a tuple with the RoutingNumber field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetRoutingNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoutingNumber, true
}

// SetRoutingNumber sets field value
func (o *DirectDepositAccountResponse) SetRoutingNumber(v string) {
	o.RoutingNumber = v
}

// GetState returns the State field value
func (o *DirectDepositAccountResponse) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *DirectDepositAccountResponse) SetState(v string) {
	o.State = v
}

// GetToken returns the Token field value
func (o *DirectDepositAccountResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *DirectDepositAccountResponse) SetToken(v string) {
	o.Token = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DirectDepositAccountResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DirectDepositAccountResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DirectDepositAccountResponse) SetType(v string) {
	o.Type = &v
}

// GetUserToken returns the UserToken field value
func (o *DirectDepositAccountResponse) GetUserToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserToken
}

// GetUserTokenOk returns a tuple with the UserToken field value
// and a boolean to check if the value has been set.
func (o *DirectDepositAccountResponse) GetUserTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserToken, true
}

// SetUserToken sets field value
func (o *DirectDepositAccountResponse) SetUserToken(v string) {
	o.UserToken = v
}

func (o DirectDepositAccountResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DirectDepositAccountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_number"] = o.AccountNumber
	toSerialize["allow_immediate_credit"] = o.AllowImmediateCredit
	toSerialize["business_token"] = o.BusinessToken
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["routing_number"] = o.RoutingNumber
	toSerialize["state"] = o.State
	toSerialize["token"] = o.Token
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["user_token"] = o.UserToken
	return toSerialize, nil
}

func (o *DirectDepositAccountResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_number",
		"allow_immediate_credit",
		"business_token",
		"created_time",
		"last_modified_time",
		"routing_number",
		"state",
		"token",
		"user_token",
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

	varDirectDepositAccountResponse := _DirectDepositAccountResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDirectDepositAccountResponse)

	if err != nil {
		return err
	}

	*o = DirectDepositAccountResponse(varDirectDepositAccountResponse)

	return err
}

type NullableDirectDepositAccountResponse struct {
	value *DirectDepositAccountResponse
	isSet bool
}

func (v NullableDirectDepositAccountResponse) Get() *DirectDepositAccountResponse {
	return v.value
}

func (v *NullableDirectDepositAccountResponse) Set(val *DirectDepositAccountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDirectDepositAccountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDirectDepositAccountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDirectDepositAccountResponse(val *DirectDepositAccountResponse) *NullableDirectDepositAccountResponse {
	return &NullableDirectDepositAccountResponse{value: val, isSet: true}
}

func (v NullableDirectDepositAccountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDirectDepositAccountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


