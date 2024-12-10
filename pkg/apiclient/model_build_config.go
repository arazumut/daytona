/*
Daytona Sunucusu API

Daytona Sunucusu API

API versiyonu: v0.0.0-dev
*/

// Bu kod OpenAPI Generator (https://openapi-generator.tech) tarafından oluşturulmuştur; DÜZENLEMEYİN.

package apiclient

import (
	"encoding/json"
)

// BuildConfig türünün MappedNullable arayüzünü derleme zamanında karşıladığını kontrol eder
var _ MappedNullable = &BuildConfig{}

// BuildConfig yapısı
type BuildConfig struct {
	CachedBuild  *CachedBuild        `json:"cachedBuild,omitempty"`
	Devcontainer *DevcontainerConfig `json:"devcontainer,omitempty"`
}

// Yeni bir BuildConfig nesnesi oluşturur
// Bu yapıcı, tanımlanmış varsayılan değerlere sahip özelliklere varsayılan değerler atar
// ve API tarafından gerekli özelliklerin ayarlandığından emin olur, ancak gerekli özelliklerin seti değiştiğinde
// argümanların seti de değişir
func NewBuildConfig() *BuildConfig {
	this := BuildConfig{}
	return &this
}

// Varsayılanlarla yeni bir BuildConfig nesnesi oluşturur
// Bu yapıcı, yalnızca tanımlanmış varsayılan değerlere sahip özelliklere varsayılan değerler atar,
// ancak API tarafından gerekli özelliklerin ayarlandığını garanti etmez
func NewBuildConfigWithDefaults() *BuildConfig {
	this := BuildConfig{}
	return &this
}

// CachedBuild alanının değerini döndürür, ayarlanmamışsa sıfır değerini döndürür.
func (o *BuildConfig) GetCachedBuild() CachedBuild {
	if o == nil || IsNil(o.CachedBuild) {
		var ret CachedBuild
		return ret
	}
	return *o.CachedBuild
}

// CachedBuild alanının değerini döndüren bir ikili döndürür, ayarlanmamışsa nil döndürür
// ve değerin ayarlanıp ayarlanmadığını kontrol etmek için bir boolean döndürür.
func (o *BuildConfig) GetCachedBuildOk() (*CachedBuild, bool) {
	if o == nil || IsNil(o.CachedBuild) {
		return nil, false
	}
	return o.CachedBuild, true
}

// CachedBuild alanının ayarlanıp ayarlanmadığını kontrol eden bir boolean döndürür.
func (o *BuildConfig) HasCachedBuild() bool {
	if o != nil && !IsNil(o.CachedBuild) {
		return true
	}

	return false
}

// Verilen CachedBuild referansını alır ve CachedBuild alanına atar.
func (o *BuildConfig) SetCachedBuild(v CachedBuild) {
	o.CachedBuild = &v
}

// Devcontainer alanının değerini döndürür, ayarlanmamışsa sıfır değerini döndürür.
func (o *BuildConfig) GetDevcontainer() DevcontainerConfig {
	if o == nil || IsNil(o.Devcontainer) {
		var ret DevcontainerConfig
		return ret
	}
	return *o.Devcontainer
}

// Devcontainer alanının değerini döndüren bir ikili döndürür, ayarlanmamışsa nil döndürür
// ve değerin ayarlanıp ayarlanmadığını kontrol etmek için bir boolean döndürür.
func (o *BuildConfig) GetDevcontainerOk() (*DevcontainerConfig, bool) {
	if o == nil || IsNil(o.Devcontainer) {
		return nil, false
	}
	return o.Devcontainer, true
}

// Devcontainer alanının ayarlanıp ayarlanmadığını kontrol eden bir boolean döndürür.
func (o *BuildConfig) HasDevcontainer() bool {
	if o != nil && !IsNil(o.Devcontainer) {
		return true
	}

	return false
}

// Verilen DevcontainerConfig referansını alır ve Devcontainer alanına atar.
func (o *BuildConfig) SetDevcontainer(v DevcontainerConfig) {
	o.Devcontainer = &v
}

func (o BuildConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CachedBuild) {
		toSerialize["cachedBuild"] = o.CachedBuild
	}
	if !IsNil(o.Devcontainer) {
		toSerialize["devcontainer"] = o.Devcontainer
	}
	return toSerialize, nil
}

type NullableBuildConfig struct {
	value *BuildConfig
	isSet bool
}

func (v NullableBuildConfig) Get() *BuildConfig {
	return v.value
}

func (v *NullableBuildConfig) Set(val *BuildConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildConfig(val *BuildConfig) *NullableBuildConfig {
	return &NullableBuildConfig{value: val, isSet: true}
}

func (v NullableBuildConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
