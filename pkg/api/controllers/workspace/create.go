// Daytona Platforms Inc. 2024 Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"fmt"
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/server/workspaces/dto"
	"github.com/gin-gonic/gin"
)

// CreateWorkspace 			godoc
//
//	@Tags			workspace
//	@Summary		Bir çalışma alanı oluştur
//	@Description	Bir çalışma alanı oluştur
//	@Param			workspace	body	CreateWorkspaceDTO	true	"Çalışma alanı oluştur"
//	@Produce		json
//	@Success		200	{object}	Workspace
//	@Router			/workspace [post]
//
//	@id				CreateWorkspace
func CreateWorkspace(ctx *gin.Context) {
	var createWorkspaceReq dto.CreateWorkspaceDTO
	err := ctx.BindJSON(&createWorkspaceReq)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %w", err))
		return
	}

	server := server.GetInstance(nil)

	w, err := server.WorkspaceService.CreateWorkspace(ctx.Request.Context(), createWorkspaceReq)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("çalışma alanı oluşturulamadı: %w", err))
		return
	}

	ctx.JSON(200, w)
}
