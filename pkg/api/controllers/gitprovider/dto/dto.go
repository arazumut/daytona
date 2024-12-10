// 2024 Daytona Platforms Inc. Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package dto

import (
	"github.com/daytonaio/daytona/pkg/gitprovider"
)

// RepositoryUrl yapısı, bir depo URL'sini temsil eder
type RepositoryUrl struct {
	URL string `json:"url" validate:"required"`
} // @name RepositoryUrl

// SetGitProviderConfig yapısı, Git sağlayıcı yapılandırma ayarlarını temsil eder
type SetGitProviderConfig struct {
	Id            string                     `json:"id" validate:"optional"`                      // Yapılandırma kimliği
	ProviderId    string                     `json:"providerId" validate:"required"`              // Sağlayıcı kimliği
	Username      *string                    `json:"username,omitempty" validate:"optional"`      // Kullanıcı adı (isteğe bağlı)
	Token         string                     `json:"token" validate:"required"`                   // Erişim belirteci
	BaseApiUrl    *string                    `json:"baseApiUrl,omitempty" validate:"optional"`    // Temel API URL'si (isteğe bağlı)
	Alias         *string                    `json:"alias,omitempty" validate:"optional"`         // Takma ad (isteğe bağlı)
	SigningKey    *string                    `json:"signingKey,omitempty" validate:"optional"`    // İmzalama anahtarı (isteğe bağlı)
	SigningMethod *gitprovider.SigningMethod `json:"signingMethod,omitempty" validate:"optional"` // İmzalama yöntemi (isteğe bağlı)
} // @name SetGitProviderConfig
