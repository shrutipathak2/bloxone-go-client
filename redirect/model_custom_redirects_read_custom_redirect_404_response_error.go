/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"encoding/json"
)

// checks if the CustomRedirectsReadCustomRedirect404ResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomRedirectsReadCustomRedirect404ResponseError{}

// CustomRedirectsReadCustomRedirect404ResponseError struct for CustomRedirectsReadCustomRedirect404ResponseError
type CustomRedirectsReadCustomRedirect404ResponseError struct {
	Code                 *string `json:"code,omitempty"`
	Message              *string `json:"message,omitempty"`
	Status               *string `json:"status,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomRedirectsReadCustomRedirect404ResponseError CustomRedirectsReadCustomRedirect404ResponseError

// NewCustomRedirectsReadCustomRedirect404ResponseError instantiates a new CustomRedirectsReadCustomRedirect404ResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomRedirectsReadCustomRedirect404ResponseError() *CustomRedirectsReadCustomRedirect404ResponseError {
	this := CustomRedirectsReadCustomRedirect404ResponseError{}
	return &this
}

// NewCustomRedirectsReadCustomRedirect404ResponseErrorWithDefaults instantiates a new CustomRedirectsReadCustomRedirect404ResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomRedirectsReadCustomRedirect404ResponseErrorWithDefaults() *CustomRedirectsReadCustomRedirect404ResponseError {
	this := CustomRedirectsReadCustomRedirect404ResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CustomRedirectsReadCustomRedirect404ResponseError) SetStatus(v string) {
	o.Status = &v
}

func (o CustomRedirectsReadCustomRedirect404ResponseError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomRedirectsReadCustomRedirect404ResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomRedirectsReadCustomRedirect404ResponseError) UnmarshalJSON(data []byte) (err error) {
	varCustomRedirectsReadCustomRedirect404ResponseError := _CustomRedirectsReadCustomRedirect404ResponseError{}

	err = json.Unmarshal(data, &varCustomRedirectsReadCustomRedirect404ResponseError)

	if err != nil {
		return err
	}

	*o = CustomRedirectsReadCustomRedirect404ResponseError(varCustomRedirectsReadCustomRedirect404ResponseError)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "message")
		delete(additionalProperties, "status")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomRedirectsReadCustomRedirect404ResponseError struct {
	value *CustomRedirectsReadCustomRedirect404ResponseError
	isSet bool
}

func (v NullableCustomRedirectsReadCustomRedirect404ResponseError) Get() *CustomRedirectsReadCustomRedirect404ResponseError {
	return v.value
}

func (v *NullableCustomRedirectsReadCustomRedirect404ResponseError) Set(val *CustomRedirectsReadCustomRedirect404ResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomRedirectsReadCustomRedirect404ResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomRedirectsReadCustomRedirect404ResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomRedirectsReadCustomRedirect404ResponseError(val *CustomRedirectsReadCustomRedirect404ResponseError) *NullableCustomRedirectsReadCustomRedirect404ResponseError {
	return &NullableCustomRedirectsReadCustomRedirect404ResponseError{value: val, isSet: true}
}

func (v NullableCustomRedirectsReadCustomRedirect404ResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomRedirectsReadCustomRedirect404ResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}