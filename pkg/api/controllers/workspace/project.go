// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"fmt"
	"net/http"
	"time"

	"github.com/daytonaio/daytona/pkg/api/controllers/workspace/dto"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/workspace/project"
	"github.com/gin-gonic/gin"
)

// SetProjectState 			godoc
//
//	@Tags			workspace
//	@Summary		Proje durumunu ayarla
//	@Description	Proje durumunu ayarla
//	@Param			workspaceId	path	string			true	"Çalışma Alanı ID veya Adı"
//	@Param			projectId	path	string			true	"Proje ID"
//	@Param			setState	body	SetProjectState	true	"Durumu Ayarla"
//	@Success		200
//	@Router			/workspace/{workspaceId}/{projectId}/state [post]
//
//	@id				SetProjectState
func SetProjectState(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")
	projectId := ctx.Param("projectId")

	var setProjectStateDTO dto.SetProjectState
	err := ctx.BindJSON(&setProjectStateDTO)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %w", err))
		return
	}

	server := server.GetInstance(nil)

	_, err = server.WorkspaceService.SetProjectState(workspaceId, projectId, &project.ProjectState{
		Uptime:    setProjectStateDTO.Uptime,
		UpdatedAt: time.Now().Format(time.RFC1123),
		GitStatus: setProjectStateDTO.GitStatus,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("çalışma alanı %s durdurulamadı: %w", workspaceId, err))
		return
	}

	ctx.Status(200)
}
