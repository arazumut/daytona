// Copyright 2024 Daytona Platforms Inc.
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

// GetRepoPRs 			godoc
//
//	@Tags			gitProvider
//	@Summary		Git deposu PR'larını al
//	@Description	Git deposu PR'larını al
//	@Param			gitProviderId	path	string	true	"Git sağlayıcı"
//	@Param			namespaceId		path	string	true	"Namespace"
//	@Param			repositoryId	path	string	true	"Repository"
//	@Param			page			query	int		false	"Sayfa numarası"
//	@Param			per_page		query	int		false	"Sayfa başına öğe sayısı"
//	@Produce		json
//	@Success		200	{array}	GitPullRequest
//	@Router			/gitprovider/{gitProviderId}/{namespaceId}/{repositoryId}/pull-requests [get]
//
//	@id				GetRepoPRs
func GetRepoPRs(ctx *gin.Context) {
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

	response, err := server.GitProviderService.GetRepoPRs(gitProviderId, namespaceId, repositoryId, options)
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
