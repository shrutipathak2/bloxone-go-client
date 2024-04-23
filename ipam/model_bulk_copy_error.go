/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the BulkCopyError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkCopyError{}

// BulkCopyError struct for BulkCopyError
type BulkCopyError struct {
	// The description of the resource that was requested to be copied.
	Description *string `json:"description,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The reason why the copy failed.
	Message *string `json:"message,omitempty"`
}

// NewBulkCopyError instantiates a new BulkCopyError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkCopyError() *BulkCopyError {
	this := BulkCopyError{}
	return &this
}

// NewBulkCopyErrorWithDefaults instantiates a new BulkCopyError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkCopyErrorWithDefaults() *BulkCopyError {
	this := BulkCopyError{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BulkCopyError) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BulkCopyError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BulkCopyError) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BulkCopyError) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyError) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BulkCopyError) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BulkCopyError) SetId(v string) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *BulkCopyError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkCopyError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *BulkCopyError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *BulkCopyError) SetMessage(v string) {
	o.Message = &v
}

func (o BulkCopyError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkCopyError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableBulkCopyError struct {
	value *BulkCopyError
	isSet bool
}

func (v NullableBulkCopyError) Get() *BulkCopyError {
	return v.value
}

func (v *NullableBulkCopyError) Set(val *BulkCopyError) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkCopyError) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkCopyError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkCopyError(val *BulkCopyError) *NullableBulkCopyError {
	return &NullableBulkCopyError{value: val, isSet: true}
}

func (v NullableBulkCopyError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkCopyError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}