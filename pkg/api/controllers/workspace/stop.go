// Daytona Platforms Inc. 2024 Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// StopWorkspace 			godoc
//
//	@Tags			workspace
//	@Summary		Çalışma alanını durdur
//	@Description	Çalışma alanını durdur
//	@Param			workspaceId	path	string	true	"Çalışma Alanı ID veya Adı"
//	@Success		200
//	@Router			/workspace/{workspaceId}/stop [post]
//
//	@id				StopWorkspace
func StopWorkspace(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")

	server := server.GetInstance(nil)

	err := server.WorkspaceService.StopWorkspace(ctx.Request.Context(), workspaceId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("çalışma alanı %s durdurulamadı: %w", workspaceId, err))
		return
	}

	ctx.Status(200)
}

// StopProject 			godoc
//
//	@Tags			workspace
//	@Summary		Projeyi durdur
//	@Description	Projeyi durdur
//	@Param			workspaceId	path	string	true	"Çalışma Alanı ID veya Adı"
//	@Param			projectId	path	string	true	"Proje ID"
//	@Success		200
//	@Router			/workspace/{workspaceId}/{projectId}/stop [post]
//
//	@id				StopProject
func StopProject(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")
	projectId := ctx.Param("projectId")

	server := server.GetInstance(nil)

	err := server.WorkspaceService.StopProject(ctx.Request.Context(), workspaceId, projectId)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("proje %s durdurulamadı: %w", projectId, err))
		return
	}

	ctx.Status(200)
}
