/*
Daytona Sunucusu API

Daytona Sunucusu API

API sürümü: v0.0.0-dev
*/

// Bu kod OpenAPI Generator (https://openapi-generator.tech) tarafından oluşturulmuştur; LÜTFEN DÜZENLEMEYİN.

package apiclient

import (
	"encoding/json"
	"fmt"
)

// CloneTarget modeli 'CloneTarget'
type CloneTarget string

// CloneTarget değerlerinin listesi
const (
	CloneTargetBranch CloneTarget = "branch"
	CloneTargetCommit CloneTarget = "commit"
)

// CloneTarget enumunun izin verilen tüm değerleri
var AllowedCloneTargetEnumValues = []CloneTarget{
	"branch",
	"commit",
}

func (v *CloneTarget) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CloneTarget(value)
	for _, existing := range AllowedCloneTargetEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v geçerli bir CloneTarget değil", value)
}

// NewCloneTargetFromValue, argüman olarak geçirilen değer için geçerli bir CloneTarget işaretçisi döndürür
// veya değer enum tarafından izin verilmiyorsa bir hata döndürür
func NewCloneTargetFromValue(v string) (*CloneTarget, error) {
	ev := CloneTarget(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("CloneTarget için '%v' geçersiz değer: geçerli değerler %v", v, AllowedCloneTargetEnumValues)
	}
}

// IsValid, değerin enum için geçerli olup olmadığını döndürür, geçerli ise true, aksi halde false
func (v CloneTarget) IsValid() bool {
	for _, existing := range AllowedCloneTargetEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr, CloneTarget değerine referans döndürür
func (v CloneTarget) Ptr() *CloneTarget {
	return &v
}

type NullableCloneTarget struct {
	value *CloneTarget
	isSet bool
}

func (v NullableCloneTarget) Get() *CloneTarget {
	return v.value
}

func (v *NullableCloneTarget) Set(val *CloneTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCloneTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCloneTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloneTarget(val *CloneTarget) *NullableCloneTarget {
	return &NullableCloneTarget{value: val, isSet: true}
}

func (v NullableCloneTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloneTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
