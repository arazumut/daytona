// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"reflect"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// defaultValidator yapısı, tek seferlik başlatma ve doğrulama nesnesini içerir.
type defaultValidator struct {
	once     sync.Once
	validate *validator.Validate
}

// defaultValidator'ın binding.StructValidator arayüzünü uyguladığını belirtir.
var _ binding.StructValidator = &defaultValidator{}

// SliceValidationError, birden fazla hatayı içeren özel bir hata türüdür.
type SliceValidationError []error

// Error metodu, SliceValidationError içindeki hataları birleştirir ve döner.
func (err SliceValidationError) Error() string {
	if len(err) == 0 {
		return ""
	}

	var b strings.Builder
	for i := 0; i < len(err); i++ {
		if err[i] != nil {
			if b.Len() > 0 {
				b.WriteString("\n")
			}
			b.WriteString("[" + strconv.Itoa(i) + "]: " + err[i].Error())
		}
	}
	return b.String()
}

// ValidateStruct metodu, verilen nesneyi doğrular.
func (v *defaultValidator) ValidateStruct(obj any) error {
	if obj == nil {
		return nil
	}

	value := reflect.ValueOf(obj)
	switch value.Kind() {
	case reflect.Ptr:
		if value.Elem().Kind() != reflect.Struct {
			return v.ValidateStruct(value.Elem().Interface())
		}
		return v.validateStruct(obj)
	case reflect.Struct:
		return v.validateStruct(obj)
	case reflect.Slice, reflect.Array:
		count := value.Len()
		validateRet := make(SliceValidationError, 0)
		for i := 0; i < count; i++ {
			if err := v.ValidateStruct(value.Index(i).Interface()); err != nil {
				validateRet = append(validateRet, err)
			}
		}
		if len(validateRet) == 0 {
			return nil
		}
		return validateRet
	default:
		return nil
	}
}

// Engine metodu, doğrulama motorunu döner.
func (v *defaultValidator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

// validateStruct metodu, verilen nesneyi doğrular.
func (v *defaultValidator) validateStruct(obj any) error {
	v.lazyinit()
	return v.validate.Struct(obj)
}

// lazyinit metodu, doğrulama nesnesini tek seferlik başlatır.
func (v *defaultValidator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New(validator.WithRequiredStructEnabled())
		v.validate.SetTagName("validate")
		_ = v.validate.RegisterValidation("optional", func(fl validator.FieldLevel) bool {
			return true
		}, true)
	})
}
