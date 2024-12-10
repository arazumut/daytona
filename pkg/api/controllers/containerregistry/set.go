// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package containerregistry

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/daytonaio/daytona/pkg/containerregistry"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// SetContainerRegistry godoc
//
//	@Tags			container-registry
//	@Summary		Konteyner kayıt defteri kimlik bilgilerini ayarla
//	@Description	Konteyner kayıt defteri kimlik bilgilerini ayarla
//	@Param			server				path	string				true	"Konteyner Kayıt Defteri sunucu adı"
//	@Param			containerRegistry	body	ContainerRegistry	true	"Konteyner Kayıt Defteri kimlik bilgilerini ayarla"
//	@Success		201
//	@Router			/container-registry/{server} [put]
//
//	@id				SetContainerRegistry
func SetContainerRegistry(ctx *gin.Context) {
	crServer := ctx.Param("server")

	decodedServerURL, err := url.QueryUnescape(crServer)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("sunucu URL'si çözülemedi: %w", err))
		return
	}

	var req containerregistry.ContainerRegistry
	err = ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %w", err))
		return
	}

	server := server.GetInstance(nil)

	cr, err := server.ContainerRegistryService.Find(decodedServerURL)
	if err == nil {
		err = server.ContainerRegistryService.Delete(decodedServerURL)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("konteyner kayıt defteri kaldırılamadı: %w", err))
			return
		}

		cr.Server = req.Server
		cr.Username = req.Username
		cr.Password = req.Password
	} else {
		cr = &req
	}

	err = server.ContainerRegistryService.Save(cr)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("konteyner kayıt defteri ayarlanamadı: %w", err))
		return
	}

	ctx.Status(201)
}
