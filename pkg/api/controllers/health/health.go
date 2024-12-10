// 2024 Daytona Platforms Inc. Tüm hakları saklıdır.
// SPDX-License-Identifier: Apache-2.0

package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SağlıkKontrolü 			godoc
//
//	@Özet			Sağlık kontrolü
//	@Açıklama		Sağlık kontrolü
//	@Üret			json
//	@Başarı			200	{object}	map[string]string
//	@Yol			/health [get]
//
//	@id				SağlıkKontrolü
func SağlıkKontrolü(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"durum": "tamam"})
}
