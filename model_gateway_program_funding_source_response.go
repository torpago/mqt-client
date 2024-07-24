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

// checks if the GatewayProgramFundingSourceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayProgramFundingSourceResponse{}

// GatewayProgramFundingSourceResponse struct for GatewayProgramFundingSourceResponse
type GatewayProgramFundingSourceResponse struct {
	// Bank account number.
	Account string `json:"account"`
	// Indicates whether the program gateway funding source is active. This field is returned if it exists in the resource.
	Active *bool `json:"active,omitempty"`
	// Password for authenticating your environment.
	BasicAuthPassword string `json:"basic_auth_password"`
	// Username for authenticating your environment.
	BasicAuthUsername string `json:"basic_auth_username"`
	// Date and time when the resource was created, in UTC.
	CreatedTime time.Time `json:"created_time"`
	// Additional custom information included in the HTTP header.
	CustomHeader map[string]string `json:"custom_header"`
	// Date and time when the resource was last modified, in UTC.
	LastModifiedTime time.Time `json:"last_modified_time"`
	// Name of the program gateway funding source.
	Name string `json:"name"`
	// Total timeout in milliseconds for gateway processing.
	TimeoutMillis int64 `json:"timeout_millis"`
	// Unique identifier of the program gateway funding source.
	Token string `json:"token"`
	// URL of the gateway endpoint hosted in your environment, to which `POST` requests are submitted by the Marqeta platform.
	Url string `json:"url"`
	// Specifies whether or not to use mutual transport layer security (mTLS) authentication for the funding request.
	UseMtls bool `json:"use_mtls"`
	// Program gateway funding source object version.
	Version string `json:"version"`
}

type _GatewayProgramFundingSourceResponse GatewayProgramFundingSourceResponse

// NewGatewayProgramFundingSourceResponse instantiates a new GatewayProgramFundingSourceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayProgramFundingSourceResponse(account string, basicAuthPassword string, basicAuthUsername string, createdTime time.Time, customHeader map[string]string, lastModifiedTime time.Time, name string, timeoutMillis int64, token string, url string, useMtls bool, version string) *GatewayProgramFundingSourceResponse {
	this := GatewayProgramFundingSourceResponse{}
	this.Account = account
	this.BasicAuthPassword = basicAuthPassword
	this.BasicAuthUsername = basicAuthUsername
	this.CreatedTime = createdTime
	this.CustomHeader = customHeader
	this.LastModifiedTime = lastModifiedTime
	this.Name = name
	this.TimeoutMillis = timeoutMillis
	this.Token = token
	this.Url = url
	this.UseMtls = useMtls
	this.Version = version
	return &this
}

// NewGatewayProgramFundingSourceResponseWithDefaults instantiates a new GatewayProgramFundingSourceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayProgramFundingSourceResponseWithDefaults() *GatewayProgramFundingSourceResponse {
	this := GatewayProgramFundingSourceResponse{}
	var useMtls bool = false
	this.UseMtls = useMtls
	return &this
}

// GetAccount returns the Account field value
func (o *GatewayProgramFundingSourceResponse) GetAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Account
}

// GetAccountOk returns a tuple with the Account field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Account, true
}

// SetAccount sets field value
func (o *GatewayProgramFundingSourceResponse) SetAccount(v string) {
	o.Account = v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GatewayProgramFundingSourceResponse) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GatewayProgramFundingSourceResponse) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GatewayProgramFundingSourceResponse) SetActive(v bool) {
	o.Active = &v
}

// GetBasicAuthPassword returns the BasicAuthPassword field value
func (o *GatewayProgramFundingSourceResponse) GetBasicAuthPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasicAuthPassword
}

// GetBasicAuthPasswordOk returns a tuple with the BasicAuthPassword field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetBasicAuthPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasicAuthPassword, true
}

// SetBasicAuthPassword sets field value
func (o *GatewayProgramFundingSourceResponse) SetBasicAuthPassword(v string) {
	o.BasicAuthPassword = v
}

// GetBasicAuthUsername returns the BasicAuthUsername field value
func (o *GatewayProgramFundingSourceResponse) GetBasicAuthUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BasicAuthUsername
}

// GetBasicAuthUsernameOk returns a tuple with the BasicAuthUsername field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetBasicAuthUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BasicAuthUsername, true
}

