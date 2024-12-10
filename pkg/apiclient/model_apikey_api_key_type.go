/*
Daytona Sunucu API'si

Daytona Sunucu API'si

API sürümü: v0.0.0-dev
*/

// Bu kod OpenAPI Generator (https://openapi-generator.tech) tarafından oluşturulmuştur; LÜTFEN DÜZENLEMEYİN.

package apiclient

import (
	"encoding/json"
	"fmt"
)

// ApikeyApiKeyType modeli 'ApikeyApiKeyType'
type ApikeyApiKeyType string

// ApikeyApiKeyType enum değerlerinin listesi
const (
	ApiKeyTypeClient    ApikeyApiKeyType = "client"
	ApiKeyTypeProject   ApikeyApiKeyType = "project"
	ApiKeyTypeWorkspace ApikeyApiKeyType = "workspace"
)

// ApikeyApiKeyType enumunun tüm izin verilen değerleri
var AllowedApikeyApiKeyTypeEnumValues = []ApikeyApiKeyType{
	"client",
	"project",
	"workspace",
}

func (v *ApikeyApiKeyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApikeyApiKeyType(value)
	for _, existing := range AllowedApikeyApiKeyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v geçerli bir ApikeyApiKeyType değil", value)
}

// NewApikeyApiKeyTypeFromValue, argüman olarak geçirilen değer için geçerli bir ApikeyApiKeyType işaretçisi döndürür
// veya geçirilen değer enum tarafından izin verilmiyorsa bir hata döndürür
func NewApikeyApiKeyTypeFromValue(v string) (*ApikeyApiKeyType, error) {
	ev := ApikeyApiKeyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("ApikeyApiKeyType için '%v' geçersiz değer: geçerli değerler %v", v, AllowedApikeyApiKeyTypeEnumValues)
	}
}

// IsValid, değerin enum için geçerli olup olmadığını döndürür, aksi takdirde false döner
func (v ApikeyApiKeyType) IsValid() bool {
	for _, existing := range AllowedApikeyApiKeyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr, apikey.ApiKeyType değerine referans döner
func (v ApikeyApiKeyType) Ptr() *ApikeyApiKeyType {
	return &v
}

type NullableApikeyApiKeyType struct {
	value *ApikeyApiKeyType
	isSet bool
}

func (v NullableApikeyApiKeyType) Get() *ApikeyApiKeyType {
	return v.value
}

func (v *NullableApikeyApiKeyType) Set(val *ApikeyApiKeyType) {
	v.value = val
	v.isSet = true
}

func (v NullableApikeyApiKeyType) IsSet() bool {
	return v.isSet
}

func (v *NullableApikeyApiKeyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApikeyApiKeyType(val *ApikeyApiKeyType) *NullableApikeyApiKeyType {
	return &NullableApikeyApiKeyType{value: val, isSet: true}
}

func (v NullableApikeyApiKeyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApikeyApiKeyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
