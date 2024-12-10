/*
Daytona Sunucusu API

Daytona Sunucusu API

API sürümü: v0.0.0-dev
*/

// OpenAPI Generator (https://openapi-generator.tech) tarafından oluşturulmuştur; DÜZENLEMEYİN.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CachedBuild türünün MappedNullable arayüzünü derleme zamanında karşıladığını kontrol eder
var _ MappedNullable = &CachedBuild{}

// CachedBuild yapısı
type CachedBuild struct {
	Image string `json:"image"`
	User  string `json:"user"`
}

type _CachedBuild CachedBuild

// NewCachedBuild yeni bir CachedBuild nesnesi oluşturur
// Bu yapıcı, tanımlanmış varsayılan değerlere sahip özelliklere varsayılan değerler atar
// ve API tarafından gerekli olan özelliklerin ayarlandığından emin olur.
func NewCachedBuild(image string, user string) *CachedBuild {
	this := CachedBuild{}
	this.Image = image
	this.User = user
	return &this
}

// NewCachedBuildWithDefaults yeni bir CachedBuild nesnesi oluşturur
// Bu yapıcı, yalnızca tanımlanmış varsayılan değerlere sahip özelliklere varsayılan değerler atar,
// ancak API tarafından gerekli olan özelliklerin ayarlandığını garanti etmez.
func NewCachedBuildWithDefaults() *CachedBuild {
	this := CachedBuild{}
	return &this
}

// GetImage Image alanının değerini döndürür
func (o *CachedBuild) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk Image alanının değerini ve değerin ayarlanıp ayarlanmadığını kontrol etmek için bir boolean döndürür.
func (o *CachedBuild) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage alan değerini ayarlar
func (o *CachedBuild) SetImage(v string) {
	o.Image = v
}

// GetUser User alanının değerini döndürür
func (o *CachedBuild) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk User alanının değerini ve değerin ayarlanıp ayarlanmadığını kontrol etmek için bir boolean döndürür.
func (o *CachedBuild) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser alan değerini ayarlar
func (o *CachedBuild) SetUser(v string) {
	o.User = v
}

func (o CachedBuild) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CachedBuild) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image"] = o.Image
	toSerialize["user"] = o.User
	return toSerialize, nil
}

func (o *CachedBuild) UnmarshalJSON(data []byte) (err error) {
	// Bu, tüm gerekli özelliklerin JSON nesnesine dahil edildiğini doğrular
	// nesneyi string anahtarlara sahip genel bir haritaya ayrıştırarak ve
	// her gerekli alanın genel haritada bir anahtar olarak mevcut olduğunu kontrol ederek.
	gerekliOzellikler := []string{
		"image",
		"user",
	}

	tumOzellikler := make(map[string]interface{})

	err = json.Unmarshal(data, &tumOzellikler)

	if err != nil {
		return err
	}

	for _, gerekliOzellik := range gerekliOzellikler {
		if _, exists := tumOzellikler[gerekliOzellik]; !exists {
			return fmt.Errorf("gerekli özellik için değer verilmedi %v", gerekliOzellik)
		}
	}

	varCachedBuild := _CachedBuild{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCachedBuild)

	if err != nil {
		return err
	}

	*o = CachedBuild(varCachedBuild)

	return err
}

type NullableCachedBuild struct {
	value *CachedBuild
	isSet bool
}

func (v NullableCachedBuild) Get() *CachedBuild {
	return v.value
}

func (v *NullableCachedBuild) Set(val *CachedBuild) {
	v.value = val
	v.isSet = true
}

func (v NullableCachedBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableCachedBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCachedBuild(val *CachedBuild) *NullableCachedBuild {
	return &NullableCachedBuild{value: val, isSet: true}
}

func (v NullableCachedBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCachedBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
