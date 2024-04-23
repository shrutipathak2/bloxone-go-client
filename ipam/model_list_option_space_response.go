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

// checks if the ListOptionSpaceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOptionSpaceResponse{}

// ListOptionSpaceResponse The response format to retrieve __OptionSpace__ objects.
type ListOptionSpaceResponse struct {
	// The list of OptionSpace objects.
	Results []OptionSpace `json:"results,omitempty"`
}

// NewListOptionSpaceResponse instantiates a new ListOptionSpaceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOptionSpaceResponse() *ListOptionSpaceResponse {
	this := ListOptionSpaceResponse{}
	return &this
}

// NewListOptionSpaceResponseWithDefaults instantiates a new ListOptionSpaceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOptionSpaceResponseWithDefaults() *ListOptionSpaceResponse {
	this := ListOptionSpaceResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListOptionSpaceResponse) GetResults() []OptionSpace {
	if o == nil || IsNil(o.Results) {
		var ret []OptionSpace
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOptionSpaceResponse) GetResultsOk() ([]OptionSpace, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListOptionSpaceResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []OptionSpace and assigns it to the Results field.
func (o *ListOptionSpaceResponse) SetResults(v []OptionSpace) {
	o.Results = v
}

func (o ListOptionSpaceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOptionSpaceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableListOptionSpaceResponse struct {
	value *ListOptionSpaceResponse
	isSet bool
}

func (v NullableListOptionSpaceResponse) Get() *ListOptionSpaceResponse {
	return v.value
}

func (v *NullableListOptionSpaceResponse) Set(val *ListOptionSpaceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOptionSpaceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOptionSpaceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOptionSpaceResponse(val *ListOptionSpaceResponse) *NullableListOptionSpaceResponse {
	return &NullableListOptionSpaceResponse{value: val, isSet: true}
}

func (v NullableListOptionSpaceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOptionSpaceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}