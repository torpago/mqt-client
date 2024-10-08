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

// checks if the SamsungPushTokenizeRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SamsungPushTokenizeRequestData{}

// SamsungPushTokenizeRequestData Contains details about a card tokenization push request.
type SamsungPushTokenizeRequestData struct {
	// Specifies the type of card.
	CardType *string `json:"card_type,omitempty"`
	// Name of the card as displayed in the digital wallet, typically showing the card brand and last four digits of the primary account number (PAN). `Visa 5678`, for example.
	DisplayName *string `json:"display_name,omitempty"`
	// Encrypted card or cardholder details.
	ExtraProvisionPayload *string `json:"extra_provision_payload,omitempty"`
	// Last four digits of the primary account number of the physical or virtual card.
	LastDigits *string `json:"last_digits,omitempty"`
	// Specifies the card network of the physical or virtual card.
	Network *string `json:"network,omitempty"`
	// Specifies the network that provides the digital wallet token service, as determined by the Samsung Wallet library.
	TokenServiceProvider *string `json:"token_service_provider,omitempty"`
}

// NewSamsungPushTokenizeRequestData instantiates a new SamsungPushTokenizeRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamsungPushTokenizeRequestData() *SamsungPushTokenizeRequestData {
	this := SamsungPushTokenizeRequestData{}
	return &this
}

// NewSamsungPushTokenizeRequestDataWithDefaults instantiates a new SamsungPushTokenizeRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamsungPushTokenizeRequestDataWithDefaults() *SamsungPushTokenizeRequestData {
	this := SamsungPushTokenizeRequestData{}
	return &this
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *SamsungPushTokenizeRequestData) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamsungPushTokenizeRequestData) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *SamsungPushTokenizeRequestData) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *SamsungPushTokenizeRequestData) SetCardType(v string) {
	o.CardType = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SamsungPushTokenizeRequestData) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamsungPushTokenizeRequestData) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SamsungPushTokenizeRequestData) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SamsungPushTokenizeRequestData) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetExtraProvisionPayload returns the ExtraProvisionPayload field value if set, zero value otherwise.
func (o *SamsungPushTokenizeRequestData) GetExtraProvisionPayload() string {
	if o == nil || IsNil(o.ExtraProvisionPayload) {
		var ret string
		return ret
	}
	return *o.ExtraProvisionPayload
}

// GetExtraProvisionPayloadOk returns a tuple with the ExtraProvisionPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamsungPushTokenizeRequestData) GetExtraProvisionPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraProvisionPayload) {
		return nil, false
	}
	return o.ExtraProvisionPayload, true
}

// HasExtraProvisionPayload returns a boolean if a field has been set.
func (o *SamsungPushTokenizeRequestData) HasExtraProvisionPayload() bool {
	if o != nil && !IsNil(o.ExtraProvisionPayload) {
		return true
	}

	return false
}

// SetExtraProvisionPayload gets a reference to the given string and assigns it to the ExtraProvisionPayload field.
func (o *SamsungPushTokenizeRequestData) SetExtraProvisionPayload(v string) {
	o.ExtraProvisionPayload = &v
}

// GetLastDigits returns the LastDigits field value if set, zero value otherwise.
func (o *SamsungPushTokenizeRequestData) GetLastDigits() string {
	if o == nil || IsNil(o.LastDigits) {
		var ret string
		return ret
	}
	return *o.LastDigits
}

// GetLastDigitsOk returns a tuple with the LastDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamsungPushTokenizeRequestData) GetLastDigitsOk() (*string, bool) {
	if o == nil || IsNil(o.LastDigits) {
		return nil, false
	}
	return o.LastDigits, true
}

// HasLastDigits returns a boolean if a field has been set.
func (o *SamsungPushTokenizeRequestData) HasLastDigits() bool {
	if o != nil && !IsNil(o.LastDigits) {
		return true
	}

	return false
}

// SetLastDigits gets a reference to the given string and assigns it to the LastDigits field.
func (o *SamsungPushTokenizeRequestData) SetLastDigits(v string) {
	o.LastDigits = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *SamsungPushTokenizeRequestData) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamsungPushTokenizeRequestData) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *SamsungPushTokenizeRequestData) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *SamsungPushTokenizeRequestData) SetNetwork(v string) {
	o.Network = &v
}

// GetTokenServiceProvider returns the TokenServiceProvider field value if set, zero value otherwise.
func (o *SamsungPushTokenizeRequestData) GetTokenServiceProvider() string {
	if o == nil || IsNil(o.TokenServiceProvider) {
		var ret string
		return ret
	}
	return *o.TokenServiceProvider
}

// GetTokenServiceProviderOk returns a tuple with the TokenServiceProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamsungPushTokenizeRequestData) GetTokenServiceProviderOk() (*string, bool) {
	if o == nil || IsNil(o.TokenServiceProvider) {
		return nil, false
	}
	return o.TokenServiceProvider, true
}

// HasTokenServiceProvider returns a boolean if a field has been set.
func (o *SamsungPushTokenizeRequestData) HasTokenServiceProvider() bool {
	if o != nil && !IsNil(o.TokenServiceProvider) {
		return true
	}

	return false
}

// SetTokenServiceProvider gets a reference to the given string and assigns it to the TokenServiceProvider field.
func (o *SamsungPushTokenizeRequestData) SetTokenServiceProvider(v string) {
	o.TokenServiceProvider = &v
}

func (o SamsungPushTokenizeRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SamsungPushTokenizeRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !IsNil(o.ExtraProvisionPayload) {
		toSerialize["extra_provision_payload"] = o.ExtraProvisionPayload
	}
	if !IsNil(o.LastDigits) {
		toSerialize["last_digits"] = o.LastDigits
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.TokenServiceProvider) {
		toSerialize["token_service_provider"] = o.TokenServiceProvider
	}
	return toSerialize, nil
}

type NullableSamsungPushTokenizeRequestData struct {
	value *SamsungPushTokenizeRequestData
	isSet bool
}

func (v NullableSamsungPushTokenizeRequestData) Get() *SamsungPushTokenizeRequestData {
	return v.value
}

func (v *NullableSamsungPushTokenizeRequestData) Set(val *SamsungPushTokenizeRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSamsungPushTokenizeRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSamsungPushTokenizeRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamsungPushTokenizeRequestData(val *SamsungPushTokenizeRequestData) *NullableSamsungPushTokenizeRequestData {
	return &NullableSamsungPushTokenizeRequestData{value: val, isSet: true}
}

func (v NullableSamsungPushTokenizeRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamsungPushTokenizeRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


