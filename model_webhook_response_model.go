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

// checks if the WebhookResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookResponseModel{}

// WebhookResponseModel struct for WebhookResponseModel
type WebhookResponseModel struct {
	// Indicates whether the webhook is active. This field is returned if you included it in your webhook.
	Active *bool `json:"active,omitempty"`
	Config WebhookConfigModel `json:"config"`
	// Date and time when the webhook event was created, in UTC.
	CreatedTime *time.Time `json:"created_time,omitempty"`
	// Specifies the types of events for which notifications are sent.  The wildcard character `\\*` indicates that you receive all webhook notifications, or all notifications of a specified category. For example, `*` indicates that you receive all webhook notifications; `transaction.*` indicates that you receive all `transaction` webhook notifications.  *NOTE:* You can only use the wildcard character with the _base_ type events, not subcategories. For example, you cannot subscribe to `cardtransition.fulfillment.\\*` events, but you can subscribe to `cardtransition.*`.
	Events []string `json:"events"`
	// Date and time when the associated object was last modified, in UTC.
	LastModifiedTime *time.Time `json:"last_modified_time,omitempty"`
	// Descriptive name of the webhook.
	Name string `json:"name"`
	// Unique identifier of the webhook.
	Token *string `json:"token,omitempty"`
}

type _WebhookResponseModel WebhookResponseModel

// NewWebhookResponseModel instantiates a new WebhookResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookResponseModel(config WebhookConfigModel, events []string, name string) *WebhookResponseModel {
	this := WebhookResponseModel{}
	var active bool = true
	this.Active = &active
	this.Config = config
	this.Events = events
	this.Name = name
	return &this
}

// NewWebhookResponseModelWithDefaults instantiates a new WebhookResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookResponseModelWithDefaults() *WebhookResponseModel {
	this := WebhookResponseModel{}
	var active bool = true
	this.Active = &active
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *WebhookResponseModel) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseModel) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *WebhookResponseModel) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *WebhookResponseModel) SetActive(v bool) {
	o.Active = &v
}

// GetConfig returns the Config field value
func (o *WebhookResponseModel) GetConfig() WebhookConfigModel {
	if o == nil {
		var ret WebhookConfigModel
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseModel) GetConfigOk() (*WebhookConfigModel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *WebhookResponseModel) SetConfig(v WebhookConfigModel) {
	o.Config = v
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *WebhookResponseModel) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseModel) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *WebhookResponseModel) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *WebhookResponseModel) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetEvents returns the Events field value
func (o *WebhookResponseModel) GetEvents() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseModel) GetEventsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *WebhookResponseModel) SetEvents(v []string) {
	o.Events = v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *WebhookResponseModel) GetLastModifiedTime() time.Time {
	if o == nil || IsNil(o.LastModifiedTime) {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseModel) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastModifiedTime) {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *WebhookResponseModel) HasLastModifiedTime() bool {
	if o != nil && !IsNil(o.LastModifiedTime) {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *WebhookResponseModel) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetName returns the Name field value
func (o *WebhookResponseModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhookResponseModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebhookResponseModel) SetName(v string) {
	o.Name = v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *WebhookResponseModel) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookResponseModel) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *WebhookResponseModel) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *WebhookResponseModel) SetToken(v string) {
	o.Token = &v
}

func (o WebhookResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	toSerialize["config"] = o.Config
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	toSerialize["events"] = o.Events
	if !IsNil(o.LastModifiedTime) {
		toSerialize["last_modified_time"] = o.LastModifiedTime
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	return toSerialize, nil
}

func (o *WebhookResponseModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"config",
		"events",
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

	varWebhookResponseModel := _WebhookResponseModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookResponseModel)

	if err != nil {
		return err
	}

	*o = WebhookResponseModel(varWebhookResponseModel)

	return err
}

type NullableWebhookResponseModel struct {
	value *WebhookResponseModel
	isSet bool
}

func (v NullableWebhookResponseModel) Get() *WebhookResponseModel {
	return v.value
}

func (v *NullableWebhookResponseModel) Set(val *WebhookResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookResponseModel(val *WebhookResponseModel) *NullableWebhookResponseModel {
	return &NullableWebhookResponseModel{value: val, isSet: true}
}

func (v NullableWebhookResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


