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

// ArticleFeedAllOf struct for ArticleFeedAllOf
type ArticleFeedAllOf struct {
	Title *string `json:"title,omitempty"`
}

// NewArticleFeedAllOf instantiates a new ArticleFeedAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArticleFeedAllOf() *ArticleFeedAllOf {
	this := ArticleFeedAllOf{}
	return &this
}

// NewArticleFeedAllOfWithDefaults instantiates a new ArticleFeedAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArticleFeedAllOfWithDefaults() *ArticleFeedAllOf {
	this := ArticleFeedAllOf{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ArticleFeedAllOf) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArticleFeedAllOf) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ArticleFeedAllOf) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ArticleFeedAllOf) SetTitle(v string) {
	o.Title = &v
}

func (o ArticleFeedAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableArticleFeedAllOf struct {
	value *ArticleFeedAllOf
	isSet bool
}

func (v NullableArticleFeedAllOf) Get() *ArticleFeedAllOf {
	return v.value
}

func (v *NullableArticleFeedAllOf) Set(val *ArticleFeedAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableArticleFeedAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableArticleFeedAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArticleFeedAllOf(val *ArticleFeedAllOf) *NullableArticleFeedAllOf {
	return &NullableArticleFeedAllOf{value: val, isSet: true}
}

func (v NullableArticleFeedAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArticleFeedAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