// SetBasicAuthUsername sets field value
func (o *GatewayProgramFundingSourceResponse) SetBasicAuthUsername(v string) {
	o.BasicAuthUsername = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *GatewayProgramFundingSourceResponse) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *GatewayProgramFundingSourceResponse) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetCustomHeader returns the CustomHeader field value
func (o *GatewayProgramFundingSourceResponse) GetCustomHeader() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.CustomHeader
}

// GetCustomHeaderOk returns a tuple with the CustomHeader field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetCustomHeaderOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomHeader, true
}

// SetCustomHeader sets field value
func (o *GatewayProgramFundingSourceResponse) SetCustomHeader(v map[string]string) {
	o.CustomHeader = v
}

// GetLastModifiedTime returns the LastModifiedTime field value
func (o *GatewayProgramFundingSourceResponse) GetLastModifiedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastModifiedTime, true
}

// SetLastModifiedTime sets field value
func (o *GatewayProgramFundingSourceResponse) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = v
}

// GetName returns the Name field value
func (o *GatewayProgramFundingSourceResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayProgramFundingSourceResponse) SetName(v string) {
	o.Name = v
}

// GetTimeoutMillis returns the TimeoutMillis field value
func (o *GatewayProgramFundingSourceResponse) GetTimeoutMillis() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimeoutMillis
}

// GetTimeoutMillisOk returns a tuple with the TimeoutMillis field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetTimeoutMillisOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeoutMillis, true
}

// SetTimeoutMillis sets field value
func (o *GatewayProgramFundingSourceResponse) SetTimeoutMillis(v int64) {
	o.TimeoutMillis = v
}

// GetToken returns the Token field value
func (o *GatewayProgramFundingSourceResponse) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GatewayProgramFundingSourceResponse) SetToken(v string) {
	o.Token = v
}

// GetUrl returns the Url field value
func (o *GatewayProgramFundingSourceResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GatewayProgramFundingSourceResponse) SetUrl(v string) {
	o.Url = v
}

// GetUseMtls returns the UseMtls field value
func (o *GatewayProgramFundingSourceResponse) GetUseMtls() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseMtls
}

// GetUseMtlsOk returns a tuple with the UseMtls field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetUseMtlsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseMtls, true
}

// SetUseMtls sets field value
func (o *GatewayProgramFundingSourceResponse) SetUseMtls(v bool) {
	o.UseMtls = v
}

// GetVersion returns the Version field value
func (o *GatewayProgramFundingSourceResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *GatewayProgramFundingSourceResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *GatewayProgramFundingSourceResponse) SetVersion(v string) {
	o.Version = v
}

func (o GatewayProgramFundingSourceResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayProgramFundingSourceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account"] = o.Account
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["basic_auth_password"] = o.BasicAuthPassword
	toSerialize["basic_auth_username"] = o.BasicAuthUsername
	toSerialize["created_time"] = o.CreatedTime
	toSerialize["custom_header"] = o.CustomHeader
	toSerialize["last_modified_time"] = o.LastModifiedTime
	toSerialize["name"] = o.Name
	toSerialize["timeout_millis"] = o.TimeoutMillis
	toSerialize["token"] = o.Token
	toSerialize["url"] = o.Url
	toSerialize["use_mtls"] = o.UseMtls
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *GatewayProgramFundingSourceResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account",
		"basic_auth_password",
		"basic_auth_username",
		"created_time",
		"custom_header",
		"last_modified_time",
		"name",
		"timeout_millis",
		"token",
		"url",
		"use_mtls",
		"version",
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

	varGatewayProgramFundingSourceResponse := _GatewayProgramFundingSourceResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGatewayProgramFundingSourceResponse)

	if err != nil {
		return err
	}

	*o = GatewayProgramFundingSourceResponse(varGatewayProgramFundingSourceResponse)

	return err
}

type NullableGatewayProgramFundingSourceResponse struct {
	value *GatewayProgramFundingSourceResponse
	isSet bool
}

func (v NullableGatewayProgramFundingSourceResponse) Get() *GatewayProgramFundingSourceResponse {
	return v.value
}

func (v *NullableGatewayProgramFundingSourceResponse) Set(val *GatewayProgramFundingSourceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayProgramFundingSourceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayProgramFundingSourceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayProgramFundingSourceResponse(val *GatewayProgramFundingSourceResponse) *NullableGatewayProgramFundingSourceResponse {
	return &NullableGatewayProgramFundingSourceResponse{value: val, isSet: true}
}

func (v NullableGatewayProgramFundingSourceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayProgramFundingSourceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


