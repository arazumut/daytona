// Copyright 2024 Daytona Platforms Inc.
// SPDX-License-Identifier: Apache-2.0

package gitprovider

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/daytonaio/daytona/pkg/api/controllers"
	"github.com/daytonaio/daytona/pkg/api/controllers/gitprovider/dto"
	"github.com/daytonaio/daytona/pkg/apikey"
	"github.com/daytonaio/daytona/pkg/gitprovider"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// ListGitProviders godoc
//
//	@Tags			gitProvider
//	@Summary		Git sağlayıcılarını listele
//	@Description	Git sağlayıcılarını listele
//	@Produce		json
//	@Success		200	{array}	gitprovider.GitProviderConfig
//	@Router			/gitprovider [get]
//
//	@id				ListGitProviders
func ListGitProviders(ctx *gin.Context) {
	var response []*gitprovider.GitProviderConfig

	server := server.GetInstance(nil)

	response, err := server.GitProviderService.ListConfigs()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("git sağlayıcılarını listeleme başarısız: %w", err))
		return
	}

	for _, provider := range response {
		provider.Token = ""
		provider.SigningKey = nil
	}

	ctx.JSON(200, response)
}

// ListGitProvidersForUrl godoc
//
//	@Tags			gitProvider
//	@Summary		URL için Git sağlayıcılarını listele
//	@Description	URL için Git sağlayıcılarını listele
//	@Produce		json
//	@Param			url	path	string	true	"Url"
//	@Success		200	{array}	gitprovider.GitProviderConfig
//	@Router			/gitprovider/for-url/{url} [get]
//
//	@id				ListGitProvidersForUrl
func ListGitProvidersForUrl(ctx *gin.Context) {
	urlParam := ctx.Param("url")

	decodedUrl, err := url.QueryUnescape(urlParam)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("sorgu parametresi çözümlenemedi: %w", err))
		return
	}

	server := server.GetInstance(nil)

	gitProviders, err := server.GitProviderService.ListConfigsForUrl(decodedUrl)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("URL için git sağlayıcısı alınamadı: %w", err))
		return
	}

	apiKeyType, ok := ctx.Get("apiKeyType")
	if !ok || apiKeyType == apikey.ApiKeyTypeClient {
		for _, gitProvider := range gitProviders {
			gitProvider.Token = ""
		}
	}

	ctx.JSON(200, gitProviders)
}

// GetGitProvider godoc
//
//	@Tags			gitProvider
//	@Summary		Git sağlayıcısını al
//	@Description	Git sağlayıcısını al
//	@Produce		plain
//	@Param			gitProviderId	path		string	true	"ID"
//	@Success		200				{object}	gitprovider.GitProviderConfig
//	@Router			/gitprovider/{gitProviderId} [get]
//
//	@id				GetGitProvider
func GetGitProvider(ctx *gin.Context) {
	id := ctx.Param("gitProviderId")

	server := server.GetInstance(nil)

	gitProvider, err := server.GitProviderService.GetConfig(id)
	if err != nil {
		statusCode, message, codeErr := controllers.GetHTTPStatusCodeAndMessageFromError(err)
		if codeErr != nil {
			ctx.AbortWithError(statusCode, codeErr)
		}
		ctx.AbortWithError(statusCode, errors.New(message))
		return
	}

	apiKeyType, ok := ctx.Get("apiKeyType")
	if !ok || apiKeyType == apikey.ApiKeyTypeClient {
		gitProvider.Token = ""
	}

	ctx.JSON(200, gitProvider)
}

// GetGitProviderIdForUrl godoc
//
//	@Tags			gitProvider
//	@Summary		URL için Git sağlayıcı ID'sini al
//	@Description	URL için Git sağlayıcı ID'sini al
//	@Produce		plain
//	@Param			url	path		string	true	"Url"
//	@Success		200	{string}	providerId
//	@Router			/gitprovider/id-for-url/{url} [get]
//
//	@id				GetGitProviderIdForUrl
func GetGitProviderIdForUrl(ctx *gin.Context) {
	urlParam := ctx.Param("url")

	decodedUrl, err := url.QueryUnescape(urlParam)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, fmt.Errorf("sorgu parametresi çözümlenemedi: %w", err))
		return
	}

	server := server.GetInstance(nil)

	_, providerId, err := server.GitProviderService.GetGitProviderForUrl(decodedUrl)
	if err != nil {
		statusCode, message, codeErr := controllers.GetHTTPStatusCodeAndMessageFromError(err)
		if codeErr != nil {
			ctx.AbortWithError(statusCode, codeErr)
		}
		ctx.AbortWithError(statusCode, errors.New(message))
		return
	}

	ctx.String(200, providerId)
}

