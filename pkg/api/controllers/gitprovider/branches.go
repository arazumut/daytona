// Daytona Platforms Inc. 2024 Tüm Hakları Saklıdır.
// SPDX-License-Identifier: Apache-2.0

package gitprovider

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/daytonaio/daytona/pkg/api/controllers"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// GetRepoBranches 			godoc
//
//	@Tags			gitProvider
//	@Summary		Git deposu dallarını getir
//	@Description	Git deposu dallarını getir
//	@Param			gitProviderId	path	string	true	"Git sağlayıcı"
//	@Param			namespaceId		path	string	true	"Namespace"
//	@Param			repositoryId	path	string	true	"Repository"
//	@Param			page			query	int		false	"Sayfa numarası"
//	@Param			per_page		query	int		false	"Sayfa başına öğe sayısı"
//	@Produce		json
//	@Success		200	{array}	GitBranch
//	@Router			/gitprovider/{gitProviderId}/{namespaceId}/{repositoryId}/branches [get]
//
//	@id				GetRepoBranches
func GetRepoBranches(ctx *gin.Context) {
	gitProviderId := ctx.Param("gitProviderId")
	namespaceArg := ctx.Param("namespaceId")
	repositoryArg := ctx.Param("repositoryId")
	options, err := getListOptions(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	namespaceId, err := url.QueryUnescape(namespaceArg)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("namespace çözümlenemedi: %w", err))
		return
	}

	repositoryId, err := url.QueryUnescape(repositoryArg)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("repository çözümlenemedi: %w", err))
		return
	}

	server := server.GetInstance(nil)

	response, err := server.GitProviderService.GetRepoBranches(gitProviderId, namespaceId, repositoryId, options)
	if err != nil {
		statusCode, message, codeErr := controllers.GetHTTPStatusCodeAndMessageFromError(err)
		if codeErr != nil {
			ctx.AbortWithError(statusCode, codeErr)
		}
		ctx.AbortWithError(statusCode, errors.New(message))
		return
	}

	ctx.JSON(200, response)
}
