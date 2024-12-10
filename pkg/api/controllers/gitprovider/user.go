// Daytona Platforms Inc. 2024 Telif Hakkı
// SPDX-License-Identifier: Apache-2.0

package gitprovider

import (
	"errors"

	"github.com/daytonaio/daytona/pkg/api/controllers"
	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// GetGitUser 			godoc
//
//	@Tags			gitProvider
//	@Summary		Git Kullanıcı Bilgilerini Getir
//	@Description	Git Kullanıcı Bilgilerini Getir
//	@Produce		json
//	@Param			gitProviderId	path		string	true	"Git Sağlayıcı Id"
//	@Success		200				{object}	GitUser
//	@Router			/gitprovider/{gitProviderId}/user [get]
//
//	@id				GetGitUser
func GetGitUser(ctx *gin.Context) {
	gitProviderId := ctx.Param("gitProviderId")

	server := server.GetInstance(nil)

	response, err := server.GitProviderService.GetGitUser(gitProviderId)
	if err != nil {
		statusCode, message, codeErr := controllers.GetHTTPStatusCodeAndMessageFromError(err)
		if codeErr != nil {
			ctx.AbortWithError(statusCode, codeErr)
		} else {
			ctx.AbortWithError(statusCode, errors.New(message))
		}
		return
	}

	ctx.JSON(200, response)
}
