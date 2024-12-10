// Daytona Platforms Inc. 2024 Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package gitprovider

import (
	"errors"
	"net/http"

	"github.com/daytonaio/daytona/pkg/api/controllers"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// GetRepositories 			godoc
//
//	@Tags			gitProvider
//	@Summary		Git depolarını getir
//	@Description	Git depolarını getir
//	@Param			gitProviderId	path	string	true	"Git sağlayıcı"
//	@Param			namespaceId		path	string	true	"Namespace"
//	@Param			page			query	int		false	"Sayfa numarası"
//	@Param			per_page		query	int		false	"Sayfa başına öğe sayısı"
//	@Produce		json
//	@Success		200	{array}	GitRepository
//	@Router			/gitprovider/{gitProviderId}/{namespaceId}/repositories [get]
//
//	@id				GetRepositories
func GetRepositories(ctx *gin.Context) {
	gitProviderId := ctx.Param("gitProviderId")
	namespaceId := ctx.Param("namespaceId")
	options, err := getListOptions(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	serverInstance := server.GetInstance(nil)

	response, err := serverInstance.GitProviderService.GetRepositories(gitProviderId, namespaceId, options)
	if err != nil {
		statusCode, message, codeErr := controllers.GetHTTPStatusCodeAndMessageFromError(err)
		if codeErr != nil {
			ctx.AbortWithError(statusCode, codeErr)
		}
		ctx.AbortWithError(statusCode, errors.New(message))
		return
	}

	ctx.JSON(http.StatusOK, response)
}
