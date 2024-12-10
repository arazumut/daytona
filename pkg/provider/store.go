// 2024 Daytona Platforms Inc. Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package provider

import "errors"

// TargetFilter, hedefleri filtrelemek için kullanılır
type TargetFilter struct {
	Isim       *string
	Varsayılan *bool
}

// TargetStore, hedefleri yönetmek için kullanılan arayüzdür
type TargetStore interface {
	List(filtre *TargetFilter) ([]*ProviderTarget, error) // Hedefleri listele
	Find(filtre *TargetFilter) (*ProviderTarget, error)   // Hedef bul
	Save(hedef *ProviderTarget) error                     // Hedef kaydet
	Delete(hedef *ProviderTarget) error                   // Hedef sil
}

var (
	ErrHedefBulunamadı = errors.New("hedef bulunamadı")
)

// IsTargetNotFound, verilen hatanın hedef bulunamadı hatası olup olmadığını kontrol eder
func IsTargetNotFound(err error) bool {
	return err.Error() == ErrHedefBulunamadı.Error()
}
