// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package prebuild

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/server/projectconfig/dto"
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
	"github.com/gin-gonic/gin"
)

// PrebuildGet godoc
//
//	@Tags			prebuild
//	@Summary		Prebuild Getir
//	@Description	Prebuild Getir
//	@Accept			json
//	@Param			configName	path		string	true	"Proje yapılandırma adı"
//	@Param			prebuildId	path		string	true	"Prebuild ID"
//	@Success		200			{object}	PrebuildDTO
//	@Router			/project-config/{configName}/prebuild/{prebuildId} [get]
//
//	@id				PrebuildGet
func PrebuildGet(ctx *gin.Context) {
	configName := ctx.Param("configName")
	prebuildId := ctx.Param("prebuildId")

	server := server.GetInstance(nil)
	res, err := server.ProjectConfigService.FindPrebuild(&config.ProjectConfigFilter{
		Name: &configName,
	}, &config.PrebuildFilter{
		Id: &prebuildId,
	})
	if err != nil {
		if config.IsPrebuildNotFound(err) {
			ctx.AbortWithError(http.StatusNotFound, errors.New("prebuild bulunamadı"))
			return
		}
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("prebuild getirilemedi: %s", err.Error()))
		return
	}

	ctx.JSON(200, res)
}

// PrebuildSet godoc
//
//	@Tags			prebuild
//	@Summary		Prebuild Ayarla
//	@Description	Prebuild Ayarla
//	@Accept			json
//	@Param			configName	path		string				true	"Yapılandırma adı"
//	@Param			prebuild	body		CreatePrebuildDTO	true	"Prebuild"
//	@Success		201			{string}	prebuildId
//	@Router			/project-config/{configName}/prebuild [put]
//
//	@id				PrebuildSet
func PrebuildSet(ctx *gin.Context) {
	configName := ctx.Param("configName")

	var dto dto.CreatePrebuildDTO
	err := ctx.BindJSON(&dto)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %s", err.Error()))
		return
	}

	server := server.GetInstance(nil)
	prebuild, err := server.ProjectConfigService.SetPrebuild(configName, dto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("prebuild ayarlanamadı: %s", err.Error()))
		return
	}

	ctx.String(201, prebuild.Id)
}

// PrebuildList godoc
//
//	@Tags			prebuild
//	@Summary		Prebuildleri Listele
//	@Description	Prebuildleri Listele
//	@Accept			json
//	@Success		200	{array}	PrebuildDTO
//	@Router			/project-config/prebuild [get]
//
//	@id				PrebuildList
func PrebuildList(ctx *gin.Context) {
	server := server.GetInstance(nil)
	res, err := server.ProjectConfigService.ListPrebuilds(nil, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("prebuildler getirilemedi: %s", err.Error()))
		return
	}

	ctx.JSON(200, res)
}

// ProjectConfigPrebuildList godoc
//
//	@Tags			prebuild
//	@Summary		Proje yapılandırması için prebuildleri listele
//	@Description	Proje yapılandırması için prebuildleri listele
//	@Accept			json
//	@Param			configName	path	string	true	"Yapılandırma adı"
//	@Success		200			{array}	PrebuildDTO
//	@Router			/project-config/{configName}/prebuild [get]
//
//	@id				ProjectConfigPrebuildList
func ProjectConfigPrebuildList(ctx *gin.Context) {
	configName := ctx.Param("configName")

	var projectConfigFilter *config.ProjectConfigFilter

	if configName != "" {
		projectConfigFilter = &config.ProjectConfigFilter{
			Name: &configName,
		}
	}

	server := server.GetInstance(nil)
	res, err := server.ProjectConfigService.ListPrebuilds(projectConfigFilter, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("prebuildler getirilemedi: %s", err.Error()))
		return
	}

	ctx.JSON(200, res)
}

// PrebuildDelete godoc
//
//	@Tags			prebuild
//	@Summary		Prebuild Sil
//	@Description	Prebuild Sil
//	@Accept			json
//	@Param			configName	path	string	true	"Proje yapılandırma adı"
//	@Param			prebuildId	path	string	true	"Prebuild ID"
//	@Param			force		query	bool	false	"Zorla"
//	@Success		204
//	@Router			/project-config/{configName}/prebuild/{prebuildId} [delete]
//
//	@id				PrebuildDelete
func PrebuildDelete(ctx *gin.Context) {
	configName := ctx.Param("configName")
	prebuildId := ctx.Param("prebuildId")
	forceQuery := ctx.Query("force")

	var err error
	var force bool

	if forceQuery != "" {
		force, err = strconv.ParseBool(forceQuery)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("zorla bayrağı için geçersiz değer"))
			return
		}
	}

	server := server.GetInstance(nil)
	errs := server.ProjectConfigService.DeletePrebuild(configName, prebuildId, force)
	if len(errs) > 0 {
		if config.IsPrebuildNotFound(errs[0]) {
			ctx.AbortWithError(http.StatusNotFound, errors.New("prebuild bulunamadı"))
			return
		}
		for _, err := range errs {
			_ = ctx.Error(err)
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(204)
}
