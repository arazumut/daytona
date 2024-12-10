// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package target

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/internal/util/apiclient/conversion"
	"github.com/daytonaio/daytona/pkg/provider"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/server/providertargets/dto"
	"github.com/gin-gonic/gin"
)

// SetTarget godoc
//
//	@Tags			target
//	@Summary		Bir hedef belirle
//	@Description	Bir hedef belirle
//	@Param			target	body	CreateProviderTargetDTO	true	"Belirlenecek hedef"
//	@Success		201
//	@Router			/target [put]
//
//	@id				SetTarget
func SetTarget(ctx *gin.Context) {
	var req dto.CreateProviderTargetDTO
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %w", err))
		return
	}

	server := server.GetInstance(nil)

	target := conversion.ToProviderTarget(req)

	err = server.ProviderTargetService.Save(target)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("hedef belirlenemedi: %w", err))
		return
	}

	ctx.Status(201)
}

// SetDefaultTarget godoc
//
//	@Tags			target
//	@Summary		Hedefi varsayılan olarak ayarla
//	@Description	Hedefi varsayılan olarak ayarla
//	@Param			target	path	string	true	"Hedef adı"
//	@Success		200
//	@Router			/target/{target}/set-default [patch]
//
//	@id				SetDefaultTarget
func SetDefaultTarget(ctx *gin.Context) {
	targetName := ctx.Param("target")

	server := server.GetInstance(nil)

	target, err := server.ProviderTargetService.Find(&provider.TargetFilter{
		Name: &targetName,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("hedef bulunamadı: %w", err))
		return
	}

	err = server.ProviderTargetService.SetDefault(target)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("hedef varsayılan olarak ayarlanamadı: %s", err.Error()))
		return
	}

	ctx.Status(200)
}
