// Daytona Platforms Inc. 2024 Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package containerregistry

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// RemoveContainerRegistry godoc
//
//	@Tags			container-registry
//	@Summary		Bir konteyner kayıt defteri kimlik bilgilerini kaldır
//	@Description	Bir konteyner kayıt defteri kimlik bilgilerini kaldır
//	@Param			server path	string	true	"Konteyner Kayıt Defteri sunucu adı"
//	@Success		204
//	@Router			/container-registry/{server} [delete]
//
//	@id				RemoveContainerRegistry
func RemoveContainerRegistry(ctx *gin.Context) {
	crServer := ctx.Param("server")

	decodedServerURL, err := url.QueryUnescape(crServer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("sunucu URL'si çözülemedi: %w", err))
		return
	}

	serverInstance := server.GetInstance(nil)

	err = serverInstance.ContainerRegistryService.Delete(decodedServerURL)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("konteyner kayıt defteri kaldırılamadı: %w", err))
		return
	}

	ctx.Status(http.StatusNoContent)
}