// SetGitProvider godoc
//
//	@Tags			gitProvider
//	@Summary		Git sağlayıcısını ayarla
//	@Description	Git sağlayıcısını ayarla
//	@Param			gitProviderConfig	body	SetGitProviderConfig	true	"Git sağlayıcı"
//	@Produce		json
//	@Success		200
//	@Router			/gitprovider [put]
//
//	@id				SetGitProvider
func SetGitProvider(ctx *gin.Context) {
	var setConfigDto dto.SetGitProviderConfig

	err := ctx.BindJSON(&setConfigDto)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("geçersiz istek gövdesi: %w", err))
		return
	}

	gitProviderConfig := gitprovider.GitProviderConfig{
		Id:            setConfigDto.Id,
		ProviderId:    setConfigDto.ProviderId,
		Token:         setConfigDto.Token,
		BaseApiUrl:    setConfigDto.BaseApiUrl,
		SigningKey:    setConfigDto.SigningKey,
		SigningMethod: setConfigDto.SigningMethod,
	}

	if setConfigDto.Username != nil {
		gitProviderConfig.Username = *setConfigDto.Username
	}

	if setConfigDto.Alias != nil {
		gitProviderConfig.Alias = *setConfigDto.Alias
	}

	server := server.GetInstance(nil)

	err = server.GitProviderService.SetGitProviderConfig(&gitProviderConfig)
	if err != nil {
		statusCode, message, codeErr := controllers.GetHTTPStatusCodeAndMessageFromError(err)
		if codeErr != nil {
			ctx.AbortWithError(statusCode, codeErr)
		}
		ctx.AbortWithError(statusCode, errors.New(message))
		return
	}

	ctx.JSON(200, nil)
}

// RemoveGitProvider godoc
//
//	@Tags			gitProvider
//	@Summary		Git sağlayıcısını kaldır
//	@Description	Git sağlayıcısını kaldır
//	@Param			gitProviderId	path	string	true	"Git sağlayıcı"
//	@Produce		json
//	@Success		200
//	@Router			/gitprovider/{gitProviderId} [delete]
//
//	@id				RemoveGitProvider
func RemoveGitProvider(ctx *gin.Context) {
	gitProviderId := ctx.Param("gitProviderId")

	server := server.GetInstance(nil)

	err := server.GitProviderService.RemoveGitProvider(gitProviderId)
	if err != nil {
		statusCode, message, codeErr := controllers.GetHTTPStatusCodeAndMessageFromError(err)
		if codeErr != nil {
			ctx.AbortWithError(statusCode, codeErr)
		}
		ctx.AbortWithError(statusCode, errors.New(message))
		return
	}

	ctx.JSON(200, nil)
}

// Sayfalama ile ilgili sorgu parametrelerini çıkar
func getListOptions(ctx *gin.Context) (gitprovider.ListOptions, error) {
	pageQuery := ctx.Query("page")
	perPageQuery := ctx.Query("per_page")

	page := 1
	perPage := 100
	var err error

	if pageQuery != "" {
		page, err = strconv.Atoi(pageQuery)
		if err != nil {
			return gitprovider.ListOptions{}, errors.New("geçersiz 'page' sorgu parametresi")
		}
	}

	if perPageQuery != "" {
		perPage, err = strconv.Atoi(perPageQuery)
		if err != nil {
			return gitprovider.ListOptions{}, errors.New("geçersiz 'per_page' sorgu parametresi")
		}
	}

	return gitprovider.ListOptions{
		Page:    page,
		PerPage: perPage,
	}, nil
}
