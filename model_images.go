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

// checks if the Images type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Images{}

// Images Specifies personalized images that appear on the card.
type Images struct {
	Card *ImagesCard `json:"card,omitempty"`
	Carrier *ImagesCarrier `json:"carrier,omitempty"`
	CarrierReturnWindow *ImagesCarrierReturnWindow `json:"carrier_return_window,omitempty"`
	Signature *ImagesSignature `json:"signature,omitempty"`
}

// NewImages instantiates a new Images object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImages() *Images {
	this := Images{}
	return &this
}

// NewImagesWithDefaults instantiates a new Images object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImagesWithDefaults() *Images {
	this := Images{}
	return &this
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *Images) GetCard() ImagesCard {
	if o == nil || IsNil(o.Card) {
		var ret ImagesCard
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Images) GetCardOk() (*ImagesCard, bool) {
	if o == nil || IsNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *Images) HasCard() bool {
	if o != nil && !IsNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given ImagesCard and assigns it to the Card field.
func (o *Images) SetCard(v ImagesCard) {
	o.Card = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *Images) GetCarrier() ImagesCarrier {
	if o == nil || IsNil(o.Carrier) {
		var ret ImagesCarrier
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Images) GetCarrierOk() (*ImagesCarrier, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *Images) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given ImagesCarrier and assigns it to the Carrier field.
func (o *Images) SetCarrier(v ImagesCarrier) {
	o.Carrier = &v
}

// GetCarrierReturnWindow returns the CarrierReturnWindow field value if set, zero value otherwise.
func (o *Images) GetCarrierReturnWindow() ImagesCarrierReturnWindow {
	if o == nil || IsNil(o.CarrierReturnWindow) {
		var ret ImagesCarrierReturnWindow
		return ret
	}
	return *o.CarrierReturnWindow
}

// GetCarrierReturnWindowOk returns a tuple with the CarrierReturnWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Images) GetCarrierReturnWindowOk() (*ImagesCarrierReturnWindow, bool) {
	if o == nil || IsNil(o.CarrierReturnWindow) {
		return nil, false
	}
	return o.CarrierReturnWindow, true
}

// HasCarrierReturnWindow returns a boolean if a field has been set.
func (o *Images) HasCarrierReturnWindow() bool {
	if o != nil && !IsNil(o.CarrierReturnWindow) {
		return true
	}

	return false
}

// SetCarrierReturnWindow gets a reference to the given ImagesCarrierReturnWindow and assigns it to the CarrierReturnWindow field.
func (o *Images) SetCarrierReturnWindow(v ImagesCarrierReturnWindow) {
	o.CarrierReturnWindow = &v
}

// GetSignature returns the Signature field value if set, zero value otherwise.
func (o *Images) GetSignature() ImagesSignature {
	if o == nil || IsNil(o.Signature) {
		var ret ImagesSignature
		return ret
	}
	return *o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Images) GetSignatureOk() (*ImagesSignature, bool) {
	if o == nil || IsNil(o.Signature) {
		return nil, false
	}
	return o.Signature, true
}

// HasSignature returns a boolean if a field has been set.
func (o *Images) HasSignature() bool {
	if o != nil && !IsNil(o.Signature) {
		return true
	}

	return false
}

// SetSignature gets a reference to the given ImagesSignature and assigns it to the Signature field.
func (o *Images) SetSignature(v ImagesSignature) {
	o.Signature = &v
}

func (o Images) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Images) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.CarrierReturnWindow) {
		toSerialize["carrier_return_window"] = o.CarrierReturnWindow
	}
	if !IsNil(o.Signature) {
		toSerialize["signature"] = o.Signature
	}
	return toSerialize, nil
}

type NullableImages struct {
	value *Images
	isSet bool
}

func (v NullableImages) Get() *Images {
	return v.value
}

func (v *NullableImages) Set(val *Images) {
	v.value = val
	v.isSet = true
}

func (v NullableImages) IsSet() bool {
	return v.isSet
}

func (v *NullableImages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImages(val *Images) *NullableImages {
	return &NullableImages{value: val, isSet: true}
}

func (v NullableImages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


