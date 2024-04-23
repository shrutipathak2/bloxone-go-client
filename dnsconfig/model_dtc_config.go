/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the DTCConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DTCConfig{}

// DTCConfig Construct for fields: _default_ttl_.
type DTCConfig struct {
	// Optional. Default TTL for synthesized DTC records (value in seconds).  Defaults to 300.
	DefaultTtl *int64 `json:"default_ttl,omitempty"`
}

// NewDTCConfig instantiates a new DTCConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDTCConfig() *DTCConfig {
	this := DTCConfig{}
	return &this
}

// NewDTCConfigWithDefaults instantiates a new DTCConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDTCConfigWithDefaults() *DTCConfig {
	this := DTCConfig{}
	return &this
}

// GetDefaultTtl returns the DefaultTtl field value if set, zero value otherwise.
func (o *DTCConfig) GetDefaultTtl() int64 {
	if o == nil || IsNil(o.DefaultTtl) {
		var ret int64
		return ret
	}
	return *o.DefaultTtl
}

// GetDefaultTtlOk returns a tuple with the DefaultTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DTCConfig) GetDefaultTtlOk() (*int64, bool) {
	if o == nil || IsNil(o.DefaultTtl) {
		return nil, false
	}
	return o.DefaultTtl, true
}

// HasDefaultTtl returns a boolean if a field has been set.
func (o *DTCConfig) HasDefaultTtl() bool {
	if o != nil && !IsNil(o.DefaultTtl) {
		return true
	}

	return false
}

// SetDefaultTtl gets a reference to the given int64 and assigns it to the DefaultTtl field.
func (o *DTCConfig) SetDefaultTtl(v int64) {
	o.DefaultTtl = &v
}

func (o DTCConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DTCConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultTtl) {
		toSerialize["default_ttl"] = o.DefaultTtl
	}
	return toSerialize, nil
}

type NullableDTCConfig struct {
	value *DTCConfig
	isSet bool
}

func (v NullableDTCConfig) Get() *DTCConfig {
	return v.value
}

func (v *NullableDTCConfig) Set(val *DTCConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDTCConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDTCConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDTCConfig(val *DTCConfig) *NullableDTCConfig {
	return &NullableDTCConfig{value: val, isSet: true}
}

func (v NullableDTCConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDTCConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}