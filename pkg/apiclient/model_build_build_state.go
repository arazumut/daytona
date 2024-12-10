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

// BuildBuildState modeli 'BuildBuildState'
type BuildBuildState string

// BuildBuildState enum değerlerinin listesi
const (
	BuildStatePendingRun          BuildBuildState = "pending-run"
	BuildStateRunning             BuildBuildState = "running"
	BuildStateError               BuildBuildState = "error"
	BuildStateSuccess             BuildBuildState = "success"
	BuildStatePublished           BuildBuildState = "published"
	BuildStatePendingDelete       BuildBuildState = "pending-delete"
	BuildStatePendingForcedDelete BuildBuildState = "pending-forced-delete"
	BuildStateDeleting            BuildBuildState = "deleting"
)

// BuildBuildState enumunun tüm izin verilen değerleri
var AllowedBuildBuildStateEnumValues = []BuildBuildState{
	"pending-run",
	"running",
	"error",
	"success",
	"published",
	"pending-delete",
	"pending-forced-delete",
	"deleting",
}

// UnmarshalJSON, JSON verisini BuildBuildState türüne dönüştürür
func (v *BuildBuildState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BuildBuildState(value)
	for _, existing := range AllowedBuildBuildStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v geçerli bir BuildBuildState değil", value)
}

// NewBuildBuildStateFromValue, verilen değerden geçerli bir BuildBuildState işaretçisi döndürür
// veya değer izin verilen enum tarafından kabul edilmiyorsa hata döner
func NewBuildBuildStateFromValue(v string) (*BuildBuildState, error) {
	ev := BuildBuildState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("BuildBuildState için '%v' geçersiz değer: geçerli değerler %v", v, AllowedBuildBuildStateEnumValues)
	}
}

// IsValid, değerin enum için geçerli olup olmadığını döner
func (v BuildBuildState) IsValid() bool {
	for _, existing := range AllowedBuildBuildStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr, BuildBuildState değerine referans döner
func (v BuildBuildState) Ptr() *BuildBuildState {
	return &v
}

type NullableBuildBuildState struct {
	value *BuildBuildState
	isSet bool
}

func (v NullableBuildBuildState) Get() *BuildBuildState {
	return v.value
}

func (v *NullableBuildBuildState) Set(val *BuildBuildState) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBuildState) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBuildState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBuildState(val *BuildBuildState) *NullableBuildBuildState {
	return &NullableBuildBuildState{value: val, isSet: true}
}

func (v NullableBuildBuildState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBuildState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
