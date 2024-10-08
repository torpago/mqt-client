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

// checks if the AuthControlExemptMidsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthControlExemptMidsRequest{}

// AuthControlExemptMidsRequest struct for AuthControlExemptMidsRequest
type AuthControlExemptMidsRequest struct {
	Association *SpendControlAssociation `json:"association,omitempty"`
	// Date and time when the exception ends, in UTC.
	EndTime *time.Time `json:"end_time,omitempty"`
	// Unique identifier of the merchant group to be exempted. This field is required if there is no entry in the `mid` field. Pass either this field or the `mid` field, not both.  For information about merchant groups, see <</core-api/merchant-groups, Merchant Groups>>.
	MerchantGroupToken *string `json:"merchant_group_token,omitempty"`
	// Merchant to be exempted. This field is required if there is no entry in the `merchant_group_token` field. Use either this field or the `merchant_group_token` field, not both.
	Mid *string `json:"mid,omitempty"`
	// Name of the merchant identifier authorization control exemption.
	Name string `json:"name"`
	// Date and time when the exception starts, in UTC.
	StartTime *time.Time `json:"start_time,omitempty"`
	// Unique identifier of the merchant identifier authorization control exemption.  If you do not include a token, the system will generate one automatically. This token is necessary for use in other API calls, so we recommend that rather than let the system generate one, you use a simple string that is easy to remember. This value cannot be updated.
	Token *string `json:"token,omitempty"`
}

type _AuthControlExemptMidsRequest AuthControlExemptMidsRequest

// NewAuthControlExemptMidsRequest instantiates a new AuthControlExemptMidsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthControlExemptMidsRequest(name string) *AuthControlExemptMidsRequest {
	this := AuthControlExemptMidsRequest{}
	this.Name = name
	return &this
}

// NewAuthControlExemptMidsRequestWithDefaults instantiates a new AuthControlExemptMidsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthControlExemptMidsRequestWithDefaults() *AuthControlExemptMidsRequest {
	this := AuthControlExemptMidsRequest{}
	return &this
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *AuthControlExemptMidsRequest) GetAssociation() SpendControlAssociation {
	if o == nil || IsNil(o.Association) {
		var ret SpendControlAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlExemptMidsRequest) GetAssociationOk() (*SpendControlAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *AuthControlExemptMidsRequest) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given SpendControlAssociation and assigns it to the Association field.
func (o *AuthControlExemptMidsRequest) SetAssociation(v SpendControlAssociation) {
	o.Association = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *AuthControlExemptMidsRequest) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlExemptMidsRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *AuthControlExemptMidsRequest) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *AuthControlExemptMidsRequest) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetMerchantGroupToken returns the MerchantGroupToken field value if set, zero value otherwise.
func (o *AuthControlExemptMidsRequest) GetMerchantGroupToken() string {
	if o == nil || IsNil(o.MerchantGroupToken) {
		var ret string
		return ret
	}
	return *o.MerchantGroupToken
}

// GetMerchantGroupTokenOk returns a tuple with the MerchantGroupToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlExemptMidsRequest) GetMerchantGroupTokenOk() (*string, bool) {
	if o == nil || IsNil(o.MerchantGroupToken) {
		return nil, false
	}
	return o.MerchantGroupToken, true
}

// HasMerchantGroupToken returns a boolean if a field has been set.
func (o *AuthControlExemptMidsRequest) HasMerchantGroupToken() bool {
	if o != nil && !IsNil(o.MerchantGroupToken) {
		return true
	}

	return false
}

// SetMerchantGroupToken gets a reference to the given string and assigns it to the MerchantGroupToken field.
func (o *AuthControlExemptMidsRequest) SetMerchantGroupToken(v string) {
	o.MerchantGroupToken = &v
}

// GetMid returns the Mid field value if set, zero value otherwise.
func (o *AuthControlExemptMidsRequest) GetMid() string {
	if o == nil || IsNil(o.Mid) {
		var ret string
		return ret
	}
	return *o.Mid
}

// GetMidOk returns a tuple with the Mid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlExemptMidsRequest) GetMidOk() (*string, bool) {
	if o == nil || IsNil(o.Mid) {
		return nil, false
	}
	return o.Mid, true
}

// HasMid returns a boolean if a field has been set.
func (o *AuthControlExemptMidsRequest) HasMid() bool {
	if o != nil && !IsNil(o.Mid) {
		return true
	}

	return false
}

// SetMid gets a reference to the given string and assigns it to the Mid field.
func (o *AuthControlExemptMidsRequest) SetMid(v string) {
	o.Mid = &v
}

// GetName returns the Name field value
func (o *AuthControlExemptMidsRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthControlExemptMidsRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthControlExemptMidsRequest) SetName(v string) {
	o.Name = v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *AuthControlExemptMidsRequest) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlExemptMidsRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *AuthControlExemptMidsRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *AuthControlExemptMidsRequest) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *AuthControlExemptMidsRequest) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthControlExemptMidsRequest) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *AuthControlExemptMidsRequest) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *AuthControlExemptMidsRequest) SetToken(v string) {
	o.Token = &v
}

func (o AuthControlExemptMidsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthControlExemptMidsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.MerchantGroupToken) {
		toSerialize["merchant_group_token"] = o.MerchantGroupToken
	}
	if !IsNil(o.Mid) {
		toSerialize["mid"] = o.Mid
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *AuthControlExemptMidsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varAuthControlExemptMidsRequest := _AuthControlExemptMidsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAuthControlExemptMidsRequest)

	if err != nil {
		return err
	}

	*o = AuthControlExemptMidsRequest(varAuthControlExemptMidsRequest)

	return err
}

type NullableAuthControlExemptMidsRequest struct {
	value *AuthControlExemptMidsRequest
	isSet bool
}

func (v NullableAuthControlExemptMidsRequest) Get() *AuthControlExemptMidsRequest {
	return v.value
}

func (v *NullableAuthControlExemptMidsRequest) Set(val *AuthControlExemptMidsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthControlExemptMidsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthControlExemptMidsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthControlExemptMidsRequest(val *AuthControlExemptMidsRequest) *NullableAuthControlExemptMidsRequest {
	return &NullableAuthControlExemptMidsRequest{value: val, isSet: true}
}

func (v NullableAuthControlExemptMidsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthControlExemptMidsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


