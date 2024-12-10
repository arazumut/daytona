// Daytona Platforms Inc. 2024 Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// GetTargetManifest godoc
//
//	@Tags			provider
//	@Summary		Provider hedef manifestosunu al
//	@Description	Provider hedef manifestosunu al
//	@Param			provider	path	string	true	"Provider adı"
//	@Success		200
//	@Success		200	{object}	ProviderTargetManifest
//	@Router			/provider/{provider}/target-manifest [get]
//
//	@id				GetTargetManifest
func GetTargetManifest(ctx *gin.Context) {
	providerName := ctx.Param("provider")

	server := server.GetInstance(nil)

	p, err := server.ProviderManager.GetProvider(providerName)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("provider bulunamadı: %w", err))
		return
	}

	manifest, err := (*p).GetTargetManifest()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("provider manifestosu alınamadı: %w", err))
		return
	}

	ctx.JSON(200, manifest)
}
