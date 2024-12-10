// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package workspace

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// GetWorkspace 			godoc
//
//	@Tags			workspace
//	@Summary		Çalışma alanı bilgisini getir
//	@Description	Çalışma alanı bilgisini getir
//	@Produce		json
//	@Param			workspaceId	path		string	true	"Çalışma Alanı ID veya Adı"
//	@Param			verbose		query		bool	false	"Ayrıntılı"
//	@Success		200			{object}	WorkspaceDTO
//	@Router			/workspace/{workspaceId} [get]
//
//	@id				GetWorkspace
func GetWorkspace(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")
	verboseQuery := ctx.Query("verbose")
	verbose := false
	var err error

	if verboseQuery != "" {
		verbose, err = strconv.ParseBool(verboseQuery)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("ayrıntılı bayrağı için geçersiz değer"))
			return
		}
	}

	server := server.GetInstance(nil)

	w, err := server.WorkspaceService.GetWorkspace(ctx.Request.Context(), workspaceId, verbose)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("çalışma alanı getirilemedi: %w", err))
		return
	}

	ctx.JSON(200, w)
}

// ListWorkspaces 			godoc
//
//	@Tags			workspace
//	@Summary		Çalışma alanlarını listele
//	@Description	Çalışma alanlarını listele
//	@Produce		json
//	@Success		200	{array}	WorkspaceDTO
//	@Router			/workspace [get]
//	@Param			verbose	query	bool	false	"Ayrıntılı"
//
//	@id				ListWorkspaces
func ListWorkspaces(ctx *gin.Context) {
	verboseQuery := ctx.Query("verbose")
	verbose := false
	var err error

	if verboseQuery != "" {
		verbose, err = strconv.ParseBool(verboseQuery)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("ayrıntılı bayrağı için geçersiz değer"))
			return
		}
	}

	server := server.GetInstance(nil)

	workspaceList, err := server.WorkspaceService.ListWorkspaces(ctx.Request.Context(), verbose)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("çalışma alanları listelenemedi: %w", err))
		return
	}

	ctx.JSON(200, workspaceList)
}

// RemoveWorkspace 			godoc
//
//	@Tags			workspace
//	@Summary		Çalışma alanını kaldır
//	@Description	Çalışma alanını kaldır
//	@Param			workspaceId	path	string	true	"Çalışma Alanı ID"
//	@Param			force		query	bool	false	"Zorla"
//	@Success		200
//	@Router			/workspace/{workspaceId} [delete]
//
//	@id				RemoveWorkspace
func RemoveWorkspace(ctx *gin.Context) {
	workspaceId := ctx.Param("workspaceId")
	forceQuery := ctx.Query("force")
	var err error
	force := false

	if forceQuery != "" {
		force, err = strconv.ParseBool(forceQuery)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("zorla bayrağı için geçersiz değer"))
			return
		}
	}

	server := server.GetInstance(nil)

	if force {
		err = server.WorkspaceService.ForceRemoveWorkspace(ctx.Request.Context(), workspaceId)
	} else {
		err = server.WorkspaceService.RemoveWorkspace(ctx.Request.Context(), workspaceId)
	}

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("çalışma alanı kaldırılamadı: %w", err))
		return
	}

	ctx.Status(200)
}
