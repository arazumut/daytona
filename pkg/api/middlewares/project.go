// Daytona Platforms Inc. 2024 Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package middlewares

import (
	"errors"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// Proje yetkilendirme ara yazılımı
func ProjeYetkilendirmeAraYazılımı() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Authorization başlığını al
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" {
			// Eğer başlık yoksa yetkisiz hatası döndür
			ctx.AbortWithError(401, errors.New("yetkisiz"))
			return
		}

		// Token'ı ayıkla
		token := ExtractToken(bearerToken)
		if token == "" {
			// Eğer token yoksa yetkisiz hatası döndür
			ctx.AbortWithError(401, errors.New("yetkisiz"))
			return
		}

		// Sunucu örneğini al
		server := server.GetInstance(nil)

		// Token geçerli bir proje veya çalışma alanı API anahtarı değilse yetkisiz hatası döndür
		if !server.ApiKeyService.IsProjectApiKey(token) && !server.ApiKeyService.IsWorkspaceApiKey(token) {
			ctx.AbortWithError(401, errors.New("yetkisiz"))
		}

		// İşleme devam et
		ctx.Next()
	}
}
