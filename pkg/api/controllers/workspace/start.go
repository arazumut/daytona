// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// StartWorkspace 			godoc
//
//	@Tags			workspace
//	@Summary		Çalışma alanını başlat
//	@Description	Çalışma alanını başlat
//	@Param			workspaceId	path	string	true	"Çalışma Alanı ID veya Adı"
//	@Success		200
//	@Router			/workspace/{workspaceId}/start [post]
//
//	@id				StartWorkspace
func StartWorkspace(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")

	server := server.GetInstance(nil)

	err := server.WorkspaceService.StartWorkspace(ctx.Request.Context(), workspaceId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("çalışma alanı %s başlatılamadı: %w", workspaceId, err))
		return
	}

	ctx.Status(200)
}

// StartProject 			godoc
//
//	@Tags			workspace
//	@Summary		Projeyi başlat
//	@Description	Projeyi başlat
//	@Param			workspaceId	path	string	true	"Çalışma Alanı ID veya Adı"
//	@Param			projectId	path	string	true	"Proje ID"
//	@Success		200
//	@Router			/workspace/{workspaceId}/{projectId}/start [post]
//
//	@id				StartProject
func StartProject(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")
	projectId := ctx.Param("projectId")

	server := server.GetInstance(nil)

	err := server.WorkspaceService.StartProject(ctx.Request.Context(), workspaceId, projectId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("proje %s başlatılamadı: %w", projectId, err))
		return
	}

	ctx.Status(200)
}
