// Daytona Platforms Inc. 2024 Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// UninstallProvider godoc
//
//	@Tags			provider
//	@Summary		Bir sağlayıcıyı kaldır
//	@Description	Bir sağlayıcıyı kaldır
//	@Accept			json
//	@Param			provider	path	string	true	"Kaldırılacak sağlayıcı"
//	@Success		200
//	@Router			/provider/{provider}/uninstall [post]
//
//	@id				UninstallProvider
func UninstallProvider(ctx *gin.Context) {
	provider := ctx.Param("provider")

	server := server.GetInstance(nil)

	err := server.ProviderManager.UninstallProvider(provider)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("sağlayıcı kaldırılamadı: %w", err))
		return
	}

	ctx.Status(200)
}
