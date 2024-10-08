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

// checks if the PinRevealRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PinRevealRequest{}

// PinRevealRequest struct for PinRevealRequest
type PinRevealRequest struct {
	// The supplemental method used to verify the cardholder's identity before revealing the card's personal identification number (PIN).  The possible cardholder verification methods are:  * *BIOMETRIC_FACE:* In-app authentication via facial recognition * *BIOMETRIC_FINGERPRINT:* In-app authentication via biometric fingerprint * *EXP_CVV:* In-app authentication by entering the card's expiration date and card verification value (CVV) * *LOGIN:* In-app authentication by re-entering the app password * *OTP:* Two-factor authentication involving a one-time password (OTP) * *OTP_CVV:* Two-factor authentication involving the card's CVV and an OTP * *OTHER:* Authentication that relies on other secure methods
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	// Unique value generated as a result of issuing a `POST` request to the `/pins/controltoken` endpoint. This value cannot be updated.
	ControlToken string `json:"control_token"`
}

type _PinRevealRequest PinRevealRequest

// NewPinRevealRequest instantiates a new PinRevealRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPinRevealRequest(cardholderVerificationMethod string, controlToken string) *PinRevealRequest {
	this := PinRevealRequest{}
	this.CardholderVerificationMethod = cardholderVerificationMethod
	this.ControlToken = controlToken
	return &this
}

// NewPinRevealRequestWithDefaults instantiates a new PinRevealRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPinRevealRequestWithDefaults() *PinRevealRequest {
	this := PinRevealRequest{}
	return &this
}

// GetCardholderVerificationMethod returns the CardholderVerificationMethod field value
func (o *PinRevealRequest) GetCardholderVerificationMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CardholderVerificationMethod
}

// GetCardholderVerificationMethodOk returns a tuple with the CardholderVerificationMethod field value
// and a boolean to check if the value has been set.
func (o *PinRevealRequest) GetCardholderVerificationMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CardholderVerificationMethod, true
}

// SetCardholderVerificationMethod sets field value
func (o *PinRevealRequest) SetCardholderVerificationMethod(v string) {
	o.CardholderVerificationMethod = v
}

// GetControlToken returns the ControlToken field value
func (o *PinRevealRequest) GetControlToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ControlToken
}

// GetControlTokenOk returns a tuple with the ControlToken field value
// and a boolean to check if the value has been set.
func (o *PinRevealRequest) GetControlTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlToken, true
}

// SetControlToken sets field value
func (o *PinRevealRequest) SetControlToken(v string) {
	o.ControlToken = v
}

func (o PinRevealRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PinRevealRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cardholder_verification_method"] = o.CardholderVerificationMethod
	toSerialize["control_token"] = o.ControlToken
	return toSerialize, nil
}

func (o *PinRevealRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cardholder_verification_method",
		"control_token",
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

	varPinRevealRequest := _PinRevealRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPinRevealRequest)

	if err != nil {
		return err
	}

	*o = PinRevealRequest(varPinRevealRequest)

	return err
}

type NullablePinRevealRequest struct {
	value *PinRevealRequest
	isSet bool
}

func (v NullablePinRevealRequest) Get() *PinRevealRequest {
	return v.value
}

func (v *NullablePinRevealRequest) Set(val *PinRevealRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePinRevealRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePinRevealRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePinRevealRequest(val *PinRevealRequest) *NullablePinRevealRequest {
	return &NullablePinRevealRequest{value: val, isSet: true}
}

func (v NullablePinRevealRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePinRevealRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


