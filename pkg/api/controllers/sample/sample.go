// Daytona Platforms Inc. 2024
// SPDX-License-Identifier: Apache-2.0

package sample

import (
	"net/http"

	"github.com/daytonaio/daytona/pkg/server"
	"github.com/gin-gonic/gin"
)

// ListSamples 			godoc
//
//	@Tags			sample
//	@Summary		Örnekleri listele
//	@Description	Örnekleri listele
//	@Produce		json
//	@Success		200	{array}	Sample
//	@Router			/sample [get]
//
//	@id				ListSamples
func ListSamples(ctx *gin.Context) {
	server := server.GetInstance(nil)

	samples, _, err := server.FetchSamples()
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(200, samples)
}
