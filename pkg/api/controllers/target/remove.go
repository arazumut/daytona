// Daytona Platforms Inc. 2024 Telif Hakkı
// SPDX-Lisans: Apache-2.0

package target

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// RemoveTarget godoc
//
//	@Tags			target
//	@Summary		Bir hedefi kaldır
//	@Description	Bir hedefi kaldır
//	@Param			target	path	string	true	"Hedef adı"
//	@Success		204
//	@Router			/target/{target} [delete]
//
//	@id				RemoveTarget
func RemoveTarget(ctx *gin.Context) {
	hedefAdi := ctx.Param("target")

	server := server.GetInstance(nil)

	hedef, err := server.ProviderTargetService.Find(&provider.TargetFilter{
		Name: &hedefAdi,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("hedef bulunamadı: %w", err))
		return
	}

	err = server.ProviderTargetService.Delete(hedef)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("hedef kaldırılamadı: %w", err))
		return
	}

	ctx.Status(204)
}
