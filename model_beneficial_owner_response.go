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
)

// checks if the BeneficialOwnerResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeneficialOwnerResponse{}

// BeneficialOwnerResponse Contains information about the beneficial owner of the business, if applicable.
type BeneficialOwnerResponse struct {
	// First name of the beneficial owner.  This field is returned if it exists in the resource.
	FirstName *string `json:"first_name,omitempty"`
	// Date of birth of the beneficial owner.  This field is returned if it exists in the resource.
	Getdob *time.Time `json:"getdob,omitempty"`
	Home *AddressResponseModel `json:"home,omitempty"`
	// Last name of the beneficial owner.  This field is returned if it exists in the resource.
	LastName *string `json:"last_name,omitempty"`
	// Middle name of the beneficial owner.  This field is returned if it exists in the resource.
	MiddleName *string `json:"middle_name,omitempty"`
	// Ten-digit phone number of the beneficial owner.  This field is returned if it exists in the resource.
	Phone *string `json:"phone,omitempty"`
	// Title of the beneficial owner.  This field is returned if it exists in the resource.
	Title *string `json:"title,omitempty"`
}

// NewBeneficialOwnerResponse instantiates a new BeneficialOwnerResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeneficialOwnerResponse() *BeneficialOwnerResponse {
	this := BeneficialOwnerResponse{}
	return &this
}

// NewBeneficialOwnerResponseWithDefaults instantiates a new BeneficialOwnerResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeneficialOwnerResponseWithDefaults() *BeneficialOwnerResponse {
	this := BeneficialOwnerResponse{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *BeneficialOwnerResponse) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerResponse) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *BeneficialOwnerResponse) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *BeneficialOwnerResponse) SetFirstName(v string) {
	o.FirstName = &v
}

// GetGetdob returns the Getdob field value if set, zero value otherwise.
func (o *BeneficialOwnerResponse) GetGetdob() time.Time {
	if o == nil || IsNil(o.Getdob) {
		var ret time.Time
		return ret
	}
	return *o.Getdob
}

// GetGetdobOk returns a tuple with the Getdob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerResponse) GetGetdobOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Getdob) {
		return nil, false
	}
	return o.Getdob, true
}

// HasGetdob returns a boolean if a field has been set.
func (o *BeneficialOwnerResponse) HasGetdob() bool {
	if o != nil && !IsNil(o.Getdob) {
		return true
	}

	return false
}

// SetGetdob gets a reference to the given time.Time and assigns it to the Getdob field.
func (o *BeneficialOwnerResponse) SetGetdob(v time.Time) {
	o.Getdob = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *BeneficialOwnerResponse) GetHome() AddressResponseModel {
	if o == nil || IsNil(o.Home) {
		var ret AddressResponseModel
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerResponse) GetHomeOk() (*AddressResponseModel, bool) {
	if o == nil || IsNil(o.Home) {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *BeneficialOwnerResponse) HasHome() bool {
	if o != nil && !IsNil(o.Home) {
		return true
	}

	return false
}

// SetHome gets a reference to the given AddressResponseModel and assigns it to the Home field.
func (o *BeneficialOwnerResponse) SetHome(v AddressResponseModel) {
	o.Home = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *BeneficialOwnerResponse) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerResponse) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *BeneficialOwnerResponse) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *BeneficialOwnerResponse) SetLastName(v string) {
	o.LastName = &v
}

// GetMiddleName returns the MiddleName field value if set, zero value otherwise.
func (o *BeneficialOwnerResponse) GetMiddleName() string {
	if o == nil || IsNil(o.MiddleName) {
		var ret string
		return ret
	}
	return *o.MiddleName
}

// GetMiddleNameOk returns a tuple with the MiddleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerResponse) GetMiddleNameOk() (*string, bool) {
	if o == nil || IsNil(o.MiddleName) {
		return nil, false
	}
	return o.MiddleName, true
}

// HasMiddleName returns a boolean if a field has been set.
func (o *BeneficialOwnerResponse) HasMiddleName() bool {
	if o != nil && !IsNil(o.MiddleName) {
		return true
	}

	return false
}

// SetMiddleName gets a reference to the given string and assigns it to the MiddleName field.
func (o *BeneficialOwnerResponse) SetMiddleName(v string) {
	o.MiddleName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *BeneficialOwnerResponse) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerResponse) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *BeneficialOwnerResponse) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *BeneficialOwnerResponse) SetPhone(v string) {
	o.Phone = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BeneficialOwnerResponse) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeneficialOwnerResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BeneficialOwnerResponse) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BeneficialOwnerResponse) SetTitle(v string) {
	o.Title = &v
}

func (o BeneficialOwnerResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeneficialOwnerResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.Getdob) {
		toSerialize["getdob"] = o.Getdob
	}
	if !IsNil(o.Home) {
		toSerialize["home"] = o.Home
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.MiddleName) {
		toSerialize["middle_name"] = o.MiddleName
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableBeneficialOwnerResponse struct {
	value *BeneficialOwnerResponse
	isSet bool
}

func (v NullableBeneficialOwnerResponse) Get() *BeneficialOwnerResponse {
	return v.value
}

func (v *NullableBeneficialOwnerResponse) Set(val *BeneficialOwnerResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBeneficialOwnerResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBeneficialOwnerResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeneficialOwnerResponse(val *BeneficialOwnerResponse) *NullableBeneficialOwnerResponse {
	return &NullableBeneficialOwnerResponse{value: val, isSet: true}
}

func (v NullableBeneficialOwnerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeneficialOwnerResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


