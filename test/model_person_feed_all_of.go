/*
 * Dummy
 *
 * Dummy
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PersonFeedAllOf struct for PersonFeedAllOf
type PersonFeedAllOf struct {
	Name *string `json:"name,omitempty"`
}

// NewPersonFeedAllOf instantiates a new PersonFeedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonFeedAllOf() *PersonFeedAllOf {
	this := PersonFeedAllOf{}
	return &this
}

// NewPersonFeedAllOfWithDefaults instantiates a new PersonFeedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonFeedAllOfWithDefaults() *PersonFeedAllOf {
	this := PersonFeedAllOf{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PersonFeedAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonFeedAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PersonFeedAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PersonFeedAllOf) SetName(v string) {
	o.Name = &v
}

func (o PersonFeedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullablePersonFeedAllOf struct {
	value *PersonFeedAllOf
	isSet bool
}

func (v NullablePersonFeedAllOf) Get() *PersonFeedAllOf {
	return v.value
}

func (v *NullablePersonFeedAllOf) Set(val *PersonFeedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonFeedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonFeedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonFeedAllOf(val *PersonFeedAllOf) *NullablePersonFeedAllOf {
	return &NullablePersonFeedAllOf{value: val, isSet: true}
}

func (v NullablePersonFeedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonFeedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

