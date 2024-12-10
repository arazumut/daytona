// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/api/controllers/provider/dto"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// ListProviders godoc
//
//	@Tags			provider
//	@Summary		Servis sağlayıcıları listele
//	@Description	Servis sağlayıcıları listele
//	@Produce		json
//	@Success		200	{array}	dto.Provider
//	@Router			/provider [get]
//
//	@id				ListProviders
func ListProviders(ctx *gin.Context) {
	server := server.GetInstance(nil)
	providers := server.ProviderManager.GetProviders()

	sonuc := []dto.Provider{}
	for _, provider := range providers {
		bilgi, err := provider.GetInfo()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("sağlayıcı alınamadı: %w", err))
			return
		}

		sonuc = append(sonuc, dto.Provider{
			Name:    bilgi.Name,
			Label:   bilgi.Label,
			Version: bilgi.Version,
		})
	}

	ctx.JSON(200, sonuc)
}
