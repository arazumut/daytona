// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package projectconfig

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/daytonaio/daytona/internal/util"
	"github.com/daytonaio/daytona/internal/util/apiclient/conversion"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/daytonaio/daytona/pkg/server/projectconfig/dto"
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Proje Konfigürasyonunu Getir godoc
//
//	@Tags			proje-konfig
//	@Summary		Proje konfigürasyon verilerini getir
//	@Description	Proje konfigürasyon verilerini getir
//	@Accept			json
//	@Param			configName	path		string	true	"Konfigürasyon adı"
//	@Success		200			{object}	ProjectConfig
//	@Router			/proje-konfig/{configName} [get]
//
//	@id				GetProjectConfig
func GetProjectConfig(ctx *gin.Context) {
	configName := ctx.Param("configName")

	server := server.GetInstance(nil)

	projectConfig, err := server.ProjectConfigService.Find(&config.ProjectConfigFilter{
		Name: &configName,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("proje konfigürasyonunu getirme başarısız: %s", err.Error()))
		return
	}

	ctx.JSON(200, projectConfig)
}

// Varsayılan Proje Konfigürasyonunu Getir godoc
//
//	@Tags			proje-konfig
//	@Summary		Git URL'ine göre proje konfigürasyonlarını getir
//	@Description	Git URL'ine göre proje konfigürasyonlarını getir
//	@Produce		json
//	@Param			gitUrl	path		string	true	"Git URL"
//	@Success		200		{object}	ProjectConfig
//	@Router			/proje-konfig/varsayilan/{gitUrl} [get]
//
//	@id				GetDefaultProjectConfig
func GetDefaultProjectConfig(ctx *gin.Context) {
	gitUrl := ctx.Param("gitUrl")

	decodedURLParam, err := url.QueryUnescape(gitUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("sorgu parametresini çözme başarısız: %s", err.Error()))
		return
	}

	server := server.GetInstance(nil)

	projectConfigs, err := server.ProjectConfigService.Find(&config.ProjectConfigFilter{
		Url:     &decodedURLParam,
		Default: util.Pointer(true),
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if config.IsProjectConfigNotFound(err) {
			statusCode = http.StatusNotFound
			ctx.AbortWithStatus(statusCode)
			log.Debugf("Git URL için proje konfigürasyonu eklenmedi: %s", decodedURLParam)
			return
		}
		ctx.AbortWithError(statusCode, fmt.Errorf("git URL ile proje konfigürasyonu bulma başarısız: %s", err.Error()))
		return
	}

	ctx.JSON(200, projectConfigs)
}

// Proje Konfigürasyonlarını Listele godoc
//
//	@Tags			proje-konfig
//	@Summary		Proje konfigürasyonlarını listele
//	@Description	Proje konfigürasyonlarını listele
//	@Produce		json
//	@Success		200	{array}	ProjectConfig
//	@Router			/proje-konfig [get]
//
//	@id				ListProjectConfigs
func ListProjectConfigs(ctx *gin.Context) {
	server := server.GetInstance(nil)

	projectConfigs, err := server.ProjectConfigService.List(nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("proje konfigürasyonlarını listeleme başarısız: %s", err.Error()))
		return
	}

	ctx.JSON(200, projectConfigs)
}

// Proje Konfigürasyonunu Ayarla godoc
//
//	@Tags			proje-konfig
//	@Summary		Proje konfigürasyon verilerini ayarla
//	@Description	Proje konfigürasyon verilerini ayarla
//	@Accept			json
//	@Param			projectConfig	body	CreateProjectConfigDTO	true	"Proje konfigürasyonu"
//	@Success		201
//	@Router			/proje-konfig [put]
//
//	@id				SetProjectConfig
func SetProjectConfig(ctx *gin.Context) {
	var req dto.CreateProjectConfigDTO
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %s", err.Error()))
		return
	}

	s := server.GetInstance(nil)

	projectConfig := conversion.ToProjectConfig(req)

	err = s.ProjectConfigService.Save(projectConfig)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("proje konfigürasyonunu kaydetme başarısız: %s", err.Error()))
		return
	}

	ctx.Status(201)
}

// Varsayılan Proje Konfigürasyonunu Ayarla godoc
//
//	@Tags			proje-konfig
//	@Summary		Proje konfigürasyonunu varsayılan olarak ayarla
//	@Description	Proje konfigürasyonunu varsayılan olarak ayarla
//	@Param			configName	path	string	true	"Konfigürasyon adı"
//	@Success		200
//	@Router			/proje-konfig/{configName}/varsayilan-olarak-ayarla [patch]
//
//	@id				SetDefaultProjectConfig
func SetDefaultProjectConfig(ctx *gin.Context) {
	configName := ctx.Param("configName")

	server := server.GetInstance(nil)

	err := server.ProjectConfigService.SetDefault(configName)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("proje konfigürasyonunu varsayılan olarak ayarlama başarısız: %s", err.Error()))
		return
	}

	ctx.Status(200)
}

// Proje Konfigürasyonunu Sil godoc
//
//	@Tags			proje-konfig
//	@Summary		Proje konfigürasyon verilerini sil
//	@Description	Proje konfigürasyon verilerini sil
//	@Param			configName	path	string	true	"Konfigürasyon adı"
//	@Param			force		query	bool	false	"Zorla"
//	@Success		204
//	@Router			/proje-konfig/{configName} [delete]
//
//	@id				DeleteProjectConfig
func DeleteProjectConfig(ctx *gin.Context) {
	configName := ctx.Param("configName")
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

	projectConfig, err := server.ProjectConfigService.Find(&config.ProjectConfigFilter{
		Name: &configName,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, fmt.Errorf("proje konfigürasyonunu bulma başarısız: %s", err.Error()))
		return
	}

	errs := server.ProjectConfigService.Delete(projectConfig.Name, force)
	if len(errs) > 0 {
		if config.IsProjectConfigNotFound(errs[0]) {
			ctx.AbortWithError(http.StatusNotFound, errors.New("proje konfigürasyonu bulunamadı"))
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
