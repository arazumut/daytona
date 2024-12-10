// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"fmt"
	"net/http"

	_ "github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// Git Olayını İşle 			godoc
//
//	@Tags			prebuild
//	@Summary		Git Olayını İşle
//	@Description	Git Olayını İşle
//	@Param			workspace	body	interface{}	true	"Webhook olayı"
//	@Success		200
//	@Router			/project-config/prebuild/process-git-event [post]
//
//	@id				ProcessGitEvent
func GitOlayiniIsle(ctx *gin.Context) {
	server := server.GetInstance(nil)

	gitSaglayici, err := server.GitProviderService.GetGitProviderForHttpRequest(ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("istek için git sağlayıcı alınamadı: %s", err.Error()))
		return
	}

	gitOlayVerisi, err := gitSaglayici.ParseEventData(ctx.Request)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("olay verisi parse edilemedi: %s", err.Error()))
		return
	}

	if gitOlayVerisi == nil {
		return
	}

	err = server.ProjectConfigService.ProcessGitEvent(*gitOlayVerisi)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("git olayı işlenemedi: %s", err.Error()))
		return
	}

	ctx.Status(200)
}
