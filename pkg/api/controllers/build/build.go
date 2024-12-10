// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package build

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/daytonaio/daytona/pkg/api/controllers/build/dto"
	"github.com/daytonaio/daytona/pkg/build"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server"
	builds_dto "github.com/daytonaio/daytona/pkg/server/builds/dto"
	"github.com/daytonaio/daytona/pkg/workspace/project/config"
	"github.com/gin-gonic/gin"
)

// CreateBuild godoc
//
//	@Tags			build
//	@Summary		Build oluştur
//	@Description	Build oluştur
//	@Accept			json
//	@Param			createBuildDto	body		CreateBuildDTO	true	"Create Build DTO"
//	@Success		201				{string}	buildId
//	@Router			/build [post]
//
//	@id				CreateBuild
func CreateBuild(ctx *gin.Context) {
	var createBuildDto dto.CreateBuildDTO
	err := ctx.BindJSON(&createBuildDto)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %s", err.Error()))
		return
	}

	s := server.GetInstance(nil)

	projectConfig, err := s.ProjectConfigService.Find(&config.ProjectConfigFilter{
		Name: &createBuildDto.ProjectConfigName,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("proje yapılandırması alınamadı: %s", err.Error()))
		return
	}

	gitProvider, _, err := s.GitProviderService.GetGitProviderForUrl(projectConfig.RepositoryUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("url için git sağlayıcısı alınamadı: %s", err.Error()))
		return
	}

	repo, err := gitProvider.GetRepositoryContext(gitprovider.GetRepositoryContext{
		Url:    projectConfig.RepositoryUrl,
		Branch: &createBuildDto.Branch,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("depo alınamadı: %s", err.Error()))
		return
	}

	newBuildDto := builds_dto.BuildCreationData{
		Image:       projectConfig.Image,
		User:        projectConfig.User,
		BuildConfig: projectConfig.BuildConfig,
		Repository:  repo,
		EnvVars:     createBuildDto.EnvVars,
	}

	if createBuildDto.PrebuildId != nil {
		newBuildDto.PrebuildId = *createBuildDto.PrebuildId
	}

	buildId, err := s.BuildService.Create(newBuildDto)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("build oluşturulamadı: %s", err.Error()))
		return
	}

	ctx.String(201, buildId)
}

// GetBuild godoc
//
//	@Tags			build
//	@Summary		Build verilerini al
//	@Description	Build verilerini al
//	@Accept			json
//	@Param			buildId	path		string	true	"Build ID"
//	@Success		200		{object}	Build
//	@Router			/build/{buildId} [get]
//
//	@id				GetBuild
func GetBuild(ctx *gin.Context) {
	buildId := ctx.Param("buildId")

	server := server.GetInstance(nil)

	b, err := server.BuildService.Find(&build.Filter{
		Id: &buildId,
	})
	if err != nil {
		statusCode := http.StatusInternalServerError
		if build.IsBuildNotFound(err) {
			statusCode = http.StatusNotFound
		}
		ctx.AbortWithError(statusCode, fmt.Errorf("build bulunamadı: %w", err))
		return
	}

	ctx.JSON(200, b)
}

// ListBuilds godoc
//
//	@Tags			build
//	@Summary		Buildleri listele
//	@Description	Buildleri listele
//	@Produce		json
//	@Success		200	{array}	Build
//	@Router			/build [get]
//
//	@id				ListBuilds
func ListBuilds(ctx *gin.Context) {
	server := server.GetInstance(nil)

	builds, err := server.BuildService.List(nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("buildler listelenemedi: %s", err.Error()))
		return
	}

	ctx.JSON(200, builds)
}

// DeleteAllBuilds godoc
//
//	@Tags			build
//	@Summary		TÜM buildleri sil
//	@Description	TÜM buildleri sil
//	@Param			force	query	bool	false	"Zorla"
//	@Success		204
//	@Router			/build [delete]
//
//	@id				DeleteAllBuilds
func DeleteAllBuilds(ctx *gin.Context) {
	forceQuery := ctx.Query("force")
	var force bool
	var err error

	if forceQuery != "" {
		force, err = strconv.ParseBool(forceQuery)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("force bayrağı için geçersiz değer"))
			return
		}
	}

	server := server.GetInstance(nil)

	errs := server.BuildService.MarkForDeletion(nil, force)
	if len(errs) > 0 {
		for _, err := range errs {
			_ = ctx.Error(err)
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(204)
}

// DeleteBuild godoc
//
//	@Tags			build
//	@Summary		Build sil
//	@Description	Build sil
//	@Param			buildId	path	string	true	"Build ID"
//	@Param			force	query	bool	false	"Zorla"
//	@Success		204
//	@Router			/build/{buildId} [delete]
//
//	@id				DeleteBuild
func DeleteBuild(ctx *gin.Context) {
	buildId := ctx.Param("buildId")
	forceQuery := ctx.Query("force")
	var force bool
	var err error

	if forceQuery != "" {
		force, err = strconv.ParseBool(forceQuery)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("force bayrağı için geçersiz değer"))
			return
		}
	}

	server := server.GetInstance(nil)

	errs := server.BuildService.MarkForDeletion(&build.Filter{
		Id: &buildId,
	}, force)
	if len(errs) > 0 {
		for _, err := range errs {
			_ = ctx.Error(err)
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(204)
}

// DeleteBuildsFromPrebuild godoc
//
//	@Tags			build
//	@Summary		Buildleri sil
//	@Description	Buildleri sil
//	@Param			prebuildId	path	string	true	"Prebuild ID"
//	@Param			force		query	bool	false	"Zorla"
//	@Success		204
//	@Router			/build/prebuild/{prebuildId} [delete]
//
//	@id				DeleteBuildsFromPrebuild
func DeleteBuildsFromPrebuild(ctx *gin.Context) {
	prebuildId := ctx.Param("prebuildId")
	forceQuery := ctx.Query("force")
	var force bool
	var err error

	if forceQuery != "" {
		force, err = strconv.ParseBool(forceQuery)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("force bayrağı için geçersiz değer"))
			return
		}
	}

	server := server.GetInstance(nil)

	// Prebuild mevcut değilse başarısız ol
	_, err = server.ProjectConfigService.FindPrebuild(nil, &config.PrebuildFilter{
		Id: &prebuildId,
	})
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("prebuild bulunamadı: %s", err.Error()))
		return
	}

	errs := server.BuildService.MarkForDeletion(&build.Filter{
		PrebuildIds: &[]string{prebuildId},
	}, force)
	if len(errs) > 0 {
		for _, err := range errs {
			_ = ctx.Error(err)
		}
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(204)
}
